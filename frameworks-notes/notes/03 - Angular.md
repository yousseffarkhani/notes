---
title: 03 - Angular
created: '2019-04-03T13:32:06.577Z'
modified: '2019-04-03T14:48:42.712Z'
---

# 03 - Angular
## Commandes pour démarrer un projet angular
```
npm install -g @angular/cli
ng new mon-premier-projet
cd mon-premier-projet
ng serve --open
```
## Description des fichiers d’un projet angular
- E2E →Fichier de test end-to-end
- tslint.json → Référence toutes les règles de syntaxe à respecter.
- tsconfig.json → Précise les settings du compiler ts
- protractor.conf.js → end to end test
- karma.conf.js → unit testing
- src → Source folder 
	- app→ comprenant les components dans l’onglet app
	- assets → Dossier comprenant des fichiers non js/ts nécessaires à l’exécution du programme
## Building blocks
### Components : 
Description : UI code (header/footer/menu/calendar/…) composé d’une partie en HTML et une autre en code.
Composition du code :
- imports : importer les fonctionnalités requises
- decorators : meta-data sur le composant (@Component)
- class : code avec la donnée dans le HTML
### Services : Reusable code (front end calculation/validation/callbacks to the server). Définit du code pouvant être accessible depuis l’ensemble de l’application.
### Modules : Organize code. 
Les components et services seront registered/déclarés dans un module.
### Raccourcis
- Créer un composant : `ng generate component ‘nom du composant’` / `ng g c ‘nom du composant`’ 
- Service : `ng g s data` (s pour service)
- Déployer une application
  - Saisir : `ng build` → Cette commande va packager le code est le mettre dans un dist folder
  - Pour packager un code de prod, saisir : `ng build –prod`
  - Ouvrir un serveur virtuel : `http-server -o`
## Manipuler de la donnée
![Icon](@attachment/databinding.png)
Le databinding représente la communication entre le code TS et le template HTML.
L’interpolation et le property binding permettent de transmettre des informations du code TS vers le HTML.
### Interpolation binding
Consiste à transmettre une donnée du fichier ts au fichier HTML.
Le fichier HTML peut exploiter les données déclarées dans la class grâce à `{{ title }}`
### Property binding
Le property binding permet de modifier dynamiquement les propriétés d’un élément du DOM en fonction des données TS (lorsque la propriété est entre crochets).
Cela sert à transmettre des données aux éléments enfants et à affecter des propriétés en fonction de condition (disabled).
```
<button [disabled]="!isDisabled">Test</button>
<app-hero-detail [hero]="selectedHero"></app-hero-detail> // Consiste à transmettre des informations à un élément enfant.
```

