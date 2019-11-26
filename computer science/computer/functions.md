# In-Place algorithm
Une fonction in-place modifie la structure de données ou d'objets en dehors de sa propre callstack.
Graĉe à cela, les changements réalisés dans la fonction restent effectifs après l'appel de la fonction.
In-place signifie que l'argument passé à la fonction est modifié directement.

A contrario une fonction out-of-place n'agit que sur ses variables locales et aucun effet n'est visible de l'extérieur.

Dans la majorité des langages, les valeurs primitives (int, float, string) sont copiées lorsqu'elles sont passées en arguments alors que les structures de donnéescomplexes (lists, heaps, hash tables) sont passées par référence.

Les fonctions in-place permettent de gagner de la complexité temps et mémoire (pas besoin d'intialiser une variable et de copier la valeur) mais ont le désavantage de créer un code très imbriqué avec des effets de bord ayant pour conséquence de modifier voir détruire des données.

En règle générale, il vaut mieux partir sur des fonctions out-of-place pour éviter les effets de bord.