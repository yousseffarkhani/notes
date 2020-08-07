# Introduction
Ce pattern permet d'avoir une autorité centrale sur un groupe d'objet en encapsulant comment ces objets interagissent.
Il restreient les communications directes entre les objets et les force à collaborer via le mediateur.
Un parrallèle pourrait être fait avec une tour de controle qui coordonne l'ensemble des arrivées et départs des avions.
# Description

# Avantages
Permet aux objets de se concentrer sur leur tâche et de laisser la logique au médiateur.
# Utilisation
Ce pattern est utilise lorsqu'il s'agit de manager plusieurs conditions complexes liées au statut de chaque object.
# Example
Une page de réservation de vol d'avion
Une règle de mediator serait d'entrer une date valide de départ, de retour, date de départ > retour, aéroport valide après quoi le bouton rechercher est disponible.
Ainsi chaque module communique directement au mediator qui s'occupe de gérer les règles.
Sinon ça devient rapidement le chaos si la date d'arrivée doit vérifier qu'elle est bien après la date de départ, ...

# Code
```TS
/**
 * The Mediator interface declares a method used by components to notify the
 * mediator about various events. The Mediator may react to these events and
 * pass the execution to other components.
 */
interface Mediator {
    notify(sender: object, event: string): void;
}

/**
 * Concrete Mediators implement cooperative behavior by coordinating several
 * components.
 */
class ConcreteMediator implements Mediator {
    private component1: Component1;

    private component2: Component2;

    constructor(c1: Component1, c2: Component2) {
        this.component1 = c1;
        this.component1.setMediator(this);
        this.component2 = c2;
        this.component2.setMediator(this);
    }

    public notify(sender: object, event: string): void {
        if (event === 'A') {
            console.log('Mediator reacts on A and triggers following operations:');
            this.component2.doC();
        }

        if (event === 'D') {
            console.log('Mediator reacts on D and triggers following operations:');
            this.component1.doB();
            this.component2.doC();
        }
    }
}

/**
 * The Base Component provides the basic functionality of storing a mediator's
 * instance inside component objects.
 */
class BaseComponent {
    protected mediator: Mediator;

    constructor(mediator: Mediator = null) {
        this.mediator = mediator;
    }

    public setMediator(mediator: Mediator): void {
        this.mediator = mediator;
    }
}

/**
 * Concrete Components implement various functionality. They don't depend on
 * other components. They also don't depend on any concrete mediator classes.
 */
class Component1 extends BaseComponent {
    public doA(): void {
        console.log('Component 1 does A.');
        this.mediator.notify(this, 'A');
    }

    public doB(): void {
        console.log('Component 1 does B.');
        this.mediator.notify(this, 'B');
    }
}

class Component2 extends BaseComponent {
    public doC(): void {
        console.log('Component 2 does C.');
        this.mediator.notify(this, 'C');
    }

    public doD(): void {
        console.log('Component 2 does D.');
        this.mediator.notify(this, 'D');
    }
}

/**
 * The client code.
 */
const c1 = new Component1();
const c2 = new Component2();
const mediator = new ConcreteMediator(c1, c2);

console.log('Client triggers operation A.');
c1.doA();

console.log('');
console.log('Client triggers operation D.');
c2.doD();
```
# Source
https://refactoring.guru/design-patterns/mediator/typescript/example
https://refactoring.guru/design-patterns/mediator
https://www.dofactory.com/javascript/design-patterns/mediator
