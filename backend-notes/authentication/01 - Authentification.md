# Introduction

Il existe plusieurs méthodes pour authentifier un utilisateur :

- Classique : Stockage des login et password (chiffré) dans une BDD. (https://www.sohamkamani.com/blog/2018/02/25/golang-password-authentication-and-storage/)
- OpenID Connect : Il s'agit d'un protocole d'authentification utilisé pour authentifier et autoriser les utilisateurs à utiliser une application à l'aide d'un service provier.

Etant donné la nature stateless de HTTP, il est important de pouvoir maintenir la connection entre le client et le serveur. Pour réaliser cela, un token est envoyé à l'utilisateur (navigateur) et ce token est envoyé à chaque requête vers le serveur.
Ce token permet ainsi d'identifier l'utilisateur.

Il existe 2 types de token :

- JWT (JSON Web Token)
- Session based token
