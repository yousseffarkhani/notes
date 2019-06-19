# Introduction
Les erreurs expriment un comportement anormal du programme.
Les erreurs en Go sont des valeurs standard de type `error`, elles peuvent être stockées dans des variables, ...

```go
func main() {  
    f, err := os.Open("/test.txt") // Signature de la fonction Open : func Open(name string) (file *File, err error)
    if err != nil {
        fmt.Println(err) // Fait appel à err.Error()
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}
```

# Définition d'une erreur
Lorsque on print une error, `fmt.Println` appelle la fonction Error() pour avoir la description de l'erreur. 
```go
type error interface {
    Error() string // N'importe quel type implémentant la fonction Error() peut être utilisé comme une erreur.
}
```
# Exploiter les erreurs
Cette partie est la suite de l'exemple en introduction avec os.Open("/test.txt")
## Récupérer le path
Il existe plusieurs solutions pour récupérer cette donnée :
### Asserting the underlying struct type with struct fields
En lisant la documentation de la fonction Open, on peut voir qu'une erreur de type *PathError est retournée.
PathError est un struct type avec l'implémentation suivante de Error() :

```go
type PathError struct {  
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() } // Correspond au message d'erreur retourné en introduction  
```

Nous allons modifier le programme initial afin de print le path:
```go
func main() {  
    f, err := os.Open("/test.txt")
    if err, ok := err.(*os.PathError); ok { // Type assertion to get the underlying value of type *os.PathError from the error interface.
        fmt.Println("File at path", err.Path, "failed to open")
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}
```

### Asserting the underlying struct type with methods
```go
type DNSError struct { ... }

func (e *DNSError) Error() string { ... }

func (e *DNSError) Timeout() bool { ... }

func (e *DNSError) Temporary() bool { ... }
```
```go
func main() {  
    addr, err := net.LookupHost("golangbot123.com") // Trying to get the IP adress
    if err, ok := err.(*net.DNSError); ok { // Type assertion to get the underlying value of type *net.DNSError from the error interface.
        if err.Timeout() {
            fmt.Println("operation timed out")
        } else if err.Temporary() {
            fmt.Println("temporary error")
        } else {
            fmt.Println("generic error: ", err)
        }
        return
    }
    fmt.Println(addr)
}
```
### Direct comparison
La 3ème méthode est d'utiliser la comparaison directe avec une variable de type error.

Nous allons utiliser la fonction Glob du package filepath pour illustrer cet exempel/
Cette fonction permet de récupérer le nom des fichiers matchant un certain pattern. Cette fonction retourne une erreur si le pattern est malformé (ErrBadPattern).
```go
func main() {  
    files, error := filepath.Glob("[")
    if error != nil && error == filepath.ErrBadPattern { // Comparaison directe avec filepath.ErrBadPattern si OK alors l'erreur est due à un malformed pattern.
        fmt.Println(error)
        return
    }
    fmt.Println("matched files", files)
}
```
# Ne jamais ignorer les erreurs
Les erreurs aident à la résolution de problème.
Par exemple si on ignore l'erreur de l'exemple précédent, on peut penser qu'il n'existe pas de fichiers portant ce nom alors qu'il s'agit d'un pattern malformé.
```go
func main() {  
    files, _ := filepath.Glob("[")
    fmt.Println("matched files", files) // Retourne "matched files []"
}
```

