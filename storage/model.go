package storage

import "time"

type Occupation struct {
	Code string `db:"code"`
	Name string `db:"name"`
}

type PersonInsert struct {
	ID              string    `db:"id"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	Email           string    `db:"email"`
	Phone           string    `db:"phone"`
	BirthDay        time.Time `db:"birth_day"`
	Sex             string    `db:"sex"`
	OccupationCode  string    `db:"occupation_code"`
	ProfileFileName string    `db:"profile_file_name"`
	ProfileBase64   string    `db:"profile_base64"`
}
