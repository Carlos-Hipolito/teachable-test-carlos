package interfaces

type Courses struct {
	Courses []Course `json:"courses"`
	Meta    Meta     `json:"meta"`
}

type Course struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	Heading      string `json:"heading"`
	Description  string `json:"description"`
	Is_published bool   `json:"is_published"`
	Image_url    string `json:"image_url"`
}
