package profile

type Service interface {
	GetBankList() (Bank, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBankList() (Bank, error) {
	return s.repository.BankList()
}
