package services

import (
	"encoding/json"
	"fmt"
	"teachable-test/src/interfaces"
)

func RetrieveCourses() (interfaces.Courses, error) {
	var courses interfaces.Courses

	responseCourses, err := TeachableHttpGet("/courses?is_published=true")
	if err != nil {
		return courses, err
	}
	err = json.Unmarshal(responseCourses, &courses)
	if err != nil {
		fmt.Println(err)
		return courses, err
	}
	return courses, nil
}

func RetrieveUsers() (interfaces.Users, error) {
	var users interfaces.Users

	responseUsers, err := TeachableHttpGet("/users")
	if err != nil {
		return users, err
	}
	json.Unmarshal(responseUsers, &users)
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	return users, nil
}

func RetrieveEnrollments(courseId int32) (interfaces.Enrollments, error) {
	var enrollments interfaces.Enrollments
	courseIdString := fmt.Sprint(courseId)
	responseEnrollments, err := TeachableHttpGet("/courses/" + courseIdString + "/enrollments")
	if err != nil {
		return enrollments, err
	}
	json.Unmarshal(responseEnrollments, &enrollments)
	if err != nil {
		fmt.Println(err)
		return enrollments, err
	}
	return enrollments, nil
}
