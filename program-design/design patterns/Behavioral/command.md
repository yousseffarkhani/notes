# Introduction
Ce pattern transforme une requête en un stand-alone object contenant toutes les informations de la requête.
Cela permet de créer un historique d'opérations et pouvoir annuler ces opérations.
Ce pattern permet d'encapsuler des operations dans des commandes disposant d'une méthode execute et undo.
# Description

# Avantages
- Permet de découpler la requête de l'objet exécutant la requête.
Exemple:
Button => Update('John', 5)
MenuItem => Update('John', 5)
Shortcut => Update('John', 5)

Button => 
MenuItem => ==> Save('John', 5)
Shortcut => 

- Open/Closed Principle: Possibilité d'ajouter des commandes sans avoir à modifier le code existant

- Permet de garder un historique des commandes et de les défaire (utile pour un éditeur de texte par exemple)
- Permet de combiner plusieurs commandes préalablement définies.
Par exemple:
save
exit
save and exit
# Utilisation
Cela peut être utilise pour ajouter et annuler des ajouts dans une BDD.
# Implémentation
1. Définir l'interface Command avec execute() et undo()
2. Extraire les requêtes en créant des classes implémentant l'interface Command
3. Identifier les classes exécutant les commandes. Créer l'historique de commandes
# Code
```TS
interface Command {
  execute(value: number): number;
  undo(value: number): number;
}

class Calculator {
  private history: Command[];
  private value: number;
  constructor() {
    this.value = 0;
    this.history = [];
  }
  executeCommand(command: Command) {
    this.value = command.execute(this.value);
    this.history.push(command);
    return this;
  }

  undo() {
    const oldCommand = this.history.pop();
    this.value = oldCommand.undo(this.value);
    return this;
  }

  getValue(): number {
    return this.value;
  }
}

class AddCommand implements Command {
  private valueToAdd;

  constructor(valueToAdd: number) {
    this.valueToAdd = valueToAdd;
  }

  execute(currentValue: number): number {
    return currentValue + this.valueToAdd;
  }

  undo(currentValue: number) {
    return currentValue - this.valueToAdd;
  }
}

class MultiplyCommand implements Command {
  private valueToMultiply;

  constructor(valueToMultiply: number) {
    this.valueToMultiply = valueToMultiply;
  }

  execute(currentValue: number): number {
    return currentValue * this.valueToMultiply;
  }

  undo(currentValue: number) {
    return currentValue / this.valueToMultiply;
  }
}

const calculator = new Calculator();
const addCommand = new AddCommand(5);
const multiplyCommand = new MultiplyCommand(5);
calculator.executeCommand(addCommand).executeCommand(multiplyCommand);
console.log(calculator.getValue());
```
# Source
https://refactoring.guru/design-patterns/command
https://www.dofactory.com/javascript/command-design-pattern
https://www.youtube.com/watch?v=GQzfF5EMD7o
