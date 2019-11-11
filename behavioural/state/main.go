package main

import "fmt"

/*

Car:
turnOn
turnOff
gas
brake

stateEngineStopped
stateMoving
stateStopped

*/

type carState interface {
	turnOn()
	turnOff()
	gas()
	brake()
}

type myCar struct {
	speed uint
	state carState
}

func (car *myCar) changeState(newstate carState) {
	car.state = newstate
}

func (car *myCar) report() {
	state := "unknown"
	switch car.state.(type) {
	case stateEngineStopped:
		state = "off"
	case stateStopped:
		state = "stopped"
	case stateMoving:
		state = "moving"
	}
	fmt.Printf("state=%s,      speed=%d ", state, car.speed)
}

type stateEngineStopped struct {
	carState
	car *myCar
}

type stateStopped struct {
	carState
	car *myCar
}

type stateMoving struct {
	carState
	car *myCar
}

////////////////////////////////////////////////////
func (st stateEngineStopped) turnOn() {
	st.car.changeState(stateStopped{car: st.car})
	fmt.Println("turnOn(): OK")
}

func (st stateEngineStopped) turnOff() {
	fmt.Println("turnOff(): NO EFFECT (already off)")
}

func (st stateEngineStopped) gas() {
	fmt.Println("gas(): NO EFFECT")
}

func (st stateEngineStopped) brake() {
	if st.car.speed > 0 {
		st.car.speed--
	}
	fmt.Println("brake(): OK")
}

////////////////////////////////////////////////////
func (st stateStopped) turnOn() {
	fmt.Println("turnOn(): NO EFFECT (already on)")
}

func (st stateStopped) turnOff() {
	st.car.changeState(stateEngineStopped{car: st.car})
	fmt.Println("turnOff(): OK")
}

func (st stateStopped) gas() {
	st.car.changeState(stateMoving{car: st.car})
	fmt.Println("gas(): OK")
}

func (st stateStopped) brake() {
	if st.car.speed > 0 {
		st.car.speed--
	}
	fmt.Println("brake(): OK")
}

////////////////////////////////////////////////////
func (st stateMoving) turnOn() {
	fmt.Println("turnOn(): NO EFFECT (already on)")
}

func (st stateMoving) turnOff() {
	st.car.changeState(stateEngineStopped{car: st.car})
	fmt.Println("turnOff(): OK")
}

func (st stateMoving) gas() {
	st.car.speed++
	fmt.Println("gas(): OK")
}

func (st stateMoving) brake() {
	if st.car.speed > 0 {
		st.car.speed--
	}
	if 0 == st.car.speed {
		st.car.changeState(stateStopped{car: st.car})
	}
	fmt.Println("brake(): OK")
}

func main() {
	var car myCar
	car = myCar{state: stateEngineStopped{car: &car}}
	car.report()
	car.state.turnOn()
	car.report()
	car.state.gas()
	car.report()
	car.state.gas()
	car.report()
	car.state.gas()
	car.report()
	car.state.brake()
	car.report()
	car.state.brake()
	car.report()
	car.state.brake()
	car.report()
	car.state.turnOff()
	car.report()
}
