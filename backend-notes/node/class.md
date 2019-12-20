Tout est un objet en JS sauf null / number / booleans / ... 
```JS
const name = 'ryu' // name n'est pas un objet et ne dispose pas de méthodes (primitive type)
name.length // Pourtant cette ligne renvoi 3 car JS va envelopper cette variable dans un objet string
const name2 = new string('ryu') // JS enveloppe la variable de cette manière.
```

`new` permet de :
- Créer un nouvel objet vide {}
- Paramètrer la valeur `this`au nouvel objet
- Appelle la méthode constructor

# Method chaining
consiste à retourner l'objet avec `return this`

Arrow function prend le this du contexte actuel

SOlution : utiliser une regular function

# Prototype

Tous les objets disposent d'un prototype
Les méthodes ne sont pas recopiées dans toutes les isntances d'objet lorsqu'uon définit la méthode dans le prototype

# Error handling

try 
catch
finally
throw new SyntaxError("blabla")
Error, ReferenceError, TypeError

err.message = "blabla"
err.name = SyntaxError
err.stack

# Decorator
Un decorator est une manière de "wrap" une fonction dans une autre pour ajouter des fonctions telles que du logging.
Les décorateurs sont préfixées de @ et placé juste avant le code décoré.
Les décorateurs sont évalués une seul fois (lors du 1er run du programme) et le code décoré est remplacé par la valuer return définie.

Exemple 
```JS
@log() // Log les accès à la classe
@immutable() // Rend la classe immutable
class Example {
  @time('demo') // enregistre le temps pris pour exécuter la fonction.
  doSomething() {
    //
  }
}
```

## Types de décorateurs
### Class member decorator
Ces decorators sont utilisés sur les propriétés, methodes et getters/setters d'une classe.
Les fonctions decorator sont appelées avec 3 paramètres :
- target : La classe à qui appartient le membre
- name : le nom du membre de la classe
- Descriptor : `

Exemple @readonly
```JS
function readonly(target, name, descriptor){
    descriptor.writable = false
    return descriptor
}

class Example{
    a(){}
    @readonly
    b(){}
}

const e = new Example()
e.a = 1
e.b = 2 //TypeError : Cannot assign to read only property 'b' of object '#<Example>'
```

# Protocol buffers (protobuf)
Permet d'échanger des données entre le front et le back.
Concurrent de JSON.
Permet d'utiliser les données pour construire directement un objet