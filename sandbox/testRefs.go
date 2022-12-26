package sandbox

import (
	"sync"
)

type Person struct {
	Name string
	Ref  string
}

// func points to Person node and assigns it field values.
func (p *Person) personField(name string) string {
	p.Name = name
	return p.Name
}

type Residence struct {
	ResidentName string
	State        string
	City         string
}

func (r *Residence) residenceField(residentName string, state string) (string, string) {
	r.ResidentName = residentName
	r.State = state
	return r.ResidentName, r.State
}

type Reference[T any] struct {
	sync.Mutex
	Data T
}

type Edge struct {
	ResidentIs *Reference[string]
	LivesIn    *Reference[[]string]
}

/*
func (e *Reference[T]) edgeResidentOf(name string) string {
	return name
}
*/

// Assigns input string to Name and ResidentName fields of Person and Residence, respectively.
func PopulateGraph(name string, state string) (Person, Residence, Edge) {
	var p = Person{}
	var r = Residence{}
	var e = Edge{}

	p.personField(name)
	r.residenceField(name, state)

	return p, r, e
}

func DrawGraph() string {
	return "test complete"
}
