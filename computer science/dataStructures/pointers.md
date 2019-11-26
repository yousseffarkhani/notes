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