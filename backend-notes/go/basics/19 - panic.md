# Introduction
La manière courante de gérer les problèmes en Go et l'utilisation des erreurs.
Cependant certaines fois le programme ne peut plus continuer après une situation anormale, 
par exemple, dans le cas des runtime errors (array out of bounds). 
Dans ce cas on fait appel à `panic` pour arrêter le programme.
Quand une fonction rencontre un `panic` l'exécution est arretée. Seules les deferred functions continuent jusqu'à leur fin.
Ce processus d'arrêt continue jusqu'à ce que toutes les fonctions de la Goroutine aient terminé.
A ce moment le programme print le panic message suivi d'une stack trace et termine.

Cependant, il est possible de regagner le contrôle après un panic à l'aide de `recover`.

**L'utilisation de panic et recover doit être évitée au maximum afin d'utiliser des erreurs.**

Il y a 2 cas où les panics sont interessantes :
- Une erreur du développeur : Imaginons une méthode qui accepte un pointeur en entrée et que le développeur appelle la méthode avec nil. A ce moment on peut panic puisque c'est la faute du développeur.
- Une erreur irrécupérable : 

# Comment utiliser panic ?
Les arguments passé à panic seront affichés dans la console.
```go
func fullName(firstName *string, lastName *string) {  
    defer fmt.Println("deferred call in fullName") // Sera affiché puisque panic laisse les deferred s'executer
    fmt.Printf("test") // Sera affiché même si panic est déclenché
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil") // lorsqu'un panic est rencontré, le programme s'arrête à ce moment
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
}

func main() {  
    defer fmt.Println("deferred call in main") // Sera affiché puisque panic laisse les deferred s'executer
    firstName := "Elon"
    fullName(&firstName, nil) // Affiche : "panic: runtime error: last name cannot be nil"
    fmt.Println("returned normally from main")
}
```

# Recover

Recover est une fonction permettant de regagner le contrôle d'une panicking Goroutine.
Recover ne fonctionne que dans la panicking goroutine.
Recover est utile uniquement lorsqu'utilisé dans une deferred function.
Exécuter un recover dans une deferred function stop la panicking sequence en restaurant l'exécution normale et en retournant les arguments passés à panic.
```go
func recoverName() {  
    if r := recover(); r!= nil { // recover() retourne la valeur passée à panic.
        fmt.Println("recovered from ", r)
    }
}

func fullName(firstName *string, lastName *string) {  
    defer recoverName() // Après l'exécution de cette fonction le contrôle revient à la goroutine parent
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
}
```

# Runtime panics
```go
func r() {  
    if r := recover(); r != nil {
        fmt.Println("Recovered", r)
        debug.PrintStack() // Permet de récupérer la stack trace et de la print
    }
}

func a() {  
    defer r()
    n := []int{5, 7, 4}
    fmt.Println(n[3]) // Provoque un panic
    fmt.Println("normally returned from a")
}

func main() {  
    a()
    fmt.Println("normally returned from main")
}
```