# Pointers

## Introduction

Un pointer retient l'adresse mémoire d'une valeur. Cela est pratique quand on veut passer une variable contenant une valeur lourde en mémoire.

## Déclarer un pointeur

```go
var p *int // Déclare un pointeur. La valeur par défaut du pointer est `nil`.
i := 42
p = &i // & génére un pointeur
fmt.Println(*p) // Permet de lire i par l'intérmédiaire du pointer p (aussi appelé "dereferencing")
*p = 21 // Modifie la valeur de i
```

# Struct

Il s'agit d'une collection de champs. C'est l'équivalent d'un objet en JS.

## Déclarer un struct

```go
type Vertex struct {
X int
Y int
}
```

## Accéder à un struct

```go
v := Vertex{1, 2}
v.X = 4
fmt.Println(v.X)
```

## Pointer vers une propriété du struct

```go
v := Vertex{1, 2}
p := &v
p.X = 1e9 // Pas besoin de *p ici
fmt.Println(v)
```

/!\ La valeur de v.X est modifié ici car p est un pointeur vers l'adresse mémoire.

## Structs literals

```go
var (
v1 = Vertex{1, 2} // has type Vertex
v2 = Vertex{X: 1} // Y:0 is implicit
v3 = Vertex{} // X:0 and Y:0
p = &Vertex{1, 2} // has type *Vertex, result: &{1 2}
)
```
