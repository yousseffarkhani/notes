---
title: 18 - API
created: '2019-04-03T09:32:41.209Z'
modified: '2019-04-03T12:07:54.732Z'
tags: [JS]
---

# 18 - API
## Introduction
L’une des fonctionnalités les plus importantes du web est d’aller chercher des informations (fetch data) d’un serveur et de l’afficher sur un site.
Dans la plupart des cas, un serveur n’existe que pour cette raison. Le serveur peut contenir des blog posts, des informations utilisateur, des scores pour un jeu, ..
Dans d’autres cas, les serveurs sont dit en open service et fournissent donc de la donnée à n’importe qui (données météo, stock prices).
Dans les deux cas, l’accès aux données est réalisé de la même manière à travers des API.
## Comment ça marche ?
Les serveurs donnant accès à leurs données sont souvent appelés API.
Dans la majorité des cas, il faudra s’inscrire au préalable pour recevoir une API key.
### Comment récupérer la donnée dans notre programme ? 
La fonction native Fetch permet de réaliser cela :
`fetch(‘url’).then(response => return response.json()).then().catch()`
Fetch retourne une promise qu’il faudra exploiter grâce à `.then` et `.catch`.
La méthode fetch renvoi uniquement le header (un objet response) dans un 1er temps, c’est pourquoi, il faut utiliser la méthode `.json()` pour extraire le corps (body) de la réponse. L’extraction (parsing) prend un certain temps, c’est pourquoi, il faut considérer cette étape comme un évenement asynchrone.
https://davidwalsh.name/fetch
### CORS (Cross Origin Resource Sharing)
CORS est un mécanisme autorisant certaines ressources sur une page web à faire appel à un autre domaine d’où provient la 1ère ressource
Pour des raisons de sécurité, les navigateurs (par défaut) restreignent les requêtes HTTP aux sources externes.
Pour résoudre ce problème, il faut ajouter l’objet {mode : ‘cors} en paramètre de fetch.
