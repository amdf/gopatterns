package main

import (
	"fmt"
	"time"
)

type commuteStrategy interface {
	commute() time.Duration
}

type commuteOnBike struct {
	commuteStrategy
}
type commuteOnPublicTransport struct {
	commuteStrategy
}
type commuteOnPersonalCar struct {
	commuteStrategy
}
type commuteByOwnLegs struct {
	commuteStrategy
}

type somePerson struct {
	commuter commuteStrategy
}

func (p somePerson) goToWork() time.Duration {
	return p.commuter.commute()
}

func (p *somePerson) setStrategy(transport commuteStrategy) {
	p.commuter = transport
}

func (cs commuteOnBike) commute() time.Duration {
	return 20 * time.Minute
}
func (cs commuteOnPublicTransport) commute() time.Duration {
	return 40 * time.Minute
}
func (cs commuteOnPersonalCar) commute() time.Duration {
	return 50 * time.Minute
}
func (cs commuteByOwnLegs) commute() time.Duration {
	return 60 * time.Minute
}

func main() {
	var worker somePerson
	worker.setStrategy(commuteOnBike{})
	fmt.Println("on bike: ", worker.goToWork())
	worker.setStrategy(commuteOnPublicTransport{})
	fmt.Println("on public transport: ", worker.goToWork())
	worker.setStrategy(commuteOnPersonalCar{})
	fmt.Println("on personal car: ", worker.goToWork())
	worker.setStrategy(commuteByOwnLegs{})
	fmt.Println("by own legs: ", worker.goToWork())

}
