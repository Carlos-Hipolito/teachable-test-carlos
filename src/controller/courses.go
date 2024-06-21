package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"teachable-test/src/interfaces"
	"teachable-test/src/services"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {

	courses, err := services.RetrieveCourses()
	if err != nil {
		fmt.Fprint(w, "error retrieving courses from teachable api: ", err)
		w.WriteHeader(500)
	}
	users, err := services.RetrieveUsers()
	if err != nil {
		fmt.Fprint(w, "error retrieving users from teachable api: ", err)
		w.WriteHeader(500)
	}

	var response_courses interfaces.ResponseCourses
	for _, s := range courses.Courses {
		var response_course interfaces.ResponseCourse
		enrollments, err := services.RetrieveEnrollments(s.Id)
		if err != nil {
			fmt.Fprint(w, "error retrieving enrollment from teachable api: ", err)
			w.WriteHeader(500)
		}
		response_course.Course_name = s.Name
		response_course.Course_heading = s.Heading
		for _, i := range enrollments.Enrollments {
			for _, u := range users.Users {
				if u.Id == i.User_id {
					response_course.Enrolled_students = append(response_course.Enrolled_students, u)
				}
			}
		}
		response_courses = append(response_courses, response_course)
	}
	fmt.Println(response_courses)
	response, _ := json.Marshal(response_courses)
	w.Write(response)
	w.WriteHeader(200)
}
