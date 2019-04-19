# PWA
## Introduction
Une PWA est une application web disposant de capacités similaires à des apps natives.
Pour cela l'application utilise les APIs des navigateurs pour délivrer une meilleure expérience utilisateur.
Avantages:
- Offline + smart networking
- Notifications

https://whatwebcando.today/
https://developers.google.com/web/progressive-web-apps/checklist#site-works-cross-browser

https://developer.mozilla.org/en-US/docs/Web/Progressive_web_apps

## Structure de fichiers
- manifest.json : C'est un fichier permettant de dire au navigateur comment se composter lorsque la PWA est installée sur le mobile ou bureau. Les données sont name, icons, start_url. 
[Plus d'infos](https://developers.google.com/web/fundamentals/web-app-manifest/)
- SW : un SW est un script qui tourne en background séparément de la page web
## Angular
Options pour mettre en place une PWA :
- Coder manuellement les service workers
- Utiliser NGSW (Angular Service Worker)
- Utiliser des librairies PWA (SW-precache, Workbox)

## MVP
Web App Manifest
App Shell : UI de notre application (Sélection des fichiers requis)
Servie via HTTPS

Au 1er chargement, mettre les fichiers dans le cache storage.

Aux chargements suivants : Les servir depuis le cache storage
Si des fichiers ont été modifiés depuis la dernière connection, updaté le cache

### NSGW
1. Scaffolding: Schematics

ng add @angular/pwa

Cette commande permet de:
- rajouter un service worker registration code au root module
- Générer un service worker configuration file
- Générer un lien Web App Manifest
- Générer des icônes
- Activer le support des SW

2. Building: Angular CLI

ng build --prod
Cette commande permet de:
3. Serving: NGSW
### Workbox
https://developers.google.com/web/tools/workbox/guides/get-started