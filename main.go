package main

import (
	"github.com/repoofideas/gorganon/explorer"
	"github.com/repoofideas/gorganon/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
