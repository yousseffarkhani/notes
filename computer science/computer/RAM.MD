# Source
https://www.interviewcake.com/article/python/data-structures-coding-interview?course=dsa
# RAM
Quand un ordinateur exécute un programme, il a besoin de garder une trace des différentes variables utilisées par le code.
Les variables sont stockées dans la RAM.

Il faut imaginer la RAM comme une très grande étagère. Chaque étage contient 8 bits (1 ou 0).
8 bit = 1 byte = 1 octet.

# Integers
1 bytes peut contenir une valeur jusqu'à 256. Cela est assez limitant c'est pourquoi en général on utilise 4 à 8 bytes (32 bits / 64 bites) pour stocker les integers.
32 bit = 2^32 valeurs = + de 4 Milliards
32 bit = 2^64 valeurs = + de 10 Milliards de Milliards

En général un integer est de taille fixée (souvent 64 bits). C'est pourquoi que le nombre soit 1 ou 8000909 ne change rien en mémoire.
Cette propriété est intéressante dans le sens où même en réalisant des opérations (addition, soustraction, ...), la mémoire utilisée sera constante (O(1)).
Les integers à taille fixée sont donc très efficient.
Cependant ils sont limités quand à leur valeur à 2^n où n représente le nombre de bits.

# Arrays
Un array correspond à des étagères côte à côte réservées pour cet élément.
Ces étagères sont numérotées, cela correspond à leur index.
Les arrays nécessitent que :
- chaque objet de l'array soit de la même taille
- l'array est ininterrompu en mémoire

Les arrays ont donc des fast lookups (O(1)) mais chaque section de l'array doit être de la même taille et il faut disposer d'un espace ininterrompu en mémoire.
## Trade offs
Positifs :
- Rapide pour la recherche
Négatifs :
- Non flexible (la taille doit être connu en avance)
- Requiert un espace réservé en mémoire pouvant être conséquent
- Les objets de l'array doivent être de la même taille

# Strings
Une série de caractères est appelée string.
Pour stocker un string, nous utilisons un array.
Mais comment stocker des caractères au lieu de nombres ?
Il faudra créer un array de 8 bits ainsi 01000001 est égale à A, 01000010 B, ...

# Pointers
Un array ne peut contenir que des éléments de la même taille.

Imaginons que l'on souhaite stocker des nombs dans un array.
Cela est possible en ayant un caractère qui signifie la fin du nom

John#___Youssef#___Johnyjohny#

_ représente de l'espace mémoire inutilisé (car les arrays sont de taille fixe)

Il existe une alternative qui est d'utiliser les pointers.
Au lieu de stocker directement les noms dans notre array, nous allons stocker des adresses vers ces derniers.

 0  1  2 3     4        5
*3 *4 *5 John# Youssef #Johnyjohny#

*3, *4, *5 correspondent aux adresses des noms.

Ainsi les noms n'ont pas besoin d'être de la même taille.
Nous n'avons pas besoin de beaucoup d'espaces ininterrompu pour générer notre array.
Le seul inconvénient est que lorsque l'on cherche un objet dans la RAM, le memory controller ramène en général également les éléments proches en mémoire dans le cache du processeur.
C'est pourquoi la lecture dans un array est rapide. Or ici, le memory controller va devoir rechercher à différents endroits les valeurs des noms.
## Trade offs
Positifs :
- Plus flexible
Négatifs :
- Recherche plus lente

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
Négatifs :
- Long à append/preprend

# Linked lists
Notre word processor aura besoin d'append des éléments.
Le problème avec les dynamic arrays est que :
- pour certains cas le fait d'append, nous obligera à recréer un nouvel array
- pour preprend, il faut décaler l'ensemble des éléments de l'array vers le bas. Rajouter le nouvel élément.

L'interêt des linked lists est qu'il est aisé d'append/prepend des éléments.
Une linked list consiste à décomposer au maximum notre array est pour chaque élément créer un array de taille 2 (node) avec l'élément à stocker et l'adresse du prochain élément.
Cela permet également d'optimiser l'espace mémoire puisque tout est utilisé et que les éléments n'ont pas besoin d'être dans l'ordre.
Le 1er node est appelé head et le dernier tail.

Ainsi pour append un élément, il suffira de prendre le dernier élément et de modifier l'adresse du prochain élément et également de modifier le tail node.

## Trade offs
Positifs :
- Rapide pour append/Preprend
- Peut aller n'importe où en mémoire
Négatifs :
- Long pour rechercher un élément
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
- Recherche rapide
- On peut rechercher via une key arbitraire (mot, ...) au lieu des indices pour les arrays
Négatifs :
- Hash collisions