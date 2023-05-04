#### Overview
A simple URL Shortener service that shortens, stores, and resolves URLs. It uses `base62` encoding with a random number to generate the short url. The random number is between `1-916132833` which guarantees a short url with max length of 5 characters (number of permutations with max 5 chars: 62<sup>5</sup> = 916132832).

#### Prerequisites
- go 1.20 installed - https://go.dev/doc/install

Check go version:
```bash 
go version
```
-  mongodb-community@6.0 installed and running on the default port `:27017` - https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-os-x/

Check that the service is running (for MacOS, using brew):
```sh
brew services list
```
``` 
Name              Status  
mongodb-community started
```
#### Run
Start the service with:
```sh
go run cmd/main.go
```
To run all tests use:
```sh
go test ./... -v
```
#### Usage
Get short url:
```sh
curl 'http://localhost:8090/' -H 'Content-Type: text/plain;charset=UTF-8' --data-raw 'https://someverylongdomainnamehere.com/some/very/very/long/path/here?foo=bar'
```
Example result:
```sh
P8gCt
```
Going to *http://localhost:8090/{shortUrl}* will result in a redirect to the original url.

Curl:
```sh
curl 'http://localhost:8090/P8gCt'
```
Example result:
```sh
<a href="https://someverylongdomainnamehere.com/some/very/very/long/path/here?foo=bar">Moved Permanently</a>.
```
