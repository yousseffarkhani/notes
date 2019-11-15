# Source
https://www.calhoun.io/using-mvc-to-structure-go-web-applications/

# Introduction
Le modèle MVC est un pattern permettant d'organiser le code.
Chaque section de code sera classifié en model, vue et controlleur.

A noter que de temps à temps, il est impossible de classifier des morceaux de code (par exemple les middlewares en charge de gérer une requête). Dans ce genre de cas, il est ok de créer de nouveaux packages.

# Model
Cette section s'occupe d'interagir avec les données de l'application.
Elle est en charge de :
- communiquer avec les BDD ou encore d'autres services ou API
- Valider et normaliser la donnée

L'idéal étant que cette section soit totalement indépendante du reste (pas d'import de package).

Exemples de logique à mettre dans model:
- les modèles de données qui seront stockés dans la BDD
- Les vérifications (Ex: vérifier que l'adresse email d'un nouvel utilisateur n'est pas prise)
- Normalisation du format de donnée (ex: ToLowerCase)
- Opérations CRUD d'une BDD
- Récupération des données relationnelles (ex: Playground, commentaires).

# View
La partie View est responsable du rendu de la donnée.
Ce rendu peut être sous différentes formes : HTML, JSON, XML.
Cette partie devrait avoir un minimum de logique.

# Controller/Handler
Le controller a un rôle similaire à celui d'un controleur de trafic aérien.
Il s'occupe de diriger les avions vers les pistes mais ne réalise jamais le pilotage de celui-ci.
Dans notre cas, cette partie s'occupe de rediriger la data entrante vers les models et views appropriés.
Cette partie devrait avoir un minimum de logique. Elle ne devrait s'occuper que de parser la donnée et l'envoyer aux autres fonctions.


# Exemples d'organisation du code

app/
    localdisk/ #models (si les données sont stockées en local)
    psql/ #models
        user.go
        course.go
    http/ #Controllers
        user.go
        course.go
    html/ #Views
        template.go # used to parse html templates
        show_user.html
        new_user.html