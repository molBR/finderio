package setup

import (
	"testing"
	"fmt"
)

func TestMainSetup(t *testing.T){
	setup:=MainSetup();
	fmt.Println(setup.DynamoClient)
	if(setup.Session== nil || setup.DynamoClient == nil){
		t.Error("Wanted setup but got null")
	}
}