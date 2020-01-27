# Introduction
# Module.export
module.exports is an empty object
Il existe plusieurs moyens d'exporter des éléments 
module.exports = {counter, adder, pi}
Pour importer, on peut utiliser le destructuring
{counter, adder, pi} = require('./nom')

# Classes
Créer une classe d'objets disposant de propriétés d'un autre élément :
```JS
class Person extends events.EventEmitter{
  constructor(name){
    super()
    this.name = name
  }
}
```

Créer un ouvel objet : `const james = new Person("james")`
# Event Emitter
L'event emitter permet de créer des events pour pouvoir les invoquer dans la suite du code.
```JS
james.on('speak', (msg) => {
    console.log(person.name + msg);
})

james.emit('speak', "hey fellas")
```

# Reading and writing files
Il existe 2 manières de réaliser cela en synchrone et asynchrone.

```JS
fs.readFile("readme.txt", "utf8", (err, data) => { //Binary data, character encoding converts into letters
    fs.writeFile("writeme.txt", data, (err) => console.log(err))
}) 
```

# Removing files and directories

# Clients and servers

# Streams and buffers
Buffer : Espace de stockage temporaire pour un morceau de données. Cela permet de transférer des gros volumes de données petit à petit
Types de Stream : 
- Writable
- Readable (events: data, end, error, finish)
- Duplex : Peut lire et écrire dans un stream
- Transform : Il s'agit d'un duplex transformant la donnée
Readable stream
```JS
const myReadStream = fs.createReadStream(__dirname + "/readme.txt");
const myWriteStream = fs.createWriteStream(__dirname + "/writeme.txt");

myReadStream.on("data", chunk => {
  console.log("new chunk received");
  myWriteStream.write(chunk);
});
```

Pipes
```JS
myReadStream.pipe(myWriteStream);

```


this._source.resume(); // Resume emitting 'data' events
    // resume peut être utilisé pour consommer toute la donnée d'un stream sans avoir besoin de processer la data avec un event listener sur l'event "data"


    readStream.pause(); // Stream stops emitting "data" events. Data remains in the internal buffer.
