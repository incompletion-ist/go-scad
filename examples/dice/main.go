package main

import (
	"log"

	"github.com/micahkemp/scad/examples/dice/die"
	"github.com/micahkemp/scad/examples/dice/dimples"
	"github.com/micahkemp/scad/pkg/scad"
)

func main() {
	die := die.Die{
		Dimple: dimples.Dimple{
			Depth:    2,
			Diameter: 10,
		},
		Width: 60,
	}

	if err := scad.Write(".", die); err != nil {
		log.Fatalf("unable to EncodeFunction: %s", err)
	}
}
