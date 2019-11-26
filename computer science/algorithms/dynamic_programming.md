# Source
https://www.geeksforgeeks.org/overlapping-subproblems-property-in-dynamic-programming-dp-1/
# overlapping subproblems
Un problème dispose d'overlapping subproblems si pour trouver sa solution il faut résoudre les mêmes sous-problèmes plusieurs fois.
Par exemple dans le cas de la séquence de fibonnaci, on peut utiliser cette simple récursion :
```go
func fib(n){
    if (n <= 1) {
        return n
    }
    return fib(n-1) + fib(n-2)
}
```
fib(n-1) et fib(n-2) correspondent aux sous-problèmes.
Si on essaie de résoudre fib(5) alors la fonction fib(2) va être appelée 3 fois.
fib(5) = fib(4) + fib(3) = (fib(3) + fib(2)) + (fib(2) + fib(1)) = ((fib(2) + fib(1)) + fib(2)) + (fib(2) + fib(1))

# Memoization
Cette technique permet de ne pas avoir à refaire 2 fois le même calcul pour les mêmes inputs en gardant la trace de tous les résultats pour tous les inputs.

Elle est utile dans le cas où les inputs sont de plus en plus petit.

Pour mettre en place la memoization, il faut créer un wrapper autour de la fonction disposant d'une hash table contenant les inputs et les outputs.
Ensuite, pour chaque calcul, il faut vérifier si la hash table contient la réponse.
Si la hash table ne contient pas le résultat alors il faut le calculer et le stocker.

```go
func memoization(fn func(n int, cache map[int]int) int) func(n int) int {
	memo := make(map[int]int)
	return func(n int) int {
		return fib(n, memo)
	}
}

func fib(n int, cache map[int]int) int {
	if result, ok := cache[n]; ok {
		return result
	}
	if n <= 1 {
		return 1
	}
	result := fib(n-1, cache) + fib(n-2, cache)
	cache[n] = result
	return result
}
```


# Bottom-up
Une autre stratégie possible est d'utiliser le bottom-up.
Au lieu de démarrer par la fin comme le fait notre fonction de manière recursive, une bottom-up algorithm démarre par le début.
En évitant d'utiliser la récursivité, on économise de l'espace mémoire qui aurait été utilisé pour stocker la callstack (O(n)).

```go
func sumToN(n int) int {
    result := 1

    for i := 1; i < n; i++ {
        result += i
    }

	return result
}
```
Cette fonction utilise O(1) space / O(n) time