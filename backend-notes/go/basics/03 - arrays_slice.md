# Arrays

## Introduction

Les arrays en Go ne peuvent pas être resizés.

## Déclarer un array

```go
var a [10]int // Les valeurs sont à 0
/* Shorthand declaration */
primes := [6]int{2, 3, 5, 7, 11, 13}
/* Assignation automatique d'une lenght */
a := [...]int{1, 2, 3}
```

En go, les arrays sont des values types et non des reference types ce qui signifie que quand on attribue un array à une nouvelle variable alors une copie de cet array est crée. Il n'y aura aucun impact sur l'array original si les valeurs sont modifiés dans cette nouvelle variable.

De même lorsque les arrays sont passés dans des fonctions en tant que paramètres l'array original n'est pas modifié

# Slices

Les slices sont des sortes d'array mais plus flexibles.
Un slice est crée à partir d'un Array en spécifiant 2 indices, une limite haute et une limite basse. Cela permet de sélectionner le 1er élément jusqu'à la limite haute (exclue).
En réalité, un slice est une référence ) un array.

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4] // ou s := primes[1:4]
```

A noter qu'un slice ne contient pas de données mais uniquement une référence vers l'array.
/!\ Changer un élément du slice modifiera également l'array parent.

## Créer un slice

```go
var s := []bool{true, true, false} // Créer un array et retourne la référence
```

### Déclarer un array de struct directement dans la définition d'une variable

```go
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
```

## Slices defaults

Les indices de limite haute et de limite basse peuventêtre omis pour utiliser les defaults à la place.
Les defaults sont 0 pour la limite basse et la longueur du slice pour la limite haute.

```go
	s := []int{2, 3, 5, 7, 11, 13}
    s = s[1:] //Affiche [3 5 7 11 13]
    s = s[:2] //Affiche [3 5]
    s = s[:] //Affiche [3 5]
```

## Propriétés d'un slice

Un slice dispose d'une length et d'une capacity.
La length est le nombre d'élément qu'il contient.
La capacité correspond aux nombres d'éléménts de l'array slicé en comptant à partir du 1ère élément du slice.

Ces propriétés sont accessibles via len(s) et cap(s).
La capacité est intéressante car elle permet de reslicé un slice et ainsi utiliser d'autres valeurs de l'array d'origine.

```go
	s := []int{2, 3, 5, 7, 11, 13}
    s = s[:0] // Slice the slice : len=0, cap=6 []
    s = s[:4] // len=4, cap=6 [2 3 5 7]
    s = s[2:] // len=2 cap=4 [5 7]
```

## Nil slices

```go
	var s []int
```

La valeur par défaut d'un slice est `nil`. Sa length et capacity valent 0.

## Créer un slice avec make

La fonction make alloue un array rempli de zéro et retourne le slice pointant vers l'array.

```go
	a := make([]int, 5) // len(a) = 5, cap(a) = 5 [0 0 0 0 0]
```

Pour spécifier une capacité rajouter un 3ème argument :

```go
    a := make([]int, 0, 5) // len(a) = 0, cap(a) = 5 []
    a = a[:cap(a)] // len(a) = 5, cap(a) = 5 [0 0 0 0 0]
    a = a[1:] // len(a) = 4, cap(a) = 4 [0 0 0 0]
```

## Slices of slices

Les slices peuvent contenir n'importe quel type, y compris d'autre slices.

```go
board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
```

## Appending to a slice

Pour ajouter des éléments à un slice, il existe une fonction append `func append(s []T, vs ...T)[]T`.
Le 1er paramètre de append est le slice de type T et le reste correpond aux valeurs à append.
Si l'array pointé par s est trop petit pour contenir les nouvelles valeurs alors un array plus grand est alloué. L'augmentation de capacité n'est pas linéaire (0 > 2 > 4 > 8 > 12).

```go
var s []int
	b := []int{
	1,
	2,
	3,
    }
    s = append(s, b...) // len=3 cap=3 [1 2 3]
```

# Optimisation de la mémoire

Etant donné que les slices gardent en mémoire une référence vers un array, cela signifie que tant que le slice existe, l'array existera et ne pourra pas être garbage collected.
Exemple : Nous avons besoin d'une petite partie d'un énorme array. Nous créons donc un slice de la partie qui nous interesse. L'énorme array continuera à exister.
La solution à ce problème est d'utiliser la fonction copy permettant de faire une copie du slice.

```go
func countries() []string {
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2] // Créer un slice
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}
func main() {
    countriesNeeded := countries()
    fmt.Println(countriesNeeded)
}
```
