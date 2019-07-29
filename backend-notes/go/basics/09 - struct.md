# Struct

Une structure est un type crée par le développeur.
Il s'agit d'une collection de champs. C'est l'équivalent d'un objet en JS.

## Déclarer un struct
- Named struct :
```go
type Vertex struct {
    X int
    Y int
}
```
- Anonymous struct :
```go
var vertex struct {
    X int
    Y int
}
```
## Créer un struct
- Avec un named struct :
```go
vertex := Vertex{12, 4}
```
- Avec un anonymous struct
```go
emp1 := emp struct{
    firstName string
    age, salary int
}{"Youssef", 28, 45000}
```
## Accéder à un struct
- Avec un named struct :
```go
v := Vertex{1, 2}
v.X = 4
fmt.Println(v.X)
```
- Avec un anonymous struct
```go
type Person struct {
    int
    string
}

func main(){
    pers := Person{22, "john"}
    fmt.Println(pers.int) // Par défaut les champs portent le nom du type
    fmt.Println(pers.string)
}
```

## Promoted fields
Il est possible de promouvoir un nested struct en le déclarant anonymement.
```go
type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age  int
    Address // Adress est déclaré anonymement
}


func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.Address = Address{
        city:  "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.city) //city is promoted field
    fmt.Println("State:", p.state) //state is promoted field
}
```

## Pointeur vers un struct

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
## Exported Structs and fields
De même que les fonctions, les struct type ayant une majuscule sont exportables et peuvent être accédés depuis d'autres packages. La même chose est valable pour les champs d'un struct.
```go
package computer

type Spec struct { //exported struct  
    Maker string //exported field
    model string //unexported field
    Price int //exported field
}
```

## Structs equality
Les structs sont des value types et sont comparables si chacun de leur champs est comparable.