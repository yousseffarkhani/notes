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

## Best practices

Définir la stratégie de cache :

- cache first strategy : Mettre en cache les éléments statiques : images, fonts, CSS, JS. Pour cela essayer d'avoir en local les fichiers.
- Network first strategy :
- Construire une app shell (Server-side rendering)

## Test d'une PWA

1. Builder le projet
2. http-server -p 8080 -c-1 dist/<project-name>
3. Visualiser l'appli en mode incognito

## Structure de fichiers

- manifest.json : C'est un fichier permettant de dire au navigateur comment se composter lorsque la PWA est installée sur le mobile ou bureau. Les données sont name, icons, start_url.
  [Plus d'infos](https://developers.google.com/web/fundamentals/web-app-manifest/)
- SW : un SW est un script qui tourne en background séparément de la page web
- ngsw-config.json :
  - AssetGroups
    - installMode :
      - prefetch : Le SW va requêter et mettre en cache les ressources
      - lazy : Le SW va mettre en cache les ressources lorsqu'elles sont demandées par l'utilisateur
    - updateMode : similaire à installMode
    - ressources :
      - files
      - urls : Ces ressources seront cachées selon leur header HTTP. Pratique pour les fonts.
  - dataGroups : Utile pour les API request
    - https://angular.io/guide/service-worker-config
    - strategy :
      - performance : utile pour les ressources qui ne changent pas du type avatar

```json
  "dataGroups": [
    {
      "name": "car-info-api",
      "urls": ["https://www.carqueryapi.com/api/0.3/"],
      "cacheConfig": {
        "strategy": "freshness",
        "maxSize": 10000,
        "maxAge": "6h",
        "timeout": "5s"
      }
    }
  ]
```

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
Si des fichiers ont été modifiés depuis la dernière connection, updater le cache

### NSGW

Qu'est ce qui est mis en cache par défaut ?(voir ngsw-config.json)

- index.html.
- favicon.ico.
- Build artifacts (JS and CSS bundles).
- Anything under assets.
- Images and fonts directly under the configured outputPath (by default ./dist/<project-name>/) or resourcesOutputPath.

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

#### Pour subscribe et envoyer des notifications

```javascript
import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {SwPush} from '@angular/service-worker';


@Injectable()
export class PushNotificationService {

  constructor(private http: HttpClient,
              private swPush: SwPush) {
  }

  addPushSubscriber() {
    this.swPush.requestSubscription({
      serverPublicKey: environment.vapid_public_key
    })
      .then(sub => {
        console.log('User accept notifications!');
        console.log(sub);

        this.http.post(environment.main_api_url + '/api/notifications', sub).subscribe(
          () => console.log('Subscription added successfully.'),
          err => console.error('Could not send subscription object to server, reason: ', err)
        );
      })
      .catch(err => console.error('Could not subscribe to notifications', err));
  }

  send() {
    console.log('Send notification request to server.');
    this.http.post(environment.main_api_url + '/api/notifications/enable', null).subscribe(
      () => console.log('Notification send successfully!'),
      err => console.error('Failed, reason: ', err)
    );
  }
}
```

### Workbox

https://developers.google.com/web/tools/workbox/guides/get-started
