package main

import "fmt"

const (
	portuguese       = "Português"
	french           = "French"
	englishPrefix    = "Hello "
	portuguesePrefix = "Olá "
	frenchPrefix     = "Bonjour "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portuguesePrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "GO"
	}
	return greetingPrefix(language) + name + "!"
}

func main() {
	fmt.Println(Hello("Filipe", "Português"))
}
