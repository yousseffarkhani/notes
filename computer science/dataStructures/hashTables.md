# Hash tables / hash map
La recherche rapide est souvent importante. C'est pour cela qu'on a tendance à utiliser des arrays plutôt que des linked lists.

Un array peut être vu comme un tableau à 2 colonnes, 1 pour les index et 1 pour les valeurs. Cependant, il est impossible de modifier les index.

Imaginons que nous souhaitions stocker un mot dans un index, une solution serait de :
- Prendre la valeur numérique de chaque lettre
 l     i     e     s
108 + 105 + 101 + 115 = 429
- Calculer le module en fonction de la taille de notre array.
429 % 30 = 9
- l'indice 9 représentera ici le mot "lies" et on peut donc ajouter une valeur à l'espace de cet indice dans notre array.

Cette structure s'appelle une hash table ou hash map.
La clé représente le mot (ici "lies") et la valeur peut représenter autre chose (le nombre de fois o`ce mot apparait).

Le calcul utilisé pour transformer le mot en index est appelé hashing function.

La recherche est rapide dans un sens mais pas dans l'autre : Si on recherche une valeur pour un mot donné cela sera rapide. L'inverse n'est pas vrai (il faudra parcourir l'ensemble des valeurs pour trouver les clés associées).
C'est la même chose pour les array. Si on connait l'indice à l'avance la recherche est rapide. S'il s'agit de trouver l'indice d'une valeur alors le programme devra parcourir tout le tableau.
## Trade offs
Positifs :
- Recherche rapide (O(1))
- On peut rechercher via une key arbitraire (mot, ...) au lieu des indices pour les arrays
- Rapide pour l'insertion et la suppression
Négatifs :
- Hash collisions