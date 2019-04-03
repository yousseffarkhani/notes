---
title: 17 - Async
created: '2019-04-02T21:32:56.115Z'
modified: '2019-04-03T12:07:50.260Z'
tags: [JS]
---

# 17 - Async
## Introduction
Etant donné que JS est le langage du web, certaines fonctions seront amenées à mettre un certain temps avant d’être complétées (toutes les fonctions d’I/O, ex : fetch de la donnée d’un serveur) ou peuvent dépendre d’évenement externes (click utilisateur).
Pour cette raison, JS intègre un support pour ces fonctions asynchrones.
>NB : Seules les fonctions accédant à des données en mémoire (RAM ou CPU) pourront être executées en synchrone.

## Callback functions
Une callback function est une fonction passée en argument d’une autre fonction qui sera invoquée par la fonction parent à un certain moment (contrairement aux fonctions classiques retournant immédiatement un résultat).
Les callbacks sont adaptés pour les appels asynchrones (on donne une fonction à executer et on ne s’en soucie plus). Dans un langage standard, il faudrait attendre le résultat avant d’enchainer la suite du programme.
Le problème des callbacks est qu’elles ne sont pas adaptées à l’enchainement d’appel dans un ordre spécifique (aussi appelé callback hell).
## Comment éviter le Callback hell ?
1. Avoir un code léger (pas beaucoup de lignes)
1. Donner des noms aux fonctions (simplifie la lecture et affiche le nom de la fonction dans la stacktrace)
1. Séparer les fonctions
1. Modulariser si possible
1. Gérer les erreurs (partir du principe : error first) → Passer l’erreur en 1er argument pour ne pas oublier de la gérer
## Inversion of control
Le problème avec les callback est qu’on prend l’hypothèse que le programme prenant en entrée la callback invoquera celle-ci au bon moment et de la bonne manière (la callback n’est appelé qu’une seule fois). Cela pose le problème d'inversion de contrôle.
Pour régler ce problème, les promises ont été créees
## Promises
(Analogie du restaurant et du bipper)
Les promises contrairement au callbacks permettent de reprendre le contrôle. En effet, les promises permettent de récupérer directement la donnée et l’envoyer aux fonctions que l’on souhaite (au lieu de devoir faire confiance à l’API).
Une promise a 3 états : pending, resolved et rejected (et settled). La promise ne peut pas revenir en arrière dans ses états. Ces états indiquent l’état de la requête.
### Déclarer une promise
Pour déclarer une promise, il faut utiliser le format suivant :
```
const promise = new Promise(function(resolve, reject){
  resolve(DONNEE);
  reject(RAISON);
}
```
### Exploiter une promise
Pour exploiter le retour de la promise, il faut utiliser .then et .catch (pour les erreurs) (et .finally si l’issue n’importe pas entre reject et resolve)
### Fonctions utiles
Créer une promise anonyme à l’aide du constructeur: `new Promise((resolve, reject) => {... })`
Renvoyer rapidement une promise sans la créer au préalable (utile dans ce cas car on attend une promise en sortie pour l’exploiter) : 
```
var userCache = {};

function getUserDetail(username) {
  // In both cases, cached or not, a promise will be returned
  if (userCache[username]) {  
    return Promise.resolve(userCache[username]); // Return a promise without the "new" keyword
  }

  return fetch('users/' + username + '.json')  // Use the fetch API to get the information, fetch returns a promise
    .then(function(result) {
      userCache[username] = result;
      return result;
    })
    .catch(function() {
      throw new Error('Could not find user: ' + username);
    });
}
```
Lorsque les promises sont chainées alors la valeur de retour est utilisée :
```
new Promise(function(resolve, reject) { 
  // A mock async action using setTimeout
  setTimeout(function() { resolve(10); }, 3000);
})
.then(function(num) { 
  console.log('first then: ', num); 
  return num * 2; 
})
.then(function(num) { 
  console.log('second then: ', num); 
  return num * 2; 
})
.then(function(num) { 
  console.log('last then: ', num);
});

// From the console:
// first then:  10
// second then:  20
// last then:  40
```
Envoyer plusieurs requêtes à la fois (contenues dans un array) et retourner un array de résultats (les résultats seront dans l’ordre des promises) : `Promise.all`
Si une promise est rejetée alors l’ensemble est rejeté.
`Promise.all([fetch('./style.css'), navigator.getBattery()]).then(results => console.log(results));`

Lancer plusieurs requêtes à la fois et récupérer le résultat de la 1ere promise (qu’elle soit rejetée ou resolved) : `Promise.race`

## Async/await
Il existe une autre façon de travailler avec des fonctions asynchrones : async/await.
### Comment déclarer une async function ?
Il faut placer le mot async avant la function :
```
async function f(){
  return 1;
}
```
Le mot async permet de déclarer qu’une fonction va retourner une promise. Même si la fonction retourne une valeur (voir exemple ci-dessus), la valeur sera enveloppée dans une promise.
Si la fonction throw une erreur alors la promise sera rejected.
### Await
Le mot await permet de faire en sorte que l’interpréteur attende le retour de la promise et retourne l résultat :
`let value = await promise;`
La fonction se met en pause au moment de await et continue quand la promise est résolue.
/!\ Le mot await ne peut être utilisée que dans une async function.
