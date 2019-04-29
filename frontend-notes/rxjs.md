# Rxjs

## Introduction

Rxjs est une librairie permettant de transformer les streams de données.
Cette librairie facilite le travail avec les streams de données asynchrones chose non possible avec les promises/async-await.

Pour utiliser rxjs:

1. Importer le script suivant : `https://unpkg.com/rxjs/bundles/rxjs.umd.min.js`
2. Importer les méthodes de cette manière :

```javascript
const { range } = rxjs;
const { map, filter } = rxjs.operators;
```

## Processus

1. Attacher un observable
2. Souscrire à l'observable
3. Ajouter des callbacks à la souscription

## Déclarer un observable

### Observable

#### Introduction

Les observables sont similaires aux promises. Angular utilise les observables afin de gérer des opérations asynchrones (AJAX, event utilisateur).
Un observable est un objet émettant des informations auxquelles on souhaite réagir.
Les observables permettent de transmettre des messages entre publishers et subscribers(/observer) dans l’application.
A cet observable, on associe un observer (un bloc de code qui sera executé à chaque fois que l’observable émet un message).
L’observable émet 3 types d’informations : des données, une erreur, un message ‘complete’. Du coup un observer peut avoir jusqu’à 3 fonctions pour réagir à chaque information.
Observable.subscribe(fonction 1, fonction 2, fonction 3). Les 2 dernières fonctions étant optionnelles.
La 1ère fonction va recevoir les données comme argument, la 2ème l’erreur (s’il y en a une), la 3ème se déclenche mais ne reçoit pas d’argument.
/!\ La souscription à un observable dure jusqu’à l’infini (fin du programme) tant que l’observable émet des données. Cela peut causer des bugs. C’est pourquoi il faut stocker cette souscription dans une variable afin que Angular la supprime lorsque nous n’en avons plus besoin. (voir subscriptions – https://openclassrooms.com/fr/courses/4668271-developpez-des-applications-web-avec-angular/5089331-observez-les-donnees-avec-rxjs).

##### Déclarer un observable

Il existe plusieurs méthodes :

- Observables de tableau

```javascript
const observable = of([1, 2, 3]); // méthode la plus utilisée
const observable = from([1, 2, 3]);
const myObservable3 = interval(2000);
```

Pour convertir une variable en observable, il faut utiliser la méthode of(‘variable’).

- Osbervables de fonction

```javascript
Observable.create(observer => {
  observer.next(42);
  observer.error(43);
  observer.complete(); //Le flux est terminé et il n’y a plus rien à émettre
});
```

- Observables de DOM

```javascript
const keyUp = Observable.fromEvent(document.querySelector("input"), "keyup");
const inputObservable = fromEvent(input, "input");
```

#### Subjects

Un subject est un type d’observable qui permet non seulement de réagir à de nouvelles informations, mais également d’en émettre. L’interêt étant de centraliser la donnée dans un service et en cas de changements le subject émettra un message à l’ensemble des observers.
Cela permet d’avoir une complète abstraction entre le service et les components. (voir Subjects - https://openclassrooms.com/fr/courses/4668271-developpez-des-applications-web-avec-angular/5089331-observez-les-donnees-avec-rxjs).

```javascript
subject.next("valeur à transmettre");
subject.error("valeur à transmettre");
subject.complete();
```

#### BehaviorSubject

- Rajoute une valeur initiale au subject
  Pour déclarer un BehaviorSubject :
  `const click = new BehaviorSubject('not clicked');`
- Peut renvoyer la dernière valeur du subject en utilisant la méthode obs.getValue() n'importe où dans le code
- A la souscription renvoie la dernière valeur émise

#### Déclarer un observer

```javascript
observable.subscribe(
  event => {
    console.log(event.target);
  },
  error => {
    console.log(error);
  },
  () => {
    console.log("Completed!");
  } // <- Gets executed once the observable completes - it doesn't receive any argument, no data
);
```

```javascript
let keyUp = fromEvent(document.querySelector("input"), "keyup");
let observer = {
  next: value => {
    console.log(value);
  },
  error: error => {
    console.log(error);
  },
  complete: () => console.log("completed")
};

keyup.subscribe(observer);


/* Pour afficher la valeur d'un input */
inputObservable.subscribe(event => {
const tar = event.target as HTMLInputElement;
console.log(tar.value);
});
```

## Operators

La puissance de Rxjs provient des operators permettant de réaliser de multiples opérations.
L'ordre de déclaration des opérateurs est important.

- map("fonction") :
  - Utilité : Applique une fonction sur un observable et renvoie ce dernier
  - Valeur de retour : Un observable avec la fonction appliquée dessus
- throttletime("temps en ms") :
  - Utilité : Permet de réguler l'envoi d'évenement d'un observable
  - Valeur de retour : Renvoi le même observable
- filter(function(value => condition)) :
  - Utilité : Permet de filtrer les valeurs entrantes selon une condition définie
  - Valeur de retour : Renvoi l'observable filtré
- debounceTime("temps en ms") :
  - Utilité : La valeur est renvoyée si rien ne se passe pendant l'intervalle de temps précisé
  - Valeur de retour : Renvoi l'observable de départ
- distinctUntilChanged() :
  - Utilité : Permet d'ignorer les requêtes identiques aux requêtes précédentes
  - Point d'attention : Utiliser map au préalable pour renvoyer la valeur totale.
- reduce :
  - Utilité : Similaire à array.reduce
  - Valeur de retour : Unique valeur
- scan :
  - Utilité : Ressemble à reduce sauf que renvoie l'ensemble des étapes de l'opération. Utile lorsque l'observable n'est jamais completé.
  - Valeur de retour : A chaque itération renvoie une valeur
- pluck :
  - Utilité : Permet d'extraire une propriété d'un objet
  - Valeur de retour : La propriété de l'objet
  - Exemple :
  ```javascript
  inputObservable
    .pipe(
      pluck("target", "value"),
      debounceTime(500),
      distinctUntilChanged()
    )
    .subscribe(event => console.log(event));
  ```
- mergeMap :
  - Utilité : Permet de combiner les valeurs de 2 observables
  - Valeur de retour : Valeurs combinées des 2 observales
  - Point d'attention : L'observable ne se déclenche que quand l'inner observable est déclenché.
  - Exemple :

```javascript
inputObservable1
  .pipe(
    pluck("target", "value"),
    mergeMap(value1 => {
      return inputObservable2.pipe(
        pluck("target", "value"),
        map(value2 => value1 + " " + value2)
      );
    })
  )
  .subscribe(event => console.log(event));
```

- switchMap :
  - Utilité : A chaque émission de données, annule la souscription au précédent observable et souscrit le nouvel observable. (Utile pour gérer des clicks events et éviter les spams). Utile pour l'auto-complétion car nous n'avons pas besoin du retour de la requête précédente.
  - Valeur de retour : La valeur de l'inner observable
  - Exemple : `obs.pipe(switchMap(() => itv)).subscribe(console.log);`

## Use-case

### Ecouter un champ input et réaliser une requête

```javascript
import { Observable } from "rxjs";

let inputValue;

let keyUp = Observable.fromEvent(document.querySelector("input"), "keyup");
let observable = keyUp
  .map(e => (inputValue = encodeURIComponent(e.target.value))) // Permet de récupérer les valeurs de l'input
  .debounceTime(500) // Permet d'attendre que l'input soit stabilisé avant de réaliser la requête (afin d'éviter les spams de touches)
  .distinctUntilChanged() //Permet d'ignorer les requêtes identiques aux requêtes précédentes
  .map(value => {
    return Observable.fromPromise(fetch("https://cataas.com/cat/" + inputValue)) // Transforme la promise en observable
      .retry(3) // Permet de recommencer la requête 3 fois avant d'émettre une erreur
      .takeUntil(keyUp); //Je récupère les valeurs à moins qu'un autre observable émette une valeur
  })
  .concatAll(); // ???
```
