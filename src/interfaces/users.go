package interfaces

type Users struct {
	Users []User `json:"users"`
	Meta  Meta   `json:"meta"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    int32  `json:"id"`
}
