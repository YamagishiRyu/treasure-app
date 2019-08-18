package models

import (
	"io/ioutil"

	"github.com/labstack/gommon/log"
)

type Mdfile struct {
	ID           int64  `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Path         string `db:"path" json:"path"`
	URL          string `db:"url" json:"url"`
	RepositoryID int64  `db:"repository_id" json:"repository_id"`
}

func (m *Mdfile) GetText() (string, error) {
	data, err := ioutil.ReadFile(m.Path)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return string(data), nil
}

func (m *Mdfile) MapToMdfileWithText() (*MdfileWithText, error) {
	text, err := m.GetText()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &MdfileWithText{
		ID:           m.ID,
		Name:         m.Name,
		Text:         text,
		URL:          m.URL,
		RepositoryID: m.RepositoryID,
	}, nil
}
