package main

import(

	"finderio/cmd/setup"
	"finderio/cmd/in/graphql"
)

func main(){
	confSetup:=setup.MainSetup()
	in_graphql.CreateServer(confSetup);
}