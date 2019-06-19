# Introduction

Un defer statement repousse l'exécution d'une fonction juste avant le return de celle-ci.
Deferpeut être utilisé sur une méthode ou une fonction.
Exemple :

```go
func main() { // Ce programme affiche "hello world"
	defer fmt.Println("world")
	fmt.Println("hello")
}
```

Les arguments d'une deferred function sont évalués lorsque le defer statement est exécuté et non lorsque l'appel de la fonction est réalisé.

La deferred function affichera 5 dans cet exemple.
```go
func printA(a int) {  
    fmt.Println("value of a in deferred function", a)
}
func main() {  
    a := 5
    defer printA(a)
    a = 10
    fmt.Println("value of a before deferred function call", a)
}
```
# Stack of defers
S'il y a plusieurs defer pushed onto a stack, elles sont lancées en last-in-first-out.
Cela permet par exemple de reverse des strings, des slices.

Le programme suivant va afficher "counting", "done", et "9 8 7 6 5 4 3 2 1 0"
```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

https://blog.golang.org/defer-panic-and-recover

# Utilité de defer
Si dans une fonction / méthode plusieurs code flow (if/switch/...) doivent appeler la/les même fonction avant un return alors l'utilisation de defer est plus interessante.
```go
func (r rect) area(wg *sync.WaitGroup) {  
    defer wg.Done()
    if r.length < 0 {
        fmt.Printf("rect %v's length should be greater than zero\n", r)
        return
    }
    if r.width < 0 {
        fmt.Printf("rect %v's width should be greater than zero\n", r)
        return
    }
    area := r.length * r.width
    fmt.Printf("rect %v's area %d\n", r, area)
}
```