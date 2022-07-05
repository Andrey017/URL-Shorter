package url_service

type Stats struct {
	Id              int    `json:"id"`
	Url_id          int    `json:"url_id" binding:"required"`
	Counttransition int    `json:"counttransition" binding:"required"`
	Date            int64  `json:"date" binding:"required"`
	DateDB          string `json:"dateDB"`
}

type StatsDetail struct {
	Statid   int    `json:"statid"`
	City     int    `json:"city" binding:"required"`
	Browser  string `json:"browser" binding:"required"`
	Os       string `json:"os" binding:"required"`
	IsMobile bool   `json:"ismobile" binding:"required"`
}

type RegionsCode struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"id" binding:"required"`
	Id   int    `json:"value" binding:"required"`
}

type StatsBrowser struct {
	Count   int    `json:"count" binding:"required"`
	Browser string `json:"browser" binding:"required"`
}

type StatsOs struct {
	Os    string `json:"os" binding:"required"`
	Count int    `json:"count" binding:"required"`
}

type StatCountType struct {
	IsMobile bool `json:"ismobile"`
	Count    int  `json:"count"`
}

type StatCountIsMobile struct {
	Os    string `json:"os"`
	Count int    `json:"count"`
}
