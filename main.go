package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Token struct
type Token struct {
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
	Attributes           []Attributes
}

// Attributes struct
type Attributes struct {
	TraitType string `json:"trait_type"`
	Value     string
}

func main() {
	response, err := http.Get("https://api-mainnet.magiceden.dev/v2/wallets/GVUAKf19vnM9c5WZxXYBLAy6hdcpaS6PuFWZfrZcTTo4/tokens?offset=0&listStatus=both")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tokenList []Token
	if err := json.Unmarshal(responseData, &tokenList); err != nil {
		log.Fatal(err)
	}

	fmt.Println(tokenList)

}
