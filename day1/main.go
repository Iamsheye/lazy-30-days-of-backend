package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(BoO("Bolu"))
	fmt.Println(BoO("Odun"))
	fmt.Println(BoO("Joga Bon"))

}

func BoO(name string) string {
	match, _ := regexp.Match("^[[:alpha:]]+$", []byte(name))

	if !match {
		return "Please remove all spaces or special characters"
	}

	if name == "Bolu" || name == "Odun" {
		return "Welcome back, " + name
	}

	return "It is nice to meet you."
}
