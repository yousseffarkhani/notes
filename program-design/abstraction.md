# Qu'est ce qu'une abstraction ?
## A quoi ça sert ?
- Simplifier la réalité en ignorant les détails non important 
- concentrer sur le plus important
## Bénéfices et désavantages
### Bénéfices
- Maintainability: Le code est bien organisé ce qui permet de savoir ou la modification de code doit avoir lieu
- Extensibilité (wrappers / interfaces)
- Testabilité
### Désavantages
- Quand on rajoute un wrapper autour d'une méthode, on rajoute une layer et donc de la complexité
### Limites
- Pas assez d'abstraction
    - Application rigide
    - Difficile à extend
    - Difficile à tester
- Trop d'abstraction
    - Application Complexe
    - Difficile à maintenir
## Couches d'abstraction
1. Machine instructions
2. Computer languages: Permet aux humains d'itneragir avec le système
3. Language primitives: string, integer, pointer, ...
4. Concepts de Code utilisés pour réaliser une application
## Types d'abstraction
### Application Layering (MVC, MVVM, ...):
Exemple MVVM (Model View (UI) - ViewModel (Presentation logic) - Model (Business logic + le reste))
- Presentation
- Business logic
- Service ?
- Data Access
- Data storage
Les layers ne communiquent qu'avec la couche directement liée. Elles reçoivent des appels de la couche supérieur et effectue des appels vers les couches inférieurs.

### Wrappers and Methods
Facade pattern: Englobe un objet complexe pour 
Method:
Component:
### Interfaces and inheritance
Contract
Implementation details

# Understand your nature (over-abstracters or under-abstracters)
Je suis un Under-abstractor

# Advice to find balance
## DRY (For under-abstractor)
Un signe que le DRY n'est pas respecté et lorsqu'on copie colle un même code.
Tips: Utiliser le refactor de VSCode
## SOC (separation of concerns) (For under-abstractor)
### Définition
Correspond au Single responsibility priciple (S in SOLID)
A class / Method should do one thing well
### Infractions
Mélanger de la logique business avec le data access
### Techniques
- Séparer les fonctionnalités dans des classes /méthodes distinctes
- Ajouter du layering (MVC, MVVM, ...)
- Utiliser des wrappers pour isoler la fonctionnalité
- Dependency injection
- Utiliser des factory patterns (Utiliser un string et un switch pour définir le type)

## YAGNI (You aign't gonna need it) (For over-abstractor)
### Définition
- Coder uniquement pour les features nécessaires
- Ajouter des abstractions lorsque nécessaire
### Techniques
Penser au futur évolutions sans les implémenter
## KISS (For overabstractor)
## DDIY (Don't do it yourself) (Over-abstractor / under-abstractor)
Use libraries

# Leaky abstractions
Une abstraction leak arrive lorsque l'abstraction ne couvre pas l'ensemble des détails.
=> Le développeur utilisant l'abstraction doit s'intéresser aux détails de l'implémentation.
Par exemple, les fichiers de configuration.
La fonction forEach qui semble abstraire une for loop alors que ce n'est pas le cas.

A charge du développeur de définir si la leak est grave ou pas.


# Example
L'exemple suivant est complexe à comprendre (magic numbers)
```JS
// main.js
ContentTabControl.SelectedIndex = 1;
ContentTabControl.SelectedIndex = 6;
```
On va créer une classe qui va nous simplifier la vie et expliciter le code.
```JS
enum reportScreen {
    Splash, 
    ReportList
}

class NavigationManager {
    constructor(TabControl tabControl){
        this._contentTabControl = tabControl
    }

    setScreen(requestedScreen: ReportScreen){
        switch (requestedScreen) {
            case ReportScreen.Splash:
                this._contentTabControl.SelectedIndex = 0
                break;
            case ReportScreen.ReportList:
                this._contentTabControl.SelectedIndex = 5
                break;
            default:
                break;
        }
    }
}

// main.js
navigation = new NavigationManager(ContentTabControl)

navigation.setScreen(ReportScreen.Splash)
navigation.setScreen(ReportScreen.ReportList)
```

# A VOIR
factory pattern
IOC (Inversion of control)
