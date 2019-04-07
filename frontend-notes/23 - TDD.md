---
title: 23 - TDD
created: '2019-04-03T12:34:45.251Z'
modified: '2019-04-03T13:03:27.664Z'
---

# 23 - TDD
## Introduction
TDD est un concept qui consiste à écrire des tests automatisés avant d’écrire le code.
Il existe énormément d’outils de test en JS (Mocha, jasmine, jest). La plupart sont similaires.

Nous allons ici parler de unit testing. 
Qu’est ce qu’une unit ? 
Une unit peut représenter un module, un component ou une fonction.
## Isolation
Il existe 2 types de fonctions/objets :
- Les fonctions/objets pures ayant pour but de retourner un résultat sans garder de trace du state.
- Les fonctions/objets ayant pour but de modifier/garder en mémoire un state et pouvant retourner un résultat. 
Un concept fondamental du test est l’isolation. Une méthode à tester devrait être indépendante et ne pas dépendre d’autres parties du code.
Les tests nous forcent à écrire des fonctions pures dans la mesure du possible.
## Mocking (& spies)
Mock functions sont des fakes functions qui sont souvent utilisées à la place de vrais appels vers une BDD ou vers internet. Ces fonctions retournent des fausses valeurs.
Pour déclarer un mock, il faut créer une class reproduisant les sorties souhaitées dans le fichier de test.
Il est également possible d’extend le service et overwrite les fonctions que nous souhaitons modifier. Cela est souhaitée lorsque la class est compliquée.
> NB : Un indice de mauvais code (non modulable) est la présence de nombreux mocks.

De même, il est possible d’utiliser un Spy. Un spy est une fonction/classe instanciée dont on modifie une variable ou une fonction de sorte à ce qu’elle retourne la valeur souhaitée.
Syntaxe :
`spyOn(‘objet’, ‘fonction à spy’).and.returnValue(false);`
Exemple: `spyOn(service, ‘isAuthenticated’).and.returnValue(false); // Dès que la fonction isAuthenticated de l’objet service sera appelée, la valeur de retour sera false.`
On peut également vérifier si la spy function est appelée grâce à `expect(service.authenticated).toHaveBeenCalled();`
## Mettre en place un test-runner (jest):
Jest est un outil de testing (utilisé par react).
Jest est composé d’une Assertion library et d’une mocking library.
Pour installer et lancer jest :
1. Installation : `npm i –save-dev jest`
1. Modifier le script de package.json par 
```     
"test": "jest",
"watch": "jest --watchAll"
```
1. Ecrire la fonction dans un fichier et l’exporter avec `module.exports = orderTotal;`
1. Ecrire les tests avec la syntaxe suivante :
```
it('Quantity', ()=> expect(orderTotal({
  items: [
    {
    'name': 'Dragon',
    price: 2,
    quantity: 1
    }
  ]
  }))
  .toBe(2)
);
```
