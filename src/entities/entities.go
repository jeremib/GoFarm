package entities

import (
	"fmt"
	"math/rand"
	"time"
)

type Ant struct {
	Name	string
	Health 	int
	Hunger 	int
}

func (a *Ant) Work() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	h := (r1.Intn(13) + 1) * -1

	// fmt.Printf("⚒\t%s\t%d\n", a.Name, h)

	a.IncreaseHealth(h)
}


func (a *Ant) Eat() {
	if ( a.canFindFood() ) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		h := r1.Intn(5) + 1


		// fmt.Printf("🍎\t%s\t%d\n", a.Name, h)

		a.IncreaseHealth(h)
	}
}

func (a *Ant) Sleep() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// s := r1.Intn(5)
	h := r1.Intn(5) + 1

	// fmt.Printf("🛏\t%s\t%d\n", a.Name, h)

	// time.Sleep(1 * time.Second)

	

	a.IncreaseHealth(h)
}

func (a *Ant) IncreaseHealth(h int) {
	a.Health += h
	if ( a.Health > 100 ) {
		a.Health = 100
	}
}

func (a Ant) Die() {
	fmt.Printf("💀\t%s", a.Name)
}

func (a Ant) Info() {
	fmt.Printf("\n%s\t🏥 %d\t🥕 %d\n", a.Name, a.Health, a.Hunger)
}

func (a Ant) canFindFood() bool {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	h := r1.Intn(100)

	return h % 10 != 0
}