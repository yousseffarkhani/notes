# Introduction

Go ne dispose pas de classes. Cependant, il est possible de définir des méthodes sur les types.
L'avantage des méthodes est qu'on peut utiliser le même nom de fonction et avoir un comportement différent en fonction du type utilisé.

Terme utilisés :
- Receiver : Type recevant la fonction
# Déclarer une méthode

Une méthode est une fonction avec un receveur précisé avant la déclaration de la fonction.

```go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X + v.Y)
}

func main(){
    v := Vertex{3,4}
    fmt.Println(v.Abs())
}
```

Il est également possible de déclarer une méthode sur un non-struct type.

```go
type MyFloat int

func (f MyFloat) Abs() int {
	return int(f*f) // Il est important de caster la valeur sinon erreur
}

func main(){
    f := MyFloat(3)
    fmt.Println(f.Abs())
}
```

/!\ La déclaration d'une méthode ne peut se faire que si le type est défini dans le même package.

# Pointer receivers
Jusqu'à présent, nous n'avons vu que des méthodes avec des value receivers.

Les pointer receiver sont intéressants car:

- La méthode peut modifier la valeur en paramètre
- Evite la copie d'une valeur à chaque appel de la méthode. Cela est efficace surtout quand il y a un large struct.

En général, toutes les méthodes sur un type doivent soit être des value receivers ou des pointer receivers mais pas un mix des 2.

```go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs()float64{
    return v.X*v.X + v.Y*v.Y
}

func (v Vertex) Scale1(s float64){
    v.X = v.X*s
    v.Y = v.Y*s
}

// Pointer receiver method
func (v *Vertex) Scale2(s float64){
    v.X = v.X*s
    v.Y = v.Y*s
}

func main (){
    v := Vertex{3,4}
    v.Scale1(10) // Cette fonction ne va pas modifier v
    v.Scale2(10) // Cette fonction va modifier v. A noter qu'il n'y a pas besoin de l'appeler avec &v.Scale2(10) car Go interprète v.Scale2(10) comme étant (&v).Scale2(10).
    fmt.Println(v.Abs())
}
```

## Accéder à la méthode d'un anonymous field
```go
type adress struct {
    city, state string
}

type person struct {
    firstName string
    adress
}

func (a adress) fullAdress(){
    fmt.Println(a.city, a.state)
}

func main(){
    pers := person{"youssef", address{"LA", "California"}}
    pers.fullAdress()
}
```
## Définir des méthodes sur des non struct types
Jusqu'à présent les méthodes ont été définies sur des struct types.
Pour définir une méthode sur un type, il faut impérativement que la définition du receiver et la définition de la méthode se situent dans le même package.

C'est pourquoi, il est impossible de définir une méthode sur les built-in types (int, bool, ...)
```go
package main

func (a int) add(b int) {  // Génère une erreur
}

func main() {

}
```
Pour réaliser cela, il faut créer un alias et ensuite créer une méthode.

```go
type myInt int

func (a myInt) add(b myInt) myInt {  
    return a + b
}
```

## Pourquoi utiliser des méthodes quand on a des fonctions ?

Il est également possible de réécrire le code qu'avec des fonctions

```go
type Vertex struct {
    X, Y float64
}

func Abs(v Vertex)float64{
    return v.X*v.X + v.Y*v.Y
}

func Scale1(v Vertex, s float64){
    v.X = v.X*s
    v.Y = v.Y*s
}

func Scale2(v *Vertex, s float64){
    v.X = v.X*s
    v.Y = v.Y*s
}

func main(){
    v := Vertex{3,4}
    Scale1(v, 10) // Cette fonction ne va pas modifier v
    Scale2(&v, 10) // Cette fonction va modifier v
    fmt.Println(Abs(v))
}
```
L'inconvénient est que
- Il est impossible de déclarer plusieurs fonctions portant le même nom en fonction du type.
- Une fonction n'est pas flexible comme une méthode en ce qui concerne l'appel avec un pointer ou une value.