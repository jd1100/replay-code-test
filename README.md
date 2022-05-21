# replay-code-test

to run the api, build comple the code with ``go build main.go`` and then run the compiled code

requires GO version 1.17.2

# Testing
you can run tests by sending a GET request to the below URL (if you are running locally)

``http://localhost:8080/api/?address={valid ethereum address}&chain={mainnet/rinkeby/kovan}``

## Examples

```
curl "http://localhost:8080/api/?address=0x158190da5b0cb33a3afc8103a4c304c86cb8410c&chain=mainnet"
curl "http://localhost:8080/api/?address=0xdf9eb223bafbe5c5271415c75aecd68c21fe3d7f&chain=mainnet"
curl "http://localhost:8080/api/?address=0xbf3aeb96e164ae67e763d9e050ff124e7c3fdd28&chain=mainnet"
```
