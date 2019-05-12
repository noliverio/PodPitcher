package main

import (
	"fmt"
	"github.com/dghubble/oauth1"
)

config := oauth1.Config{
	
}

func main() {
	fmt.Println("vim-go")
	requestToken, requestSecret, err = config.RequestToken()
}
