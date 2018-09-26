package main

import (
    "fmt"
    "log"

   T "github.com/chewxy/gorgonia"
	"gorgonia.org/tensor"
)

func main() {
    g := T.NewGraph()

    var x, y, z *T.Node
    var err error

    // define the expression
    x = T.NewScalar(g, T.Float64, T.WithName("x"))
    y = T.NewScalar(g, T.Float64, T.WithName("y"))
    z, err = T.Add(x, y)
    if err != nil {
        log.Fatal(err)
    }

    // create a VM to run the program on
    machine := T.NewTapeMachine(g)

    // set initial values then run
    T.Let(x, 2.0)
    T.Let(y, 2.5)
    if machine.RunAll() != nil {
        log.Fatal(err)
    }

    fmt.Printf("%v\n", z.Value())
    // Output: 4.5

	fmt.Println("go deepLearning")
	a := tensor.New(tensor.WithShape(2,2), tensor.WithBacking([]int{1,2,3,4}))
	fmt.Printf("a: \n%v\n",a)
	b := tensor.New(tensor.WithBacking(tensor.Range(tensor.Float32,0,24)),tensor.WithShape(2,3,4))
	fmt.Printf("b: \n%v\n",b)
}
