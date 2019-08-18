package models

type Repository struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Path string `db:"path" json:"path"`
	URL  string `db:"url" json:"url"`
}

type RepositoryDetail struct {
	ID      int64            `json:"id"`
	Name    string           `json:"name"`
	Path    string           `json:"path"`
	URL     string           `json:"url"`
	Mdfiles []MdfileWithText `json:"mdfiles"`
}

func (r *Repository) MapToDetail(ms []MdfileWithText) *RepositoryDetail {
	return &RepositoryDetail{
		ID:      r.ID,
		Name:    r.Name,
		Path:    r.Path,
		URL:     r.URL,
		Mdfiles: ms,
	}
}
