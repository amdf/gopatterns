package main

import "fmt"

type bike interface {
	travelOnBike()
}

type car interface {
	travelByCar()
}

type bus interface {
	travelOnBus()
}

type diesel interface {
	refuelDiesel()
}

type gas interface {
	refuelGas()
}

type gasBike struct {
	gas
	bike
}

type gasBus struct {
	gas
	bus
}

type dieselBus struct {
	diesel
	bus
}

type gasCar struct {
	gas
	car
}

type dieselCar struct {
	diesel
	car
}

type abstractFactory interface {
	createCar() car
	createBus() bus
	createBike() bike
}

type dieselFactory struct {
	abstractFactory
}

type gasFactory struct {
	abstractFactory
}

func (f gasFactory) createCar() car {
	return gasCar{}
}
func (f gasFactory) createBus() bus {
	return gasBus{}
}
func (f gasFactory) createBike() bike {
	return gasBike{}
}

func (f dieselFactory) createCar() car {
	return dieselCar{}
}
func (f dieselFactory) createBus() bus {
	return dieselBus{}
}
func (f dieselFactory) createBike() bike {
	return nil
}

func (x gasBike) travelOnBike() {
	fmt.Println("travel on gas bike")
}
func (x gasCar) travelByCar() {
	fmt.Println("travel by gas car")
}
func (x gasBus) travelOnBus() {
	fmt.Println("travel on gas bus")
}

func (x dieselCar) travelByCar() {
	fmt.Println("travel by diesel car")
}
func (x dieselBus) travelOnBus() {
	fmt.Println("travel on diesel bus")
}

func main() {

	var di dieselFactory
	var ga gasFactory

	dieselCar := di.createCar()
	gasCar := ga.createCar()
	dieselBus := di.createBus()
	gasBus := ga.createBus()
	gasBike := ga.createBike()
	dieselCar.travelByCar()
	gasCar.travelByCar()
	dieselBus.travelOnBus()
	gasBus.travelOnBus()
	gasBike.travelOnBike()
	oddBike := di.createBike()
	fmt.Println("diesel bike is ", oddBike)
}
