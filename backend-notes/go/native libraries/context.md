Source : https://www.sohamkamani.com/blog/golang/2018-06-17-golang-using-context-cancellation/

# Introduction

Cette librairie est utilisé pour annuler les opérations sous-jacentes d'une opération principale.
Ex : Un client réalise une requête, le serveur process la requête et interroge la BDD. Le client ferme son navigateur alors il n'est plus nécessaire d'interroger la BDD pour pouvoir fournir une réponse.
Le package context permet de manager ces processus chronophages.