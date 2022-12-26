package types

type Graph struct {
	Entities              []Project
	Services              []Service
	Projects              []Project
	Initiatives           []Initiative
	Regulations           []Regulation
	Places                []Place
	GovernmentAuthorities []GovernmentAuthority
	Edges                 []EDGES
}
