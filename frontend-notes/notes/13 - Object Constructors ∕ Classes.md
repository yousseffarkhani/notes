---
title: 13 - Object Constructors / Classes
created: '2019-04-02T21:32:09.706Z'
modified: '2019-04-03T12:07:45.807Z'
tags: [JS]
---

# 13 - Object Constructors / Classes
## Déclarer un objet
1. 1ère méthode: Object litteral
```
const playerOne = {
  name: "tim",
  marker: "X"
}
```
2. 2ème méthode: Object constructor
Un constructeur démarre toujours par une majuscule
```
function Player(first, last) {
this.first = first,
this.last = last
}
const playerOne = new Player("Youssef", "fkh");
```


## Vérifier l’existence d’une propriété sur un objet
Pour savoir si une propriété appartient à l’objet initial ou au prototype, il existe 2 méthodes:
 - l’opérateur « in »
```
var school = {schoolName:"MIT"};
console.log("schoolName" in school);  // true
console.log("schoolType" in school);  // false car schooltype n’a pas été défini dans l’objet ni le prototype
console.log("toString" in school);  // true car cette méthode  été héritée du prototype.
```
 - La méthode hasOwnProperty
```
var school = {schoolName:"MIT"};
console.log(school.hasOwnProperty ("schoolName"));  // true
console.log(school.hasOwnProperty ("toString"));  // false car la méthode est héritée du prototype.
```
Pour énumérer les propriétés de l’objet, on peut utiliser :
```
for(let key in leo){
	if(leo.hasOwnProperty(key)){ // Cette condition permet de filtrer les propriétés appartenant au prototype (Pas sûr que cela soit nécessaire)
		console.log(`Key : ${key}, value : ${leo[key]}`); 
	} 
}
```
Vérifier qu’un objet est une instance d’une fonction :
`leo instanceof Animal; // retourne true ou false`

