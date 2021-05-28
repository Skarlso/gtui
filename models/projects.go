package models

// Project defines details about a github project.
type Project struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

// ProjectColumn .
type ProjectColumn struct {
	ID                 int64
	Name               string
	ProjectColumnCards []*ProjectColumnCard
}

// ProjectColumnCard .
type ProjectColumnCard struct {
	ID       int64
	IssueID  int64
	Content  string
	Title    string
	Name     string
	Note     *string
	Author   string
	Assignee string
}

// ProjectData .
type ProjectData struct {
	ProjectColumns []*ProjectColumn
}