/!\ `hero="selectedHero"` permet de binder le string `"selectedHero"` tandis que `[hero]="selectedHero"` permet de binder la variable `"selectedHero"` 
Pour cela, il faut utiliser le décorateur `@Input` en l’important (`import { Component, OnInit, Input } from '@angular/core'`)
Puis saisir `@Input() hero: Hero;`
Ce décorateur va créer une propriété hero qu’on peut fixer depuis la balise app-hero-detail.
Propriété intéressante : disabled, innerText
### Event binding
La convention est de nommer les méthodes avec un préfixe on.
`<button (click)="onClick($event)">`
Il est possible de passer en paramètre `$event`  référant à l’évenement. En fonction de l’évenement suivi, l’event sera différent.
Le type de l’event est Event, `onClick(event: Event){..}`
### Two way data binding
Ce type de binding permet de synchroniser la donnée et ce qui est montré à l’écran.
Ce binding est pratique pour les inputs.
`[(ngModel)]= ‘hero.name’` (ngModel est ce qu’on appelle une directive)
Pour l’utiliser, il faut importer le `FormsModule` dans l’AppModule en l’ajoutant aux imports (tableau contenant les modules externes requis) et en avec `import { FormsModule } from "@angular/forms";`
### Directive
Il existe 2 types de directives :
Les directives structurelles qui modifient la structure du document :
- `*ngFor` : répéte l’élement sur lequel il est appliqué pour tous les éléments d’une liste. 
`<app-appareil *ngFor="let appareil of appareils; let i = index"  [index] ="i"></app-appareil>`
- `*ngIf` : Affiche l’élément selon une condition
Les directives par attribut qui modifient le comportement d’élément existant :
- `ngModel` en est un exemple car elle modifie la valeur du `<input>`
- `ngStyle` permet d’appliquer des styles dynamiquement
`<button [ngStyle]="{backgroundColor:getColor()}">`
**/!\ A noter qu’il faut des {} autour de l’attribut car on passe un objet. Les attributs doivent être en CamelCase.**
- NgClass permet d’attribuer une ou plusieurs classe(s) à un élément. Class (utilisé dans l’exemple en dessous) ne permet de binder qu’une seule classe.
```
[class.’nom de la classe’] =’condition’ // La même syntaxe fonctionne avec style.
<li [ngClass]="{'list-group-item': true,
                'list-group-item-success': appareilStatus === 'allumé',
                'list-group-item-danger': appareilStatus === 'éteint'}">
[ngClass]="myClassObject" //  myClassObject est un objet déclaré dans le code TS.
<div [ngClass]="a === b ? 'class1' : 'class2'">
```
## Accès aux données Input/output
#### Inputs : 
Il est possible pour un élément enfant d’accèder aux données fournies par un élément parent. La syntaxe est la suivante :
`<app-appareil *ngFor="let appareil of appareils; let i = index"  [index] ="i"></app-appareil>`
/!\ Pour pouvoir l’utiliser, il faut déclarer `@Input()` index dans l’élément enfant
#### Outputs : 
Il est possible de partager de l’information de 2 manières :
- Event Emitter: Permet de remonter des informations d’un élément enfant vers un élément parent. Pour cela, il faut :
        ◦ Importer Output dans l’élément enfant (/!\ @angular/core)
        ◦ Déclarer l’output → `@Output() alert = new EventEmitter<string>();`
        ◦ Emettre l’output grâce à `this.alert.emit(‘hello’);`
        ◦ Dans le template de l’élément parent, déclarer 
          `<app-server (alert)="attention($event)"></app-server>`
        ◦ Déclarer une fonction à déclencher sur réception de l’output (ici attention).
        ◦ Exploiter la variable émise par l’enfant comme un event.
- Subjects: Il est également possible d’utliser des subjects pour transmettre des informations qui ne seront pas limitées à la relation parent/enfant, c’est d’ailleurs le pattern recommandé lorsqu’on a besoin de transmettre l’information à plusieurs composants (via un service).
        ◦ Déclarer un subject dans le service: `librarySubject = new Subject<Book[]>();`
        ◦ Déclarer une fonction permettant d’émettre l’information : `emitBook(){ this.librarySubject.next(this.library) } // this.library représente l’élément que nous souhaitons transmettre.`
        ◦ Souscrire au subject (utilisé dans un component ou un autre service, il faut donc l’importer via le constructor): `this.libraryService.librarySubject.subscribe(library => console.log(library);`
        ◦ Emettre le subject: `this.libraryService.emitBook(); // Tous les éléments ayant subscribe recevront l’information.`

