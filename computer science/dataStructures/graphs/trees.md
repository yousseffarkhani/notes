# Binary tree
Un tree est un graphe spécifique disposant de propriétés particulières :
- Il dispose d'un unique Root node
- Les Edges sont directed (root -> vers le bas) -> implique qu'il n'y a pas de cycles (on appelle ce type de graphe DAG - Directed Acyclical Graphs)
- Il est équilibré : Le nombre d'enfants à gauche et droite de chaque noeud sont égaux (ou presque)
- Il est trié -> Cela rend la recherche d'en élément plus rapide (Log(N))

# Générer un arbre à partir d'un fichier
Fichier d'entrée (dfs.in):
```bash
5 # Nombre de nodes
1 -1 -1 # index 0
# 1 : Valeur, -1 (pas de node): index du node à gauche, index du node à droite : -1
3 -1 -1 # index 1
2 0 1 # index 2
6 -1 -1 # index 3
4 2 3 # index 4
```
```go
type Node struct {
	Value int
	left  *Node // l'interêt des pointers est que pour des valeurs vides, cette valeur sera égale à nil
	right *Node
}

func read() []Node {
	var N int
	fmt.Scanf("%d", &N)
	var nodes []Node = make([]Node, N)

	for i := 0; i < N; i++ {
		var value, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &value, &indexLeft, &indexRight)
		nodes[i].Value = value
		if indexLeft >= 0 {
			nodes[i].left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[i].right = &nodes[indexRight]
		}
	}
	return nodes
}
```

Pour lancer ce programme : `go run main.go < dfs.in`