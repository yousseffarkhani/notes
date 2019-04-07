---
title: 09 - Arrays
created: '2019-04-02T21:15:24.549Z'
modified: '2019-04-03T12:07:34.296Z'
tags: [JS]
---

# 09 - Arrays
Accéder à l’ensemble des éléments de l’array 
```
var cars = ["Saab", "Volvo", "BMW"];
document.getElementById("demo").innerHTML = cars; // affiche l’ensemble de l’array séparé par des virgules.
```
Les array ne sont pas un objet dans JS. 
Pour accéder à une valeur contenue dans un array, il faut utiliser un nombre représentant son index.
(Pour accéder à une valeur contenue dans un objet, il faut utiliser un nom (clé) pour y accéder. Dans d’autres langages cela est permis, cela s’appelle hashes/associative arrays).
On peut tout stocker dans les arrays : fonctions, objet, array, nombre, string, …
#### Spread operator
`let arrayWord = [...string]` → permet d’attribuer chaque lettre du string à un élément du tableau
#### Propriétés
`.length`

#### Methodes
`.sort()`
`.forEach()`
`.push(x)` > Retourne la longueur du nouvel array
`.pop()` > Retourne la valeur de l’élément « popé »
`.shift()` > Retire le 1er élément et retourne la valeur supprimée
`.unshift(x)` > Rajoute x en 1er élément et retourne la longueur du nouvel array
`.splice(a,b,c,d)` > a défini la position de départ. B défini la position de fin (supprime). C et d indique les nouveaux éléments à rajouter. 
`.slice(1,2)` > Cette méthode ne modifie pas l’array sur lequel elle est appliquée
La méthode retourne un array avec les éléments supprimés.
`.concat(x,x,x,x)` > Cette méthode ne modifie pas l’array sur lequel elle est appliquée
`delete fruits[0]` > Supprime le 1er élément et laisse un trou (undefined). Il est préférable d’utiliser shit ou pop à la place.
`.toString()` > Sépare les valeurs par une virgule / Pour spécifier le séparateur il faut utiliser .join(« x »)
`.isArray()` renvoie true lorsque l’objet est un array
`.indexOf(x[,y])` > renvoie l’indice pour lequel on trouve x dans le tableau. Y permet de spécifier à partir de quel index la recherche démarre. Cette méthode renvoie -1 si rien n’a été trouvé.

###### Filter
La méthode filter crée et retourne un nouveau tableau contenant tous les éléments du tableau d’origine qui remplissent une condition déterminée.
**/!\ Filter ne modifie pas le tableau sur lequel il est appliqué.**
Exemple :
```
var words = ['spray', 'limit', 'elite', 'exuberant', 'destruction', 'present'];
const result = words.filter(word => word.length > 6);
```
###### map
La méthode map crée un nouveau tableau avec les résultats de l’appel d’une fonction sur chaque élément du tableau.
Exemple 1 - Multiplier tous les éléments du tableau par 2 :
```
var array1 = [1, 4, 9, 16];
const map1 = array1.map(x => x * 2);
```
Exemple 2 - Créer un tableau ne contenant que les valeurs souhaitées:
```
const fullNames = inventors.map(inventor => [inventor.first, inventor.last])
console.table(fullNames);
```
Exemple 3 - (Generic map) Créer un tableau ne contenant que les valeurs souhaitées:
```
var elems = document.querySelectorAll('select option:checked');
var values = Array.prototype.map.call(elems, function(obj) {
  return obj.value;});
```
###### sort
Sort trie les éléments d’un tableau directement sur ce même tableau et renvoie le tableau.
Exemple 1 – Trie sur les mots :
```
var months = ['March', 'Jan', 'Feb', 'Dec'];
months.sort();
```
Exemple 2 – Trie sur les chiffres
```
var array1 = [1, 30, 4, 21];
array1.sort();
```
Exemple 3 – Trie sur les chiffres
```
var nombres = [4, 2, 5, 1, 3];
nombres.sort(function(a, b) {
  return a - b;});
```
Exemple 4 – Trier les éléments d’un tableau :
```
let liste = ['Delta', 'alpha', 'CHARLIE', 'bravo'];
let mapped = liste.map(function(string, index) {
	return {index: index, value: string.toLowerCase()}})
mapped.sort((a,b) => a.value > b.value? 1: -1)
let result = mapped.map(obj => liste[obj.index]);
```
###### Reduce
La méthode reduce applique une fonction callback à chaque élément de la liste. Le retour de cette fonction est stocké dans un accumulateur. L’objectif est de réduire la liste à une seule valeur. 

Syntaxe : `arr.reduce(callback[, valeurInitiale])`
La fonction callback dispose de 4 éléments en entrée : l’accumulateur (la valeur précédemment retournée par la callback ou valeur initiale si elle est définie), la valeur courante, l’index et le tableau.
La valeur initiale est utilisé comme 1er argument lors du 1er appel de la fonction callback

Exemple 1 – Somme d’un tableau d’objets :
```
var valeurInitiale = 0;
var somme = [{x: 1}, {x:2}, {x:3}].reduce(
    (accumulateur, valeurCourante) => accumulateur + valeurCourante.x
    , valeurInitiale);
```
Exemple 2 – Somme d’un tableau :
```
let sumYearsLived = inventors.reduce((prev, curr) => prev + (curr.passed - curr.year), 0)
```
Exemple 3 – Somme des éléments d’un tableau :
```
let count = data.reduce((prev, curr) => {
if(!prev[curr]) {
prev[curr] = 0;
}
prev[curr] ++;
return prev;
}, {})
```
###### Some
La méthode some() teste si au moins un élément du tableau remplit la condition énoncée dans la fonction fournie.
La valeur de retour sera true si le test est ok.
La méthode some ne modifie pas le tableau sur lequel elle est appliquée
Exemple 1 : 
```
    let olderThan19 = people.some(person => {
      let date = new Date().getFullYear();
      return date - person.year >= 19;
    })
```
###### Every
Cette méthode permet de vérifier que l’ensemble des éléments du tableau vérifie une condition
Exemple : 
```
    let allAdults = people.every(person => {
      let date = new Date().getFullYear();
      return date - person.year >= 19;
    })
    console.log({allAdults});
```
###### Find
Cette méthode ressemble à filter sauf qu’elle retourne un élément et non un tableau contenant le 1er élément remplissant la condition.
Exemple :
```
  let id = comments.find(comment => comment.id === 823423)
```
###### findIndex
La méthode findIndex ressemble à find mais donne en retour l’index de l’élément remplissant la condition.
Exemple :
```
    let index = comments.findIndex(comment => comment.id === 823423);
    comments.splice(index,1);
```
###### Destructuring an array
Spread operator: Cela permet de tout « éclater » dans un tableau
Exemple : 
```
	let item = "youssef"
	let array = [...item]
```
array Destructuring:permet de choisir ce que l’on souhaite extraire.
Exemple :
```
	const [a, b] = [1, 2, 3, 4] // a = 1, b = 2
	const [a, b,,c] = [1, 2, 3, 4] // a = 1, b = 2, c = 4
```
Cela peut permettre d’inverser des valeurs 
```
	let a = 2;
	let b = 3;
	[a, b] = [b, a];
```
Destructuring with spread operator :
```
	const [a,,...arr] = [1, 2, 3, 4, 5] // a = 1, arr = [3, 4, 5]
```
###### Astuce
 - Console.table
