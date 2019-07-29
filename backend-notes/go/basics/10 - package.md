# Introduction

Quand un programme Go atteint un import statement, il va tout d'abord chercher dans la librairie standard (GOROOT/go/src ou usr/local/go/src). Si le package n'est pas présent alors go va se référer au GOPATH (/home/ysf/go/src).

# Créer un package

Le nom que le package portera correspond au nom du directory dans GOPATH/src.
Par exemple, le package s'appelera logger si le code est situé dans \$GOPATH/src/github.com/yourname/logger.
Toues les fichiers du package doivent avoir le même package declaration.

# Types de packages

Il existe 2 types de packages :

- Utility package : n'est pas self executable mais améliore les fonctionnalités d'un package exécutable en permettant d'exploiter ses propres fonctions.
- Executable package : Ces packages déclarent en en-tête package main

# Exporter des variables

Pour cela, il faut mettre en majuscule la 1ère lettre de la variable.

# Créer un utility package

Il faut mettre en en-tête `package nom_du_package`

# Go install

Cette commande va voir si un fichier à l'intérieur du package directory contient un main declaration. Si un fichier est trouvé alors Go sait que c'est un exécutable et va créer un fichier binaire.
Si un package ne contient pas de main alors Go va créer un package archive (.a) dans le pkg directory.

# Go get

`go get -u github.com/jinzhu/gorm` permet d'importer les fichiers de l'URL et de les sauvegarder au sein de src/github.com/jinzhu/gorm.
Pour l'importer dans un projet, il faut utiliser la syntaxe : `import "github.com/jinzhu/gorm"`.
Si le package est un executable alors les gens peuvent l'utiliser depuis la CLI autrement, ils peuvent l'importer dans un projet.
