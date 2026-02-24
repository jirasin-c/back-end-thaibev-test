package storage

import "github.com/jmoiron/sqlx"

type OccupationStore interface {
	List() ([]Occupation, error)
	Exists(code string) (bool, error)
}

type occupationStore struct{ db *sqlx.DB }

func NewOccupationStore(db *sqlx.DB) OccupationStore {
	return &occupationStore{db: db}
}

func (s *occupationStore) List() ([]Occupation, error) {
	var out []Occupation
	err := s.db.Select(&out, `SELECT code, name FROM thaibev_test.occupations ORDER BY name`)
	return out, err
}

func (s *occupationStore) Exists(code string) (bool, error) {
	var n int
	err := s.db.Get(&n, `SELECT COUNT(1) FROM thaibev_test.occupations WHERE code=$1`, code)
	return n > 0, err
}
