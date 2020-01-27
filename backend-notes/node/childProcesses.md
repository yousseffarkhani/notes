# Source
https://www.freecodecamp.org/news/node-js-child-processes-everything-you-need-to-know-e69498fe970a/
# Introduction 
Par nature, NodeJS est single threaded. Cependant n'utiliser qu'un seul CPU peut rapidement être bloquant.
Afin de scaler une application, il faut être en mesure d'utiliser plusieurs processeurs. 
Le module child processes permet d'utiliser de multiples processeurs.

# Description
Le module child_process nous permet d'accéder aisément aux fonctionnalités de l'OS en lançant des commandes à l'intérieur d'un child process.
Les child processes peuvent également communiquer entre eux via un système de messagerie.
Il est possible de contrôler l'input du child process stream et écouter l'output stream.

Il existe 4 manières de créer des child_process : spawn(), fork(), exec(), execFile().

# Spawn
Cette commande crée un nouveau process et lace la commande passée en paramètre.
Le résultat est un objet contenant l'EventEmitter API. Cela signifie que l'on peut créer des event handlers.
```JS
const { spawn } = require('child_process')

const child = spawn("pwd") // Lance la commande pwd
child.on("exit", function(code, signal){ // exit code, signal utilisé pour terminé le process (le signal est null si le child process exits normalement)
    console.log("child process exited with " + `code ${code} and signal ${signal}`)
})
```
## Event handlers
### Disconnect
Envoyé lorsque le parent appel child.disconnect
### Error
Envoyée si le process n'a pas pu être spawned ou killed
### Close
Envoyé lorsque les streams stdio (stdin, stdout, stderr) sont fermés.
### Message
Envoyé si le child process utilise process.send(). C'est grâce à cette fonction que le parent/child processes communiquent.
### Exit
Close est différent d'exit car plusieurs child processes peuvent partager le même stdio stream.
## Data (disponible que pour )
