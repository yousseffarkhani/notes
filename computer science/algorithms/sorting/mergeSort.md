# Source
https://www.geeksforgeeks.org/merge-sort/
# Utilisation
Merge Sort est décomposé en 2 fonctions :
- La 1ère permet de découper en 2 un input
- La 2ème permet de fusionner les arrays en les replaçant en ordre.
# Avantages
- Est préféré à quickSort pour les linked Lists
# Inconvénients
# Complexité
Time : n*log(n)
Space : O(n)
# Fonctionnement
## Out-of-place
La fonction merge accepte 2 arrays déjà triés au préalable. C'est pourquoi on éclate au maximum l'input afin de n'avoir qu'un élément dans chaque tableau ce qui vérifie la condition 2 arrays triés.
Ensuite on remonte au fur et à mesure puisque nous sommes sûr que les arrays en entrée de merge seront triés.
### Code
```go
func quickSort(input []int) []int {
	if len(input) == 1 {
		return input
	}
	pivot := input[len(input)-1]
	input = input[:len(input)-1]
	var left, right []int
	for _, val := range input {
		if val > pivot {
			right = append(right, val)
		} else {
			left = append(left, val)
		}
	}
	var result []int
	if len(left) > 0 {
		result = append(result, quickSort(left)...)
	}
	result = append(result, pivot)
	if len(right) > 0 {
		result = append(result, quickSort(right)...)
	}
	return result
}
```
## In-place

### Code
```go

```