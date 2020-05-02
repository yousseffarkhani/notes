Proven solutions to common programming problems.
# Introduction
## Définition
Un design pattern décrit la solution de manière générale à une problématique récurrente.
Il existe plusieurs manières d'utiliser un design pattern car il ne s'agit que d'une guideline. 
Un pattern décrit:
- Le problème
Exemples de problèmes récurrents:
Recevoir une notification lorsqu'un process (asynchronous call) ou un statut d'une propriété change
Réaliser une API complexe facile à utiliser
Ajouter une foctionnalité à un objet (même si on ne dispose pas d'accès à l'objet original)
State machine qui réagit différement selon le mode actuel
Loop over elements of a collection or sequence
- Le core de la solution (et non l'implémentation)
 Pattern name: Observer, Decorator, Repository, MVC
 Solution: Core elements and interactions
 Consequences: Everything has a cost

## Gang of four
Ecrivains du livre design patterns.
Ils ont décrit 23 Patterns:
- Creational Patterns (abstrait l'instanciation d'un objet)
- Structural patterns (comment les classes sont composées pour former des structures larges)
- Behavorial patterns (Algorithms)
## Importance
# Patterns utilisés
- Observer
    - Problem: Getting a notification when a process completes or a state changes
    - Description:
Publish / Subscribe relationship
Publisher (Subject) 
Subscriber (Observer)
    - Consequences:
        - Les notifications peuvent ne jamais arriver (si le subject n'envoit rien)
        - Les notifications peuvent être envoyées à répétition
        - Les notifications peuvent arriver très rapidement (les observers doivent être en mesure de gérer le workload)
        - Des strong references peuvent empêcher le garbage collector de supprimer les observers
    - Tips: 
        - Le subscriber doit filtrer les inputs
        - Le subscriber doit faire attention aux inputs (surtout s'il y a des process lourds réalisés à la suite)
        - Le subscriber doit penser au cas où aucune notification arrive

- Proxy
    
- Iterator
    - Problem: Looping through elements of a collection

Getting items sequentially from a an aggregate (collection, array, ...)
Permet d'accéder aux éléments d'un objet aggrégés sans exposer sa représentation interne
Changer de chaine sur une télécommande
Aggregate: chaines de télé
Iterator: Télécommande / provide next channel
Consumer: Nous
ForEach loop
```JS
const it = iterable[Symbol.iterator](); // it contient l'énumérateur. Cet énumérateur dispose des propriétés next() (renvoie true s'il y a un item à la suite) et value (contient la valeur de l'iterable)
// it contient également un reset()
let next = it.next(); 
while (!next.done) {
    // eslint-disable-next-line no-await-in-loop
    await func(next.value)
    .then((res) => arr.push(res))
    // eslint-disable-next-line no-loop-func
    .catch((error) => log.error({ error }));
    next = it.next();
}
```
Iterator Consequences 
Adding or removing items from aggregate may cause an exception
Select / WHERE
- Chain of responsiblity
Tech support
Message is sent to first member of the chain. If the member can handle it then it stop otherwise passed to next member
Error handling javascript
If none of the receivers handle the message then the message will fall off the end of the chain
- Facade
    - Problem: Making a complex API easier to work with
Facade defines a higher level interface that makes the subsystem easier to use
Interface
- Factory Method
Define an interface for creating an object but let subclasses decide which class to instantiate

- Decorator
    - Problem: Adding functionality to an existing object
Attach une resposabilité additionnelle à un objet de manière dynamique
Example: Caching repository
- Adapter
- State
    - Problem: Behaving in distinct ways based on a current mode or state

# Autres patterns


