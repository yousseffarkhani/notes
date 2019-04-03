---
title: 14 - OOP Principles
created: '2019-04-02T21:32:26.194Z'
modified: '2019-04-03T12:07:47.153Z'
tags: [JS]
---

# 14 - OOP Principles
5 principes, appelés SOLID Design principles, régissent la programmation orientée objet :
## Single responsibility
Ce principe présente le fait qu’une classe (ou un module ou un objet) ne devrait avoir qu’une seule responsabilité.
Par exemple la logique de l’application devrait être gérée à part, les fonctions de MAJ et d’écriture dans le DOM également.
« Do one thing and do it well »
Où dessiner la ligne entre trop de choses dans une fonction et pas assez ?


Pour appliquer ce théorème et aider à l’organisation du code, plusieurs approches existent :
## Object Role Stereotypes
C’est une collection de rôles qui sont souvent utilisés dans une application. En établissant ces stéréotypes, nous avons accès à des templates d’objets permettant de décomposer le comportement d’une application et ainsir disposer d’une carte mentale.
Liste d’object rôle stereotypes :
- Information holder – Représente un objet désigné pour contenir des informations et les fournir aux autres objets
- Structurer - 
- Service provider - 
- Controller - 
- Coordinator -
- Interfacer – Un objet transformant l’information ou requête entre 2 parties distinctes du système.
  
Structurer – an object that maintains relationships between objects and information about those relationships.
Service provider – an object that performs specific work and offers services to others on demand.
Controller – an object designed to make decisions and control a complex task.
Coordinator – an object that doesn’t make many decisions but, in a rote or mechanical way, delegates work to other objects.
Interfacer – an object that transforms information or requests between distinct parts of a system.
## The Open/Closed Principle
Ce principe présente le fait que les modules JS devrait être extensibles mais fermés à la modification.
Il y a une règle simple à suivre, qui est que si je dois ouvrir le code du module pour faire des modifications alors ce module ne respecte pas le principe.
Exemple :
```
let iceCreamFlavors = ['chocolate', 'vanilla'];
let iceCreamMaker = {
  makeIceCream(flavor) {
    if (iceCreamFlavors.indexOf(flavor) > -1) {
      console.log('Great success. You now have ice cream.');
    } else {
      console.log('Epic fail. No ice cream for you.');
    }
  },
};
export default iceCreamMaker;
```
Ici, il est impossible de rajouter une flavor sans modifier le module donc il ne respecte pas le principe. Pour se faire :
```
let iceCreamFlavors = ['chocolate', 'vanilla'];
let iceCreamMaker = {
  makeIceCream(flavor) {
    if (iceCreamFlavors.indexOf(flavor) > -1) {
      console.log('Great success. You now have ice cream.');
    } else {
      console.log('Epic fail. No ice cream for you.');
    }
  },
  addFlavor(flavor) {
    iceCreamFlavors.push(flavor);
  },
};
export default iceCreamMaker;
```
## The Liskov Substitution Principle
## The Interface Segregation Principle
## The Dependency Inversion Principle
