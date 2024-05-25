package model

import "time"

// Report struct model
type Report struct {
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"date_created"`
}

type ReportCreatedResponse struct {
	Code        string    `json:"code"`
	DateCreated time.Time `json:"date_created"`
}
