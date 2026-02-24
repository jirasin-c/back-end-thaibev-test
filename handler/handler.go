package handler

import (
	"thaibev-test/service"
	"thaibev-test/storage"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	PeopleService     service.PeopleService
	OccupationService service.OccupationService
}

func New(db *sqlx.DB) *Handler {
	occStore := storage.NewOccupationStore(db)
	peopleStore := storage.NewPeopleStore(db)

	return &Handler{
		PeopleService:     service.NewPeopleService(peopleStore, occStore),
		OccupationService: service.NewOccupationService(occStore),
	}
}
