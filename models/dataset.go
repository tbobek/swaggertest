package models

type Dataset struct {
	// The ID of a dataaset
	// example: 130670
	Id string `json:"id"`
	// Actual data
	// example: { "key": "dataset1" }
	Data int `json:"data"`
}

type ErrorMsg struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}
