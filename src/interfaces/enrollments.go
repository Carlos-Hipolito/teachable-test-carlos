package interfaces

type Enrollments struct {
	Enrollments []Enrollment `json:"enrollments"`
	Meta        Meta         `json:"meta"`
}

type Enrollment struct {
	User_id int32 `json:"user_id"`
}
