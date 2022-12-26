package main

import (
	"github.com/wyrth-io/go-test/basic"
	"github.com/wyrth-io/go-test/yaml"
)

func main() {
	yaml.Convert()
	basic.SetEDGE("Simon", "13 Poplar", "RESIDENT_OF_Residence")
}
