package service

import "thaibev-test/storage"

type OccupationService interface {
	List() ([]OccupationDTO, error)
}

type occupationService struct {
	store storage.OccupationStore
}

func NewOccupationService(s storage.OccupationStore) OccupationService {
	return &occupationService{store: s}
}

func (s *occupationService) List() ([]OccupationDTO, error) {
	list, err := s.store.List()
	if err != nil {
		return nil, err
	}
	out := []OccupationDTO{}

	for _, o := range list {
		out = append(out, OccupationDTO{Code: o.Code, Name: o.Name})
	}
	return out, nil
}
