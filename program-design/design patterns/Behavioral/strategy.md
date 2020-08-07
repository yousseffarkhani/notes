# Introduction
Définit une famille d'algorithmes et encapsule chacun d'entre eux dans une classe afin de les rendre interchangeable.
# Description
Ce pattern permet de disposer de différents algorithmes définis au runtime.
Le code client choisi quel algorithme il souhaite utiliser.
Il existe dans ce pattern 2 objets:
- Context:
    - Maintient la référence de l'objet représentant la stratégie
    - Permet au client de choisir la stratégie
- Strategy:
    - Implémente l'algorithme
# Avantages
Utile lorsqu'il s'agit de comparer plusieurs algorithmes (app de navigation)
# Utilisation
- L'utilisateur choisi l'algorithme à appliquer (voir exemple)
- Comparer différents algorithmes en les appelant de manière séquentielle.
# Code
```TS
interface Strategy {
    calculateTimeOfDelivery() int
}
class Context {
  constructor(strategy Strategy) {
    this.strategy = strategy;
  }

  setStrategy(strategy Strategy) {
    this.strategy = strategy;
  }

// Délègue certains calculs aux classes de stratégie
  applyStrategy() {
    return this.strategy.calculateTimeOfDelivery();
  }
}

class UPS implements Strategy{
  constructor() {}

  calculateTimeOfDelivery() {
    return 1;
  }
}

class DHL implements Strategy {
  constructor() {}

  calculateTimeOfDelivery() {
    return 2;
  }
}
const ups = new UPS();
const dhl = new DHL();

const userChoice = 'UPS';

const shipping = new Context();
if (userChoice === 'UPS') {
  shipping.setStrategy(ups);
} else {
  shipping.setStrategy(dhl);
}

const timeOfDelivery = shipping.applyStrategy();
console.log(timeOfDelivery);
```
