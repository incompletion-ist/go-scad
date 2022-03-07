// Copyright 2022 Micah Kemp
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
