package storage

import "github.com/jmoiron/sqlx"

type PeopleStore interface {
	Insert(p PersonInsert) (string, error)
}

type peopleStore struct{ db *sqlx.DB }

func NewPeopleStore(db *sqlx.DB) PeopleStore {
	return &peopleStore{db: db}
}

func (s *peopleStore) Insert(p PersonInsert) (string, error) {
	var id string
	q := `
	INSERT INTO thaibev_test.people_profiles (
	  first_name,last_name,email,phone,birth_day,sex,occupation_code,
	  profile_file_name,profile_base64
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id;
	`
	err := s.db.Get(&id, q,
		p.FirstName, p.LastName, p.Email, p.Phone, p.BirthDay, p.Sex, p.OccupationCode,
		p.ProfileFileName, p.ProfileBase64,
	)
	return id, err
}
