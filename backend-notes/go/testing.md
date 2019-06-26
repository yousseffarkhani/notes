Source : https://quii.gitbook.io/learn-go-with-tests/

# Introduction

En Go, il n'y a pas besoin de framework de test, tout est fourni directement.
**Avant de démarrer les tests, il faut décomposer notre programme en petites fonctionnalités que nous pourrons tester séparément.**

# Etapes de test

1. Ecrire le test
2. Faire en sorte que le programme compile
3. lancer le test et voir si ça fail bien. Vérifier que la description du texte explique de manière simple le problème. Se poser la question que vérifie mon test ?
4. Ecrire le minimum de code pour passer le test
5. Refactorer (même le code de test)
6. Retester

# Tips

- La définition de refactoring est le code change mais le comportement reste le même. En théorie le test n'a pas besoin de changer.
  - Est ce que je teste le comportement ou l'implémentation du code ?
  - Si je refactore le code, est ce que j'aurais beaucoup de choses à changer dans les tests ?
- Ne pas tester les private functions car elles ont avoir avec l'implémentation.
- Si un test a besoin de 3 mocks pour fonctionner alors il faut repenser le design de la fonction.
- Utiliser les spies avec précaution (etre sur que les détails sont intéressants)

# Règles à suivre

- Le fichier de test doit se nommer xxx_test.go
- La fonction de test doit démarrer avec le mot Test.
- La fonction de test ne prend qu'un argument : `t *testing.T`
  - Ce type permet de faire des appels aux méthodes suivantes :
    - `t.Fail()`
    - `t.Errorf("got '%s' want '%s'", got, want)` : permet de print un msg et de faire le test
    - `t.Fatal("didn't get an error but wanted one")` : permet de stopper le test si appelé

# Syntaxe

Dans la CLI écrire `go test` ou `go test -v` pour avoir le détail des tests.

```go
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) { // Il est possible de déclarer des fonctions dans les fonctions
		t.Helper() // Permet de spécifier que cette méthode est un helper. En faisant cela, la ligne d'erreur affichée dans la console sera celle de la fonction appelée et non du fichier test.go
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) { // t.Run correspond à un subtest
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want) // t est transmis pour dire au code de fail quand le test est KO
	})

	t.Run("Saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Adding ES as 2nd parameter should define the greeting in spanish", func(t *testing.T) {
		got := Hello("Carlos", "ES")
		want := "Hola, Carlos"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Adding FR as 2nd parameter should define the greeting in french", func(t *testing.T) {
		got := Hello("Carlos", "FR")
		want := "Bonjour, Carlos"
		assertCorrectMessage(t, got, want)
	})
}
```

# Table based tests

```go
func TestArea(t *testing.T) {
	// Table based tests are useful for testing various implementations of an interface or if the data being passed in to a function has lots of different requirements.
	areaTests := []struct { // Anonymous struct
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 2.0, Height: 4.0}, hasArea: 8.0},
		{name: "Circle", shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, v := range areaTests {
		t.Run(v.name, func(t *testing.T) {
			t.Helper()
			got := v.shape.Area()
			if got != v.hasArea {
				t.Errorf("%#v got %.2f want %.2f", v.shape, got, v.hasArea) // %#v pour afficher un struct : geom.Triangle{Base:12, Height:6}
			}

		})
	}
}
```

# Benchmark

Les benchmarks permettent de tester la vitesse des fonctions.

Ajouter dans le fichier de test :

