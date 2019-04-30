package main

import "github.com/noliverio/podpitcher/src/oauth1"

func main() {
	endpoint := oauth1.Endpoint{
		RequestTokenEndpoint: "https://api.twitter.com/oauth/request_token",
		AccessTokenEndpoint:  "https://api.twitter.com/oauth/access_token",
		AuthorizeEndpoint:    "https://api.twitter.com/oauth/authorize",
	}
}
