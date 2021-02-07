package models

//Biodata is ...
type Biodata struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Profile  profile   `json:"profile"`
	Articles []Article `json:"articles:"`
}

type profile struct {
	Fullname string   `json:"full_name"`
	Birthday string   `json:"birthday"`
	Phones   []string `json:"phones"`
}

//Article is ...
type Article struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	PublishedAt string `json:"published_at"`
}
