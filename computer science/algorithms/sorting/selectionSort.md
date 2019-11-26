# Source
https://www.geeksforgeeks.org/selection-sort/
# Utilisation
Cet algorithme permet de mettre en ordre une liste (nombre, strings, ...)
# Avantages
- In-place
# Inconvénients
# Complexité
Time : O(n²)
Space : O(1)
# Fonctionnement
Consiste à prendre le plus petit élément d'un tableau et à le mettre devant ainsi de suite.
Il y a 2 subArrays : 1 sorted et 1 unsorted

# Code
```go
func selectionSort(input []int) []int {
	minIndexValue := 0
	for sortedIndex := 0; sortedIndex < len(input); sortedIndex++ {
		minIndexValue = sortedIndex
		for i := sortedIndex + 1; i < len(input); i++ {
			if input[minIndexValue] > input[i] {
				minIndexValue = i
			}
		}
		input[sortedIndex], input[minIndexValue] = input[minIndexValue], input[sortedIndex]
	}

	return input
}
```