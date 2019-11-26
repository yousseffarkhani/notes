# Introduction
(Nous utilisons un log base 10)

> 100 = 10^x

On multiplie donc par le log des 2 côtés ce qui nous donne :

> log(100) = log(10^2) = log(10^x)

Le logarithme permet de répondre à la question : "A quelle puissance doit on élever la base pour arriver au résultat ?" => A quelle puissance doit on élever la base 10 pour avoir 10^x ? Le résultat est x.

> x = 2

**Le logarithme nous permet de trouver x lorsque x est un exposant.**

# Règles générales du logarithme
- Simplification : (log de base b)
log(b^x) = x
-Multiplication :
log(b^x*c^x) = log(b^x) + log(c^x)
- Division :
log(b^x/c^x) = log(b^x) - log(c^x)
- Puissance
log(x^y) = y*log(x)
- Changement de base :
logb​(x)=logc​(x)/logc​(b)


# Binary search
Pour calculer le temps nécessaire pour un binary search (voir /algorithms/examples/binarySearch.md), il faut résoudre l'équation :
n * (1/2)* (1/2) ... * (1/2) = 1
n*(1/2)^x = 1
n/2^x = 1
n = 2^x
==> log2(n) = x

# Binary tree
Trouver la taille d'un binary tree parfait (Tous les enfants ont un élément gauche et droite, sauf les leafs).
On considère n le nombre total de noeuds, h la hauteur de l'arbre (le nombre de niveaux).

Cela revient à poser la question combien de fois on double 1 pour arriver au nombre de noeuds de la dernière ligne (n+1)/2

h ~ log2((n+1)/2)
Prenons le cas d'un tree parfait à 2 étages.
Ici n = 2, le résultat devrait être h = 2
h ~ log2(2) = log2(2^1) = 1
Or le résultat est h = 2.
Pour corriger ce problème on rajoute donc 1, la formule devient donc
h = log2((n+1)/2) + 1
h = log2(n+1) - log2(2) + 1
h = log2(n+1) - 1 + 1
h = log2(n+1)
