package main

import (
	"fmt"

	"github.com/excing/goflag"
)

// Config is gd2b config
type Config struct {
	Owner string `flag:"github repository owner"`
	Name  string `flag:"github repository name"`
	Token string `flag:"github authorization token (see https://docs.github.com/zh/graphql/guides/forming-calls-with-graphql)"`
}

func main() {
	var config Config
	goflag.Var(&config)
	goflag.Parse("config", "Configuration file path.")
	fmt.Println(config)
}
