# Source
https://www.geeksforgeeks.org/stability-in-sorting-algorithms/
# Stability in sorting algorithms
La notion de stabilité est importante lorsque nous avons plusieurs paires clé-valeur possédant la même clé (comme des prénoms en clé et des détails en valeur).

Une algorithme de tri est dite stable lorsque 2 objets dsposant de la même clé garde le même ordre que dans le tableau d'origine.

La stabilité n'est pas importante dans les cas où :
- les valeurs sont non distinctes (integers)
- Il n'y a pas de doublon de clé

Prenons le cas d'un tableau d'élève avec leurs notes :
Alice : B
Carole : A
Dave : A
Eric : B
Ken : A

Si on tri les notes avec un algorithme instable, cela (peut) nous donner :
Carole : A
Dave : A
Ken : A
Eric : B <-
Alice : B <-

On peut voir au niveau des 2 derniers éléments que nous avons perdu l'ordre du tableau d'origine. Cela est du à l'algorithme instable.

# Quels algorithmes sont stables ?
Bubble Sort, Insertion Sort, Merge Sort, Count Sort.
Tous les algorithmes basés sur des comparaisons car ces derniers vérifient que un élément A est STRICTEMENT supérieur/inférieur à un élément B pour agir sur le tableau.

# Quels algorithmes sont instables par nature ?
Quick sort, Heap sort
Ces derniers peuvent devenir stable en prenant également en compte la position de l'élément en considération (Cela prendra un peu plus de mémoire).