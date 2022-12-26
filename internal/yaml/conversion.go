// Package yaml adds support for conversion of manually entered data stories to Go structs
// and creation of edges based on keys extracted from struct fields.
package yaml

import (
	"fmt"
	"log"

	"github.com/ghodss/yaml"
	nodes "github.com/wyrth-io/go-test/internal/types"
)

func Convert() string {
	// Marshal an Entity struct to YAML.
	DTE := nodes.Entity{Name: "DTE"}
	y, err := yaml.Marshal(DTE)
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(string(y))

	// Unmarshal the YAML back into an Entity struct.
	var DelawareTotalPower nodes.Entity
	err = yaml.Unmarshal(y, &DelawareTotalPower)
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(DelawareTotalPower)
	return "test complete"
}
