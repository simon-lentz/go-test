// Package yaml adds support for conversion of manually entered data stories to Go structs
// and creation of edges based on keys extracted from struct fields.
package yaml

import (
	"fmt"

	"github.com/ghodss/yaml"
)

type Person struct {
	Name string `json:"name"` // Affects YAML field names too.
	Age  int    `json:"age"`
}

type Entity struct {
	KnownAs            string `json:"knownAs"`
	Ownership          string `json:"ownership"`
	OrganizationalForm string `json:"organizationalForm"`
	Description        string `json:"description"`
	Comment            string `json:"comment"`
	Governmental       string `json:"governmental"`
	Domain             string `json:"domain"`
	Name               string `json:"name"`
	EndUser            bool   `json:"endUser"`
	ServiceProvider    bool   `json:"serviceProvider"`
}

func Convert() string {
	// Marshal a Person struct to YAML.
	p := Person{"John", 30}
	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(string(y))
	/* Output:
	age: 30
	name: John
	*/

	// Unmarshal the YAML back into a Person struct.
	var p2 Person
	err = yaml.Unmarshal(y, &p2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(p2)
	/* Output:
	{John 30}
	*/
	return "test complete"
}
