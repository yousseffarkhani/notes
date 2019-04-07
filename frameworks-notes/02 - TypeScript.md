---
title: 02 - TypeScript
created: '2019-04-03T13:31:34.516Z'
modified: '2019-04-03T13:58:40.052Z'
---

# 02 - TypeScript
## Commandes 
Pour transpiler le fichier .ts en .js, il faut utiliser le compilateur tsc ‘fichier.ts’.
Pour que le fichier se mette à jour automatiquement : `tsc -watch ‘file.ts’`
Pour mettre à niveau le compilateur pour un fichier : `tsc ‘fichier.ts’ --target ES2015`
Pour automatiser cela à la création d’un projet (non angular), créer un fichier `tsconfig.json` et ajouter (il faudra p-t remplacer les ‘ par des ‘’ :
```
{
	‘compilerOptions’ : {
		‘target’: ‘esnext’,
		‘pretty’: true,
		‘outDir’: ‘./build’, //Dossier de destination
		‘watch’: true
	},
	‘files’: [
		‘demo.ts’ //Préciser les fichiers à transpiler
	],
	‘excluse’:[‘node_modules’]
}
```
Puis saisir: `tsc`
## Type annotations
### Standard
#### Déclaration du type d’une variable simple
En ajoutant :string derrière le paramètre d’une fonction, le compilateur va vérifier si le type de donnée passé en argument est le bon.
```
function greeter(person: string) {
return "Hello, " + person;
}
```
Cependant, même si le compilateur detecte une erreur le fichier .js est généré.
Lors de la déclaration d’une variable, il est possible de le faire implicitement également : 
`let num = 2;`
Si on change la valeur num à ‘hello’, on aura une erreur.
Pour que la variable puisse prendre différents types au cours de sa vie, on peut utiliser `let num: any`.
Pour signifier que le paramètre est optionnel il faut ajouter un ? → person?: string
Il est possible de déclarer 2 types avec | → x : number | string
#### Déclaration du type pour une fonction et callbacks
Pour les fonctions :
```
function pow(x: number, y: number): string (définit le format de sortie){
	return Math.pow(x,y)
}
```
Pour les callbacks: 
function forEach(arr: string[], cb: (arg: string) => void): void{…}
Il est possible d’utiliser … pour exploiter les arguments d’une fonction (cette opérateur retournera un  tableau donc il faut utiliser join):
```
function greet(prenom: string, ...nomDeFamille: string[]):string {
  let output = `Hello ${prenom}`;
  if(nomDeFamille) output += ` ${nomDeFamille.join(' ')}` 
  return output;
}
```
#### Déclaration du type pour un array
Pour les arrays:
```
const arr: number[] (définit le format des entrées) = [];
arr.push(1) //ok
arr.push(‘2’) //ko
```
Caster un type → `<string>x`
### Interfaces
Il est possible de déclarer le format d’objets en paramètre à l’aide d’interfaces:
```
interface Person {
    firstName: string;
    lastName: string;
}
```
```
function greeter(person: Person) {
    return "Hello, " + person.firstName + " " + person.lastName;
}
```
### Classes
```
class Student {
    public fullName: string;
    constructor(public firstName: string, public middleInitial: string, public lastName: string) {
        this.fullName = firstName + " " + middleInitial + " " + lastName;
    }
}
```
Il est possible de créer des propriétés/méthodes  public / private / protected permettant de l’utiliser dans le code.
Pour affecter/afficher une valeur à une propriété privée, il faut utiliser des getters / setters.
### Generics
Les génériques permettent de créer des composants qui peuvent fonctionner avec une variété de types.
Pour utiliser les génériques, il faut un moyen de capturer le type de l’argument de sorte à pouvoir l’utiliser pour la valeur de retour. 
#### Generic identity functions
Dans TS, il faut utiliser un type variable déclaré de la manière suivante: function identity<’nom de la variable’>(id : ’nom de la variable’) :’nom de la variable’{ ..}. 
Après la déclaration de la fonction, il est possible de l’utiliser de plusieurs manières :
    • `identity<string>(‘youssef’)` → Ici le type est assigné explicitement, l’entrée et la sortie doivent donc être de type string.
    • `identity(‘youssef’)` → Ici le type est assigné implicitement, la sortie devra être un string
Les observables par exemple sont génériques.
### Namespaces
Permet de créer un namespace, à l’intérieur de ce namespace il faudra exporer les classes à utiliser.
## Tips & tricks
Sélectionner une invocation de fonction et appuyer sur <kbd>F12</kbd> permet d’aller à la définition de la fonction.
Il est également possible de renommer un item en cliquant droit et ‘renommer le symbole’.

