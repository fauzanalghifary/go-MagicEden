package token

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Token, error)
	Create(token Token) ([]Token, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Token, error) {
	var tokens []Token
	err := r.db.Find(&tokens).Error
	return tokens, err
}

func (r *repository) Create(token Token) (Token, error) {
	err := r.db.Create(&token).Error

	return token, err
}
