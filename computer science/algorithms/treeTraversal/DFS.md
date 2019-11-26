# DFS (Depth First Search)
Cet algorithme est utilisé sur un binary tree. (voir /dataStructures/graphes/trees)
Cet algorithme permet souvent de répondre à la question : "Est-ce que je peux aller du noeud A vers le noeud B du graphe ?".
Il consiste à aller le plus loin possible dans un graphe dans une direction choisie.
Dès qu'il n'y a plus d'enfants dans la direction choisie alors on regarde dans l'autre direction.
Lorsqu'il n'y a plus d'enfants dans l'autre direction alors on remonte d'un niveau.
Ainsi de suite, jusqu'à trouver l'élément.
Si on revient au root node alors cela signifie que l'élement n'existe pas.