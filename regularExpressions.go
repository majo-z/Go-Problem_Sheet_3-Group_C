//Adapted from & researched:
//https://golang.org/pkg/math/rand/
//https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
//https://techterms.com/definition/wildcard
//https://github.com/google/re2/wiki/Syntax
//https://stackoverflow.com/questions/15145659/what-do-i-and-i-in-regex-mean
//https://www.regular-expressions.info/wordboundaries.html
//https://stackoverflow.com/questions/16702924/how-to-explain-1-2-in-javascript-when-using-regular-expression

package main

import (
	"fmt"
	"math/rand"
	"regexp" //Package regexp implements regular expression search
	"time"
)

func ElizaResponse(input string) string { //input string, output string

	//MatchString checks whether a textual regular expression matches a string
	//func MatchString(pattern string, s string) (matched bool, err error)
	//if matched, _ := regexp.MatchString(".*father.*", input); matched { //.* matches zero or more characters (any) before and after father
	//if matched, _ := regexp.MatchString(".*[Ff]ather.*", input); matched { //[]inside square brackets = anything...logical or
	//if matched, _ := regexp.MatchString(`(?i).*father.*`, input); matched { //(?i) starts case-insensitive mode, (?-i) turns off case-insensitive mode
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched { //\b matches at a position that is called a "word boundary"...ex. only searching for "father", not grandfather
		return "Why don’t you tell me more about your father?"
	}

	//Capture "I am"

	//re := regexp.MustCompile(`(?i)I am (.*)`)
	re := regexp.MustCompile(`(?i)I am (.*)`) //fail fast so you find your mistake early

	//if matched, _ := regexp.MatchString(`(?i)I am (.*)`, input); matched { //Anything after I am + space
	if matched := re.MatchString(input); matched {
		//return "How do you know you are _?"
		return re.ReplaceAllString(input, "How do you know you are $1?") //$1 refers to a first match
	}

	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	return answers[rand.Intn(len(answers))]
}

func main() {
	//rand.Seed(42) //rovided seed value to initialize the default Source to a deterministic state
	//rand.Seed(time.Now().UTC().UnixNano())
	rand.Seed(time.Now().Unix()) //seed rand with Unix time

	/* answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	fmt.Println("Eliza replies:", answers[rand.Intn(len(answers))]) */

	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println()

	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println()

	fmt.Println("I was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println()

	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println()

	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println()

	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println()

	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println()

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println()

	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
	fmt.Println()

} //main
