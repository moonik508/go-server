package request

import "time"

type TestDataFormat struct {
	Id    string `json:"id" binding:"required"`
	Title string
	Date  time.Time
}