## Constructeurs
*“I want to create an object whose prototype is this other object”*
Avant ECMAScript 5, pour créer un objet possédant un prototype hérité d’un autre objet, il fallait créer une fonction constructeur pour réaliser cela :
```
var proto = {protoprop: "protoprop"};
function C() {} //Dummy throwaway constructor.
C.prototype = proto;
var obj = new C();
proto.isPrototypeOf(obj); //=> true
```
Depuis ECMAScript 5, cela s’est considérablement simplifié :
```
var proto = {protoprop: "protoprop"};
var obj = Object.create(proto);
proto.isPrototypeOf(obj); //=> true
```
Le comportement de ces constructeurs étant très incertains (voir https://tsherif.wordpress.com/2013/08/04/constructors-are-bad-for-javascript/), un pattern de création d’objet est nait : les factory functions.
## Factory function
Les factory functions sont similaires aux constructeurs mais au lieu d’avoir à utiliser le terme new pour créer un objet, les factory functions définissent et renvoient un nouvel objet à l’appel de la fonction.
Déclarer une factory function
```
function myObject(data) {
  var obj = Object.create(myObject.proto);
  obj.data = data;
  return obj;
}
myObject.proto = {
  getData: function() {
    return this.data;
  }
}
var o = myObject("data");
```
L’avantage de cette méthode est qu’elle ne nécessite pas l’utilisation de new et que les fonctions sont utilisées comme des fonctions.
Autre exemple de l’utilisation du factory function pattern :
```
const personFactory = (name, age) => {
  const sayHello = () => console.log('hello!');
  return { name, age, sayHello };
};
const jeff = personFactory('jeff', 27);
```
**/!\Object shorthand : return {name, age, sayHello} est équivalent à return {name: name, age: age, sayHello: sayHello}**
Cela est utile au moment de faire un console.log pour voir les clés et valeurs :
```
console.log(name, color, number, food) // Maynard red 34 rice
console.log({name, color, number, food}) // { name: 'Maynard', color: 'red', number: 34, food: 'rice' }
```

## Héritage
```
const Person = (name) => {
  const sayName = () => console.log(`my name is ${name}`)
  return {sayName}
}

const Nerd = (name) => {
  // simply create a person and pull out the sayName function!
  const {sayName} = Person(name)
  const doSomethingNerdy = () => console.log('nerd stuff')
  return {sayName, doSomethingNerdy}
}

const jeff = Nerd('jeff')
```
ou (si on veut hériter de toutes les méthodes)
```
const Nerd = (name) => {
  const prototype = Person(name)
  const doSomethingNerdy = () => console.log('nerd stuff')
  return Object.assign({}, prototype, {doSomethingNerdy})// L’opérateur assign permet de merger des éléments de la source à la target suivant ce pattern .assign(target, source, source2 ...)
}
```
## The module pattern
Un module est exactement la même chose qu’une factory function sauf que celle-ci est wrapped dans une IIFE. Ce pattern est à utiliser lorsque nous savons que nous n’allons pas créer plusieurs instances d’un objet :
```
const calculator = (() => {
  const add = (a, b) => a + b;
  const sub = (a, b) => a - b;
  const mul = (a, b) => a * b;
  const div = (a, b) => a / b;
  return {
    add,
    sub,
    mul,
    div,
  };
})();

calculator.add(3,5) // 8
```
## IIFE (Immediately Invoked Function Expressions
La raison principale d’utiliser des IIFE est d’avoir de la data privacy. Comme javascript fonctionne en block-scope, toutes les variables déclarées dans l’IIFE ne seront pas accessibles par l’extérieur.
```
(function () {
    var foo = "bar";
    console.log(foo); // Outputs: "bar"
})();

console.log(foo); // ReferenceError: foo is not defined
```
Il est également possible de passer un argument à la fonction :
```
var foo = "foo";
(function (innerFoo) {
    console.log(innerFoo); // Outputs: "foo"
})(foo);
```
Bien sûr il serait possible de réaliser la même chose avec une déclaration classique et un appel de la fonction. Cependant, l’IIFE présente l’avantage de :
1. ne pas polluer le global namespace (réduisant le risque de name collisions)
1. Montrer la volonté de n’utiliser qu’une seule fois cette fonction
1. Comme il n’y a pas de nom, elle ne peut pas être invoquée par erreur.



## Comment créer une Class ?
Les classes ont été introduites avec ES6. Il s’agit en réalité de syntaxical sugar (extensions du langage permettant de rendre plus agréable la lecture sans modifier la fonctionnalité).
Pour déclarer une classe :
- class declaration :
```
class Animal {
  constructor(name, energy){
    this.name = name;
    this.energy = energy; 
  }
  eat(amount){
    console.log(`${this.name} is eating`);
    this.energy += amount;
  }
}
let leo = new Animal("leo", 25);
```
La fonction eat est instanciée dans le prototype.
- Class expression:
```
let Rectangle = class {
  constructor(height, width) {
    this.height = height;
    this.width = width;
  }};
```
Une différence importante entre les functions declarations et class declarations/expressions est le hoisting. Cela veut dire qu’il faut absolument déclarer la classe avant de l’instancier.
/!\ Il n’y a pas de virgules entre les différents éléments de la classe.
Il est possible de créer de l’héritage grâce à extends. Exemple avec la classe Animal au-dessus:
```
class Dog extends Animal {
	constructor(name){
		super(name); // Appelle le constructeur de la classe parent
}
```
Extends ne permet d’étendre que les objets constructibles. Pour hériter des fonctions, il faut utiliser `Object.setPrototypeOf()` :
```
const Animal = {
  speak() {
    console.log(this.name + ' makes a noise.');
  }
};

class Dog {
  constructor(name) {
    this.name = name;
  }
}

Object.setPrototypeOf(Dog.prototype, Animal); // If you do not do this you will get a TypeError when you invoke speak
```
L’opérateur super permet d’appeler la méthode correspondante de la super-classe et permet ainsi de l’enrichir:
```
class Cat { 
  constructor(name) {
    this.name = name;
  } 
  speak() {
    console.log(`${this.name} makes a noise.`);
  }
}

class Lion extends Cat {
  speak() {
    super.speak();
    console.log(`${this.name} roars.`);
  }
}
```

