package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

// Hello returns a personalized greeting in a given language.
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

// Hello returns a greeting prefix in a given language.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
