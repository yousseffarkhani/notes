---
title: SPA vs SSR
created: '2019-04-03T13:13:39.510Z'
modified: '2019-04-03T13:26:19.801Z'
---

# SPA vs SSR
Critères permettant de choisir entre les deux :
- SEO : Avantage SSR
- Performance : 
  - Initial Page Load : Avantage SSR
  - Par la suite : Avantage SPA car tous les assets sont téléchargés
- Population : 
  - Mobile : Avantage SSR

Traditionnellement, le client réalise une requête auprès d'un serveur et reçoit le HTML rempli avec les données nécessaires.
Avec les SPA cependant, le serveur envoie l'ensemble des fichiers nécessaires puis le client réalise des requêtes AJAX et reçoit du JSON qui est ensuite affiché à l'écran.

Architecture : 
- Home / About / Pricing / ... : SSR (website.com)
- Application : SPA (app.website.com)
