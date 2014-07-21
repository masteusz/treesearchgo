package main

import "fmt"

type Bus struct {
	l, b, h           int
	rows, seatsPerRow int
}

// type Cuboider interface {
// 	CubicVolume() int
// }

func (bus Bus) CubicVolume() int {
	return bus.l * bus.b * bus.h
}

// type PublicTransporter interface {
// 	PassengerCapacity() int
// }

func (bus Bus) PassengerCapacity() int {
	return bus.rows * bus.seatsPerRow
}

func main() {
	b := Bus{
		l: 10, b: 6, h: 3, rows: 10, seatsPerRow: 5}
	fmt.Println("Cubic volume of bus:", b.CubicVolume())
	fmt.Println("Maximum number of passengers:", b.PassengerCapacity())
}
