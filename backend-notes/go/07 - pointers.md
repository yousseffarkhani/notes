# Pointers
http://www.cplusplus.com/doc/tutorial/pointers/

## Introduction

Un pointer retient l'adresse mémoire d'une valeur. Cela est pratique quand on veut passer une variable contenant une valeur lourde en mémoire.

## Déclarer un pointeur

```go
var p *int // Déclare un pointeur de type int. La valeur par défaut du pointer est `nil`.
size := new(int) // Autre manière de déclarer un pointeur de type int
i := 42
p = &i // & Permet de récupérer l'adresse de la variable et d'assigner l'adresse à la variable
fmt.Println(*p) // Permet de lire i par l'intérmédiaire du pointer p (aussi appelé "dereferencing")
*p = 21 // Modifie la valeur de i
```

## Retourner un pointer d'une fonction
```go
func hello() *int {
    i := 55
    return &i
}

func main(){
    num := hello()
    fmt.Println(*num)
}
```
## Modifier un array dans une fonction à l'aide de l'adresse
Il est possible d'envoyer l'adresse mémoire de l'array est ensuite modifier mais le plus simple est d'utiliser les slices.

```go
func modify(slice []int) {
    slice[0] = 55
}

func main(){
    num := [3]int{1, 2, 34}
    modify(num[:])
    fmt.Println(num)
}
```
