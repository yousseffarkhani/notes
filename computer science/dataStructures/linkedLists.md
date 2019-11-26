# Linked lists
Notre word processor aura besoin d'append des éléments.
Le problème avec les dynamic arrays est que :
- pour certains cas le fait d'append, nous obligera à recréer un nouvel array
- pour preprend, il faut décaler l'ensemble des éléments de l'array vers le bas. Rajouter le nouvel élément.

L'interêt des linked lists est qu'il est aisé d'append/prepend des éléments.
Une linked list consiste à décomposer au maximum notre array et pour chaque élément créer un array de taille 2 (node) avec l'élément à stocker et l'adresse du prochain élément.
Cela permet également d'optimiser l'espace mémoire puisque tout est utilisé et que les éléments n'ont pas besoin d'être dans l'ordre.
Le 1er node est appelé head et le dernier tail.

Ainsi pour append un élément, il suffira de prendre le dernier élément et de modifier l'adresse du prochain élément et également de modifier le tail node.

# Doubly linked lists
Les doubly linked lists correspondent à des linked lists pour lesquelles chaque node retient la valeur suivante et précédente.
Cela nous permet de traverser la liste dans le sens inverse.

## Trade offs
Positifs :
- Rapide pour append/Preprend/Insert/delete O(1)
- Peut aller n'importe où en mémoire
- Taille flexible (on peut rajouter des éléments à tout moment)
Négatifs :
- Long pour rechercher un élément O(n) (Il faut parcourir toute la liste pour trouver l'élément N)
- La recherche est encore plus longue que dans un array car les systèmes de cache lisent plus rapidement les éléments organisés de manière séquentielle que des éléments dispersés.