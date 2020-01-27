---
title: 11 - Scope / Closure
created: '2019-04-02T21:30:23.831Z'
modified: '2019-04-03T12:07:40.630Z'
tags: [JS]
---

# 11 - Scope / Closure
https://www.freecodecamp.org/news/this-is-why-we-need-to-bind-event-handlers-in-class-components-in-react-f7ea1a6f93eb/
## Contexte / This : 
Le mot `this` permet de réutiliser une fonction avec différents contextes (quel objet sera focus).
`this` est régie par 4 règles :
 - Implicit binding : This représente le 1er élément à gauche (de la fonction) du point au moment de l’appel (même s’il y a plusieurs éléments à gauche).
 - Explicit binding : voir call/bind/apply ci-dessous
 - New and window binding 
Le contexte représente ce que la valeur this vaut au moment de l’appel.
This peut poser problème lorsqu’une fonction est déclarée dans une autre :
```
var nav = document.querySelector('.nav'); // <nav class="nav">
var toggleNav = function () {
  console.log(this); // <nav> element
  setTimeout(function () {
    console.log(this); // [object Window]
  }, 1000);
};
nav.addEventListener('click', toggleNav, false);
```
On peut voir que this perd la référence à l’objet nav. Pour résoudre ce problème, on peut mettre la référence en cache et la réutiliser :
```
var nav = document.querySelector('.nav'); // <nav class="nav">
var toggleNav = function () {
  var that = this;
  console.log(that); // <nav> element
  setTimeout(function () {
    console.log(that); // <nav> element
  }, 1000);
};
nav.addEventListener('click', toggleNav, false);
```
Changer de contexte avec .call(), apply() et bind()
```
let obj = {
  foo: function(){
    console.log(this)
  }
}

obj.foo.call(window); // permet de changer le contexte lors de l’appel de la méthode.
```
/!\ Les arguments d’une méthode viennent après la valeur de this. Exemple pour une méthode foo : function (one, two, three) :
obj.foo.call(window, 1, 2, 3) ;
call et apply font exactement la même chose. La différence étant que apply nécessite un array en 2ème argument (si il y en a besoin) : obj.foo.apply(window, [1, 2, 3]) ;
Bind quand à lui ne modifie pas directement le contexte de la méthode appelée mais retourne une méthode avec un contexte modifié :
```
let myBoundFoo = obj.foo.bind(window);
myBoundFoo();
```
De la même manière que call on peut rajouter des arguments à bind qui seront transmis à la fonction.


## Execution contexts
Voir https://www.youtube.com/watch?v=Nt-qa_LlUH0
Voir https://tylermcginnis.com/javascript-visualizer/
## Global execution context
Le 1er contexte lancé par JS est le global execution context.
Chaque execution context aura 2 phases, la phase de création et la phase d’exécution.
La phase de création consiste en :
1. La création de window (global object dans node)
1. La création de this se référant à window
1. Attribution d’un espace mémoire pour les variables et les fonctions
1. Attribution de la valeur par défaut undefined aux variables.
1. Attribution des déclarations de fonctions en mémoire
1. Fin de la phase de création

La phase d’execution, c’est dans cette phase que JS execute le code ligne par ligne :
Attribution des valeurs aux variables.

## Hoisting (hisser/lever)
Le processus de hoisting représente la phase où les variables et fonction sont définies en tant que undefined/déclaration de fonctions dans la phase de création (étape 4/5).

## Function execution context
Ce contexte est crée lorsqu’une fonction est invoquée.
La différence par rapport au global execution context est qu’au lieu de créer le global object (étape 1), le function execution context va créer un arguments objects car lorsqu’une fonction est invoquée elle peut contenir des arguments (Lorsque argument[0] est utilisé, on fait appel à cet objet arguments).
De plus lorsqu’un argument est passé alors le function execution context va créer cette variable et lui attribuer sa valeur dès la phase de création.
Dès que le execution context est terminé, il est retiré de l’execution context parent.
## Execution stack
Représente l’ensemble des execution context notamment en cas de function nested dans d’autres fonctions.
## Scope : 
Le scope désigne l’endroit où les variables et fonctions peuvent être utilisées dans le code.
Cet endroit correspond en réalité à l’execution context dans lequel est stocké la variable.
Toute fonction/boucle/conditions(if/switch) permet de créer un local scope. Si une variable y est déclarée, elle n’est utilisable qu’à l’intérieur de ce dernier. On dit que la variable est block scoped (c’est à dire dès que des {} apparraissent). 
/!\Le mot var ne faisait que du function scope, ce qui veut dire que déclarée dans un if statement, la variable était accessible partout, c’est pourquoi, il faut utiliser let et const.

/!\ Il est possible d’ajouter « use strict » ; au début du code JS pour informer JS que nous ne souhaitons pas pouvoir polluer le root scope (window).
```
"use strict";

function foo(){
a = 2; //let/const a été oublié
}
foo(); //Cet appel va créer une ReferenceError car JS tente d’affecter a au root scope. 
console.log(a);
```
## Closure
Dès qu’on a une fonction dans une autre fonction, la fonction à l’intérieur va retenir les variables du contexte d’execution de la variable parent : ce process est appelé closure.
Le concept de closure est que chaque fonction retient son scope et ses variables même si elle est appelée en dehors de la fonction mère. La fonction retient son scope jusqu’à ce qu’il n’y ai plus du tout de possibilité d’appel de la fonction et des fonctions internes.
L’endroit où sont stockés le scope et les variables associés s’appelle une closure.
Les closures nous permettent de sécuriser une partie des données en ne donnant accès qu’à la fonction souhaitée (qui elle pourra accéder à ces données)
Exemple :
```
const FactoryFunction = string => {
  const capitalizeString = () => string.toUpperCase();
  const printString = () => console.log(`----${capitalizeString()}----`);
  return { printString };
};

const taco = FactoryFunction('taco');
taco.printString(); // Utilise le concept de closure et print"----TACO----"
```
Autre exemple :
```
const counterCreator = () => {
  let count = 0;
  return () => {
    console.log(count);
    count++;
  };
};

const counter = counterCreator();

counter(); // 0
counter(); // 1
counter(); // 2
counter(); // 3
```
**/!\ Pour sécuriser l’appel de fonction, le pattern factory function, nous permet de déclarer des private functions non utilisables à l’extérieur de la fonction (voir capitalizeString()). Seules les fonctions spécifiées dans le return sont utilisables à l’extérieur.**
