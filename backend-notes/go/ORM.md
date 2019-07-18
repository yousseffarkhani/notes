Source : https://tutorialedge.net/golang/golang-orm-tutorial/

# Introduction

Les ORM permettent de cacher la complexité des commandes SQL.

# Inconvénients / Avantages

Inconvénients :

- Les développeurs ne savent pas vraiment lors des appels vers la BDD car toute cette couche est abstraite.
  Dans une petite application il n'y a pas trop de danger.
  Cependant pour une grosse application, il faut faire attention aux queries qu'on demande à la BDD.
  La motivation principale d'écrire en pur SQL les requêtes est que les développeurs doivent connaitre la complexité de leur queries et donc de savoir quand créer des index (en cas de join ou de filter).  
  Avantages :
- Permet au débutant de ne pas avoir à apprendre un nouveau language (SQL).
- Permet de ne pas avoir à se prendre la tête avec les requêtes SQL et se focus sur la business logic.

# Comment utiliser les ORM ?

Il faut créer un package séparé contenant l'ensemble des fonctions qui vont faire appel à la BDD.
