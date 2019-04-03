---
title: 10 - Objects
created: '2019-04-02T21:31:08.659Z'
modified: '2019-04-03T12:07:36.935Z'
tags: [JS]
---

# 10 - Objects

Un objet contient de nombreuses données. 2  types de données existent soit des données pures, soit des fonctions. Ces types sont appelés propriétés et méthodes.

Un objet est crée avec des {..} avec une liste de propriétés. 
Les propriétés sont au format "key:value" où key est un string (aussi appelé property name / identifier) et value peut être n’importe quoi.
##### Déclarer un objet
Pour déclarer un objet, il existe 2 méthodes :
 - `let user = new Object(); // "object constructor" syntax`
 - `let user = {};  // "object literal" syntax`
La 2ème méthode est recommandée.
##### Ajouter des propriétés 
```
let user = {     // an object
  	name: "John",  // by key "name" store value "John"
  	age: 30,      // by key "age" store value 30
    "likes birds": true
};
user.isAdmin = true;
```
 - Il est possible d’ajouter des property names composé à l’aide des « » dans la déclaration de l’objet
      
##### Supprimer une propriété
 - `delete user.age;`
##### Cibler une propriété
 - `user["likes birds"]`
 - `let key = "likes birds"  `
      `user[key] = true`
##### Computed properties
```
let fruit = prompt("Which fruit to buy?", "apple");
let bag = {
 	 [fruit]: 5, // the name of the property is taken from the variable fruit
};
alert( bag.apple ); // 5 if fruit="apple"
```
```
let bag = {
 	 [fruit + 'Computers']: 5 // bag.appleComputers = 5
};
```
##### Property value shorthand
Au lieu d’écrire :
```
function makeUser(name, age) {
  return {
    name: name,
    age: age
    // ...other properties
  };
}
```
Il est possible en javascript de simplement écrire :
```
function makeUser(name, age) {
  return {
    name, // same as name: name
    age   // same as age: age
    // ...
  };
}
```
ou 
```
let user = {
  name,  // same as name:name
  age: 30
};
```
##### Destructuring object
Permet d’assigner directement à une variable une propriété d’un objet.
```
const dog = {
	name: 'john',
	age: 25,
	hobbies:{
		sport: 'basketball',
		food: 'chewing gum'
  }
}
const {name} = dog // permet de récupérer uniquement le name.
const {name: value, age: old, hobbies: { sport: thisOne}} = dog
```
##### Existence check
En Javascript , si on tente d’accéder à une propriété qui n’existe pas, aucune erreur ne sera remonté (il sera juste indiqué que la valeur est undefined).
Pour vérifier qu’une propriété existe vraiment dans un objet, il faut utiliser la syntaxe :
`"key" in object` ("key" étant le nom de la propriété) /!\ les guillemets sont requis sinon JS va rechercher le nom dans une variable.
Cetre syntaxe retourne vraie ou faux.
Pour parcourir l’ensemble des objets, il existe la boucle `for .. in ..` > `for(let key in object)`.
##### Ordre des clés
Les clés contenues dans un objet sont triées dans l’ordre de création.
Cependant, si les clés contenues dans un objet correspondent à des integer ou des string ne contenant que des chiffres alors ces clés seront triées par ordre croissant.
Pour contourner ce problème il faut rajouter un + devant le nombre (« +39 »)
##### Copie par référence
La différence fondamentale entre un objet et une valeur primitive est que les primitives sont assignées/copiées entièrement. (`let message= "hello"`) tandis que dans la déclaration d’un objet seule l’adresse mémoire est assignée.
Un objet déclaré avec const peut être modifié.

Pour cloner un objet, il existe plusieurs méthodes :
```
let clone = {}; // the new empty object
// let's copy all user properties into it
for (let key in user) {
  clone[key] = user[key];
}
```
`Object.assign(dest[, src1, src2, src3…])` permet de copier tous les objets dans dest et retourne dest.
`let clone = Object.assign({}, user);`

Un problème apparaît lorsque l’objet contient un autre objet car on ne souhaite pas copier sa référence mais son contenu. Les méthodes vues jusqu’à présent ne permettent que de cloner un objet contenant des primitives. 
Le processus de cloner un objet dans sa globalité (même s’il contient d’autres objets) s’appellent le deep cloning. La librairie lodash permet cela grâce à la méthode `.cloneDeep(obj)`.

