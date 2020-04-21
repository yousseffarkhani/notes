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
- Chain of responsiblity
- Facade
- Factory Method
- Decorator
- Adapter
# Autres patterns
