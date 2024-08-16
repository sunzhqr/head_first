package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64
type Title string

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Liters(40)
	busFuel := Gallons(240)
	fmt.Printf("For car we have %0.1f gallons fuel\n", ToGallons(carFuel))
	fmt.Printf("For bus we have %0.1f liters fuel\n", ToLiters(busFuel))

	fmt.Println(len())
}
