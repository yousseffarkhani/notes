# Utilisation
Cet algorithme permet de chercher un objet dans une liste triée.

Cet algorithme peut être utilisé pour :
- Chercher la racine carré d'un nombre
- Déterminer si une liste contient un nombre
# Avantages
- Permet de rechercher un objet rapidement dans une liste (Si 1M de données en entrée alors temps d'exécution = 19,93s)
# Inconvénients
# Complexité
O(log(N))
# Pré-requis
La liste en entrée doit être triée
# Fonctionnement
L'algorithme va regarder la valeur au milieu du tableau et la comparer à la valeur recherchée.
Si elle est égale alors OK.
Si la valeur à trouver est inférieure à la valeur du milieu alors on va chercher dans la partie gauche du tableau.
Si la valeur à trouver est supérieure à la valeur du milieu alors on va chercher dans la partie droite du tableau.
# Code
## Recursive
```go
func binarySearch(numList []int, valueToFind, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	middleVal := numList[mid]
	if middleVal == valueToFind {
		return mid
	}
	if middleVal > valueToFind {
		return binarySearch(numList, valueToFind, low, mid-1)
	}
	return binarySearch(numList, valueToFind, mid+1, high)
}
```
## Loop
```go
func binarySearch(list []int, val int) int {
	hi := len(list) - 1
	lo := 0

	for lo <= hi {
		middleIndex := lo + (hi-lo)/2
		middleVal := list[middleIndex]
		fmt.Println(middleIndex)
		fmt.Println(list)
		if middleVal == val {
			return middleIndex
		} else if middleVal > val {
			hi = middleIndex - 1
		} else {
			lo = middleIndex + 1

		}
	}
	return -1
}
```
