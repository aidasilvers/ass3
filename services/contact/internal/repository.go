package internal

import (
	"aida.net/internal/repository"
)

type ContactService struct {
	ContactRepository repository.ContactRepository
	GroupRepository   repository.GroupRepository
}

func NewContactService(contactRepository repository.ContactRepository, groupRepository repository.GroupRepository) *ContactService {
	return &ContactService{
		ContactRepository: contactRepository,
		GroupRepository:   groupRepository,
	}
}
