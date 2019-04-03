---
title: 15 - DOM Manipulation
created: '2019-04-03T09:31:43.463Z'
modified: '2019-04-03T12:07:48.428Z'
tags: [JS]
---

# 15 - DOM Manipulation
## Introduction
L’interêt principal de javascript est qu’il est capable d’accèder et modifier les éléments du DOM.
Le DOM est une représentation en forme d’arbre du contenu d’une page web avec différentes relations entre chacun des éléments de l’arbre (parent, child, sibling).
En réalité le code HTML est parsé par le navigateur  et converti en DOM. L’interêt étant que les nœuds deviennent des objets disposant de propriétés et méthodes. Ce sont ces propriétés et méthodes qui vont nous permettre de manipuler une page web avec javascript.

## Selectors
Pour pouvoir manipuler un élément, il faut le sélectionner. 
A la manière de CSS, il existe différents moyens de cibler un nœud :
- sélecteur CSS : `div.className`
- relational selectors : `querySelector(container.firstchild / controls.previousSibling/ ...);`

Query Selectors :
- `document.querySelector(selector)` retourne le HTML de l’objet
- `document.querySelectorAll(selector)` retourne une nodelist contenant l’ensemble des propriétés de l’élément sélectionné. 

/!\ une nodelist n’est pas un array et ne dispose donc pas des mêmes propriétés / méthodes. (Il faut donc le convertir en array si  l’on souhaite)
`console.dir(x)` : permet d’afficher les propriétés d’un éléments x.

## Methods
- `document.createElement(‘div’);`
- `parentNode.appendChild(childNode)`
- `parentNode.insertBefore(newNode, referenceNode)`: ajoute newNode avant referenceNode
- `.remoteChild(child)`
- `.replaceChild(newNode, nodeToBereplaced)`
- `div.setAttribute('style', 'color: blue; background: white');`: permet de modifier les attributs d’un élément.
- `div.setAttribute('id', 'theDiv');` : if id exists update it to 'theDiv' else create an id with value "theDiv"
- `div.getAttribute('id');` : returns value of specified attribute, in this case "theDiv"
- `div.removeAttribute('id');` : removes specified attribute
voir https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes pour la liste des attributs.

Modifier les classes
- `div.classList.add('new');` : adds class "new" to your new div
- `div.classList.remove('new');` : remove "new" class from div
- `div.classList.toggle('active');` : if div doesn't have class "active" then add it, or if it does, then remove it

La dernière solution est souhaitée.
- Ajouter du texte à l’élément : `div.textContent = ‘heelo’;`
- Ajouter du html dans un élément : `div.innerHTML = '<span>Hello World!</span>';`   
## Events
Les events sont des actions telles que des clicks / dbl clicks. En utilisant Javscript, nous pouvons faire en sorte d’écouter ces éléments sur une page web et réagir à ça.
Les events trackables sont multiples: (liste exhaustive: https://www.w3schools.com/jsref/dom_obj_event.asp)
- click
-     dblclick
-     keypress
-     keydown
-     keyup
## Event listeners
### Individuel
Il existe 3 moyens de créer des event listeners:
- `<button onclick="alert('Hello World')">Click Me</button>`

Solution pas souhaitée car du js est mélangé à du html
- HTML: `<button id="btn">Click Me</button>`
JS: 
```
var btn = document.querySelector('#btn');
btn.onclick = () => alert("Hello World");
```
Solution plus optimal. Cependant, un élément ne pourra contenir qu’1 onclick property
- HTML: `<button id="btn">Click Me</button>`
JS: 
```
var btn = document.querySelector('#btn');
btn.addEventListener('click', () => {
 	alert("Hello World");
});
```
**Solution recommandée**

A noter qu’on peut directement appeler des fonctions pour les 3 méthodes :
- `<button onclick="alertFunction()">CLICK ME BABY</button>`
- `btn.onclick = alertFunction`
- `btn.addEventListener('click', alertFunction);`
### Multiple
HTML :
```
<div id="container">
    <button id="1">Click Me</button>
    <button id="2">Click Me</button>
    <button id="3">Click Me</button>
</div>
```
JS :
```
// buttons is a node list. It looks and acts much like an array.
const buttons = document.querySelectorAll('button');

// we use the .forEach method to iterate through each button
buttons.forEach((button) => {
  // and for each one we add a 'click' listener
  button.addEventListener('click', (e) => {
    alert(button.id);
  });
});
```
### Accéder aux informations de l’event
Il faut placer un paramètre dans la fonction appelée : 
```
btn.addEventListener('click', function(e){
  console.log(e)
})
```
e dispose de nombreuse propriétés et fonctions. Les plus utiles sont :
- `.target` > renvoie le node clické (il est possible de changer les attributs avec `e.target.setAttribute(…);`



