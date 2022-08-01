package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Response Type
type Response []Input

// Input struct
type Input struct {
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
	// Attributes           []Attributes
}

// Attributes struct
type Attributes struct {
	TraitType string `json:"trait_type"`
	Value     string
}

// GetDataFromAPI Func
func GetDataFromAPI(walletAddress string) (Response, error) {

	url := "https://api-mainnet.magiceden.dev/v2/wallets/" + walletAddress + "/tokens?offset=0&listStatus=both"

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	var tokenList Response
	err = json.Unmarshal(responseData, &tokenList)
	return tokenList, err
}
