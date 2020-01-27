# Source
https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
# Introduction
Les enums permettent de :
- grouper



# Examples
Timezones : EST, CST
T-shirt sizes : Small, medium, large

# Implementation
```Go
type Weekday int

const (
    Sunday Weekday = 0
    Monday Weekday = 1
    Tuesday Weekday = 2
    Wednesday Weekday = 3
    Thursday Weekday = 4
    Friday Weekday = 5
    Saturday Weekday = 6
)

fmt.Println(Sunday) // Prints 0

func (day Weekday) String() string {
    names := [...]string{
        "Sunday",
        "Monday", 
        "Tuesday", 
        "Wednesday",
        "Thursday", 
        "Friday", 
        "Saturday"
    }
    if day < Sunday || day > Saturday {
        return "Unknown"
    }

    return names[day]
}

fmt.Printf("Which day it is ? %s\n", Sunday) // prints "Sunday"
```

# Implémentation d'enum en Golang

Dans Golang les enums sont représentés par iota.
Un iota est un compteur démarrant à 0 qui n'est utilisé qu'avec des valeurs constantes.

```Go
const (
    Sunday Weekday = iota + 1 // = 1
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
// Chaque constante sera de type Weekday et exécutera iota + 1

type Timezone int

const (
    // iota = 0, EST = -5
    EST Timezone = - (5+iota)
    // iota = 1, EST = -6
    CST
    // iota = 2, EST = -7
    MST
    // iota = 3, EST = -8
    PST
)
```

Il est possible de sauter des valeurs grâce à `_`.
