//Adapted from https://golang.org/pkg/math/rand/
//Adapted from https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//rand.Seed(42) //try changing this number
	//rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UTC().UnixNano())
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	fmt.Println("Eliza replies:", answers[rand.Intn(len(answers))])
} //main
