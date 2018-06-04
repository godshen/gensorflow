package main

import (
    "fmt"
    "gorgonia.org/tensor"
)

func main() {
    fmt.Println("go deepLearning")
    a := tensor.New(tensor.WithShape(2,2), tensor.WithBacking([]int{1,2,3,4}))
    fmt.Printf("a: \n%v\n",a)
    b := tensor.New(tensor.WithBacking(tensor.Range(tensor.Float32,0,24)),tensor.WithShape(2,3,4))
    fmt.Printf("b: \n%v\n",b)
}
