package repository

import (
	"aida.net/internal/domain"
	"context"
)

type ContactRepository interface {
	CreateContact(ctx context.Context, contact *domain.Contact) error
	GetContactByID(ctx context.Context, id int) (*domain.Contact, error)
	UpdateContact(ctx context.Context, contact *domain.Contact) error
	DeleteContact(ctx context.Context, id int) error
}

type GroupRepository interface {
	CreateGroup(ctx context.Context, group *domain.Group) error
	GetGroupByID(ctx context.Context, id int) (*domain.Group, error)
	UpdateGroup(ctx context.Context, group *domain.Group) error
	DeleteGroup(ctx context.Context, id int) error
}
