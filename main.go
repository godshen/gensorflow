package main

import (
	"fmt"
	"log"

	"github.com/chewxy/gorgonia"
	"gorgonia.org/tensor"
)

func main() {
	g := gorgonia.NewGraph()
	var x, y, z *gorgonia.Node
	var err error
	// define the expression
	x = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))
	z, err = gorgonia.Add(x, y)
	if err != nil {
		log.Fatal(err)
	}
	// create a VM to run the program on
	machine := gorgonia.NewTapeMachine(g)
	// set initial values then run
	gorgonia.Let(x, 2.0)
	gorgonia.Let(y, 2.5)
	if machine.RunAll() != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", z.Value())

	fmt.Println("go deepLearning")
	a := tensor.New(tensor.WithShape(2, 2), tensor.WithBacking([]int{1, 2, 3, 4}))
	fmt.Printf("a: \n%v\n", a)
	b := tensor.New(tensor.WithBacking(tensor.Range(tensor.Float32, 0, 24)), tensor.WithShape(2, 3, 4))
	fmt.Printf("b: \n%v\n", b)
}
