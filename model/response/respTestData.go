package response

import "time"

type TestDataFormat struct {
	Key   string    `json:"key"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}
