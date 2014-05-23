package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	var target string

	flag.StringVar(&target, "t", "100100", "the target to try to get to")
	flag.Parse()

	generation := spawnGeneration([2]string{&target, &target})

	fittest := getFittest(generation)

	fmt.Println(fittest)

	// for _, value := range generation {
	// 	fmt.Println(value)
	// }

}

func createChild() string {

	child := ""

	for i := 0; i < 6; i++ {
		child = child + strconv.Itoa(rand.Intn(2))
	}

	return child

}

func spawnGeneration(parents [2]string) [100]string {

	var generation [100]string

	for i := 0; i < len(generation); i++ {
		generation[i] = createChild()
	}

	return generation

}

// Gets the fitness of a child. It will return
// a int which is how far off it is from the
// target
func getFitness() int {

	return 1

}

func getFittest(generation [100]string) [2]string {

	return [2]string{"1", "0"}

}
