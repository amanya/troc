package mocks

import (
	"time"

	"troc.amanya/internal/models"
)

var mockTroc = models.Troc{
	ID: 1,
	Title: "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type TrocModel struct{}

func (m *TrocModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *TrocModel) Get(id int) (models.Troc, error) {
	switch id {
	case 1:
		return mockTroc, nil
	default:
		return models.Troc{}, models.ErrNoRecord
	}
}

func (m *TrocModel) Latest() ([]models.Troc, error) {
	return []models.Troc{mockTroc}, nil
}