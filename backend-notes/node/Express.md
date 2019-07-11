---
title: Express
created: '2019-04-04T09:42:52.038Z'
modified: '2019-04-04T14:57:18.515Z'
---

# Express
## Introduction
Express est un framework basé sur le module http de node.js.
Express permet de:
- Délivrer du contenu statique
- Construire des API
- Se connecter à des sources de données

Express facilite:
- Le parsing du corps des requêtes HTTP
- Le parsing des cookies
- La gestion des sessions
- L’organisation des routes (aussi appelés end point)
- La détermination du data type à utiliser dans le header de la réponse
- Le parsing des querystring et paramètres d’URL
- La mise en place de headers dans les réponses de manière automatiques
- L’organisation du code
- Donne accès à pleins de plugins (appelés middleware)
- L’authentification, la validation
- La construction de web applications
- La construction de RESTful API servers

Les components sont appelés middleware et sont les pierres angulaires de ce framework.
## Commandes
### Initialisation:
```
const express= require(‘express’);
const app = express()
app.get(‘/’, (request, response) => { … }
app.listen(3000);
```
### Structure d’un projet express: (le module express-generator permet de créer cette structure automatiquement)
`express ‘nom du projet’` → Génère la structure du projet automatiquement.
#### Fichiers/dossiers:
- `app.js`: fichier principal, le serveur et la logique de l’application se situent ici.
- `/public`: contient tous les fichiers statiques du serveur. Peut être initialisé grâce à express.static(‘PATH’)(voir dans NPM module middlewares)
- `/routes`:contient toutes les routes API (s’il s’agit d’un REST API server). Permet d’améliorer la lisibilité du projet en exportant et regroupant toutes les routes ayant le même objet.
      - `App.js`: `app.use('/api/members', require('./routes/api/members'))`
      - `members.js`: `const router = express.Router();`
      - `router.get('/', (req, res) => res.json(members))` // Il n’est pas nécessaire de repréciser api/members
      - `module.exports = router;`
- `routes/users.js`: contient les endpoints pour les ressources (comme les users)
- `/views`: contient des templates
- `package.json`
- `/www`: script de démarrage
- `config`: contient les éléments de configuration (adresse de la DB, MDP, ...)

/!\ On différencie les web applications nécessitant du server-side-rendering et les REST APIs avec un serveur de données (le rendering a lieu en front directement chez le client).
#### Structure de `app.js` (fichier principal):
1. Imports
`const express = require(‘express’);`
2. Instantiations
`const app = express();`
3. Configurations
4. Middleware
`app.use(express.json());`
5. Routes
`app.get(...)`
6. Error handlers
7. Server bootup / server export
`app.listen(3000);`

Le serveur express peut être configuré via la méthode set:
```
app.set(‘port’, process.env.PORT || 3000);
app.set(‘views’, ‘templates’); // Ici on paramètre le template directory (dossier avec les server-side templates) pour que ce soit templates au lieu de views
app.set(‘view engine’, ‘jade’);
```
#### Request
L’objet request dispose de plus de propriétés dans le cadre d’express que pour son équivalent avec http:
- `request.params`: URL parameters (members/:id → `request.params.id`)
      /!\ Cette valeur est retournée sous format string. Il faut la convertir en nombre si besoin.
- `request.query`: query string parameters
- `request.route`: current route as a string
- `request.cookies`: cookies, requires cookieParser
- `request.signedCookies`: signed cookies, requires cookie-parser
- `request.body: body/payload`, requires body-parser (`app.use(express.json())` / Si utilisation d’un formulaire: `app.use(express.urlencoded({extended: false})`)
- `request.headers`: headers

