package main

import "fmt"

type Shaper interface {
	Area() int
	Sth() string
}

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Sth() string {
	return "I am a rectangle!"
}

type Square struct {
	side int
}

func (sq Square) Area() int {
	return sq.side * sq.side
}

func (sq Square) Sth() string {
	return "I am a square!"
}

func main() {
	r := Rectangle{length: 5, width: 3}
	q := Square{side: 5}
	shapesArr := [...]Shaper{r, q}

	fmt.Println("Looping!")
	for n, _ := range shapesArr {
		fmt.Println("Shape details: ", shapesArr[n])
		fmt.Println("Area of shape: ", shapesArr[n].Area())
	}
}
