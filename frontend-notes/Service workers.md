---
title: Service workers
created: '2019-04-06T13:59:35.581Z'
modified: '2019-04-06T16:03:24.850Z'
---

# Service workers
## Introduction
### Description
Les service workers sont des scripts JS qui permettent de mettre en cache des pages web. 
Le viewer peut ainsi visualiser les pages sans disposer de connexion à la suite du 1er chargement.
Ils sont utilisés dans le cadre des PWA (Progressive Web App).

Le service worker intervient entre le navigateur et le serveur. Il va intercepter les requêtes du navigateur et décider de l'action à réaliser : Montrer la version online du site ou offline. il s'agit donc d'une sorte de proxy.
Au préalable les pages auront été stockées dans le cache storage du navigateur.

Il existe une librairie pour la gestion des SW : workbox.
### Use cases
- Offline browsing and caching
- Notifications
- Background sync : Reporte les actions réalisées par un utilisateur (ex: like d'un post) à un moment ultérieur ou la connection sera rétablie.

### Propriétés
- Les services workers n'agissent pas sur le DOM, ils ne font que rerouter vers un fichier HTML.
- Les service workers ont besoin d'https sauf sur localhost.

## Commandes
Il existe 2 manières de mettre en cache les pages :
- Au chargement de la page en spécifiant tous les fichiers à mettre en cache :

  0. (Optionnel) Configuration 
```
  const cacheName = 'v1'; //Permet de versionner le cache
  const cacheAssets = [
      'index.html',
      'about.html',
      '/css/style.css',
      '/js/main.js'
  ];
```
  1. Register the worker (main.js)
```
  // Make sure SW are supported
  if('serviceWorker' in navigator){ //navigator représente le browser object, il est également possible d'utiliser : if(navigator.serviceWorker)
      window.addEventListener('load', () => {
          navigator.serviceWorker
              .register('../sw_cached_pages.js')
              .then(reg => console.log('Service Worker: Registered'))
              .catch(err => console.log(`Service Worker : Error : ${err}`));
      }) // Register when the window load
  }
```
  2. Install (sw_cached_pages.js)
```
  self.addEventListener('install', (event) => {
      console.log('Service Worker: Installed');
      event.waitUntil(
          caches
              .open(cacheName)
              .then(cache => {
                  console.log('Service worker : Caching Files');
                  cache.addAll(cacheAssets);
              })
              .then(() => self.skipWaiting())
      )
  });
```
  3. Activate (sw_cached_pages.js)
```
  self.addEventListener('activate', (event) => {
      console.log('Service Worker: Activated');
      // Remove unwanted caches
      event.waitUntil(
          caches.keys() // Loop through the caches, if the cache of the iteration is not the current cache, delete it
              .then(cacheNames => {
                  return Promise.all(
                      cacheNames.map(cache => {
                          if(cache !== cacheName){
                              console.log('Service Worker: Clearing old cache');
                              return caches.delete(cache);
                          }
                      })
                  )
              })
      )
  });
```
  4. Reroute on fetch (sw_cached_pages.js)
```
  self.addEventListener('fetch', event => { //Catch the request
      console.log('Service Worker: Fetching');
      event.respondWith(
          fetch(event.request)
              .catch(() => caches.match(event.request)) //Loads the request from the cache
      )
  })
```
  - Mise en cache au fur et à mesure de la navigation :

    0. (Optionnel) Configuration 
```
  const cacheName = 'v1'; //Permet de versionner le cache
```
  1. Register the worker (main.js)
```
  // Make sure SW are supported
  if('serviceWorker' in navigator){ //navigator représente le browser object, il est également possible d'utiliser : if(navigator.serviceWorker)
      window.addEventListener('load', () => {
          navigator.serviceWorker
              .register('../sw_cached_site.js') //Changement de SW
              .then(reg => console.log('Service Worker: Registered'))
              .catch(err => console.log(`Service Worker : Error : ${err}`));
      }) // Register when the window load
  }
```
  2. Install (sw_cached_site.js)
```
  self.addEventListener('install', (event) => {
      console.log('Service Worker: Installed');
  });
```
  3. Activate (sw_cached_site.js)
```
  self.addEventListener('activate', (event) => {
      console.log('Service Worker: Activated');
      // Remove unwanted caches
      event.waitUntil(
          caches.keys() // Loop through the caches, if the cache of the iteration is not the current cache, delete it
              .then(cacheNames => {
                  return Promise.all(
                      cacheNames.map(cache => {
                          if(cache !== cacheName){
                              console.log('Service Worker: Clearing old cache');
                              return caches.delete(cache);
                          }
                      })
                  )
              })
      )
  });
```
  4. Reroute on fetch (sw_cached_site.js)
```
self.addEventListener('fetch', event => { //Catch the request
    console.log('Service Worker: Fetching');
    event.respondWith(
        fetch(event.request)
            .then(res => {
                // Make clone of response
                const resClone = res.clone();
                // Open cache
                caches
                    .open(cacheName)
                    .then(cache => {
                        // Add response to cache
                        cache.put(event.request, resClone);
                    });
                return res;
            })
            .catch(err => caches.match(event.request).then(res => res)) //Loads the request from the cache
    )
})
```
