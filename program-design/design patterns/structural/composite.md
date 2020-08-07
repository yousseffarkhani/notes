# Introduction
Permet de composer un objet sous forme d'arbre et ensuite pouvoir l'utiliser comme s'il ne s'agissait que d'un unique objet.
C'est utile dans le cas de commandes pouvant contenir d'autres commandes ou des produits.
Si on veut connaitre le prix total, un moyen serait de traverser l'ensemble de l'arbre et d'ajouter les prix des produits.
Un autre moyen est d'utiliser le composite pattern.
Ce pattern met à disposition une interface qui fait que lorsqu'on l'appel sur l'objet parent alors il l'appelle à l'ensemble de ces enfants qui eux même à leur tours font la même chose.

# Description
Il s'agit de créer une classe composite qui transmet les demandes à tout ces enfants.
# Avantages

# Utilisation

# Code
```TS
interface Test {
  getPrice(): number;
}

class Product implements Test {
  private price: number;
  constructor(price) {
    this.price = price;
  }
  getPrice() {
    return this.price;
  }
}

class Box implements Test {
  private products: Array<Test>;

  constructor() {
    this.products = [];
  }

  add(product: Test) {
    this.products.push(product);
  }

  remove(index) {
    // remove
  }

  getPrice() {
    return this.products.reduce((acc, curr) => acc + curr.getPrice(), 0);
  }
}

const product1 = new Product(1);
const product2 = new Product(10);
const box1 = new Box();
box1.add(product1);
const command = new Box();
command.add(box1);
command.add(product2);

console.log(command.getPrice());
```
# Source
https://refactoring.guru/design-patterns/composite
