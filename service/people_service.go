package service

import (
	"errors"
	"strings"
	"thaibev-test/storage"
	"time"
)

type PeopleService interface {
	Create(req CreatePersonRequest) (CreatePersonResponse, error)
}

type peopleService struct {
	peopleStore storage.PeopleStore
	occStore    storage.OccupationStore
}

func NewPeopleService(p storage.PeopleStore, o storage.OccupationStore) PeopleService {
	return &peopleService{peopleStore: p, occStore: o}
}

var (
	ErrBadRequest = errors.New("bad_request")
	ErrNotFound   = errors.New("not_found")
)

func (s *peopleService) Create(req CreatePersonRequest) (CreatePersonResponse, error) {
	// Required fields check (เพราะ handler อาจจะ bind ได้แต่ไม่ validate)
	if strings.TrimSpace(req.FirstName) == "" ||
		strings.TrimSpace(req.LastName) == "" ||
		strings.TrimSpace(req.Email) == "" ||
		strings.TrimSpace(req.Phone) == "" ||
		strings.TrimSpace(req.BirthDay) == "" ||
		strings.TrimSpace(req.OccupationCode) == "" ||
		strings.TrimSpace(req.ProfileFileName) == "" ||
		strings.TrimSpace(req.ProfileBase64) == "" ||
		(req.Sex != "M" && req.Sex != "F") {
		return CreatePersonResponse{}, ErrBadRequest
	}

	if !phoneRe.MatchString(req.Phone) {
		return CreatePersonResponse{}, errors.New("invalid phone format")
	}

	bd, err := parseBirthDay(req.BirthDay)
	if err != nil {
		return CreatePersonResponse{}, errors.New("invalid birthDay format (dd/MM/yyyy)")
	}

	if err := validateBase64(req.ProfileBase64); err != nil {
		return CreatePersonResponse{}, err
	}

	ok, err := s.occStore.Exists(req.OccupationCode)
	if err != nil {
		return CreatePersonResponse{}, err
	}
	if !ok {
		return CreatePersonResponse{}, errors.New("occupation not found")
	}

	id, err := s.peopleStore.Insert(storage.PersonInsert{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Email:           req.Email,
		Phone:           req.Phone,
		BirthDay:        bd,
		Sex:             req.Sex,
		OccupationCode:  req.OccupationCode,
		ProfileFileName: req.ProfileFileName,
		ProfileBase64:   req.ProfileBase64,
	})
	if err != nil {
		return CreatePersonResponse{}, err
	}

	return CreatePersonResponse{ID: id}, nil
}

// ถ้าต้องใช้ต่อ: convert time to date only
func dateOnly(t time.Time) time.Time { return t.Truncate(24 * time.Hour) }
