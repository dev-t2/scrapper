package main

import "fmt"

type Gallons float64

type Liters float64

func toGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func toLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)

	carFuel += toGallons(Liters(40.0))
	busFuel += toLiters(Gallons(30.0))

	fmt.Printf("Car Fuel: %0.1f Gallons\n", carFuel)
	fmt.Printf("Bus Fuel: %0.1f Liters\n", busFuel)
}