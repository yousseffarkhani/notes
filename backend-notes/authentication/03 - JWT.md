Source : https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/

# Introduction

JWT est un format de token permettant de transmettre des données entre 2 parties de manière sécurisée.
Ce token contient divers informations comme la date d'expiration, le username, ...
Ce token est pratique puisque grâce à lui il n'est pas nécessaire de garder en mémoire la session d'un utilisateur.

# Où le stocker ?

Dans les cookies.

# Avantages

- Simple à mettre en oeuvre
- Pas de problème de scaling

# Inconvénients

- Difficile d'invalider le JWT d'un utilisateur puisqu'il détient les cookies
- Pas de suivi du client (pas de state/stateless) à moins de le rajouter au token
