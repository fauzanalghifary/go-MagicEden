package token

// Service interface
type Service interface {
	FindAll() ([]Token, error)
	Create(token Token) ([]Token, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Token, error) {
	tokens, err := s.repository.FindAll()
	return tokens, err
}

func (s *service) Create(token Token) (Token, error) {
	newToken, err := s.repository.Create(token)
	return newToken, err
}
