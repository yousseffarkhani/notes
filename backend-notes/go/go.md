# Introduction

Tous les programmes Go sont composés de packages.

Les programmes démarrent dans le package main.

# Variables

## Déclarer une variable

var c bool
var i int
var i, j = 1, 2

Dans les fonctions, il est possible d'utiliser := pour assigner une valeur à une variable sans déclarer le type.
k := 3

La zero value (0 pour les numériques, false pour les boolean et "" pour les strings) est attribuée aux variables déclarées mais sans valeur initiale

## Types de variables

bool, string, int, float32

## Convertir des variables

Pour convertir une valeur v au type T, il faut utiliser T(v)
i:= 42
f := float64(i)

## Constantes

const Pi = 3.14
Les constantes ne peuvent pas être déclarées avec :=

# Boucles / conditions

## For

for i := 0; i < 0; i++ {}

## While

Les init et post statements sont optionnels :
sum := 3
for ; sum < 100; {
sum = sum + 5
}

Ce type de for représente le while en Go est peut-être écrit :

for sum < 100 {}

## Boucle infini

for {}

## If

### Syntaxe

if x < 0 {}

### If avec un statement

Comme dans la boucle for, le if statement peut réaliser une opération avant la lecture de la condition. Les variables ainsi déclarées sont dans le scope du if.
if v := 20\*x; v < 40 {}

# Functions

## Déclarer une fonction

func add(x int, y int) int {
return x + y
}

Si les 2 paramètres sont de même types alors on peut écrire :
func add(x, y int) int {
return x + y
}

Il est également possible de retourner plusieurs résultats :
func swap(x, y string)(string, string){
return y, x
}

## Naked return

Consiste à créer des variables dès le début de la fonction (ici x et y) qui seront retournées automatiquement s'il n'y a rien derrière return.

func split(sum int) (x, y int){
x = 3\*sum
y = sum -x
return
}

# Packages

Pour importer un package : import ("nom_du_package")

Listes de packages :

- fmt :
  - fmt.Println("string")
    - %g, %T pour afficher le type, %v pour la valeur
- math/rand :
  - rand.Intn(10)
- math :
  - math.Sqrt(7)
