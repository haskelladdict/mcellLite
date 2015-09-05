// Copyright 2015 Markus Dittrich
// Licensed under BSD license, see LICENSE file for details
package main

import (
	"fmt"
	"math/rand"

	"github.com/haskelladdict/mcellLite/engine"
	"github.com/haskelladdict/mcellLite/geom"
	"github.com/haskelladdict/mcellLite/mol"
	"github.com/haskelladdict/mcellLite/vec"
)

// Simulation contains all information related to simulation control such as
// timestep, number of iterations
// type Simulation struct {
// 	dt    float64
// 	iters uint64
// }

// // Partition contains all state information for a single patch in the simulation
// type Partition struct {
// 	Neighbors   []*Partition
// 	Mols        []*types.Molecule
// 	Faces       []*geometry.Face
// 	Events      scheduler.Events
// 	MolReceiver <-chan *types.Molecule
// 	Messager    chan struct{}
// 	Commander   chan int
// }

// Run is in charge of running all events occurring within the given
// partition. The main loop consists of checking for events on a cmd channel,
// several input channels or otherwise picking events of off the scheduler.
// func (p *Partition) Run() {
// 	// timer used for testing
// 	timer := time.NewTimer(time.Second * 2)

// 	// create scheduler
// 	sched := scheduler.Create()

// 	start := 0.0
// 	//	end := 1.0
// 	for {

// 		// add diffusion events
// 		for _, m := range p.Mols {
// 			sched.Add(&scheduler.Diffusion{m, start})
// 		}

// 		select {
// 		case c := <-p.Commander:
// 			fmt.Println("got command ", c)
// 		case m := <-p.MolReceiver:
// 			fmt.Println(m)
// 		case <-timer.C:
// 			fmt.Println("partition expiring")
// 			p.Messager <- struct{}{}
// 		}
// 	}
// }

func main() {
	// define vertices of geometry
	//  cube := makeCube()

	// // prepare partition
	// aMol := types.Species{6e-6}
	// mol1 := types.Molecule{Species: aMol, Pos: geometry.Vec3{0.5, 0.5, 0.6}}
	// mol2 := types.Molecule{Species: aMol, Pos: geometry.Vec3{0.3, 0.2, 0.4}}
	// mols := []*types.Molecule{&mol1, &mol2}

	// molCh := make(chan *types.Molecule)
	// cmdCh := make(chan int)
	// msgCh := make(chan struct{})
	// p := Partition{Neighbors: nil,
	// 	Mols:        mols,
	// 	Faces:       cube.Faces,
	// 	Events:      scheduler.Create(),
	// 	MolReceiver: molCh,
	// 	Messager:    msgCh,
	// 	Commander:   cmdCh,
	// }

	// go p.Run()
	// p.Commander <- 1
	// <-msgCh

	a := mol.Species{"A", 600}
	var mols []mol.Mol3
	molMap := make(map[string][]mol.Mol3)
	for i := 0; i < 10000; i++ {
		mols = append(mols, mol.Mol3{&a, vec.Vec3{0.0, 0.0, 0.0}, 0.0})
	}
	molMap[a.Name] = mols

	fmt.Println("Hi")
	rng := rand.New(rand.NewSource(99))
	m := geom.CreateRect(vec.Vec3{-1.0, -1.0, -1.0}, vec.Vec3{1.0, 1.0, 1.0})
	engine.Diffuse(&mol.Mol3{&a, vec.Vec3{0.0, 0.0, 0.0}, 0.0}, 1e-6, m, rng)
}

/*















*/
