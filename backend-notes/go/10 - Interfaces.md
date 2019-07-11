# Introduction

Une interface est un type définissant une liste de méthodes.
**Lorsque qu'un type définit toutes les méthodes de la liste on dit que le type implémente l'interface.**
L'interface spécifie les méthodes que doivent contenir les types et les types décident de l'implémentation de ces méthodes.

# Déclarer et implémenter une interface

```go
/* Interface definition */
type VowelsFinder interface{
    FindVowels() []rune
}

type myString string

/* myString implements VowelsFinder. Cela est très différent des autres langages car en Go il n'y a pas besoin d'expliciter le fait qu'un struct utilise une interface */
func (ms myString) FindVowels()[]rune{
    vowels := []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

func main(){
    name := myString("hello")
    var v VowelsFinder
    v = name // Possible car myString implémente VowelsFinder
    fmt.Println(v.FindVowels()) // Il n'est possible d'accèder qu'aux méthodes définies sur les interfaces
    fmt.Println(name.FindVowels())
}
```

# A quoi ça sert ?

```go
package main

import (
	"fmt"
)

type Permanent struct {
    empId    int
    basicpay int
    pf       int
}

type Contract struct {
    empId  int
    basicpay int
}

type SalaryCalculator interface {
    CalculateSalary() int
}

func (perm Permanent) CalculateSalary() int{
    return perm.basicpay + perm.pf
}
func (cont Contract) CalculateSalary() int{
    return cont.basicpay
}

func sumSalaries(employees []SalaryCalculator) int{
    sum := 0
    for _, v := range employees {
        sum += v.CalculateSalary()
    }
    return sum
}

func main() {
    perm1 := Permanent{1, 20000, 100}
    cont1 := Contract{11, 2000}
    employees := []SalaryCalculator{perm1, cont1} // Il est possible d'ajouter différents types dans un même slice
    fmt.Println(sumSalaries(employees)) // Cette méthode prend en paramètre une interface donc plusieurs types
}
```

Le principal interêt est que sumSalaries peut être étendu à n'importe quel nouveau type d'employé (du moment qu'il implémente l'interface en définissant les méthodes).

# Interface internal representation

Une interface peut être vue comme une combinaison d'une valeur avec un type (type, value).

# Empty interface

Une interface peut ne pas spécifier de méthodes mais contenir des valeurs. Le type de cette interface est interface{}.

```go
func main(){
    var i interface{}
    i = 42
    i = "hello"
}
```

# Type assertions

Une type assertion permet d'accéder aux valeurs de l'interface : `t := i.(T)`.
Cette déclaration affirme que l'interface i contient un type T et assigne les valeurs de T à la variable t.
Si i ne contient pas le type T alors une erreur (panic) apparait.

Pour tester si l'interface contient le type, il est possible d'utiliser `t, ok := i.(T)`.

```go
var i interface{} = "hello"

s := i.(string) // affecte hello à s
s, ok := i.(string)

f, ok := i.(float64) // Retourne 0 et false
f := i.(float64) // Retourne une erreur panic
```

# Type switches

Un type switch permet d'enchainer plusieurs type assertions.
Un type assertion peut également permettre de comparer avec une interface au lieu d'un type.

```go
function do(i interface{}){
    switch v := i.(type){ // type est le keyword utilisé
        case int :
            fmt.Println(v, v*2)
        case string :
            fmt.Println(v, len(v))
        default:
            fmt.Printf("I don't know about the type %T", v)
    }
}

func main(){
    do(24)
    do("hello")
    do(true)
}
```

# Implémenter des interfaces en utilisant des pointer receivers

```go
type Describer interface {
    Describe()
}

type Person struct{
    name string
    age int
}

func (p Person) Describe(){ // Implements Describer using value receiver
    fmt.Println(p.name, p.age)
}

type Adress struct {
    city, state string
}

func (a *Adress) Describe(){
    fmt.Println(a.city, a.state)
}

func main() {
    pers := Person{"youssef", 22}
    var d Describer
    d = pers
    d.Describe()
    d = &pers
    d.Describe()
    adress := Adress{"LA", "California"}
    var a Describer
    a = adress
    a.Describe() // Provoque une erreur car l'interface est implémenté en utilisant un pointer receiver et non un value receiver
    a = &adress
    a.Describe()
}
```

# Implémenter plusieurs interfaces

Un type peut avoir plusieurs interfaces s'il implémente les méthodes de ces interfaces.

# Embedded interfaces

Il est possible de créer de nouvelles interfaces à partir d'autres interfaces.
Les types implémentant l'ensemble des méthodes implémenteront l'interface.

```go
type Describer interface {
    Describe()
}

type Test interface {
    test()
}

type Embedded interface {
    Describer
    Test
}
```

# Nil interface values

Les interfaces ne détiennent ni valeurs ni types.
Si une méthode est appelée sur une nil interface alors qu'aucun type n'est déclaré dans l'interface alors cela provoque une run-time error.

```go
type I interface {
	M()
}

func main() {
	var i I
	i.M() // Provoque une run-time error
}
```

<!-- A voir si je garde -->

Une valeur de type interface peut contenir n'importe quelle valeur implémentant ces méthodes.

```go
type Abser interface {
    Abs() float64
}

type MyFloat float64

func (m MyFloat) Abs()float64 {
    if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs()float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
    var a Abser
    v := Vertex{1, 3}
    m := MyFloat(1.3)
    a = &v
    fmt.Println(a.Abs()) // a *Vertex implements Abser

    a = v
    fmt.Println(a.Abs()) // Throws an error because Vertex doesn't implement Abser

    a = m
    fmt.Println(a.Abs()) // a MyFloat implements Abser
}
```

Autre manière de déclarer une interface

```go
type I interface{
    M()
}

type T struct {
    S string
}

func (t T) M() { // Cette méthode signifie que T implémente l'interface I
    fmt.Println(t.S)
}

func main() {
    var i I = T{"Hello"}
    i.M()
}
```

# Interface values

Les valeurs d'interface peuvent être vues comme une combinaison d'une valeur avec un type.
Lorsqu'une méthode est appelée sur une valeur d'interface, la méthode appelée est celle portant le même nom et le même type.

```go
type T struct {
    S string
}

func (t *T) M(){
    fmt.Println(t.S)
}

type F float64

func (f F) M(){
    fmt.Println(f)
}

type I interface{
    M()
}

func main() {
    var i I
    i = &T{"Hello"} // Car M est implémenté sur le pointeur. i est de type *main.T
    i.M()

    i = F(3.1) // main.F
    i.M()
}
```

# Nil

```go
type I interface{
    M()
}

type T struct{
    S string
}

func (t *T) M() {
    if(t == nil){
        fmt.Println("nil")
        return
    }
    fmt.Println(t.S)
}


func main(){
    var i I
    var t *T
    i = t
    i.M() // La méthode est appelée avec nil
}
```

# Stringers

Il est possible de reformater la relation entre certaines fonctions et certains types.
Par exemple, ici, le string renvoyé par les variables a et z pour être utilisé dans fmt.Println et reformaté à l'aide de fmt.Sprintf("%v (%v years)", p.Name, p.Age).

```go
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```
