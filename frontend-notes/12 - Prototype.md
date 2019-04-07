---
title: 12 - Prototype
created: '2019-04-02T21:31:42.074Z'
modified: '2019-04-03T12:07:44.181Z'
tags: [JS]
---

# 12 - Prototype
## Introduction
Un prototype est une propriété d’une fonction (toutes les fonctions JS dispose d’un prototype) pointant vers un objet listant les méthodes associées. Cela permet de partager des méthodes à travers toutes les instances d’une fonction.
## Propriété d’un prototype
Toute fonction JS possède une propriété prototype (vide par défaut).
Pour créer de l’héritage, il faut attacher des propriétés et méthodes à ce prototype
## Pourquoi ajouter une propriété à un prototype plutôt qu’à l’objet lui même ? (explications des instanciation patterns)
Référence : https://tylermcginnis.com/beginners-guide-to-javascript-prototype/
Le problème en ajoutant des méthodes à une fonction de cette manière est qu’à chaque instanciation de l’objet les fonctions seront recopiées utilisant ainsi de la mémoire inutilement.
```
function Animal(name, energy) {
  let animal = {}; // Cette ligne permet de créer l’objet en lui même
  animal.name = name;
  animal.energy = energy;

  animal.eat = function(amount){
    console.log(`${this.name} is eating`);
    this.energy += amount;
  }
  return animal; // Permet de récupérer l’objet crée
}

const leo = Animal("Leo", 7);
const elo = Animal("Elo", 8);
```
Afin de résoudre ce problème, la 1ère méthode pourrait être de créer un objet spécifique répertoriant les méthodes à créer :
```
const AnimalMethods = { // équivalent d’un prototype (peut également contenir des propriétés)
  eat : function(amount){
    console.log(`${this.name} is eating`);
    this.energy += amount;
  }
}

function Animal(name, energy) {
  let animal = {};
  animal.name = name;
  animal.energy = energy;
  animal.eat = AnimalMethods.eat;
  return animal;
}
```
Le problème de cette solution est qu’elle n’est pas très évolutive, lorsqu’une nouvelle méthode est ajoutée à AnimalMethods, il faut également la rajouter dans la fonction Animal.

Pour résoudre cela, on utilise la méthode Object.create(X). La méthode Object.create(X) fonctionne de sorte à ce que quand on essaie de lire une propriété/méthode d’un objet et que ce dernier ne contient pas la propriété/méthode alors l’interpreteur JS va aller voir dans l’objet X si cette propriété/méthode existe.
/!\ Il ne faut pas créer d’héritage de cette manière Animal.prototype = Animal2.prototype car cela créerait un lien direct entre ces 2 constructeurs. Pour créer une copie d’un objet, il faut utiliser Object.create(X)
```
function Animal(name, energy) {
  let animal = Object.create(AnimalMethods); // permet de déléguer les failed look ups à AnimalMethods.
  animal.name = name;
  animal.energy = energy;
  return animal;
}
```
Cette solution présente le défaut de devoir gérer un objet contenant toutes les méthodes d’un objet. C’est là que le prototype vient résoudre le problème.
Un prototype est une propriété d’une fonction (toutes les fonctions JS dispose d’un prototype) pointant vers un objet listant les méthodes associées. Cela permet de partager des méthodes à travers toutes les instances d’une fonction.
Pour déclarer un prototype, la syntaxe est la suivante :
```
//Constructor function permettant de créer un objet
function Animal(name, energy) {
  let animal = Object.create(Animal.prototype);
  animal.name = name;
  animal.energy = energy;
  return animal;
}

Animal.prototype.eat = function(amount){
  console.log(`${this.name} is eating`);
  this.energy += amount;
}
```
L’opérateur new réalise pour nous les 2 étapes suivantes en remplacant animal par this.:
1. `let this = Object.create(AnimalMethods); // permet de créer l’objet et de déléguer les failed look ups à AnimalMethods.`

2. `return this; // Permet de récupérer l’objet crée`

Après avoir utilisé le mot « new » , l’utilisation de this devient obligatoire.
```
function Animal(name, energy) {
  this.name = name;
  this.energy = energy;
}
```
Astuces :
    • Object.getPrototypeOf(X) → Pour récupérer le prototype sur lequel est basé un objet X

/!\ Il est impossible de créer une fonction constructeur avec une arrow function. C’est pourquoi les arrow functions n’ont pas de propriété prototype.
