# Utilisation
Cette structure de données est utilisée dans le domaine des graph theory.
Un algorithme utilisant cette structure est le BFS (Breadth first search). (voir https://gophercises.com/exercises/sitemap)

Les queues peuvent être implémentées de différentes manières (en utilisant des linked lists, slices, foxed size array).
On peut également fixer une capacité maximum à une queue.

Nous allons ici utiliser un slice car c'est plus simple (même s'il ne s'agit pas de l'option la plus optimal).
# Avantages
- Opérations rapides (Enqueue, Dequeue, peek prennent O(1))
# Inconvénients
# Complexité
# Pré-requis
# Fonctionnement
# Terminologie
Enqueue : Ajouter un objet à la fin de la queue
Dequeue : Retirer un élément du début d'une queue
# Code
## Recursive
```go
type Queue struct {
	slice []int
}

func (q *Queue) String() string {
	return fmt.Sprint(q.slice)
}

func (q *Queue) enqueue(num int) {
	q.slice = append(q.slice, num)
}

func (q *Queue) dequeue() int {
	if len(q.slice) == 0 {
		return -1
	}
	topItem := q.slice[0]
	q.slice = q.slice[1:]
	return topItem
}
```
