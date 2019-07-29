Source : https://www.sohamkamani.com/blog/golang/2018-06-24-oauth-with-golang/

# Introduction

Dans un flow OAuth, il existe 3 acteurs :

- Le client : L'utilisateur souhaitant se logger
- Le consomateur : L'application à laquelle l'utilisateur souhaite accéder.
- Le service provider : L'application externe à travers laquelle le client s'authentifie.

Il existe une multitude de service provider : Google, github, facebook, ...

Avec ce protocole, il est possible de récupérer les informations du service provider.

# Processus

1. Le client émet une demande d'autorisation au service provider (https://github.com/login/oauth/authorize?client_id=6e93cc54ddsdse1208dca9ac&redirect_uri=http://localhost:8080/oauth/redirect)
2. Le service provider autorise le client et le redirige vers l'uri de redirection du 1er lien avec en plus un request token (http://localhost:8080/oauth/redirect?request_token=121233).
3. Le client émet donc la requête due à la redirection auprès du consommateur avec le request token.
4. Le consomateur parse la requête pour récupérer le request token et l'envoyer au service provider en échange d'un access token (https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s).
5. Le consomateur ou l'utilisateur peuvent requêter les API avec l'access token afin de récupérer des informations client.

# Sécurité

- L'access token ne doit pas être transmis au client. Il faut créer un session token à la place qui est envoyé à l'utilisateur sous forme de cookie.
  Le consomateur va maintenir une table avec les session tokens et les access tokens.
  Au lieu de faire un requête directement à github, l'utilisateur va faire une requête au serveur avec le session token. Le serveur va se charger de réaliser la requête auprès de github et renvoyer la réponse à l'utilisateur.
- Dans l'uri de redirection fournit à github il faut ajouter un state à la requête(https://github.com/login/oauth/authorize?client_id=myclientid123&redirect_uri=http://localhost:8080/oauth/redirect&state=somerandomstring). Cela permet à l'application de comparer la valeur au moment de la redirection avec celle qu'elle aura créée au préalable. S'il ne s'agit pas des mêmes valeurs cela signifie que la requête ne provient pas de l'utilisateur et doit être rejetée.
