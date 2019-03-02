package main

// docker container run -it --rm -p 6060:6060 -v `pwd`:/go golang:1.8

import (
	// "fmt"
	"entities"
	// "runtime"
	"sync"
	// "time"
	// "math/rand"
)





func main() {
	ants := []entities.Ant{
		{Name: "🐛", Health: 100, Hunger:100, Age: 1},
		{Name: "🐜", Health: 100, Hunger:100, Age: 1},
		{Name: "🐞", Health: 100, Hunger:100, Age: 1},
	}

	var wg sync.WaitGroup
	
	wg.Add(len(ants))

	for _,a := range ants {
		go birth(a, &wg)
	}

	wg.Wait()
}

func birth(a entities.Ant, wg *sync.WaitGroup) {
	defer wg.Done()

	d := 0
	for ok := true; ok; ok = (a.Health > 0) {
		a.Eat()
		a.Work()
		a.Eat()
		a.Sleep()

		if ( a.Health <= 0 ) {
			// a.Die()
			a.Info()
		}

		a.IncreaseAge(1)

		// if ( d % 5000 == 0) {
			// a.Info()
		// }
		d++
	}
}