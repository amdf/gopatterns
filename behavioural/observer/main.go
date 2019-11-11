package main

import "fmt"

type publisher struct {
	srcinfo string
	subs    map[string]subscriber
}

func (pub *publisher) subscribe(sub *mySubscriber) {
	if pub.subs == nil {
		pub.subs = make(map[string]subscriber)
	}
	pub.subs[sub.name] = sub
}

func (pub *publisher) unsubscribe(sub *mySubscriber) {
	delete(pub.subs, sub.name)
}

func (pub *publisher) notify() {
	for _, sub := range pub.subs {
		sub.update(pub)
	}
}

type subscriber interface {
	update(*publisher)
}

type mySubscriber struct {
	subscriber
	name string
	info string
}

func (sub *mySubscriber) update(pub *publisher) {
	sub.info = pub.srcinfo
}

func (sub mySubscriber) doStuff() {
	fmt.Println(sub.name, ": ", sub.info)
}

func main() {

	pub := publisher{srcinfo: "ADADADAD"}
	sub1 := mySubscriber{name: "sub1"}
	sub2 := mySubscriber{name: "sub2"}
	sub3 := mySubscriber{name: "sub3"}

	pub.subscribe(&sub1)
	pub.subscribe(&sub2)
	pub.subscribe(&sub3)

	pub.notify()
	sub1.doStuff()
	sub2.doStuff()
	sub3.doStuff()

	pub.unsubscribe(&sub2)
	pub.srcinfo = "BBCCBBCC"
	pub.notify()
	sub1.doStuff()
	sub2.doStuff()
	sub3.doStuff()
}
