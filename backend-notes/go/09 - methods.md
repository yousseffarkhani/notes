# Introduction

Go ne dispose pas de classes. Cependant, il est possible de définit des méthodes sur les types.

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

Il est également possible de décalrer une méthode sur un non-struct type.

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

Il est également possible de réécrire ce code qu'avec des fonctions

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

A l'inverse :

```go
type Vertex struct {
    X, Y float64
}

func Abs(v Vertex)float64{
    return v.X*v.X + v.Y*v.Y
}

func (v Vertex) Abs()float64{
    return v.X*v.X + v.Y*v.Y
}

func main(){
    v := &Vertex{3,4} // Pointeur
    fmt.Println(Abs(*v))
    fmt.Println(v.Abs()) // A noter qu'il n'y a pas besoin de l'appeler avec *v car Go interprète v.Abs() comme étant (*v).Abs().
}
```
