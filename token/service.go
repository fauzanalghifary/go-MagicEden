package token

type Service interface {
	FindAll() ([]Token, error)
	Create(token Token) ([]Token, error)
}

type service struct {
	repository Repository
}
