package usecase

import (
	"aida.net/internal/domain"
	"context"
)

type ContactUseCase interface {
	CreateContact(ctx context.Context, contact *domain.Contact) error
	GetContactByID(ctx context.Context, id int) (*domain.Contact, error)
	UpdateContact(ctx context.Context, contact *domain.Contact) error
	DeleteContact(ctx context.Context, id int) error
}

type GroupUseCase interface {
	CreateGroup(ctx context.Context, group *domain.Group) error
	GetGroupByID(ctx context.Context, id int) (*domain.Group, error)
	CreateAndAddContactToGroup(ctx context.Context, groupID int, contact *domain.Contact) error
}