## Services
Les services permettent de partager du code ou de la donnée entre plusieurs classes.
Les services ont été construit de manière à créer une couche d’accès à la donnée indépendante du modèle (pattern MVC).
Pour utiliser les services, il faut les injecter au code (voir Définitions-Dependecy Injection).
Les services peuvent être injectés à 3 niveaux :
- dans AppModule : la même instance du service sera utilisée par tous les composants de l’application et par les autres services
- dans AppComponent : tous les components auront accès au service mais pas tous les autres services
- dans un autre component: Le component lui même et tous ses enfants (tous les components qu’il englobe) aurant accès à la même instance du service mais le reste de l’application n’y aura pas accès.
Pour injecter un service, l’importer et l’intégrer dans le constructeur (cela permet de créer une instance de ce service).
## Routing
Les routes permettent de dire quelle vue (composant) afficher quand un utilisateur click sur une URL.
#### Ajouter une route dans le app-routing-module :
Une route dans angular est composé de : (ex : `{ path: 'heroes', component: HeroesComponent }`)
- un path : chemin d’accès
- component : component affiché quand l’utilisateur accède à ce path
`RouterModule.forRoot(routes)` → La méthode est appelée `forRoot()` car le routeur est configuré au niveau de l’application root. 
En parrallèle, il faut également ajouter dans l’app.component.html :`<router-outlet></router-outlet>`  
Cela permet au routeur de savoir où afficher les components.
#### Pour ajouter des liens de navigation : 
`<a routerLink="/heroes">Heroes</a>`
> routerLink est utilisé à la place de href car href a pour comportement de recharger la page. Or, on perd tout l’interêt d’une SPA.

#### Ajouter une classe à l’activation d’un lien (ici ‘active’)
`<a routerLinkActive="active" routerLink="/realisations">Réalisations</a>`

