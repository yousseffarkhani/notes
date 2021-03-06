# Source
https://www.geeksforgeeks.org/bubble-sort/
# Utilisation
Cet algorithme permet de mettre en ordre une liste (nombre, strings, ...)
# Avantages
- Simplicité
- Permet de stopper le code très tôt si la liste est déjà triée.
- Stable
- In-place
# Inconvénients
- Il ne s'agit pas de l'algorithme de triage la plus optimisée
# Complexité
Time : O(n^2)
Space : O(1)
# Fonctionnement
Il s'agit de prendre une paire de nombre et de les comparer.
Si le 2ème élément est plus grand alors les places sont échangées.
Ainsi de suite jusqu'à la fin du tableau.
# Code
```go
func bubbleSort(input []int) []int {
	var swapped bool
	for j := 1; j < len(input); j++ {
		for i := 0; i < len(input)-j; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return input
}
```
```go
type User struct {
	firstname string
	lastname  string
}

func bubbleSort(input []User) {
	for i := 0; i < len(input)-1; i++ {
		if !sweep(input[:len(input)-i]) {
			break
		}
	}
}

func sweep(input []User) bool {
	var didSwap bool
	for i := 0; i < len(input)-1; i++ {
		if i >= len(input)-1 {
			break
		}
		if greater(input[i], input[i+1]) {
			tmp := input[i]
			input[i] = input[i+1]
			input[i+1] = tmp
			didSwap = true
		}
	}
	return didSwap
}

func greater(A, B User) bool {
	if strings.ToLower(A.lastname) == strings.ToLower(B.lastname) {
		return strings.ToLower(A.firstname) == strings.ToLower(B.firstname)
	}
	return strings.ToLower(A.lastname) > strings.ToLower(B.lastname)
}
```