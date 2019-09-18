https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters

```go
const englishHelloPrefix = "Hello, " // Utilisation de constantes
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "ES"
const french = "FR"

func Hello(s string, lang string) string {
	suffix := s
	if s == "" {
		suffix = "World"
	}
	return greetingPrefix(lang) + suffix
}

func greetingPrefix(lang string) (prefix string) { // la fonction démarre avec une minuscule donc elle est privée
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix // A noter qu'il n'y a pas besoin de déclarer prefix avec :=
	case french:
		prefix = frenchHelloPrefix
	default: // Ne jamais oublier default
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", "")) // Good behavior to separate the code from the outside world (side-effects). Printing to stdout is a side effet.
}
```
