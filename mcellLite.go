// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/haskelladdict/mcellLite/engine"
	"github.com/haskelladdict/mcellLite/geom"
	"github.com/haskelladdict/mcellLite/mol"
	"github.com/haskelladdict/mcellLite/output"
	"github.com/haskelladdict/mcellLite/vec"
)

const outPath = "/Users/markus/programming/go/src/github.com/haskelladdict/mcellLite/tests/viz_data"

func main() {

	var numIters int64 = 100
	a := mol.Species{"A", 600}
	var mols []*mol.Mol3
	molMap := make(mol.MolMap)
	for i := 0; i < 10000; i++ {
		mols = append(mols, &mol.Mol3{&a, vec.Vec3{0.0, 0.0, 0.0}, 0.0})
	}
	molMap[a.Name] = mols

	fmt.Println("Hi")
	rng := rand.New(rand.NewSource(99))
	m := geom.CreateRect(vec.Vec3{-0.1, -0.1, -0.1}, vec.Vec3{0.1, 0.1, 0.1})

	if err := output.WriteCB(molMap, outPath, "test", 0); err != nil {
		log.Println(err)
	}
	for i := int64(1); i <= numIters; i++ {
		fmt.Printf("iteration %d\n", i)
		for _, mols := range molMap {
			for _, mol := range mols {
				engine.Diffuse(mol, 1e-6, m, rng)
			}
		}
		if i%10 == 0 {
			if err := output.WriteCB(molMap, outPath, "test", i); err != nil {
				log.Println(err)
			}
		}
	}
}

/*















*/
