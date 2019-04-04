---
title: mongoDB
created: '2019-04-04T07:34:07.975Z'
modified: '2019-04-04T08:48:22.882Z'
---

# mongoDB
## Introduction
MongoDB est une BDD No SQL (non-relationnelle).
Cela signifie que cette BDD ne supportent pas des relations complexes comme le font les BDD SQL.
Il existe différents types de BDD no SQL:
- key value store (Redis → hash tables)
- document store (mongoDB → JSON)
- Graphs store
### Avantages du no-SQL:
- Comme la donnée n’est pas structurée, il ne faut pas obligatoirement la même structure de donnée pour toutes les entrées. Ce qui est pratique quand on souhaite récupèrer des données de différentes sources et les stocker au même endroit.
- Rapide à prototyper
Comme la BDD ne stocke pas les relations entre entités (pas de join queries), ces relations sont stockées au niveau de l’application (node.js code), plus précisément dans un ORM/ODM.

Comparaison entre no SQL et SQL: https://www.youtube.com/watch?v=uD3p_rZPBUQ
## CAP theorem
CAP:
- Consistency
- Availability
- Partition tolerance

Le théorème précise qu’on ne peut pas avoir les 3 en même temps.
Dans les DB SQL, on a le C + A, en no SQL, on a A+P.

La consistance n’est pas cruciale dans la majorité des applications. Cela signifie par exemple qu’un statut facebook ne sera pas forcément à jour dans la seconde. Un FB friend au Japon ne verra peut-être pas ma MAJ aussi vite qu’un FB friend aux USA.
Cependant en ce qui concerne par exemple la banque, je veux que mon solde soit consistant peu importe l’endroit où je me connecte pour vérifier mon solde.
## Commandes
### Installation/configuration
Emplacement:
Fichiers de logs: `var/lib/mongodb`
Fichiers de Données: `var/lib/mongodb/data/db`
Configuration: `/etc/mongodb.conf`
### Utilisation
- Lancer le REPL: `mongo`
- Voir toutes les BDD: `show dbs`
- Démarrer mongoDB: `sudo service mongod start`
- Arrêter mongoDB: `sudo service mongod stop`
- Redémarrer mongoDB: `sudo service mongod restart`
- Vider la console: `cls`
- Créer une nouvelle DB: `use ‘nom de la BDD’`
- Voir dans quel DB je me situe: `db`
- Créer un utilisateur:
```
db.createUser({
    user:"ysf",
    pwd:"123",
    roles:["readWrite", "dbAdmin"]
})
```
- Créer une collection: `db.createCollection(‘nom de la collection’)` (ex:`db.createCollection(‘customers’)`)
- Afficher les collections: `show collections`
- Ajouter des données à une collection: `db.customers.insert({first_name:’Youssef’, last_name:’Farkhani’})`

