# Introduction
Ce pattern fait en sorte de modifier le state intere d'un objet et ainsi son comportement.
Ce pattern permet de clairifier le code utilisant une machine d'états.
# Description
Il consiste à extraire en une classe chaque state que peut avoir un objet.
# Avantages
Separation of concern: La logique est dans l'objet State.
# Utilisation

# Code
```TS
class Context {
  private state: State;

  constructor(state: State) {
    this.changeState(state);
  }

  public changeState(state: State) {
    console.log('changing state to ', (<any>state).constructor.name);
    this.state = state;
    this.state.setContext(this);
  }

  public request1() {
    this.state.handle1();
  }

  public request2() {
    this.state.handle2();
  }
}

abstract class State {
  protected context: Context;

  public setContext(context: Context) {
    this.context = context;
  }

  public abstract handle1(): void;

  public abstract handle2(): void;
}

class ConcreteStateA extends State {
  public handle1(): void {
    console.log('Concrete State A handle 1');
    this.context.changeState(new ConcreteStateB());
  }

  public handle2(): void {
    console.log('Concrete State A handle 2');
  }
}

class ConcreteStateB extends State {
  handle1(): void {
    console.log('Concrete State A handle 1');
  }

  handle2(): void {
    console.log('Concrete State A handle 2');
    this.context.changeState(new ConcreteStateA());
  }
}

const context = new Context(new ConcreteStateA());
context.request1();
context.request2();
```
# Source

https://refactoring.guru/design-patterns/state
