package main

import (
	"github.com/wyrth-io/go-test/internal/yaml"
	methods "github.com/wyrth-io/go-test/sandbox"
)

func main() {
	yaml.Convert()
	methods.SetEDGE("Simon", "13 Poplar", "RESIDENT_OF_Residence")
}
