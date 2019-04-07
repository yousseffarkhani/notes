---
title: 20 - ES6 Modules
created: '2019-04-03T09:33:27.360Z'
modified: '2019-04-03T12:24:38.935Z'
tags: [JS]
---

# 20 - ES6 Modules
## Introduction
L’interêt d’écrire son code via des modules est la réutilisation simple de certains éléments.
Les modules permettent également d’éviter la pollution du global namespace.

### Processus :
1. Export : // a file called utils.js
- Named exports :
```
export function first (array) {
		return arr[0]
}
```
```
const functionOne = () => console.log('FUNCTION ONE!')
export { functionOne }
```
- Default exports :
```
export default function first (array) {
  return arr[0]
}
  ```
**Il y a 2 manières d’exporter : les named exports (si plusieurs fonctions doivent être exportées) ou default exports (export default functionOne)**
2. Import : // demo.js file
- Tout importer :
  `import * as utils from ‘./utils.js’ //utils.first([1,2,3])`
- Importer une fonction :
  `import {first} from ‘./utils.js‘ // first([1,2,3])`
- Importer depuis un default export :
  `import first from ‘./utils.js’;`
- Importer des named et default exports :
  `import first, { first } from ‘./utils.js’`
      
      
Il faut absolument déclarer `<script type="module" src="demo.js"></script>` dans le fichier html

## Histoire
ES6 modules est très récent. Jusqu’à présent, on utilisait des external module bundler.
1. Pour utiliser une librairie externe, il fallait la télécharger et la mettre dans le repo du site web afin de mettre un lien vers le script dans la page html.
       Le problème avec ce système est qu’il était chronophage de rechercher et MAJ la librairie.
2. Sont alors arrivés les package managers permettant d’automatiser le processus de téléchargement et de MAJ des librairies depuis un repo central. 
       En ce moment les package managers les plus populaires sont yarn et npm.
       A la base npm était un  package manager conçu spécifiquement pour node.js

#### Processus de création d’un package avec npm :
A) Se placer dans le repo et saisir : `npm init` → Cela permet de générer un fichier package.json qui stocke toutes les informations sur le projet
B) Installer un package (librairie) : `npm install "nom de la librairie" --save` → Cette commande permet de télécharger le code du package dans un fichier appelé node_modules et de créer une entrée dans package.json spécifiant la nouvelle dépendance (grâce à save).
C) Package.json est pratique car pour partager le projet, il n’y aura pas besoin d’envoyer tous les node_modules mais seulement package.json qui permettra à une autre personne d’installer les packages via la commande : `npm install`
D) Insérer  `<script src="node_modules/moment/min/moment.min.js"></script>` dans le fichier html.
E) Le problème de cette solution est qu’il faut ajouter toutes les librairies 1 par 1 dans le fichier html. C’est là qu’interviennent les module bundler (webpack)

3. Tous les langages de programmation permettent d’importer du code d’un fichier à un autre. Cependant, javascript ne permet pas cette fonctionnalité car JS a été designé en vue d’être utilisé depuis un navigateur et on ne souhaitait pas autoriser l’accès au système de fichier du client (pour des raisons de sécurité). 
       Dans un 1er temps la fonctionnalité a été développée pour node.js, grâce à au script suivant :
       `var moment = require('moment');` (node.js connaît la localisation des packages donc pas besoin d’écrire `require('./node_modules/moment/min/moment.min.js`)
       Cela fonctionne bien pour node.js car il a accès au fichier de l’ordinateur mais pas pour le JS standard. 
       C’est là qu’interviennent les module bundler (webpack). Le module bundler va recréer le fichier JS (ayant besoin de librairie externe) mais en y ajoutant toutes les dépendances directement dans le code.
       Pour installer webpack, il faut saisir : `npm install webpack webpack-cli --save-dev //save-dev sauvegarde cette dépendance dans l’environnement de développement spécifiquement`
#### Processus d’utilisation de webpack :
A) Pour utiliser webpack, saisir : `./node_modules/.bin/webpack index.js --mode=development`
B) Cette commande va lancer webpack sur notre fichier index.js et rechercher les « require » et les remplacer par le code approprié.
C) Le fichier final sera dans `dist/main.js`, il faut mettre à jour la source du script dans le html.
**/!\ `--mode=development` permet de garder le fichier lisible. `--mode=production` aurait pour effet de générer un fichier minifié.**
Avec l’utilisation de webpack, il faudra relancer les commandes décrites ci-dessus à chaque changement de index.js.
#### Pour éviter cela :
A) Créer un fichier webpack.config.js ayant pour contenu :
```
// webpack.config.js
module.exports = {
  mode: 'development',
  entry: './index.js',
  output: {
    filename: 'main.js',
    publicPath: 'dist'
  }
};
```
B) Lancer webpack avec `./node_modules/.bin/webpack`, cela va mettre à jour automatiquement le fichier `dist/main.js`
Pour éviter d’avoir à relancer cette commande à chaque modification, il est possible d’automatiser ce processus en utilisant des task runner.
Processus d’utilisation d’un task runner :
Les différentes tâches à réaliser sont : minifier le code, jouer les tests.
De nos jours, l’utilisation du script npm est le plus utilisé (par rapport à grunt et gulp).
Pour pouvoir lancer ces tâches, il faut rajouter à package.json :
```
"scripts": {
    "build": "webpack --progress --mode=production",
    "watch": "webpack --progress --watch"
}
```
`--progress` : permet de voir l’avancée en %
`--mode=production` : permet de minifier le code
`–watch` relance webpack à chaque modification d’un fichier JS.
Pour lancer les scripts :
- `npm run build`
- `npm run watch`

Il est également possible de mettre en place un serveur local en :
1. installant webpack-dev-server : `npm install webpack-dev-server –save-dev`
1. Ajoutant au fichier package.json : `"server": "webpack-dev-server --open`"
1. Lançant le serveur avec : `npm run server`

## Transpiler
Il s’agit d’un logiciel traduisant du code d’un langage à un autre. 
Les transpilers sont importants car souvent les navigateurs mettent du temps à implémenter de nouvelles fonctionnalités. De nouveaux langages sont crées afin de combler ces fonctionnalités. Le rôle du transpiler et de convertir ce nouveau langage en un langage compréhensible par le navigateur.
Exemples de transpiler :
- CSS : Sass, Less
- JS : CoffeeScript, babel (permet de convertir du « nouveau » javascript en ancien JS pour la compatibilité des anciens navigateurs), TypeScript
## NPM (Node package manager)
Npm est une base de données de packages.
https://docs.npmjs.com/


