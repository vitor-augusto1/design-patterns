package main

import "fmt"

// Factory method is a creational design pattern which solves the problem of
// creating product objects without specifying their concrete classes.

// The Factory Method defines a method, which should be used for creating
// objects instead of using a direct constructor call (new operator). Subclasses
// can override this method to change the class of objects that will be created.

// It’s impossible to implement the classic Factory Method pattern in Go due to
// lack of OOP features such as classes and inheritance. However, we can still
// implement the basic version of the pattern, the Simple Factory.

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

// Concrete Product: AK47

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 7,
		},
	}
}

// Concrete Product:  Musket

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 3,
		},
	}
}

// Here is the gun factory

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type...")
}

// Client code here

func printGunDetails(g IGun) {
  fmt.Printf("Gun: %s\n", g.getName())
  fmt.Printf("Power: %d\n\n", g.getPower())
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

  printGunDetails(ak47)
  printGunDetails(musket)
}