#### Interagir avec le navigateur :
Importer (`@angular/common`) et injecter le service suivant :
- `private location: Location` →Service permettant d’interagir avec le navigateur.
Pour revenir en arrière, `this.location.back();`
#### Rediriger le navigateur vers une vue :
Importer (`@angular/router`) et injecter le service router → `constructor(private router:Router)`.
Pour rediriger → `this.router.navigate([‘path’]);`
#### Créer des paramètres de routes
Pour configurer des routes disposant de paramètres, il faut ajouter dans les Routes `{ path: ‘appareil/:id’, component: SingleAppareilComponent}`.
L’utilisation des: ici signifie que tous les chemins en appareils/* seront redirigés vers notre component.
#### Extraire le paramètre d’une route:
Importer (`@angular/router`) et injecter le service suivant:
- `private route: ActivatedRoute` →Détient les informations concernant une route. L’ID sera extrait d’ici.
      `+this.route.snapshot.paramMap.get('id');` //+permet de convertir en nombre
      `route.snapshot` est une image statique de la route après que le composant ait été crée.
- Alternative: `+this.route.snapshot.params['id'];`
Pour attribuer l’id au routerLink: `<a [routerLink]="[id]">`
#### Rediriger l’utilisateur (erreur 404):
Créer un composant affichant le message voulu puis dans les Routes, ajouter une route directe et une route ‘wildcard’ redirigant toute route inconnue vers la route directe.
Route wildcard → `{path: ‘**’, redirectTo: ‘not-found’}`
#### Executer du code avant qu’un utilisateur accède à une route
Par exemple pour vérifier l’identité ou authentifier l’utilisateur et le rediriger vers la page d’authentification en cas d’utilisateur non-authentifié.
Une guard est un service qu’Angular executera au moment où l’utilisateur essaye d’accéder à une route.
Voir https://openclassrooms.com/fr/courses/4668271-developpez-des-applications-web-avec-angular/5088826-gerez-la-navigation-avec-le-routing
## Rxjs
*Voir définitions - Observables*
Pour utiliser les observables de la classe Rxjs, il faut 
`import { Observable, of } from 'rxjs';`
Ensuite créer la méthode :
```
getHeroes(): Observable<Hero[]>{
  this.messageService.add('HeroService: fetched heroes');
  return of(HEROES);
}
```
Puis pour exploiter cette méthode :
```
this.heroService.getHeroes()
  .subscribe(heroes => this.heroes = heroes);
```
## Forms
En Angular, il y a 2 grandes méthodes pour créer des formulaires :
- La méthode template
- la méthode reactive

Voir https://openclassrooms.com/fr/courses/4668271-developpez-des-applications-web-avec-angular/5090131-ecoutez-lutilisateur-avec-les-forms-methode-reactive
## Http Services
Le module HttpClient est le service permettant à angular de communiquer avec un serveur via HTTP. (voir définitions – Observables pour plus de détails)
1. `import { HttpClientModule }    from '@angular/common/http'; //app.module.ts`
1. Déclarer le module dans le constructeur : `constructor(private http: HttpModule)`
1. Pour l’utiliser:`this.http.get<type de donnée>(‘URL’)` cela retourne un observable.
Pour attraper les erreurs, il faut importer
`import { catchError, map, tap } from 'rxjs/operators';`
et utiliser la méthode pipe() sur l’observable avec l’opérateur catchError()→
```
this.http.get<Hero[]>(this.heroesUrl)
    .pipe(
      catchError(this.handleError('getHeroes', []))
    );
}
```

## AppModule
Correspond au NgModule principal.
Angular a besoin de ce module pour savoir comment l’application est construite, quelles sont les fichiers et librairies requis pour le fonctionnement.
Les modules sont importés, les composants/pipes/directives sont déclarés, les services sont ‘providés’ (fournis).
## Life Cycle hooks
`ngOnInit` → Cette méthode est lancé peu de temps après la création d’un objet
Très souvent utilisé pour initialiser les données une fois le component crée.
## Pipes
Les pipes sont utiles pour formatter les strings, currency, dates. (https://angular.io/api?type=pipe)
Il existe des fonctions utilisables directement sur les données importées (interpolation binding) dans le HTML (ex: | uppercase/date/currency)
Les pipes sont également utile pour gérer les éléments asynchrones grâce au pipe async:
```
<p>{{lastDate | async}}</p> // Affiche la date au bout de 2 sec
lastDate = new Promise((resolve, reject) =>
  setTimeout(() => {
    resolve(new Date())
  }, 2000));
}
```

## Définition
### Décorateurs
Les décorateurs permettent d’attacher des méta-données à un objet (classes/propriétés/paramètres)
Decorator functions: `@Component` /` @Injectable`
### Bootstrapping
Le champs bootstrap (dans le NgModule principal) permet de dire à angular quel sera le composant inséré dans le HTML

### NgModules
`@NgModule` est un décorateur de classes NgModules.
Ces classes permettent de dire quels sont les éléments nécessaires au fonctionnement de l’app: modules requis, composants/directives/pipes, services.
Un NgModule ne peut exporter que les classes dont il dispose ou sont importés.
### Browser module
Ce module donne accès à des services essentiels pour lancer l’app sur un browser (DOM rendering). 
Ce module permet d’importer avec lui le CommonModule permettant au composant d’utiliser les directives ngIf/ngFor.
Le browser module module ne doit être importé qu’une seule fois.
### Dependency injection (DI)
Les dependecy sont des services ou des objets dont une classes a besoin pour que ses méthodes soient valides.
`@Injectable()` est le décorateur d’injection. Il est placé dans le service avant la déclaration de la classe et signifie que le service peut être injecté.
Pour injecter ce service dans un service ou component, il faut l’importer et le déclarer dans le constructeur (ce qu’on appelle aussi injection). → constructor(heroService: HeroService)
Quand angular va créer le composant, le dependecy injection system va associer heroService au singleton HeroService.
/!\Pour exploiter le service dans le template HTML, il faut le déclarer en public dans le constructeur.
ProvidedIn : ‘root’ précise la portée de l’injection. Au niveau ‘root’, seule une instance de ce service sera créée et injecté.
### In Memory Web Api
Il s’agit d’un module permettant d’intercepter les requêtes http de l’appli (réalisée avec httpClient) afin de renvoyer une réponse similaire à celle d’un serveur.
Ce module est pratique pour simulé un serveur en début de projet.
Plus d’informations : https://github.com/angular/in-memory-web-api

### Observable
#### Introduction
Les observables sont similaires aux promises. Angular utilise les observables afin de gérer des opérations asynchrones (AJAX, event utilisateur).
Un observable est un objet émettant des informations auxquelles on souhaite réagir.
Les observables permettent de transmettre des messages entre publishers et subscribers(/observer) dans l’application.
A cet observable, on associe un observer (un bloc de code qui sera executé à chaque fois que l’observable émet un message).
L’observable émet 3 types d’informations : des données, une erreur, un message ‘complete’. Du coup un observer peut avoir jusqu’à 3 fonctions pour réagir à chaque information.
Observable.subscribe(fonction 1, fonction 2, fonction 3). Les 2 dernières fonctions étant optionnelles.
La 1ère fonction va recevoir les données comme argument, la 2ème l’erreur (s’il y en a une), la 3ème se déclenche mais ne reçoit pas d’argument.
/!\ La souscription à un observable dure jusqu’à l’infini (fin du programme) tant que l’observable émet des données. Cela peut causer des bugs. C’est pourquoi il faut stocker cette souscription dans une variable afin que Angular la supprime lorsque nous n’en avons plus besoin. (voir subscriptions – https://openclassrooms.com/fr/courses/4668271-developpez-des-applications-web-avec-angular/5089331-observez-les-donnees-avec-rxjs).
#### Subjects
Un subject est un type d’observable qui permet non seulement de réagir à de nouvelles informations, mais également d’en émettre. L’interêt étant de centraliser la donnée dans un service et en cas de changements le subject émettra un message à l’ensemble des observers.
Cela permet d’avoir une complète abstraction entre le service et les components. (voir Subjects - https://openclassrooms.com/fr/courses/4668271-developpez-des-applications-web-avec-angular/5089331-observez-les-donnees-avec-rxjs)

Pour la suite, il faut
`import { Observable, of, from } from 'rxjs';`
#### Déclarer un observable
Il existe plusieurs méthodes :
- Observables de tableau
```
const  observable = of([1, 2, 3]) // méthode la plus utilisée
const  observable = from([1, 2, 3])
const myObservable3 = interval(2000);
```
Pour convertir une variable en observable, il faut utiliser la méthode of(‘variable’).
- Osbervables de fonction
```
Observable.create(observer => { 
  observer.next(42);
  observer.error(43);
  observer.complete(); //Le flux est terminé et il n’y a plus rien à émettre
});
```
- Observables de DOM

Use-case : Ecouter un champ input et réaliser une requête
```
import { Observable } from 'rxjs/Rx';

const img = document.querySelector('img');

let inputValue;

let keyUp = Observable.fromEvent(document.querySelector("input"), 'keyup');
let observable = keyUp
  .map(e => inputValue = encodeURIComponent(e.target.value)) // Permet de récupérer les valeurs de l'input
  .debounceTime(500) // Permet d'attendre que l'input soit stabilisé avant de réaliser la requête (afin d'éviter les spams de touches)
  .distinctUntilChanged() //Permet d'ignorer les requêtes identiques aux requêtes précédentes
  .map(value => {
  return Observable.fromPromise(fetch("https://cataas.com/cat/"+inputValue))// Transforme la promise en observable
  .retry(3) // Permet de recommencer la requête 3 fois avant d'émettre une erreur
  .takeUntil(keyUp) //Je récupère les valeurs à moins qu'un autre observable émette une valeur
  }).concatAll(); // ???

observable
  .subscribe(
    nb => img.src = nb.url, 
    err => console.log(err), 
    () => console.log("completed")
  );
```

Il existe de nombreux autres opérateurs (throttletime(), scan(), reduce(), filter(), ..).
!\ Pour les utiliser, il faut les importer (from ‘rxjs/operators’).

#### Exploiter les données d’un observable
Pour exploiter les données des observables 
```
obervable.subscribe(
  nb => console.log(nb), 
  err => console.log(err), 
  () => console.log("completed")
);
```
Dans cet exemple, la 1ère fonction sera appelée tant qu’il y aura des valeurs à afficher. La 2ème sera appelée en cas d’erreur et la 3ème lorsqu’il n’y aura plus de valeur à afficher.
