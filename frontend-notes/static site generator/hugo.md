https://www.staticgen.com/
https://davidwalsh.name/introduction-static-site-generators

# Pourquoi utiliser un static site generator ?

L'interêt d'un générateur de website est d'interpréter du contenu (au format .md par exemple) dans un format HTML.
La plupart des générateurs est dit dynamique car le serveur HTTP (le programme envoyant les fichiers au navigateur) lance le générateur pour créer un fichier HTML à chaque requête.

Avec le temps, des systèmes de cache ont été mis en place pour optimiser le temps de livraison des pages. Une page mis en cache correspond à une version statique de cette page.

C'est pourquoi les static site generator ont vu le jour.
Un static site generator va interpréter l'ensemble du contenu en HTML. Ces pages HTML seront ensuite copiées sur le serveur HTTP.

Cette manière de fonctionner permet un gain de performance.

# JAM stack

- Javascript : Permet de gérer le contenu dynamique directement chez le client
- APIs : Tous les processus server side ou BDD sont accessibles via des APIs appelées avec JS.
- Markup

# Avantages

1. Rapidité : Comme il n'y a aucune requête à réaliser pour générer la page (appel à une BDD ou autre)
2. Version Control du contenu : Etant donné que le contenu ne se situe pas dans une BDD mais dans des fichiers, il est aisé d'apporter des modifications et de collaborer sur ces derniers.
3. Sécurité : Il n'y a pas de user input ou d'authentification dans ce genre de site.
4. Simplicité de gestion de l'infrastructure : Pas de gestion de plusieurs serveurs (webserver, BDD)

# Inconvénients

1. Pas de contenu temps réel : Le contenu sera le même pour tout le monde (pas de recommandations personnalisées)
   - Solution : Ajouter du Javascript afin de réaliser des appels et générer le contenu.
2. Pas de user input : Pas de système de commentaire possible.
   - Solution : Des services existent afin de gérer les POST requests et les envoyer par mail.
3. Impossible de publier un blog sans installer Hugo

# Introduction

Hugo est un générateur de static site.
Les sites générés avec Hugo n'ont pas besoin de BDD ou de runtimes comme Ruby, Python ou PHP.

# Pourquoi utiliser Hugo ?

- Edition offline : Hugo est proche d'un CMS. Cependant avec Hugo l'édition est réalisée dans un text editor et peut être réalisée offline.
- Rapidité de build : <1ms par page>

# Fonctionnement

https://gohugo.io/getting-started/quick-start/

# Hugo Modules

Les modules sont les buildings blocks d'Hugo.

# Thèmes

https://themes.gohugo.io/

# Contact Form

https://formspree.io/#setup
