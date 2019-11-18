https://www.calhoun.io/lets-learn-algorithms-an-intro-to-bubble-sort/
https://www.hiredintech.com/classrooms/algorithm-design/lesson/78
http://discrete.gr/complexity/


https://www.bigocheatsheet.com/
# Cours
## Vidéos
https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-fall-2011/lecture-videos/
https://www.youtube.com/watch?v=0JUN9aDxVmI&list=PL2SOU6wwxB0uP4rJgf5ayhHWgw7akUWSf&index=2&t=0s
https://www.youtube.com/watch?v=A2bFN3MyNDA&list=PLOtl7M3yp-DX32N0fVIyvn7ipWKNGmwpp
## MOOC
https://www.coursera.org/specializations/algorithms
## Paid MOOC
https://www.udemy.com/course/js-algorithms-and-data-structures-masterclass/
## Exercices
https://www.topcoder.com/community/competitive-programming/tutorials/how-to-find-a-solution/
# Introduction
L'interêt de l'algorithmie est de pouvoir résoudre des problèmes tout en optimisant le temps d'exécution du programme et ainsi économiser de la mémoire et du temps.
En effet, un problème peut être résolu de différentes manières.
Il faudra donc évaluer le temps d'exécution du programme pour fair eun choix entre différentes idées.
La notion de complexité cherche à décrire ce temps par rapport à la taille de l'input et le nombre d'instructions à réaliser.
Une instruction correspond à :
- Attribution d'une valeur à une variable
- Comparaison entre 2 valeurs
- Incrémenter une valeur
- Opération arithmétique (addition, soustraction, ...)
- Rechercher une valeur dans un tableau
Les
Par exemple, ce programme permet de trouver la valeur maximum dans un tableau :
```go
input := []int{3, 5, 8, 1, 20, 15}
max := input[0] // 2 instructions exécutées au début du code : 1 pour rechercher input[0] et l'autre pour l'assigner à max

for i := 0; i < len(input)-1; i++ { // 2 instructions exécutées au début du code : i := 0 + la comparaison i < len(input)-1
    if max < input[i] { // 2 instructions : 1 pour rechercher input[0] + la comparaison
        max = input[i] // 2 instructions : 1 pour rechercher input[0] et l'autre pour l'assigner à max
    }
} // 2 instructions exécutées à chaque boucle : i++ et la comparaison i < len(input)-1
```
Sans compter l'intérieur de la boucle, la fonction décrivant cette complexité s'écrit f(N) = 4 + 2*N.
Estimer la complexité d'une fonction avec une condition if dépend de l'input et est difficile à évaluer. C'est pourquoi on réalise souvent un worst case analysis.
En comptant l'intérieur de la boucle, f(N) = 4 + 6N.

# Comportement asymptotique
Définir le Comportement asymptotique consiste à retirer toutes les constantes et les plus petits facteurs car ces derniers dépenderont du langage et du hardware pour ne se concentrer que sur la valeur de N.
Dans notre exemple, f(N) = N.
> Dès lors, pour déterminer la complexité d'un programme, il suffira de regarder le nombre de boucles imbriquées dans le programme et les récursions éventuelles.
> Si on dispose d'une série de boucle, il faudra choisir la plus lente pour déterminer le comportement asymptotique du programme.

Le comportement asymptotique est noté Θ(N).
Les différents comportements peuvent être :
- Θ(1) : constant => size of input doesn't matter
- Θ(log(n)) : Logarithmic
- Θ((log(n))^c) : polylogarithmic
- Θ(n) : linear => Un array 2 fois plus grand prendra 2 fois plus de temps
- Θ(n²) : Quadratic => Un array 2 fois plus grand prendra 4 fois plus de temps
- Θ(n^c) : polynomial
- Θ(c^n) : exponential

/!\ Le comportement asymptotique est défini par rapport à la longueur des inputs. Dans le cas suivants la complexité est de O(len(a)+len(b)) et non O(len(a)).
```go
func test(a, b ){
	for _, _ range a {
		for _, _ range b {
			fmt.Println("test)
		}
	}
}
```
# Big-O notation

Selection sort
```go
func sort(input []int) []int {
	var sorted []int
	for _ = range input {
		max, index := Max(input)
		sorted = append(sorted, max)
		input = append(input[:index], input[index+1:]...)
	}
	return sorted
}
```
Ici on peut voir que notre fonction dispose d'une boucle dans une boucle. Cependant la boucle intérieur diminue en taille au fur et à mesure : n, n-1, n-2, ..., 1.
La complexité serait égale à n*(n+(n-1)+(n-2)+...+1).
Pour simplifier la mesure, on prend la limite supérieur (worst case scenario), la complexité devient donc O(N)=N².
# Recursive complexity
Calcul du factoriel d'un nombre
```go
func Fact(number int) int {
	fact := 1
	if number == 1 {
		return fact
	}
	return number * Fact(number-1)
}
```
Cette fonction ne contient pas de boucle mais n'a pas ue complexité linéaire pour autant.
O(N) = N
# Logarithmic complexity
L'algorithme de binary search permet de trouver un nombre dans un tableau rapidement.
```go
func binarySearch(input []int, value int) bool {
	middleIndex := int(math.Ceil(float64(len(input) / 2)))
	if input[middleIndex] == value {
		return true
	}
	if len(input) == 1 {
		return false
	}
	if input[middleIndex] < value {
		return search2(input[:middleIndex], value)
	}
	return search2(input[middleIndex:], value)
}
```
L'algorithme consiste à comparer la valeur du milieu avec la valeur recherchée puis couper en 2 le tableau et faire une récursion.
On remarque ici que pour chaque itération, on aura la moitié de la taille précédente :
- Itération 0 : n
- Itération 1 : n/2
- Itération 2 : n/4
- Itération 3 : n/8
- Itération 4 : n/2^4
- Itération i : n/2^i
- Dernière Itération : 1

Il faut alors résoudre l'opération 1 = n/2^i
n = 2^i
log(N) = i
Cela parait logique car pour un tableau avec N=32 éléments, cela revient à faire 32 -> 16 -> 8 -> 4 -> 2 -> 1 soit 5 opérations ce qui correspond au log(32).
La compléxité est donc ici de log(N).
# Mergesort

# Définitions
## Notations
Θ(N) = Complexité de la solution
O(N) = Worst case (tight) de complexité
o(N) = Worst case not tight
Omega(N) = Best case (tight) de complexité
omega(N) = Best case not tight

# Méthodologie
Utilisation du algorithm design canvas
### Constraints
Contient toutes les contraintes liées au problème
Exemple :
- Quelle taille pour l'input array ?
https://www.hiredintech.com/the-common-constraints-handout.pdf
### Ideas
Lister idéalement 2-3 idées pour résoudre le problème.
#### Simplify the tasks
Prendre un exemple simple pour pouvoir le résoudre facilement
#### Try a few examples
#### Think of suitable data structures
#### Think about related problems you know
### Computation Complexities
Chaque idée contient 2 zones de complexité : Le temps d'exécution et l'espace mémoire utilisé.
Pour chaque algorithme, il faut être en capacité à évaluer la complexité de celle-ci.
#### Time complexity
Cette valeur est mesurée en fonction de la taille de l'input.

#### Space
### Code
Après avoir discuter et choisit une idée, il faut la coder.
### Tests
Il faut tester la solution à cette étape.