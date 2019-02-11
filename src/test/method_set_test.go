package test

import (
	"fmt"
	"testing"
)

type duration int

func (d *duration) pretty() {
	fmt.Printf("%d", d)
}

func TestNoPointer(t *testing.T) {
	d := duration(42)
	d.pretty()

	// cannot call pointer method on duration(42)
	// cannot take the address of duration(42)
	// duration(42).pretty()
}

type actor struct{}

func (a actor) actA() {
	fmt.Println("acting A")
}

func (a *actor) actB() {
	fmt.Println("acting B")
}

func TestPointerValue(t *testing.T) {
	var a *actor
	a.actA()
}

func TestPointerPointer(t *testing.T) {
	var a *actor
	a.actB()
}

func TestValue(t *testing.T) {
	(actor{}).actA()
	// (actor{}).actB()
}

type action interface {
	do()
}

type actorA struct{}

type actorB struct{}

func (a actorA) do() {
	fmt.Println("a acting")
}

func (b *actorB) do() {
	fmt.Println("b acting")
}

func TestActorAValue(t *testing.T) {
	var a actorA
	a.do()
}

func TestActorAPointer(t *testing.T) {
	var a *actorA
	a.do()
}

func TestActorBValue(t *testing.T) {
	var b actorB
	b.do()
}

func TestActorBPointer(t *testing.T) {
	var b *actorB
	fmt.Println(b == nil)
	b.do()
}

func TestNilInvoke(t *testing.T) {
	(*actorB)(nil).do()
	fmt.Printf("%p\n", (*actorB)(nil))
}

var nilB *actorB

func bAction() action {
	return nilB
}

func bActor() *actorB {
	return nilB
}

func TestNilB(t *testing.T) {
	bAction := bAction()
	fmt.Println(bAction == nil)

	res, ok := bAction.(*actorB)
	fmt.Printf("res:%v, ok:%v\n", res, ok)

	bActor := bActor()
	fmt.Println(bActor == nil)
}
