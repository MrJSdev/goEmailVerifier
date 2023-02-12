package main

import (
	"fmt"

	"github.com/MrJSdev/goEmailVerifier/service"
)

func main() {
	if service.HasDomainWorkingEmailSystem("aspireapp.com") {
		fmt.Println("Email id is valid and record is exists")
	} else {
		fmt.Println("email id record is not exist")
	}
}
