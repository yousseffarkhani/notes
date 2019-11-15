https://dave.cheney.net/practical-go/presentations/qcon-china.html
https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters
https://peter.bourgon.org/go-best-practices-2016/#repository-structure

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
# Rendre les dépendances explicites
Cela rend le code plus lisible et facilite les tests.
## Types de dépendances possibles
- Output : Utiliser fmt.FPrintf(...)
- Input :
- database handles :
- Loggers : log.FPrintf
- commandline flags
- Requests :
Dans le code suivant http.Get dispose d'une dépendance interne vers DefaultClient, ce qui rend le test compliqué à réaliser.
```go
func foo() {
    resp, err := http.Get("http://zombo.com")
    // ...
}
```
Un moyen d'externaliser cette dépendance est :
```go
func foo(client *http.Client) {
    resp, err := client.Get("http://zombo.com")
    // ...
}
```
Cependant cette forme est trop restrictive puisqu'un struct de type http.Client est attendu, on peut rendre cela plus flexible en créant une interface :
```go
type Doer interface {
    Do(*http.Request) (*http.Response, error)
}

func foo(d Doer) {
    req, _ := http.NewRequest("GET", "http://zombo.com", nil)
    resp, err := d.Do(req)
    // ...
}
```
Désormais pour effectuer notre test, il faudra créer une implémentation de l'interface Doer.
# Make the zero value useful
Vérifier les nil values et les initialiser