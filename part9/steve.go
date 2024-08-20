package main

import "fmt"

type Liters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (l Liters) String() string {
	return fmt.Sprintf("%.2f liters", l)
}

type Milliliters float64

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (m Milliliters) String() string {
	return fmt.Sprintf("%.2f milliliters", m)
}

type Gallons float64

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (g Gallons) String() string {
	return fmt.Sprintf("%.2f gallons", g)
}

func main() {
	fmt.Println(Gallons(1.0))
	fmt.Println(Liters(1.0))
	bytes, err := fmt.Println(Milliliters(1.0))
	if err == nil {
		fmt.Println(bytes)
	}
}
