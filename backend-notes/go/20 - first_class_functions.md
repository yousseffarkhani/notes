# Introduction
Un langage est dit first class functions lorsqu'il permet aux fonctions d'être assignées à des variables, d'être passées en argument d'autres fonctions et d'être retournées depuis d'autres fonctions.

# Anonymous function
## Stockage de la fonction dans une variable
```go
a := func(){
    fmt.Println("anonymous function")
}
a() // Appel de la fonction anonyme
fmt.Printf("%T", a) // Affiche func()
```
## Appel direct de la fonction anonyme
```go
func(s string){
    fmt.Println(s, "anonymous function")
}("hello") // Affiche "hello anonymous function"
```


# User define function types
Il est également possible de définir des types de fonctions.
```go
type add func(a, b int) int

func main(){
    var addFunction add = func(a, b int) int { // Définition d'une fonction de type add
        return a+b
    }
    s := addFunction(5 + 6)
    println(s)
}
```

# Higher order functions
Une higher order function est une fonction acceptant une autre fonction en argument ou retournant une fonction.
## Fonction acceptant une fonction en paramètre
```go
func simple(f func(a,b int) int){
    println(f(60, 7))
}

func main(){
    f := func(a, b int) int{
        return a + b
    }

    simple(f)
}
```
## Fonction retournant une fonction 
```go
func simple() func(a, b int) int{ // Le return type est de type func
	f := func(a, b int) int{
       		return a + b
    	}
	return f
}

func main(){
    s := simple()
	println(s(60, 7))
}
```
# Closures

Les closures sont des cas spécial de fonctions anonymes. 
Celles-ci sont capables d'accéder aux variables en dehors du scope de leur fonction.
```go
func appendStr() func(string) string {  
    t := "Hello"
    c := func(b string) string {
        t = t + " " + b // Utilisation de la variable t en dehors du scope de la fonction donc la fonction sera une closure liée à t
        return t
    }
    return c
}

func main() {  
    a := appendStr() // Création de closure lié à la variable t
    b := appendStr()
    fmt.Println(a("World")) // Affiche "Hello World"
    fmt.Println(b("Everyone"))

    fmt.Println(a("Gopher")) // Affiche "Hello World Gopher" car la fonction garde en mémoire les changements de t
    fmt.Println(b("!"))  // Affiche "Hello Everyone !" car les fonctions a et b ne partagent pas la même variable t
}
```
Autre exemple
```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

# Usage pratique des first class functions
Cela permet de créer des filter ou bien des map functions (fonction appliquée à l'ensemble d'une collection (slice/array)).
Filter :
```go
type student struct {
	firstName, lastName, grade, country string
}

func filter(students []student, f func(student student) bool) []student {
	var result []student
	for _, v := range students {
		if f(v) == true {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "B",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}

	students := []student{s1, s2}
	filterResult := filter(students, func(student student) bool {
		if student.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(filterResult)
}
```

Map :
```go
func multiplyMap(slice []int, f func(n int) int) []int { // Les fonctions agissant sur tous les éléments d'un slice sont appelées map functions.
	var result []int
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}

func main() {
	intSlice := []int{1, 2, 3, 4}
	result := multiplyMap(intSlice, func(n int) int {
		return n * 2
	})
	fmt.Println(result)
}
```