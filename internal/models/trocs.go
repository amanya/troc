package models

import (
	"database/sql"
	"errors"
	"time"
)

type TrocModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (Troc, error)
	Latest() ([]Troc, error)
}

type Troc struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type TrocModel struct {
	DB *sql.DB
}

func (m *TrocModel) Insert(title string, content string, expires int)(int, error) {
	stmt := `INSERT INTO trocs (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TrocModel) Get(id int)(Troc, error) {
	stmt := `SELECT id, title, content, created, expires FROM trocs
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	var t Troc

	err := row.Scan(&t.ID, &t.Title, &t.Content, &t.Created, &t.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Troc{}, ErrNoRecord
		} else {
			return Troc{}, nil
		}
	}
	return t, nil
}

func (m *TrocModel) Latest() ([]Troc, error) {
	stmt := `SELECT id, title, content, created, expires FROM trocs
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var trocs []Troc

	for rows.Next() {
		var t Troc

		err = rows.Scan(&t.ID, &t.Title, &t.Content, &t.Created, &t.Expires)
		if err != nil {
			return nil, err
		}

		trocs = append(trocs, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trocs, nil
}