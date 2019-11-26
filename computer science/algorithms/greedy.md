# Source
https://www.interviewcake.com/question/python/stock-price
https://www.geeksforgeeks.org/activity-selection-problem-greedy-algo-1/
# Description
Ce paradigme permet de trouver la meilleure solution lorsque l'on tente différentes combinaisons.
Cet algorithme consiste à construire une solution en choisissant la meilleur option à chaque passage et en la sauvegardant pour la comparer avec le passage suivant.
Elle est souvent utilisée dans dans les problèmes d'optimisation.

Par exemple, un caissir veut donner 67 cts € à un client.
Il va procéder de la manière suivante :
- 50 cts
- 10 cts
- 5 cts
- 2 cts
# Complexité
- Time : O(n)
- Space : O(1)
# Avantages
- Rapide
# Inconvénients
# Utilisation
- Placer un maximum de meeting dans une salle de réunion.
A chaque passage, organiser le meeting finissant le plus tôt.
# Méthodologie
- Stocker la meilleure option.
- Définir les autres variables à garder pour pouvoir définir cette meilleure option.
# Code
```go
func get_max_profit(stock_prices []int) int {
	if len(stock_prices) < 2 {
		return -100
	}

	min_price := stock_prices[0]
	max_profit := stock_prices[1] - stock_prices[0]

	for i := 1; i < len(stock_prices); i++ {
		current_price := stock_prices[i]
		potential_profit := current_price - min_price
		max_profit = int(math.Max(float64(potential_profit), float64(max_profit)))
		min_price = int(math.Min(float64(min_price), float64(current_price)))
	}

	return max_profit
}
```