package models

import "time"

type Token struct {
	ID                   int
	MintAddress          string
	Owner                string
	Supply               int
	Collection           string
	CollectionName       string
	Name                 string
	UpdateAuthority      string
	PrimarySaleHappened  bool
	SellerFeeBasisPoints int
	Image                string
	ListStatus           string
	TokenAddress         string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
