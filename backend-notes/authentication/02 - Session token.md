Source : https://www.sohamkamani.com/blog/2018/03/25/golang-session-authentication/

# Introduction

Le session based token consiste à stocker la session utilisateur côté back-end et à créer un token correspondant qui sera transmis à l'utilisateur.
L'utilisateur pourra inclure ce token dans ses requêtes afin que le serveur puisse l'identifier par la suite.

Pour stocker les sessions, Redis est souvent utilisé car il s'agit d'une BDD key-value permettant de définir une date d'expiration.

# Avantages

- Possibilité de stocker la session utilisateur (pratique pour les sites de e-commerce)
- Il est possible de facilement révoquer l'accès à un utilisateur.

# Inconvénients

- Il faut gérer une BDD cache.
- Pas pratique pour les cross platform app car les mobile apps ne sont pas des navigateurs et ne supportent pas nativement les cookies.
