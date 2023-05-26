package internal

type ContactUsecase struct {
	ContactService *ContactService
}

func NewContactUsecase(contactService *ContactService) *ContactUsecase {
	return &ContactUsecase{
		ContactService: contactService,
	}
}
