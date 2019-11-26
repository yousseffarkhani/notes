# Nombres négatifs
Les nombres négatifs sont souvent représentés en utilisant le two's complement.
Dans ce système d'encodage le chiffre tout à gauche correspond à la valeur négative et le reste est positif :
 1 + 0 + 1
-4 + 0 + 1 = -3
Si la variable est unsigned alors elle utilisera le système classique et le nombre ne pourra qu'être positif.
Si la variable est signed alors elle utilise l'encodage two's complement.
# Shift
Un bit shift correspond au déplacement des bits vers la gauche ou la droite.
## Left shifts
0010 << 1 -> 0100
0010 << 2 -> 1000
Un simple left shift permet de multiplier par 2
0101 = 5
0101 << 1 -> 1010 = 10
## Logical Right shifts
0010 >> 1 -> 0001
0010 >> 2 -> 0000
Un simple right shift permet de diviser par 2 en supprimant le reste
0101 = 5
0101 >> 1 -> 0010 = 2
## Arithmetic Right shifts
Le most-significant bit est copié et sa valeur est ajouté.
0010 >> 2 -> 0000
1010 >> 2 -> 1110
Si le nombre est encodé avec le two's complement alors ce type de shift permet de conserver le signe du chiffre.
# Bitwise
## AND
Consiste à comparer bit par bit et appliquer AND
5 = 0 1 0 1
6 = 0 1 1 0
& = 0 1 0 0 = 4
## OR
Consiste à comparer bit par bit et appliquer OR
5 = 0 1 0 1
6 = 0 1 1 0
| = 0 1 1 1 = 7
## XOR
Consiste à comparer bit par bit et appliquer XOR
5 = 0 1 0 1
6 = 0 1 1 0
^ = 0 0 1 1 = 3
## NOT
Consiste à comparer bit par bit et appliquer NOT
5 = 0 1 0 1
~ = 1 0 1 0 = -6 (Utilisation du two's complement)