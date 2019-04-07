---
title: 19 - JSON
created: '2019-04-03T09:33:13.723Z'
modified: '2019-04-03T12:07:56.566Z'
tags: [JS]
---

# 19 - JSON
## Introduction
Le format JSON est souvent utilisé pour les échanges de données sur le web.
Lors de l’échange la donnée est toujours sous forme de texte.
Pour exploitée pleinement la donnée avec JS, il faut convertir le JSON en un objet JS.
## JSON.parse()
Pour transformer directement en JSON après une requete HTTP :
`request.responseType = 'json';`
Pour transformer en texte :
` request.responseType = 'text';`
Pour convertir un JSON en un objet JS : `JSON.parse()`
/!\Si le JSON représente un array alors il sera converti en array.
Pour convertir une data depuis un JSON, 2 méthodes existent :
- Convertir le JSON en objet puis convertir la date via `date = new Date(obj.date)`
- En utilisant le 2nd paramètre de parse appelé le reviver (le reviver est une fonction qui vérifie les propriétés avant de retourner une valeur)
```
obj = JSON.parse(text, function(key, value){
    if(key ≡ ‘date’){
      return new Date(value)
    } else {
      return value
    }});
```
## JSON.stringify()
Lorsque les données sont envoyées à un serveur web, elle doivent impérativement être au format string.
Il est possible de convertir un objet JS ou un array.
Les date objects seront converties en string (`new Date()` → 2019-02-14T09:20:21.814Z).
Les fonctions sont supprimés du JSON :
```
var obj = { name: "John", age: function () {return 30;}, city: "New York"};
var myJSON = JSON.stringify(obj);
→  {"name":"John","city":"New York"}
```
Pour intégrer des fonctions, il faut d’abord les convertir en string :
`obj.age = obj.age.toString();`
/!\ Il est très peu recommandé d’envoyer des fonctions à travers un JSON
