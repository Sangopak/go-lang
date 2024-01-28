//package src
package main

import "fmt"

type Rectangle struct {
    width, height int
}

type Geometry interface {
    area() int
    perimeter() int
}

func (r Rectangle) area() int {
    return r.width * r.height
}

func (r Rectangle) perimeter() int {
    return 2*r.width + 2*r.height
}

func measure(g Geometry) {
    fmt.Println("geometry: ", g)
    fmt.Println("area: ",g.area())
    fmt.Println("perim: ",g.perimeter())
}

func main() {
    r := Rectangle{width: 10, height: 5}
    measure(r)
}