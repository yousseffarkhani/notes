# Utilisation
Cette structure de données est utilisée dans le domaine des graph theory.
Un algorithme utilisant cette structure est le DFS (depth first search).

Une stack peut être comparée à une pile de brique.
Quand on souhaite rajouter une brique, on la place au sommet des briques déjà empilées.
Quand on souhaite retirer une brique, on prend la dernière à avoir été ajoutée.

Une stack est similaire à une queue, la seule différence vient de la terminologie et de l'ordre de sortie des éléments.
# Différence avec les queues
Queues = FIFO
Stacks = LIFO
# Avantages
# Inconvénients
# Complexité
# Pré-requis
# Fonctionnement
# Terminologie
Push : Ajouter un objet au sommet de la stack
Pop : Retirer un élément du sommet de la stack
Peek : Voir mais ne pas retirer l'élement au sommet de la stack
# Code
## Recursive
```go
type Stack struct {
	slice []int
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}

func (s *Stack) push(num int) {
	s.slice = append(s.slice, num)
}

func (s *Stack) pop() int {
	s.slice = s.slice[:len(s.slice)-1]
	return s.peek()
}

func (s *Stack) peek() (int, error) {
    if len(s.slice) == 0 {
		return 0, errors.New("No items left in the stack")
	}
	return s.slice[len(s.slice)-1], nil
}
```
