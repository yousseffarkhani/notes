voir : https://www.ardanlabs.com/blog/2013/10/manage-dependencies-with-godep.html

# Introduction

Un package manager permettra de consigner et gérer les dépendances du projet vis à vis de librairies externes.

Il en existe plusieurs, les plus connus étant dep, ...

# Processus

1. Initialiser un projet
   dep init
2. Ajouter les libraires
   dep ensure -add github.com/gorilla/mux
