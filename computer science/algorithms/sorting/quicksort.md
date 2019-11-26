# Source
https://www.geeksforgeeks.org/quick-sort/
https://www.youtube.com/watch?v=8hHWpuAPBHo
https://www.youtube.com/watch?v=ywWBy6J5gz8
# Utilisation
Merge Sort est décomposé en 2 fonctions :
- La 1ère permet de découper en 2 un input
- La 2ème permet de fusionner les arrays en les replaçant en ordre.
# Avantages
- Ne requiert pas d'espace de stockage supplémentaire (in-place)
- Plus intéressant que MergeSort sur les arrays car cet algorithme n'a pas besoin des O(n) stockage supplémentaire.
# Inconvénients
- Quicksort n'est pas stable par défaut
# Complexité
Time : n*log(n)
# Fonctionnement
La fonction merge accepte 2 arrays déjà triés au préalable. C'est pourquoi on éclate au maximum l'input afin de n'avoir qu'un élément dans chaque tableau ce qui vérifie la condition 2 arrays triés.
Ensuite on remonte au fur et à mesure puisque nous sommes sûr que les arrays en entrée de merge seront triés.
# Code
```go
func mergeSort(input []int) []int {
    if len(input) == 1 {
        return input
    }

	middleIndex := len(input) / 2

	return merge(mergeSort(input[:middleIndex]), mergeSort(input[middleIndex:]))
}

func merge(left, right []int) []int {
	i := 0
	var result []int
	for 0 < len(left) && 0 < len(right) {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
		i++
	}

	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
```