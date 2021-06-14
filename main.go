package main

import (
	"finderio/cmd/in/graphql"
	"finderio/cmd/setup"
)

func main() {
	confSetup := setup.MainSetup()
	in_graphql.CreateServer(confSetup)
}
