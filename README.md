# replay-code-test

# Installation

clone the repository

``git clone https://github.com/jd1100/replay-code-test``

cd into repository and build

``cd replay-code-test``

``go build``

run the compiled binary

``./replay-code-test``

# Testing
you can test by sending a GET request to the below URL (if you are running locally)

``http://localhost:8080/api/?address={valid ethereum address}&chain={mainnet/rinkeby/kovan}``

## Unit Tests

>There are two unit test functions (provided below) in ``replay-code-test/replayCode_test.go``

run all unit tests 

``go test -v``

To run unit tests individually use 

``go test -v -run <Test Function Name>``


---

**TestAPIHandlerResponseStatusWithInvalidChain**

>tests the API with an invalid chain parameter. Test will fail if status code returned is not 400

**TestAPIHandlerResponseObject**

>validates that the object returned in the API response is a valid object

---

## Examples using curl

```
user@computer:~/replay-code-test$ curl "http://localhost:8080/api/?address=0x158190da5b0cb33a3afc8103a4c304c86cb8410c&chain=mainnet"
{
        "ownedNfts": [
                {
                        "contractAddress": "0x0889ca523437a8b952a89ca7b402ef23561378eb",
                        "tokenId": "0x0000000000000000000000000000000000000000000000000000000000000001",
                        "title": "Mona Lisa #T1-NB1-1"
                }
        ]
}
```
---
```
user@computer:~/replay-code-test$ curl "http://localhost:8080/api/?address=0xdf9eb223bafbe5c5271415c75aecd68c21fe3d7f&chain=mainnet"
{
        "ownedNfts": null
}
```
---
```
user@computer:~/replay-code-test$ curl "http://localhost:8080/api/?address=0xbf3aeb96e164ae67e763d9e050ff124e7c3fdd28&chain=mainnet"
{
        "ownedNfts": [
                {
                        "contractAddress": "0xf12243191cead2d6a25a21b1151c269557bc3fb4",
                        "tokenId": "0x0000000000000000000000000000000000000000000000000000000000000015",
                        "title": "Weathereport #21"
                },
                {
                        "contractAddress": "0xf12243191cead2d6a25a21b1151c269557bc3fb4",
                        "tokenId": "0x0000000000000000000000000000000000000000000000000000000000000003",
                        "title": "Weathereport #3"
                },
                {
                        "contractAddress": "0xf12243191cead2d6a25a21b1151c269557bc3fb4",
                        "tokenId": "0x0000000000000000000000000000000000000000000000000000000000000012",
                        "title": "Weathereport #18"
                },
                {
                        "contractAddress": "0xfaafdc07907ff5120a76b34b731b278c38d6043c",
                        "tokenId": "0x1000000000000428000000000000000000000000000000000000000000000000",
                        "title": "Protocol of Quick Response"
                },
                {
                        "contractAddress": "0x3a3b0dbdc0f6bc77421dcd2f55cfa087b0db9aec",
                        "tokenId": "0x00000000000000000000000000000000000000000000000000000000000004b5",
                        "title": "MAX OSIRIS/MARTIN FISCHER 'CYBERPUNK FUTURE IS NOW'/JAY DELAY AND ROBNESS CYBERPOP - A SHAREWARE CRYPTO ART 'SPLIT' #4"
                }
        ]
}
```
---
```
user@computer:~/replay-code-test$ curl "http://localhost:8080/api/?address=0xbf3aeb96e164ae67e763d9e050ff124e7c3fdd28&chain=mfewwewfef"
invalid ethereum network
```
