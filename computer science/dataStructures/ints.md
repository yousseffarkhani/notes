# Integers
1 bytes peut contenir une valeur jusqu'à 256. Cela est assez limitant c'est pourquoi en général on utilise 4 à 8 bytes (32 bits / 64 bites) pour stocker les integers.
32 bit = 2^32 valeurs = + de 4 Milliards
32 bit = 2^64 valeurs = + de 10 Milliards de Milliards

En général un integer est de taille fixée (souvent 64 bits). C'est pourquoi que le nombre soit 1 ou 8000909 ne change rien en mémoire.
Cette propriété est intéressante dans le sens où même en réalisant des opérations (addition, soustraction, ...), la mémoire utilisée sera constante (O(1)).
Les integers à taille fixée sont donc très efficient.
Cependant ils sont limités quand à leur valeur à 2^n où n représente le nombre de bits.
