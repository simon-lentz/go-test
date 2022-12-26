package basic

import (
	"fmt"
	"log"

	"github.com/ghodss/yaml"
)

type Person struct {
	Name                  string `json:"name"`
	Age                   int    `json:"age"`
	Occupation            string `json:"occupation"`
	LikesChocolate        bool   `json:"likesChocolate"`
	RESIDENT_OF_Residence string
}

// setName receives a pointer to Person so it can set field values.
func (p *Person) setName(name string) {
	p.Name = name
}

func (p Person) getName() string {
	return p.Name
}

type Residence struct {
	ResidentName string `json:"residentName"`
	State        string `json:"state"`
	City         string `json:"city"`
	Address      string `json:"address"`
	Zip          int    `json:"zip"`
	Prewar       bool   `json:"prewar"`
}

func (r *Residence) setResidentName(name string) {
	r.ResidentName = name
}

type EDGE struct {
	RESIDENT_OF_Residence string
}

func SetEDGE(person string, residence string, edge string) {
	// Create struct instances of a new person, a new residence,
	// and a new edge connecting the person and residence.
	newPerson := Person{}
	newResidence := Residence{}
	newEdge := EDGE{}

	// Use the setName function to assign a name to the new person.
	newPerson.setName("Simon")

	// Use the getName function to return the new person's name as set in the step above.
	name := newPerson.getName()

	// Use the setResidentName function with the new person's name variable (name = "Simon")
	// to set the new residence's resident name field to the new person's name.
	newResidence.setResidentName(name)

	// Link the person and residence node with their common feature: newPerson.Name == newResidence.ResidentName
	newEdge.RESIDENT_OF_Residence = newResidence.ResidentName
}

func Unmarshal() string {
	// Marshal an Entity struct to YAML.
	simon := Person{}
	y, err := yaml.Marshal(simon)
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(string(y))

	// Unmarshal the YAML back into an Entity struct.
	err = yaml.Unmarshal(y, &simon)
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(simon)
	return "test complete"
}