Il existe également des raccourcis pour accéder aux propriétés du header d’une requête:
- `request.get(headerKey)`: value for the header key
- `request.get(‘host’)`: value for the header key
- `request.accepts(type)`: checks if the type is accepted
- `request.acceptsLanguage(language)`: checks language
- `request.acceptsCharset(charset)`: checks charset
- `request.is(type)`: checks the type
- `request.ip`: IP address
- `request.ips`: IP addresses (with trust-proxy on)
- `request.path`: URL path
- `request.hostname`: host without port number
- `request.fresh`: checks freshness
- `request.stale`: checks staleness
- `request.xhr`: true for AJAX-y requests
- `request.protocol`: returns HTTP protocol
- `request.secure`: checks if protocol is https
- `request.subdomains`: array of subdomains
- `request.originalUrl`: original URL
### Response
L’objet response dispose dans le cadre d’express de certaines méthodes en plus de celle de http (statusCode/writeHeader/end/write):
- `response.redirect(url)`: redirect request
- `response.send(data)`: send response, the content-type is automatically determined
- `response.set(‘Content-Type’, ‘text/plain’)`
- `response.json(data)`: send JSON and force proper headers
- `response.sendfile(path, options, callback)`: send a file
- `response.render(templateName, locals, callback)`: render a template
- `response.locals`: pass data to template
### Status Code
Pour spécifier un status Code, il est possible de modifier la propriété `response.status(200)` ou renvoyer la réponse directement via `response.sendStatus(204)`.
### Routing
**Express interdit automatiquement l’accès aux autres routes (ex: localhost:3000/transaction)**
Get
```
app.get('/profile', (req, res) => {
	if(req.query.id) res.send(profile[req.query.id]); 
	else res.send(profile);
})
```
Delete
```
app.delete('/profile/:id', (req, res) => {
	console.log('deleted', profile[req.params.id]);
	profile.splice(req.params.id,1);
	res.sendStatus(204);; //Permet de renvoyer une réponse sans avoir besoin de spécifier le corps de la réponse (204: No Content)
})
```
Put
```
app.put('/profile/:id', (req, res) => {
  Object.assign(profile[req.params.id], req.body); // Permet de remplacer les valeurs des champs de profile par ceux de req.body
  console.log('updated', req.body);
  res.sendStatus(201); 
})
```
Post
```
app.post('/profile', (req, res) => {
  profile.push(req.body);
  console.log('added', req.body);
  res.sendStatus(201); // = created
})
```
!\ En théorie, POST sert à ajouter une nouvelle ressource et PUT à remplacer complètement une ressource ou à MAJ partiellement une ressource(PATCH sert à ça également).
`App.head()`: ces requêtes sont équivalentes à GET mais sans corps dans la requête.
`App.all(‘*’, fn)` est souvent utilisé après toutes les autres routes pour gérer tous les cas et renvoyer vers une erreur 404.
### URL parameters
Il existe 2 méthodes pour récupérer des paramètres d’une URL:
- `req.query`: Cette méthode est utilisée lorsque les paramètres sont spécifiés dans l’URL de la manière suivante http://google.com/profile?api_key=123&id=1.

Pour les récupérer, il faut utiliser `req.query[‘api_key’]` / `req.query.id`
Cette méthode est intéressante dans le cadre de la méthode GET car elle permet de renvoyer 2 résultats différents en fonction de la présence ou non du paramètre (l’équivalent avec l’autre méthode n’est pas possible car si le paramètre n’existe pas alors nous avons une erreur):
```
app.get('/profile', (req, res) => {
	if(req.query.id) res.send(profile[req.query.id]); 
	else res.send(profile);
})
```
- req.params: Si cette méthode est utilisée alors il faut spécifier dans l’url (`app.put(‘url’, …)`) où se situera le paramètre avec: (exemple: `app.put(‘/profile/:id’)`).

