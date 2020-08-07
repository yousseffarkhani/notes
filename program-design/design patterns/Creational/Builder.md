# Introduction
Permet de créer des objets disposant de multiples types et représentations qui utilise le même code de construction. 
Par exemple la création d'une voiture consiste en l'assemblage de plusieurs pièces. 
Le builder pattern permet de rendre plus clair notre code.

Au lieu d'avoir certains paramètres à null (Car(windows, seats, radio, airConditioner) new Car(1, 2, 3, null)), on préfèrera utiliser les méthodes d'un builder lorsque l'on en a besoin
Car builder:
 - buildSeat
 - buildRadio
 - buildAirConditioner

On peut également avoir recours à un Director qui est une entité qui contient la logique de construction.
```python
class Car is

class Manual is

interface Builder is
    method reset()
    method setSeats(...)
    method setEngine(...)
    method setTripComputer(...)
    method setGPS(...)

class CarBuilder implements Builder is
    private field car:Car

    constructor CarBuilder() is
        this.reset()

    method reset() is
        this.car = new Car()

    method setSeats(...) is

    method setEngine(...) is

    method setTripComputer(...) is

    method setGPS(...) is

    method getProduct():Car is
        product = this.car
        this.reset()
        return product

class CarManualBuilder implements Builder is
    private field manual:Manual

    constructor CarManualBuilder() is
        this.reset()

    method reset() is
        this.manual = new Manual()

    method setSeats(...) is

    method setEngine(...) is

    method setTripComputer(...) is

    method setGPS(...) is

    method getProduct():Manual is

class Director is
    private field builder:Builder

    method setBuilder(builder:Builder)
        this.builder = builder

    method constructSportsCar(builder: Builder) is
        builder.reset()
        builder.setSeats(2)
        builder.setEngine(new SportEngine())
        builder.setTripComputer(true)
        builder.setGPS(true)

    method constructSUV(builder: Builder) is
        // ...


class Application is

    method makeCar() is
        director = new Director()

        CarBuilder builder = new CarBuilder()
        director.constructSportsCar(builder)
        Car car = builder.getProduct()

        CarManualBuilder builder = new CarManualBuilder()
        director.constructSportsCar(builder)

        // The final product is often retrieved from a builder
        // object since the director isn't aware of and not
        // dependent on concrete builders and products.
        Manual manual = builder.getProduct()
```
# Description

# Avantages

# Utilisation

# Code
```TS

```
# Source
https://www.youtube.com/watch?v=M7Xi1yO_s8E&list=PLZlA0Gpn_vH_CthENcPCM0Dww6a5XYC7f&index=3
