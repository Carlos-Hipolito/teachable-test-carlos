package interfaces

type ResponseCourses []ResponseCourse

type ResponseCourse struct {
	Course_name       string `json:"course_name"`
	Course_heading    string `json:"course_heading"`
	Enrolled_students []User `json:"enrolled_students"`
}