Ensuite, on pourra accèder à ce paramètre dans le request handler grâce à `req.params.id`.
En cas de multiples paramètres, il suffit de spécifier entièrement la route et les paramètre (/profile/:id/user/:userId/:filter) → `request.params.id` /  `request.params.userId` / `request.params.filter`
### Input Validation
Il est très important de ne jamais se fier aux données envoyées par le client.
C’est pourquoi il est nécessaire de mettre en place des inputs validation dans notre code (notamment au niveau de PUT et POST).
L’implémentation de POST suivante est très dangereuse car il n’y a pas de vérification au préalable de req.body:
```
app.post('/profile', (req, res) => {
	profile.push(req.body);
	console.log('added', req.body);
	res.sendStatus(201); // = created
})
```
C’est pourquoi pour vérifier certaines informations, on va utiliser cette forme:
```
app.post('/profile', (req, res) => {
  if(!req.body.username || !req.body.email || !req.body.url) return res.sendStatus(400);
  let obj = { //whitelisting des valeurs acceptées
    username: req.body.username,
    email: req.body.email,
    url: req.body.url
  };
  profile.push(obj);
  console.log('added', obj);
  res.sendStatus(201); 
})
```
Les vérifications ici sont insuffisantes (on peut ajouter un champs rempli d’espace, ...). Cependant, il existe des modules permettant de réaliser cela (express-validator).
### Middleware
Le middleware pattern est une série de petites unités de code connectées entre elles où la sortie d’une unité et l’entrée d’une autre.
Les middlewares ressemblent aux services d’angular car ils permettent la réutilisation et la modularisation du code.

#### Processus:
Request →Middleware 1 → Middleware 2 → route →response 

Pour utiliser un middleware, il faut utiliser `app.use(middleware1)`. 
**/!\ L’ordre est très important puisque le programme sera lu de haut en bas en ce qui concerne les middlewares.**

Il existe 2 types de middlewares:
- Les modules npm nécessitant un import au préalable `app.use(bodyParser.json())` // bodyParser doit être importé au préalable.
- Les custom middleware utilisant `app.use((req, res, next) => {next()})` ou `app.use(middleware)` → middleware étant une fonction définie au préalable.

Le middleware peut être utilisé également pour décorer l’objet request et ainsi pouvoir transmettre des données à l’ensemble des middlewares ou endpoints. Ici on récupère un tableau de posts et on décore l’objet request avec ce tableau afin qu’il puisse être utilisé partout.
```
app.use((req, res, next) => {
	req.posts = require('./Posts');
	next();
})
```
#### Custom middlewares
Les middlewares disposent de l’argument next car ils se situent au milieu, les routes (app.get) n’en disposent pas car elles sont à la fin (voir processus ci-dessus).

#### Exemples de middlewares:
##### Vérification de la présence d’api key
```
app.use((req, res, next) => {
  if(req.query['api_key']) {
    next(); 
  } else {
    res.status(401).send({msg:'You need an API key to connect'});
  } 
});
```
/!\ Il est important de créer une condition en if-else sinon une erreur côté serveur apparaitra!
/!\ la méthode send renvoie directement la réponse, peu importe où elle se situe (middleware ou route).

##### Middleware injecté dans une route
Un middleware est utilisé avec app.use lorsque le code de ce dernier doit s’appliquer à toutes les routes. Cependant si il n‘y a besoin du middleware que dans une route spécifique, il est possible d’en injecter un de la manière suivante: 
```
app.get('/', (req, res, next) => {
  console.log('inline middleware'); //Inline middleware
  next(); //Inline middleware
  }, 
  (req, res) => {
    res.send({msg:'Youssef'});
});
```
#### Gestion des erreurs
Pour gérer les erreurs, il faut créer un middleware spécifique:
```
App.get(‘/’, (req, res, next) => {
	next(new Error(‘Oops’))
}, (req, res) => {
	res.send(‘hello’)
})
App.use((error, req, res, next) => {
	res.status(500).send(error));
})
```
#### NPM module middlewares
Il existe de nombreux middlewares, évitant ainsi de devoir réinventer la roue:
- body-parser: Permet de convertir le corps de la requête en JSON. Il est maintenant intégré nativement avec express.
      Ce dernier a 2 principaux usages:
    - Usage: AJAX/XHR provenant de SPA
          Syntaxe: `app.use(express.json());` 
          Cela va créer une propriété body à l’objet request accessible via request.body
          /!\ Pour que ce middleware parse correctement (reconnaisse le type de payload), il faut rajouter dans le header de la requête client le content-type (`curl -d ‘JSON ou autre’ localhost:3000 -i -H ‘Content-type: application/json’`)
    - Usage: HTML web forms avec un attribut action permettant de poster la requête
          Syntaxe: `app.use(express.urlencoded({extended: false}))`
