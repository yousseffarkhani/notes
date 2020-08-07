# Introduction
Consiste à retourner un objet vide disposant des mêmes propriétés et méthodes cela permet d'éviter de devoir checker si l'objet retourné est null.
# Description
Souvent en JS, il est obligatoire de checker si un objet n'est pas null avant de pouvoir utiliser une méthode ou une propriété.
# Avantages
Pas besoin de vérifier les valeurs de retours.
# Utilisation
User Accounts (guest, member)
# Code
```TS
interface IUser {
  getId(): number;
  getName(): string;
  hasAccess(): boolean;
}

class User implements IUser {
  private id: number;
  private name: string;
  constructor(id: number, name: string) {
    this.id = id;
    this.name = name;
  }

  getId(): number {
    return this.id;
  }
  getName(): string {
    return this.name;
  }
  hasAccess(): boolean {
    return this.name === 'bob';
  }
}

class NullUser implements IUser {
  private id: number;
  private name: string;
  constructor() {
    this.id = -1;
    this.name = 'Guest';
  }

  getId(): number {
    return this.id;
  }
  getName(): string {
    return this.name;
  }
  hasAccess(): boolean {
    return false;
  }
}

const users = [new User(1, 'bob'), new User(2, 'john')];

const getUser = (id: number): IUser => {
  const user = users.find((user) => user.getId() === id);
  if (user == null) {
    return new NullUser();
  }
  return user;
};

const printUser = (id: number) => {
  const user = getUser(id);
  console.log('hello', user.getName());

  if (user.hasAccess()) console.log('You have access to this resource');
  else console.log("You don't have access to this resource");
};

printUser(1);
printUser(2);
printUser(3);
```
# Source
https://www.youtube.com/watch?v=D4Dja5WSZoA&list=PLZlA0Gpn_vH_CthENcPCM0Dww6a5XYC7f&index=2
