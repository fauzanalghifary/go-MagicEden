package token

import "gorm.io/gorm"

// Repository Interface
type Repository interface {
	FindAll(walletAddress string) ([]Token, error)
	FindOne(tokenMintAddress string) (Token, error)
	Create(token Token) (Token, error)
	Delete(token Token) (Token, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository Func
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(walletAddress string) ([]Token, error) {
	var tokens []Token
	err := r.db.Where("owner = ?", walletAddress).Find(&tokens).Error
	return tokens, err
}

func (r *repository) FindOne(tokenMintAddress string) (Token, error) {
	var token Token
	err := r.db.Where("mint_address = ?", tokenMintAddress).Find(&token).Error
	return token, err
}

func (r *repository) Create(token Token) (Token, error) {
	err := r.db.Create(&token).Error
	return token, err
}

func (r *repository) Delete(token Token) (Token, error) {
	err := r.db.Delete(&token).Error
	return token, err
}
