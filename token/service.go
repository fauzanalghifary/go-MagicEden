package token

// Service interface
type Service interface {
	FindAll(walletAddress string) ([]Token, error)
	Create(tokens Token) (Token, error)
	Delete(tokenMintAddress string) (Token, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(walletAddress string) ([]Token, error) {
	tokens, err := s.repository.FindAll(walletAddress)
	return tokens, err
}

func (s *service) Create(token Token) (Token, error) {
	newToken, err := s.repository.Create(token)
	return newToken, err
}

func (s *service) Delete(walletAddress string) (Token, error) {
	deleteToken, err := s.repository.FindOne(walletAddress)
	token, err := s.repository.Delete(deleteToken)
	return token, err
}
