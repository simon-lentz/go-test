package main

import (
	"fmt"

	methods "github.com/wyrth-io/go-test/sandbox"
)

func main() {
	personNode, residenceNode, edge := methods.PopulateGraph("Simon", "NY")
	fmt.Println(personNode, residenceNode, edge)
	fmt.Println(methods.DrawGraph())
}
