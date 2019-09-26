Source :

# Introduction
Nginx est un serveur web.
Il peut servir à réaliser du load balancing, du caching http ou servir de reverse proxy*.

Ce dernnier peut herberger des applications web.
# Installation
sudo apt-get install nginx
# Processus
1. Nginx reçoit une requête d'une ressource
2. Nginx envoie une "proxied" request au serveur spécifié par la 1ère requête
3. Après avoir reçu la réponse, NGinx va retourner le résultat à la requête originale.


# Définition
## Proxy serveur :
Il existe 2 types de proxy serveurs : forward proxy (aussi appelé proxy) et reverse proxy.

### Différences
La principale différence est que le proxy est utilisé par le client (navigateur) alors que le reverse proxy est utilisé par le serveur (web server).
### Proxy
Le proxy peut être utilisé par le client pour bypasser des restrictions du firewall pour visiter des sites bloqués.
Si un site web a bloqué un intervalle d'adresse IP de visiter ce site alors la personne dans cet intervalle peut utiliser un proxy pour masquer son IP pour pouvoir visiter le site.
### Reverse proxy
Le reverse proxy est principalement utilisé par les admin serveurs pour :
- réaliser du load balancing : Un 1er serveur reçoit la requête puis la forward à un serveur disponible.
- accélerer le temps de réponse : Les proxy peuvent mettre en cache des ressources souvent demandées, compresser les données entrantes et sortantes, réaliser le cryptage SSL pour éviter cette surcharge au webserver.
- augmenter la sécurité et l'anonymité : En interceptant les requêtes le proxy peut protéger les identités.