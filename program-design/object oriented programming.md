# Définition
Approche pour designer et construire une application.

1. Identifier les classes
2. Séparer les responsabilités des classes
3. Etablir les relations entre les classes
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
Cette section consiste à voir si les classes définies dans la 1ère partie disposent de trop de responsabilités. 
### Maximize Cohesion
La classe customer contient des adresses. Or la gestion d'adresses est un sujet complexe => Les adresses nécessitent une classe à part.
### Minimize Coupling
Lorsque notre class est dépendante de ressource externes (Ex: Database) alors il est souvent intéressant de l'extraire dans sa propre classe.
La plupart du temps pour les opérations de sauvegarde ou de retrieve d'une base de données une class appelée customer repository est ajoutée (Design pattern: repository).

## Etablir les relations entre les classes
Comment les objets interagissent entre eux ?
### Types of relationships
#### Collaboration (uses a)
Un objet utilise un autre objet. (Customer repository => Customer: Le repository utilise le customer pour renvoyer de la donnée sous une forme cohérente (retrieve) et la sauvegarder)
Autre exemple: Customer repository => adress repository

new est utilisé
#### Compositions (has a)
Une classe est dite composée lorsqu'elle contient des références à d'autres classes.
Il en existe 2 types:
- Aggregation: Chaque entité est indépendante. Les 2 objets peuvent exister indépendamment (Order <=> Customer)
Créer une propriété référencant l'ID de l'autre objet.
- Composition: Une entité dépend de l'autre. Si l'entité mère est détruite alors l'autre n'a plus de sens (Order => Order items)

Les classes appelant d'autres classes ont généralement des propriétés référencant une autre classe (La classe customer dispose d'une propriété home address => Customer has an adress).

Ces compositions disposent de cardinalité (1-1, 1-many, many-many)

Propriétés faisant référence à d'autres classes
#### Inheritance (is a)
L'inheritance est utilisée lorsqu'on veut créer une version plus spécialisée d'une autre classe (customer => Business Customer)
L'inheritance est à utiliser lorsqu'il y a quelques différences entre la classe de base et la version spécialisée autrement il est plus intéressant d'utiliser une propriété type sur la classe de base.

## Réutilisation du code
Identifier les fonctionnalités communes et en créer des classes.
Il existe différentes méthodes pour réutiliser le code:
- Collaboration / Composition: Utilisation de l'adress dans la class Customer et Order.
- Inheritance: Extraire le code dans une base class et le réutiliser en créant des class enfants.
- Construire des components réutilisables (librairie d'utilitaires)
- Interfaces
### Inheritance
Possibilité d'override des fonctionnalités de la classe de base
### Polymorphism
Inheritance-based polymorphism
Concept représentant le fait d'utiliser la même signature de méthode mais avec différents retours
### Building a base class
Il existe 2 types de classes:
- Abstract class: Ne peut pas être instanciée, ne sert que de base à d'autres classes
- Concrete class: peut être instanciée

### Building a reusable components
Lorsque l'on souhaite créer des general purpose classes (String operations: String Handler, logging: Logging Service, notifications: Email Service), il devient intéressant de créer une classe spécifique qui puisse être réutilisable à différents endroits.

Third party system
Static classes/methods are often used
/!\ Ne créer des classes statiques que quand elles ne disposent pas de state.
### Interfaces
Les interfaces fournissent une manière propre d'interagir avec plusieurs classes d'une même manière.

Il existe en général 2 types d'interface sur une classe:
- La public interface d'une classe
- L'implémentation d'une interface
Les interfaces ressemblent à des rôles (Métaphore de l'humain: IParent, IEmployee)
L'interface est un contrat qui spécifit les propriétés et méthodes que chaque classe doit implémenter l'interface

```JS
interface Iloggable {
    log() string
}

class Customer implements Iloggable{
    log(){
        return "hello"
    }
}

class Product implements Iloggable{
    log(){
        return "hello"
    }
}

class LogginService {
    logToPubSub(item Iloggable){
        sendToPubSub(item.log())
    }
}
```

Image de la lollipop pour les interfaces et la public interface
# Vocabulaire
Business object 
Class defined to solve a business problem (Customer)

Entity
An element from the real world abstracted as a class (Customer)
