# Arrays
Un array correspond à des étagères côte à côte réservées pour cet élément en mémoire.
Ces étagères sont numérotées, cela correspond à leur index.
Les arrays nécessitent que :
- chaque objet de l'array soit de la même taille
- l'array est ininterrompu en mémoire

Les arrays ont donc des fast lookups (O(1)) mais chaque section de l'array doit être de la même taille et il faut disposer d'un espace ininterrompu en mémoire.
## Trade offs
Positifs :
- Look ups rapide O(1)
- Append rapide
Négatifs :
- Non flexible (la taille doit être connu en avance)
- Requiert un espace réservé en mémoire pouvant être conséquent
- Insert et deletes lent (O(N))
- Les objets de l'array doivent être de la même taille
- Recherche lente (O(n)) / Binary Search permet d'arriver à O(log(n))

# Dynamic arrays
Imaginons que nous souhaitions développer un word processor.
Il est impossible d'utiliser un array pour stocker les strings car il s'agirait de devoir dire en avance quelle serait la taille de notre array (pour pouvoir réserver l'espace en mémoire).

C'est pourquoi un autre type d'array a été crée le dynamic array.
Lorsque nous créons un dynamic array, ce dernier crée en réalité un array classique.
La taille correspond à celle utilisée à l'initialisation.
Exemple:

1) Création d'un array de taille 10
__________
2) 4 lettres sont ajoutées
Hello_____
A ce stade notre dynamic array a une taille de 4 et une capacité de 10 (correspond à la taille de l'array réservé)
Le dynamic array garde en mémoire l'index de fin.
3) Si on rajoute des lettres en plus, à un moment la capacité max sera atteinte et il va se passer plusieurs étapes :
- Créer un plus grand array (généralement 2 fois plus grand)
Pourquoi ne pas étendre notre array ? Car l'espace mémoire à la suite de ce dernier peut être occupé par autre chose.
- Copier chaque élément de l'ancien array vers le nouveau
- Libérer l'espace mémoire de l'ancien array
- Ajouter la nouvelle lettre

L'avantage des dynamic arrays est que nous n'avons pas besoin de spécifier la taille en avance. Cependant l'extension de notre dynamic array peut prendre du temps.
## Trade offs
Positifs :
- Recherche rapide
- Taille variable
Négatifs :
- Long à append/preprend (Peut même être pire si on décide d'append et qu'il faut agrandir l'array)