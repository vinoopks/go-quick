package main

import (
	"fmt"
	"math/rand"
	"time"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

//Hello returns a greeting in the specified language
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefixText string) {

	switch language {
	case french:
		prefixText = frenchHelloPrefix
	case spanish:
		prefixText = spanishHelloPrefix
	default:
		prefixText = englishHelloPrefix
	}

	return
}

//init sets the initial values for variables used in function.
func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	lang := []string{"English", "French", "Spanish"}
	radomLang := lang[rand.Intn(len(lang))]
	fmt.Println(Hello("World", radomLang))

}
