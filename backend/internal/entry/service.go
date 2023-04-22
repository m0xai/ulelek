// entry/service.go

package entry

import (
	"errors"
)

type EntryService struct {
	repo *EntryRepository
}

func NewEntryService(repo *EntryRepository) *EntryService {
	return &EntryService{
		repo: repo,
	}
}

func (s *EntryService) GetAllEntries() ([]Entry, error) {
	entries, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, errors.New("no entries found")
	}

	var result []Entry
	for _, e := range entries {
		result = append(result, *e)
	}

	return result, nil
}

func (s *EntryService) GetEntryByID(id int) (*Entry, error) {
	entry, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if entry == nil {
		return nil, errors.New("entry not found")
	}

	return entry, nil
}

func (s *EntryService) CreateEntry(entry_req *CreateEntryRequest) (*Entry, error) {
	entry, err := s.repo.Create(entry_req)
	if err != nil {
		return nil, err
	}

	return entry, err
}

func (s *EntryService) UpdateEntry(entry_req *UpdateEntryRequest) (*Entry, error) {
	entry, err := s.repo.Update(entry_req)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *EntryService) DeleteEntry(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
