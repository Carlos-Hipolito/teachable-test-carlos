package interfaces

type Meta struct {
	Total           int32 `json:"total"`
	Page            int32 `json:"page"`
	From            int32 `json:"from"`
	To              int32 `json:"to"`
	Per_page        int32 `json:"per_page"`
	Number_of_pages int32 `json:"number_of_pages"`
}
