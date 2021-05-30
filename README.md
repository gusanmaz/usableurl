## UsableURL Library, CLI and Microservice

usable-url is a minuscule suite that resolves potentially unusable urls. If you need expansion for a 
shortened url or url with a missing prefix (e.g wikipedia.org instead of https://www.wikipedia.org) usable-url
could help

### Requirements 

Behind the scenes this program use [cURL](https://curl.se/). cURL should have been installed in your computer to use this program.

### Installation

1. `go get github.com/gusanmaz/usable-url`
2. `go install ${GOPATH}/src/github.com/gusanmaz/usable-url/cmd/cli/usable-url-cli.go`
3. `go install ${GOPATH}/src/github.com/gusanmaz/cmd/usable-url/server/usable-url-server.go`

### Usage

###### API

1. Import `github.com/gusanmaz/usableurl` into your project
2. Call usableURL.Sanitize("your unprocessed url goes here")

###### CLI

```usable-url-cli --url google.com``` This command should return https://www.google.com

###### SERVER

1. ```usable-url-server --port 8081``` If you don't specify the port, it defaults to 9001
2. type `localhost:8081/?url=google.com` into your browser's address bar.
3. Upon enter you could receive a JSON response similar to:

```json
{
  "In": "google.com",
  "Out": "http://www.google.com/"
}
```

### AUTHOR

Güvenç Usanmaz

### LICENSE

MIT License


