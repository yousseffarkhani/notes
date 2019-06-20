# Introduction

Les fonctions sont des valeurs comme en JS.

```go
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main(){
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }

    fmt.Println(compute(hypot))
}
```

# init functions

Les init functions sont des fonctions lancées automatiquement permettant d'initialiser ou de vérifier des conditions avant le lancement de la func main().
Chaque package peut contenir une init function.
L'ordre de lancement est :
0. (optionnel) Si un package importe un autre package alors les init functions du package importé seront executées en 1er. (Un package sera initialisé une seule fois même s'il est importé plusieurs fois)
1. Initialisation des package level variables
2. Init function

# Variadic function

Une fonction variadique est une fonction pouvant accepter un nombre variable de paramètres.
Seul le dernier paramètre d'une fonction peut être variable.

## Syntaxe 
```go
func hello(a int, b... int){} // Le paramètre b est variadique.

hello(1) // OK
hello(1, 2) // OK
hello(1, 2, 4) // OK
```
## Fonctionnement
La fonction va convertir les arguments en slice du type donné dans la déclaration de la fonction (b... int).
## Pourquoi utiliser un variadic argument au lieu d'un slice ?
Il serait en effet possible de réecrire la fonction de cette manière :
```go
func hello(a int, b []int){} // Le paramètre b est variadique.

hello(1, []int // OK
hello(1, []int{2}) // OK
hello(1, []int{2, 4}) // OK
```

Cependant, la méthode variadic a des avantages :
- Plus facile à écrire et comprendre
- Pas besoin de créer un slice à chaque appel de fonction
- Plus facile de gérer les cas où l'argument est vide

## Passer un slice en argument
Il n'est pas possible de passer directement un slice à une fonction variadic (ne respecte pas la signature de la fonction). 
Pour cela, il faut utiliser ... : 
```go
func hello(a int, b ...int){} // Le paramètre b est variadique.

nums := []int{2, 4}
hello(1, nums...) // OK
```

**Ici le slice est passé en argument (pas de création de nouveau slice), toutes modifications entrainera une modification sur l'array original**

## Complément
Append est une fonction variadic.
### Fonctionnement de append
Si le slice passé en argument est de taille suffisante alors celui-ci est directement modifié sinon un nouveau slice est crée et le slice original n'est pas touché.
https://play.golang.org/p/THiBHCyjCZH
