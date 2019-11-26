# Heaps
Un heap est en réalité un graphe.
On peut avoir soit un min heap soit un max heap.
Le min heap consiste à avoir :
- l'élément le plus petit en root node
- Les enfants de chaque node deviennent de plus en plus grand
Dans la suite, nous allons prendre le cas d'un min heap.
# Représentation
On peut représenter cette structure sous forme d'array.

# Opérations
- Trouver le minimum/maximum : 1er élément de notre array
- Insérer un élément : On insère la valeur au niveau du 1er endroit disponible de notre graph. S'il s'agit :
    - d'une valeur plus grande que le parent alors c'est OK
    - d'une valeur plus petite que le parent alors il faut faire un bubble sort et swap les valeurs des nodes jusqu'à ce que le nouvel élément se trouve à la bonne place
- Supprimer un élément :
    - S'il s'agit d'une leaf alors il suffit de la supprimer
    - S'il s'agit du root node alors on swap le root node avec le dernier élément.
        - On supprime le dernier élément (correspond à l'ancien root node)
        - Si le nouveau root node est plus petit que ses enfants alors OK
        - Si le nouveau root node est plus grand que l'un de ses enfants alors on le swap avec la plus petite valeur.

# Complexité
- Trouver le minimum : O(1)
- Insertion : O(log(N))
- SUppression : O(log(N))