package handler

type Urls struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Real_url    string `json:"real_url" binding:"required"`
	Shorter_url string `json:"shorter_url"`
}
