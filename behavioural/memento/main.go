package main

import (
	"fmt"
	"time"
)

type fields struct {
	field1 string
	field2 int
	field3 bool
}

type someclass struct {
	data fields
}

//Snapshot snapshot
type Snapshot struct {
	data fields
	name string
	time time.Time
}

func (s someclass) GetSnapshot(newname string) Snapshot {
	return Snapshot{
		data: s.data,
		name: newname,
		time: time.Now(),
	}
}

func (s *someclass) restore(snap Snapshot) {
	s.data.field1 = snap.data.field1
	s.data.field2 = snap.data.field2
	s.data.field3 = snap.data.field3
}

//GetTime GetTime
func (s Snapshot) GetTime() time.Time {
	return s.time
}

//GetName GetName
func (s Snapshot) GetName() time.Time {
	return s.time
}

func main() {

	var smth someclass

	var history []Snapshot

	smth.data.field1 = "123"
	smth.data.field2 = 123
	smth.data.field3 = false
	history = append(history, smth.GetSnapshot("snap1"))
	smth.data.field1 = "999"
	smth.data.field2 = 999
	smth.data.field3 = true
	history = append(history, smth.GetSnapshot("snap2"))

	fmt.Println("last: ", smth.data)
	fmt.Println("history 0: ", history[0].data)
	fmt.Println("history 1: ", history[1].data)

	smth.restore(history[0])
	fmt.Println("restored: ", smth.data)
}
