# Introduction
Permet de créer un unique objet partagé entre différentes ressources.
# Description
L'objet ne doit être instancié qu'une seule fois dans l'application.
# Avantages
# Désavantage
Crée énormément de coupling entre les différentes parties de l'application.
Peut créer des problèmes de race conditions étant donné que plusieurs parties du code accèdent à la même ressources au même moment.
# Utilisation
Logger
Services
# Code
```TS
class FancyLogger {
  static instance: FancyLogger;
  private logs: string[];

  constructor() {
    // Permet d'éviter les doublons de FancyLogger
    if (FancyLogger.instance == null) {
      // static property to store the singleton
      this.logs = [];
      FancyLogger.instance = this;
    }
    return FancyLogger.instance;
  }

  log(message: string) {
    this.logs.push(message);
    console.log('Fancy: ', message);
  }

  printLogCount() {
    console.log(this.logs.length);
  }
}

const logger = new FancyLogger();
Object.freeze(logger); // Empêche l'ajout de propriété ou de méthodes
module.exports = logger;
```
# Source
https://www.youtube.com/watch?v=sJ-c3BA-Ypo&list=PLZlA0Gpn_vH_CthENcPCM0Dww6a5XYC7f&index=4