```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

Pour lancer le benchmark : `go test -bench=.`

Résultat :

goos: linux
goarch: amd64
pkg: iteration
BenchmarkRepeat-4 10000000 184 ns/op

Cela signifie que le test a été lancé 1000000 fois pour un temps moyen de 184ns.

# Coverage
Le coverage permet de voir combien de % du code source est testé.
A noter qu'il n'y a pas non plus besoin d'avoir 100% de coverage.
Pour vérifier : `go test -cover`
Résultat : coverage: 100.0% of statements

Il est également possible de voir via une présentation en html :

1. `go test -coverprofile=coverage.out`
2. `go tool cover -html=coverage.out`

# Errcheck

Errcheck permet de vérifier que toutes les erreurs renvoyées par le programme ont été testées.

1. Installer errcheck : `go get -u github.com/kisielk/errcheck`
2. Puis saisir : `errcheck` dans la CLI

# Dependency injection

La dependency injection permet de contrôler la dépendance d'une fonction à un élément.
Par exemple, si on veut tester :

```go
func Greet(name string) {
    fmt.Printf("Hello, %s", name)
}
```

`fmt.Printf` va print ça sur `stdout`. Cet évenement est impossible à capturer pour nous dans le cadre des tests.
Ce qu'il nous faut c'est pouvoir injecter la dépendance du print.

Notre fonction n'a pas besoin de savoir où sera printée la phrase c'est pourquoi elle devrait accepter une interface plutôt qu'un type spécifique.

En faisant cela, on peut contrôler l'endroit où sera print la phrase.

En analysant la fonction `Printf`, on peut voir qu'en réalité on utilise `Fprintf` avec comme 1er argument, l'endroit où la phrase va être print (ici `os.Stdout`):

```go
func Printf(format string, a ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}
```

Qu'est ce que `os.Stdout` ? Pour cela analysons la fonction `Fprintf` :

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintf(format, a)
    n, err = w.Write(p.buf)
    p.free()
    return
}
```

Il s'agit en réalité d'un `io.Writer`
`Writer` est une interface du package `io`.

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

## Code

Dans ce code on peut voir que le fait de rendre générique notre fonction, nous a permis de l'utiliser dans plusieurs contextes : CLI, http, testing.

Pour cela, il faut créer une interface comportant la méthode que nous souhaitons injecter.

```go
// Modification de la dépendance à stdout
func Greet(writer io.Writer, name string) { // En utilisant l'interface io.Writer notre fonction devient générique
	fmt.Println("Hello, %s", name) // Test is KO because it prints the value to stdout
	fmt.Fprintf(writer, "Hello, %s", name) // Similaire à fmt.Printf mais prend un writer en 1er argument (fmt.Printf envoi par défaut à stdout)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) { // ResponseWriter implements io.writer
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Elodie")
	http.HandleFunc("/", MyGreeterHandler)
	http.ListenAndServe(":5000", nil)
}
```

```go
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{} // Buffer type from bytes package implements Writer interface
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
```

# Mocking

Le mocking est intéressant car en tant que développeur, nous voulons avoir des feedback loops rapides.
Le mocking permet :

- de simuler un appel de WS / des timers / ... et donc de ne pas avoir à attendre
- Tester un state particulier du système
- Eviter la dépendance envers des services / databases

## Spies

Pour mocker, il faut :

1. Créer une interface comportant la méthode que nous souhaitons injecter.
2. Dans le fichier de test, nous allons créer un spy c'est à dire créer un struct comportant les indicateurs souhaités (nombre d'appels, arguments passés) et implémentant l'interface définie au préalable.

**Source code**

```go
func main() {
sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
Countdown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration // Permet de configurer le time slept
	sleep    func(time.Duration) // Permet de donner une fonction ayant pour paramètre un type time.Duration
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration) // Intéressant, utilisation d'une fonction donnée en paramètre
}

func Countdown(writer io.Writer, sleeper Sleeper) {
for i := countdownStart; i > 0; i-- {
sleeper.Sleep()
fmt.Fprintln(writer, i)
}
sleeper.Sleep()
fmt.Fprint(writer, finalWord)
}

```

**fichier de test**

```go
type CountdownOperationsSpy struct { // Implements both io.Writer and Sleeper
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {

	t.Run("Prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{}) // Nous permet de ne pas avoir à attendre pendant le test en faisant appel à CountdownOperationsSpy.Sleep()

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
```

# Elements interessants

`reflect.DeepEqual(got, want)` : Permet de comparer 2 types (pratique pour les slices)
