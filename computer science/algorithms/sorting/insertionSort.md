# Source
https://www.geeksforgeeks.org/insertion-sort/
# Utilisation
A utiliser lorsque le nombre d'éléments est petit ou que la majorité du tableau est déjà triée.
# Avantages

# Inconvénients
# Complexité
Time : O(n²)
Space : O(1)
# Fonctionnement
Prendre le 1er élément du tableau et le marquer en tant que sorted.
Prendre le 2ème élément du tableau, si :
- Il est supérieur au 1er élément alors marquer en tant que sorted
- Il est inférieur au 1er élément alors swap les 2 éléments.
Prendre le 3ème élément et faire une boucle sur la partie sorted pour le placer.
...

# Code
```go
func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
				j = j - 1
				continue
			}
			break
		}
	}
}
```