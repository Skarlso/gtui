package models

// Project defines details about a github project.
type Project struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}
