---
title: 07 - How does the web works ?
created: '2019-04-02T20:05:42.555Z'
modified: '2019-04-03T12:07:24.713Z'
tags: [Web]
---

# 07 - How does the web works ?
## Définitions
 - Modem : Le modem permet de transformer une information du format internet au format téléphone permettant de le transférer à travers les lignes téléphoniques déjà déployées.
 - Internet : Internet représente l’infrastructure mise en place. Le web est un service construit pour être utilisé sur cette infrastructure.
 - ISP  (internet service provider) : Fournit la connexion à internet aux utilisateurs
 - Paquets : 
 - Routeur : permet de délivrer un paquet à la bonne adresse IP
 - RNS : Resolving name server
 - DNS :
En réalité, la vraie adresse est www.google.com. Le dernier point représente le root de l’internet namespace. La recherche va d’abord passer par un root name server, puis un TLD (top level domain) name server puis un authoritative name server
 - web server, est composé de 2 parties :
    - hardware : un web server contenant un logiciel de serveur web ainsi que les fichiers composant le site web
    - software : permet de gérer comment les utilisateurs vont interagir avec le serveur web
      - serveur http : permet de comprendre les requêtes http
 - HTTP : C’est un protocole de requête / réponse
Afin de communiquer avec un ordinateur, ce dernier est caractérisé par une adresse IP et un port.
Les adresses IP sont fournies par les ISP.
Le protocle HTTP est stateless (ie aucune information n’est retenue par le serveur). Chaque requête / réponse sont indépendantes les unes des autres

## Commande 
Whois mozilla.org : permet de trouver toutes les informations d’un site

## A voir

livre sur http : https://launchschool.com/books/http/read/background
How the web works :
    • https://twit.tv/shows/security-now/episodes/25?autostart=false  (à partir de 8 min 30)
    • https://twit.tv/shows/security-now/episodes/26?autostart=false

