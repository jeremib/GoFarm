package entities

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Ant struct {
	Name	string
	Health 	int
	Hunger 	int
	Age		int
}

func (a *Ant) Work() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	r := r1.Intn(12)
	m := float64(a.GetAge() - 100) / float64(100) * -1

	h := int(math.Floor(float64(r) / float64(m)))

	// w := r - int(math.Ceil(float64(r) - m))

	// fmt.Printf("%v", w)

	he := h * -1

	// fmt.Printf("⚒\t%s\t%d\n", a.Name, h)

	a.IncreaseHealth(he)
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


func (a *Ant) IncreaseAge(h int) {
	a.Age += h
}

func (a Ant) Die() {
	fmt.Printf("💀\t%s", a.Name)
}

func (a Ant) GetAge() int {
	return a.Age
}

func (a Ant) Info() {
	fmt.Printf("\n%s\t🏥 %d\t🥕 %d\t ❤️ %d\n", a.Name, a.Health, a.Hunger, a.GetAge())
}

func (a Ant) canFindFood() bool {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	h := r1.Intn(100)

	return h % 10 != 0
}