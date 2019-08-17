package models

type GitResponse struct {
	TotalCount int        `json:"total_count"`
	Items      []GitItems `json:"items"`
}

type GitItems struct {
	ID          int     `json:"id"`
	FullName    string  `json:"full_name"`
	HtmlURL     string  `json:"html_url"`
	Description string  `json:"description"`
	Language    string  `json:"language"`
	Owner       GitUser `json:"owner"`
}

type GitUser struct {
	ID   int    `json:"id"`
	Name string `json:"login"`
	URL  string `json:"url"`
}
