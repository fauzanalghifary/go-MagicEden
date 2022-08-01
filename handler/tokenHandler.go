package handler

import (
	"goMagicEden/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type tokenHandler struct {
	tokenService token.Service
}

type RequestBody struct {
	TokenMintAddress string `json:"token_mint_address"`
}

// NewTokenHandler Func
func NewTokenHandler(tokenService token.Service) *tokenHandler {
	return &tokenHandler{tokenService}
}

// Root Handler func
func (h *tokenHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Fauzan",
	})
}

func (h *tokenHandler) GetWalletContents(c *gin.Context) {

	walletAddress := c.Param("wallet_address")

	tokens, err := h.tokenService.FindAll(walletAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tokens,
	})
}

func (h *tokenHandler) PostWalletContents(c *gin.Context) {

	walletAddress := c.Query("wallet_address")
	tokenList, err := token.GetDataFromAPI(walletAddress)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": "Wallet Not Found, Can't Post Anything",
		})
		return
	}
	var listedTokens []token.Token

	for _, t := range tokenList {

		theToken := token.Token{
			MintAddress:          t.MintAddress,
			Owner:                t.Owner,
			Supply:               t.Supply,
			Collection:           t.Collection,
			CollectionName:       t.CollectionName,
			Name:                 t.Name,
			UpdateAuthority:      t.UpdateAuthority,
			PrimarySaleHappened:  t.PrimarySaleHappened,
			SellerFeeBasisPoints: t.SellerFeeBasisPoints,
			Image:                t.Image,
			ListStatus:           t.ListStatus,
			TokenAddress:         t.TokenAddress,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		}

		if theToken.ListStatus == "listed" {
			token, err := h.tokenService.Create(theToken)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"errors": err,
				})
				return
			}
			listedTokens = append(listedTokens, token)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"data": listedTokens,
	})
}

func (h *tokenHandler) DeleteToken(c *gin.Context) {

	var RequestBody RequestBody

	if err := c.BindJSON(&RequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	token, err := h.tokenService.Delete(RequestBody.TokenMintAddress)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": "Can't find token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})

}
