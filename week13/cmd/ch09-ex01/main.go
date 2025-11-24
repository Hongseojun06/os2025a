package main

import "fmt"

type Meters float64
type Kilometers float64
type Miles float64

func (l Kilometers) ToMiles() Miles {
	return Miles(l * 0.62137)
}
func (m Meters) ToMiles() Miles {
	return Miles(m * 0.000621371)
}

func main() {
	kmph := Kilometers(150)
	fmt.Printf("%0.2f Kilomters equals %0.2f Miles\n", kmph, kmph.ToMiles())
	meter := Meters(500)
	fmt.Printf("%0.2f Meters %0.2f Miles\n", meter, meter.ToMiles())
}
