# Définition
Approche pour designer et construire une application.

1. Identifier les classes
2. Séparer les responsabilités des classes
3. Etablier les relations entre les classes
4. Réutiliser en extrayant les fonctionnalités communes

# Workflow
## Identifier les classes depuis les requirements
Cette section consiste à construire les classes à partir des business requirements.

1. Analyser le business problem
2. Commencer avec les noms pour extraire les classes
- Identifier les classes représentant les entities (Customer, Product, Orders)
3. Definir les caractéristiques des classes
- Définir les propriétés et méthodes de ces classes (validate(), retrieve(), save(), ...)
Il faut que les classes soient logiques (ie order has many order item)
- Considérer les évolutions futures (changement de prix, d'adresse, ...)
- Définir l'abstraction appropriée (Joe et sa famille) en ne gardant que les informations utiles au programme.

Application layer (UI, Business logic, Data access layer, common code (logging, emailing, ...))
Business logic
Propriétés
Getters / Setters sur les propriétés (sauf si read ou write only)
Les méthodes statiques peuvent être utilisés pour traquer des informations globales (nombre de fois où l'instance a été crée) 

Méthodes
Les méthodes sont les fonctions exercant la business logic
Ne pas oublier l'overloading (consiste à avoir le même nom pour les méthodes mais avec des signatures différentes)
Ex: Retrieve() => Renvoie une liste, Retrieve(int customerId) => Renvoie un customer

## Séparer les responsabilités des classes
Cette section cosiste à voir si les classes définies dans la 1ère partie disposent de trop de responsabilités. 
### Maximize Cohesion
La classe customer contient des adresses. Or la gestion d'adresses est un sujet complexe => Les adresses nécessitent une classe à part.
### Minimize Coupling
Lorsque notre class est dépendante de ressource externes (Ex: Database) alors il est souvent itnéressant de l'extraire dans sa propre classe.
La plupart du temps pour les opérations de sauvegarde ou de retrieve d'une base de données une class appelée customer repository est ajouté (Design pattern: repository).

## Etablier les relations entre les classes
Comment les objets interagissent entre eux ?
1. Définir  les relations

### Types of relationships
#### Collaboration (uses a)
Un objet utilise un autre objet. (Customer repository => Customer: Le repository utilise le customer pour renvoyer de la donnée sous une forme cohérente (retrieve) et la sauvegarder)
Autre exemple: Customer repository => adress repository

#### Compositions (has a)
Il en existe 2 types:
- Aggregation: Chaque entité est indépendante. Les 2 objets peuvent exister indépendamment (Order <=> Customer)
Créer une propriété référencant l'ID de l'autre objet.
- Composition: Une entité dépend de l'autre. Si l'entité mère est détruire alors l'autre n'a plus de sens (Order => Order items)
Une voiture a moteur et des roues

Les classes appelant d'autres classes ont généralement des propriétés référencant une autre classe (La classe customer dispose d'une propriété home address => Customer has an adress).

Ces compositions disposent de cardinalité (1-1, 1-many, many-many)

#### Inheritance (is a)


## Réutiliser en extrayant les fonctionnalités communes

# Vocabulaire
Business object 
Class defined to solve a business problem (Customer)

Entity
An element from the real world abstracted as a class (Customer)
