Source : https://www.callicoder.com/golang-packages/
https://github.com/golang-standards/project-layout
https://www.calhoun.io/structuring-web-applications-in-go/
# Introduction
L'organisation du code est une question fréquente en Go car contrairement à d'autres languages (ruby, python) certains frameworks ont été massivement adoptés (ruby on rails, django).
Ces frameworks permettant d'éviter de se poser la question de l'organisation du code.


Il existe plusieurs types de layout :
- cmd layout pattern : Ce layout est utile lorsque le projet à besoin de plusieurs binaries
- pkg directories
- MVC

Le type de layout choisi va dépendre de plusieurs facteurs :
Quelle est la taille de l'application ? de l'équipe ? S'agit-il d'un projet Open source ? Est-ce que les tests sont importants ? Si oui, quelle couverture est demandée ?
# Processus

# Fonctionnement

# Commandes

You generally create a new subdirectory inside your repository for every separate Go package.