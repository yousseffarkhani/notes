Source : https://www.sohamkamani.com/blog/golang/2018-06-24-oauth-with-golang/

# Introduction

Oauth est un protocole permettant d'accéder de manière temporaire à une ressource située sur un serveur externe.

Exemple : Spotify qui demande d'accéder aux informations FB.

# Acteurs

Dans un flow OAuth, il existe 3 acteurs :

- Le client : L'utilisateur souhaitant utliser une application
- Le consommateur : L'application à laquelle l'utilisateur souhaite accéder.
- Le service provider : L'application externe disposant des informations client. Ce service provider dispose d'un serveur d'authentification et d'un autre contenant les ressources.

Il existe une multitude de service provider : Google, github, facebook, twitter, ...

# OpenID Connect

OpenID Connect est une couche construite sur la base de OAuth 2.0 permettant de réaliser l'authentification d'un utilisateur.
En effet, la particularité est que l'application ne va demander que les informations de profil.
Ces informations peuvent être envoyées sous forme de JWT.

# Avantages

# Inconvénients

# Processus

0. L'utilisateur se connecte au consommateur (application).
1. L'application redirige le client vers un service provider afin d'émettre une demande d'autorisation à ce dernier (https://github.com/login/oauth/authorize?client_id=6e93cc54ddsdse1208dca9ac&redirect_uri=http://localhost:8080/oauth/redirect)
1. Le service provider autorise le client et le redirige vers l'uri de redirection du 1er lien avec en plus un request token (http://localhost:8080/oauth/redirect?request_token=121233).
1. Le client émet donc la requête due à la redirection auprès du consommateur avec le request token.
1. Le consommateur parse la requête pour récupérer le request token et l'envoyer au service provider en échange d'un access token (https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s).
1. Le consommateur ou l'utilisateur peuvent requêter les API avec l'access token afin de récupérer des informations client.

**Les passwords utlisateurs ne sont jamais transmis !**

> L'access token ne donne accès qu'aux données énoncées lors de l'écran d'authentification.

# Sécurité

- L'access token ne doit pas être transmis au client. Il faut créer un session token à la place qui est envoyé à l'utilisateur sous forme de cookie.
  Le consommateur va maintenir une table avec les session tokens et les access tokens.
  Au lieu de faire un requête directement à github, l'utilisateur va faire une requête au serveur avec le session token. Le serveur va se charger de réaliser la requête auprès de github et renvoyer la réponse à l'utilisateur.
- Dans l'uri de redirection fournit à github il faut ajouter un state à la requête(https://github.com/login/oauth/authorize?client_id=myclientid123&redirect_uri=http://localhost:8080/oauth/redirect&state=somerandomstring). Cela permet à l'application de comparer la valeur au moment de la redirection avec celle qu'elle aura créée au préalable. S'il ne s'agit pas des mêmes valeurs cela signifie que la requête ne provient pas de l'utilisateur et doit être rejetée.