- morgan: Permet de logger des informations sur la requête entrante (server logs). Il faut le configurer avec `app.use(morgan(‘dev’));`
Résultat: POST /?api_key=1 200 55.230 ms – 17
- compression gzip
- connect-timeout set request timeout
- cookie-parser Cookies
- cookie-session Session via Cookies store
- csurf CSRF
- errorhandler error handler
- `express.static(‘PATH’);` → Permet de créer un dossier statique. Il suffit de saisir /’nom du fichier’.html pour y accéder.
      /!\ Le fichier index.html fera office de fichier par défaut lorsque aucun path spécifique n’est mentionné.
      /!\ Il est également possible de rajouter du CSS à ces fichiers.
- express-session session via in-memory or other store
- method-override HTTP method override
- response-time
- serve-favicon favicon
- serve-index
- serve-static static content
- vhost: virtual host
- cookies and keygrip: analogous to cookieParser
- raw-body: raw request body
- connect-multiparty, connect-busboy: file upload
- qs: analogous to query
- st, connect-static analogous to staticCache
- express-validator: validation
- less: LESS CSS
- passport: authentication library
- helmet: security headers
- connect-cors: CORS
- connect-redis
- moment: Donne l’heure et la date du jour → `moment().format()`. Il existe de base `Date()` mais moment permet de personnaliser l’affichage.

#### Templating engine: Handlebars
1. Installer `express-handlebars`
2. Importer `const exphbs = require('express-handlebars');`
3. Créer les dossiers et fichier suivant views > layouts > main.handlebars
   main.handlebars contiendra le boilerplate html. Dans le body, il faudra rajouter {{body}} afin de pouvoir inclure des données
4. Initialiser handlebars dans app.js
```
// Handlebars Middleware
app.engine('handlebars', exphbs({defaultLayout: 'main'}));
app.set('view engine', 'handlebars');
```
5. Créer une route pour accéder à un fichier handlebar (le fichier devra être nommé index.handlebars)
```
// Homepage Route
app.get('/', (req, res) => res.render('index', {
  title: 'Member App',
  members
}))
```
6. Créer le fichier index.handlebars dans le dossier views
```
<form action="/api/members" method="POST" class="mb-4">
  <div class="form-group">
    <label for="name">Name</label>
    <input type="text" name="name" class="form-control">
  </div>
  <div class="form-group">
    <label for="email">Name</label>
    <input type="email" name="email" class="form-control">
  </div>
  <input type="submit" value="Add member" class="btn btn-primary btn-block">
</form>
<h4>Members</h4>
<ul class="list-group">
  {{#each members}}
  <li class="list-group-item">{{this.name}}: {{this.email}}</li>
  {{/each}}
</ul>
```
7. Ajouter dans le endpoint gérant les post request `res.redirect('/');`
#### Alternative PUG (Jade) templating engine
Cela permet de renvoyer des pages html prédéfinies.
```
/* Load view engine */
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'pug');

/* Home route */
app.get('/', (req, res) => {
	res.render('index', {
	title: 'Hello'
});
})
```
Voir https://www.youtube.com/watch?v=Ad2ngx6CT0M&index=3&list=PLillGF-RfqbYRpji8t4SxUkMxfowG4Kqp

