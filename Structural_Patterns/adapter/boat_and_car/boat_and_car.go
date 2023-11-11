package main

import "fmt"

type Transport interface {
  navigateToDestination()
}

type Captain struct {}

type Boat struct {}

type Car struct {}

type CarAdapter struct {
  car *Car
}

func (c *Captain) startNavigation(transport Transport) {
  fmt.Println("Captain starting the navigation.")
  transport.navigateToDestination()
}

func (b *Boat) navigateToDestination() {
  fmt.Println("Boat is navigating to the island.")
}

func (c *Car) driveToDestination() {
  fmt.Println("Car is going to the destination.")
}

func (ca *CarAdapter) navigateToDestination() {
  fmt.Println("Modifying the car to drive in water.")
  ca.car.driveToDestination()
}

func main() {
  captain := &Captain{}
  boat := &Boat{}

  captain.startNavigation(boat)

  car := &Car{}
  carAdapter := &CarAdapter{
    car: car,
  }

  captain.startNavigation(carAdapter)
}
