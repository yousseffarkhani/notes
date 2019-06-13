# Introduction

Tous les programmes Go sont composés de packages.

Les programmes démarrent dans le package main.

# Variables

## Déclarer une variable

```go
var c bool

var i int

var i, j = 1, 2
```

Dans les fonctions, il est possible d'utiliser := pour assigner une valeur à une variable sans déclarer le type.

```go
k := 3
```

La zero value (0 pour les numériques, false pour les boolean et "" pour les strings) est attribuée aux variables déclarées mais sans valeur initiale

## Types de variables

```go
bool, string, int, float32
```

## Convertir des variables

Pour convertir une valeur v au type T, il faut utiliser T(v)

```go
i:= 42
f := float64(i)
```

## Constantes

```go
const Pi = 3.14
```

Les constantes ne peuvent pas être déclarées avec :=

# Boucles / conditions

## For

```go
for i := 0; i < 0; i++ {}
```

## While

Les init et post statements sont optionnels :

```go
sum := 3
for ; sum < 100; {
sum = sum + 5
}
```

Ce type de for représente le while en Go est peut-être écrit :

```go
for sum < 100 {}
```

## Boucle infini

```go
for {}
```

## If

### Syntaxe

```go
if x < 0 {}
```

### If avec un statement

Comme dans la boucle for, le if statement peut réaliser une opération avant la lecture de la condition. Les variables ainsi déclarées sont dans le scope du if.

```go
if v := 20*x; v < 40 {}
```

Les variables déclarées dans les if sont également disponibles dans les else.

## Switch

Le switch est différent dans le Go :

- Il n'y a pas besoin de déclarer de break cela est fait automatiquement

# Functions

## Déclarer une fonction

```go
func add(x int, y int) int {
return x + y
}
```

Si les 2 paramètres sont de même types alors on peut écrire :

```go
func add(x, y int) int {
return x + y
}
```

Il est également possible de retourner plusieurs résultats :

```go
func swap(x, y string)(string, string){
return y, x
}
```

## Naked return

Consiste à créer des variables dès le début de la fonction (ici x et y) qui seront retournées automatiquement s'il n'y a rien derrière return.

```go
func split(sum int) (x, y int){
x = 3*sum
y = sum -x
return
}
```

# Defer

Un defer statement repousse l'exécution d'une fonction à plus tard (jusqu'à ce que les fonctions autour retournent une valeur)
Exemple :

```go
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
```

S'il y a plusieurs defer pushed onto a stack, elles sont lancées en last-in-first-out.

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