Il est possible d’ajouter des données sous forme de tableau. Elles seront "remise à plat" en format JSON dans la collection. (ex: `db.customers.insert([{first_name:’Clélia’, last_name:’Timsit’}, {first_name:’Bequille’, last_name:’Doe’}])`
- Afficher les données d’une collection: `db.customers.find().pretty()` (`pretty()` permet de rendre l’affichage plus agréable).
- Filtrer la recherche: `db.customers.find({first_name:"Youssef"})` → va remonter tous les résultats correspondants.
- Filtrer la recherche avec plusieurs paramètres: `db.customers.find({$or:[{first_name:"Youssef"},{first_name:"Clélia "})` 
- Filtrer la recherche avec une condition: `db.customers.find({age:{$lt:40}})`
      Il existe également gt/gte/lte (equal)
- Filtrer la recherche dans une propriété: `db.customers.find({"adress.city": "Boston"})`
- Trier la recherche: `db.find().sort({last_name:1})`
      Lorsque le paramètre est à 1, l’ordre est ascendant si on veut l’inverse, utiliser -1.
- Compter le nombre de retour: `db.find().count()`
- Limiter le nombre de résultats: `db.find().limit(4)`
- Remplacer un document: `db.customers.update({first_name:"Youssef"}, {first_name: "John", last_name:"Doe", gender: "male"})`
      Le premier paramètre permet de filtrer et le 2ème remplace le document.
- Mettre à jour un document: `db.customers.update({first_name:"Youssef"}, {$set:{gender: "male"})`
      Si la valeur est présente alors elle est mise à jour sinon elle est crée.
- `db.customers.update({first_name:"Youssef"}, {$inc:{age:5}})`
      Si la valeur "age" est présente alors on y ajoute 5  sinon elle est crée.
- Retirer un champs: `db.customers.update({first_name:"Youssef"}, {$unset:{age:1}}})`
- Mettre à jour un document si il n’existe pas le créer: `db.customers.update({first_name:’Jean’}, {first_name:’Jean’, last_name:’Michel’}, {upsert: true})`
- Renommer une valeur: `db.customers.update({first_name:"Youssef"}, {$rename:{"gender": "sex"}})`
- Supprimer un/des documents: `db.customers.remove({first_name:"Youssef"})`
- Supprimer un unique document: `db.customers.remove({first_name:"Youssef"}, {justOne: true})`
- Itérer sur des éléments: `db.customers.find().forEach(function(doc){print("Customer Name:" + doc.first_name)})`
## Mongoose:
Mongoose est une librairie donnant accès à des connecteurs permettant de requêter la BDD depuis le code JS. On l’appelle le client de la BDD.
Mongoose permet de créer une abstraction de la BDD dans le code JS
De plus, comme aucun contrôle de cohérence n’est fait lorsque de la donnée est injectée dans mongoDB, cette vérification est déplacée du côté de l’application et plus précisément du coté de mongoose avec la création de models.
### Initialiser mongoose
```
const mongoose = require('mongoose');
mongoose.connect('mongodb://localhost/nodekb', {useNewUrlParser: true}); //nodekb est le nom de notre BDD
let db = mongoose.connection;
/* Check connection */
db.once('open', () => console.log('Connected to MongoDB'));

/* Check for DB errors */
db.on('error', (err) => console.log(err));
```
### Créer un model
Les modèles permettent:
- Mettre à disposition des méthodes pour sauvegarder, récupérer, .. de la donnée
- Définir le schéma de notre document (type des données) : /!\ Mongoose ignorera les données non définies dans le Schema
```
const mongoose = require('mongoose');
/* Article Schema */
const articleSchema = mongoose.Schema({
  _id: mongoose.Schema.Types.ObjectId,
  title: {
    type: String,
    required: true //Oblige la présence de cette propriété
    default: ‘Harry Potter’ //Rempli par défaut la valeur de la propriété
  },
  author: {
    type: String,
    required: true
  },
  body: String
});
module.exports = mongoose.model('Article', articleSchema);
```
Ensuite, il faut importer ce modèle dans le programme principal:
```
/* Bring in models */
const Article = require('./models/article');
```
### Requêter la BDD
A noter que la méthode exec() permet de retourner une promise sur laquelle on peut réaliser des actions.
Mongoose peut également être utilisé avec des callbacks au lieu de promise.
#### Récupérer des données
```
/* Home route */
app.get('/', (req, res) => {
  Article.find()
    .exec()
    .then(docs => {
      res.status(200).render('index', {
        title: 'Articles',
        articles: docs
    });
  })
    .catch(err => console.log(err));
})
```
/!\ Il est possible de chainer les commandes limit, where, sort à Article.find()
/!\ Si il n’y a pas de données alors la valeur renvoyée est un tableau vide.
#### Recherche par id
```
app.get('/article/:id/edit', (req, res) => {
  Article.findById(req.params.id)
    .exec()
    .then(article => {
      res.render('edit_article', {
        title: 'Edit',
        article: article
    })
  })
  .catch(err => console.log(err));
})
```
FindById est utilisée car elle ramène un élément unique au lieu d’un array, ce qui nous évite de préciser article: article[0].
#### Mettre à jour les données
```
app.post('/article/:id/edit', (req, res) => {
  const id = req.params.id;
  Article.update({_id:id}, {$set: {
    title = req.body.title;
    author = req.body.author;
    body = req.body.body;
  }})
    .exec()
    .then(result => res.redirect('/'))
    .catch(err => console.log(err));
})
```
#### Mettre à jour uniquement les données du document de manière spécifique
```
app.patch('/article/:id/edit', (req, res) => {
  const id = req.params.id;
  const updateOps = {};
  for(const ops of req.body){
    update[ops.propName] = ops.value
  }
  Article.update({_id:id}, {$set: updateOps })
    .exec()
    .then(result => res.redirect('/'))
    .catch(err => console.log(err));
  })
```
Dans ce cas ci, la requête devra envoyer un array avec la valeur de ops.propsName et de ops.value:
`[{"propName":"title", "value": "Harry Potter"}]`
Plus de détails, voir https://www.youtube.com/watch?v=WDrU305J1yw à 34’.
#### Ajouter / Supprimer les données
Les balises HTML n’intègre pas de tag permettant de soumettre une DELETE request (ni une PUT) pour des raisons de sécurité. En effet, quelqu’un de mal intentionné pourrait attaquer un site en supprimant ou modifiant ses ressources (si le website est mal codé).
##### Ajouter des données
Pour PUT, il existe une solution de contournement grâce à l’utilisation de POST.
```
app.post('/articles/add', (req, res) => {
  const article = new Article({
    _id: new mongoose.Types.ObjectId(), //Permet de créer une ID unique (non obligatoire, mongoDB le fait automatiquement)
    title = req.body.title;
    author = req.body.author;
    body = req.body.body;
  });
  article.save().then(result => {
    console.log(result);
  })
  .catch(err => console.log(err));
  res.redirect('/');
})
```
##### Supprimer des données
Pour DELETE, il faut créer un script permettant d'ajouter un event Listener à notre page qui fera appel à notre endpoint :
```
const deleteArticleBtn = document.querySelector(".delete-article");
deleteArticleBtn.addEventListener('click', function(e){
    const id = e.target.getAttribute('data-id');
    const url = "/article/"+id
    const myInit = { method: 'DELETE' }; 
    fetch(url, myInit)
        .then(() => {
            window.location.href='/';
        })
        .catch(err => console.dlogir(err));
});
```
Puis créer notre endpoint
```
app.delete('/article/:id', (req,res) => {
    const query = { _id: req.params.id};
    Article.deleteOne(query, (err) => {
        if(err) console.log(err);
        res.send('Success');
    })
});
```
## Définitions:
- REPL: Read Evaluate Print Loop (shell) → Il s’agit d’un environnement sur la CLI permettant de lancer des commandes spécifique à un programme et d’exécuter les instructions. Par exemple, si on saisit node alors une > (node console) va apparaître nous signifiant qu’on peut entre une expression JS.
	Dans le cadre de mongoDB, il s’agit d’un client permettant de se connecter à sa DB et de 	réaliser l’ensemble des actions (CRUD).
- Collections: Equivalent des tables en SQL.
## MongoDB Atlas
Version cloud de mongoDB (Database AAS) 
Cluster, réplication de la donnée
## A voir:
Pour les BDD SQL, voir **sequelize**.

