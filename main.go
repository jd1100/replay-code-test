
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
	"io/ioutil"
	"time"
)

/*type alchemyGetNFTSRequestObject struct {
    Owner string `json:"owner"`
    PageKey       string `json:"pageKey,omitempty"`
}
*/

type AlchemyGetNFTsResponseObject struct {
	OwnedNfts []struct {
		Contract struct {
			Address string `json:"address"`
		} `json:"contract"`
		ID struct {
			TokenID       string `json:"tokenId"`
			TokenMetadata struct {
				TokenType string `json:"tokenType"`
			} `json:"tokenMetadata"`
		} `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		TokenURI    struct {
			Raw     string `json:"raw"`
			Gateway string `json:"gateway"`
		} `json:"tokenUri"`
		Media []struct {
			Raw     string `json:"raw"`
			Gateway string `json:"gateway"`
		} `json:"media"`
		Metadata struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Image       string `json:"image"`
			ExternalURL string `json:"external_url"`
			Attributes  []struct {
				Value     string `json:"value"`
				TraitType string `json:"trait_type"`
			} `json:"attributes"`
		} `json:"metadata"`
		TimeLastUpdated time.Time `json:"timeLastUpdated"`
	} `json:"ownedNfts"`
	TotalCount int    `json:"totalCount"`
	BlockHash  string `json:"blockHash"`
}


type APIResponseObject struct {
	OwnedNfts []struct {
		ContractAddress string `json:"contractAddress"`
		TokenID         string `json:"tokenId"`
		Title           string `json:"title",omitempty`
	} `json:"ownedNfts"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var alchemyURL string
	var alchemyGetNFTsResponseObject AlchemyGetNFTsResponseObject
	var apiResponseObject APIResponseObject

	// listen for GET request only
	if (r.Method == "GET") {
		fmt.Println(r.URL)


		// parse the http request parameters into string variables
		parsedURL, err := url.Parse(r.URL.String())

		if err != nil {
			fmt.Println(err)
			return
		}
		urlParameters := parsedURL.Query()

		ethereumAddress := strings.Join(urlParameters["address"], "")
		ethereumChain := strings.Join(urlParameters["chain"], "")

		//fmt.Println(ethereumAddress)
		//fmt.Println(ethereumChain)

		// valid that the ethereum chain is one of 3 values: mainnet, rinkeby, kovan
		switch ethereumChain {
		case "mainnet":
			fmt.Println("mainnet")
			alchemyURL = "https://eth-mainnet.alchemyapi.io/v2/demo/getNFTs/"
		case "rinkeby":
			fmt.Println("rinkeby")
			alchemyURL = "https://eth-rinkeby.alchemyapi.io/v2/demo/getNFTs/"
		case "kovan":
			fmt.Println("kovan")
			alchemyURL = "https://eth-kovan.alchemyapi.io/v2/demo/getNFTs/"
		default:
			w.WriteHeader(400)
			fmt.Fprintf(w, "invalid ethereum network")
			return

		}
		
		
		// construct http request to external api
		resp, err := http.Get(alchemyURL + "?owner=" + ethereumAddress)
		if err != nil {
			fmt.Println(err)
		}
		buffer, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		resp.Body.Close()
		// unmarshal json byte array into struct
		json.Unmarshal(buffer, &alchemyGetNFTsResponseObject)

		
		// loop through the response from alchemy getNFTs api and parse the below variables into the apiResponseObject
		for _, ownedNFT := range alchemyGetNFTsResponseObject.OwnedNfts {

			// create an anonymous struct with and append to the apiResponseObject.OwnedNfts array
			nextOwnedNFT := struct {
				ContractAddress string `json:"contractAddress"`
				TokenID         string `json:"tokenId"`
				Title           string `json:"title",omitempty`
			}{
				ContractAddress: ownedNFT.Contract.Address,
				TokenID: ownedNFT.ID.TokenID,
				Title: ownedNFT.Title,
			}

			apiResponseObject.OwnedNfts = append(apiResponseObject.OwnedNfts, nextOwnedNFT)
		}
		
		// encode the apiResponseObject struct into a json byte array
		apiResponseObjectByteArray, err := json.MarshalIndent(apiResponseObject, "", "\t")
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(apiResponseObjectPrettyPrint)
		//fmt.Fprintf(w, "Address: %s\nChain: %s\n%s", ethereumAddress, ethereumChain, apiResponseObjectPrettyPrint)
		fmt.Fprintf(w, "%s", apiResponseObjectByteArray)

	} else{
		fmt.Fprintf(w, "Invalid Request")
	}

	fmt.Print(r.Method)
}



func main() {

	// initialize http router
	mux := http.NewServeMux()

	// add handlers to the http router
	mux.HandleFunc("/api/", apiHandler)

	// start http server with given http router
	http.ListenAndServe("localhost:8080", mux)
}