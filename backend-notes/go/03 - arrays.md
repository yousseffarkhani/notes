# Arrays

## Introduction

Les array en Go ne peuvent pas être resizé.

## Déclarer un array

var a [10]int // Les valeurs sont à null
primes := [6]int{2, 3, 5, 7, 11, 13}

# Slices

Les slices sont des sortes d'array mais plus flexibles.
Un slice est crée à partir d'un Array en spécifiant 2 indices, une limite haute et une limite basse. Cela permet de sélectionner le 1er élément jusqu'à la limite haute (exclue).

primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4]

A noter qu'un slice ne contient pas de données mais uniquement une référence vers l'array.
Changer un élément du slice modifiera également l'array parent.
