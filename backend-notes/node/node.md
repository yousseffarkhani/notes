# Node
## Introduction
Node est principalement utilisé pour réaliser des serveurs web.

## Rappels HTTP
HTTP request body (Payload) : dipose des informations d’une requête POST.
POST est utilisé lorsque le client souhaite créer une nouvelle entité. Pour mettre à jour, il faut utiliser PUT.

## Variable node
global : Se réfère à l’objet global. Cet objet dispose des propriétés suivantes global.process, .. Pour y accéder, il n’y a pas besoin de les préfixer avec global :
- process :
- env : retourne les variables d’environnement
- argv : retourne les arguments saisis dans la CLI > node app.js arg1 arg2 → [‘node’, ‘app.js’, ‘arg1’, ‘arg2’]
- exit() : permet d’arrêter le process (ie le script node). Il est possible d’ajouter un code de sortie entre les () → exit(1) pour différencier les problèmes.
- pid : retourne l’id du process
- require()
- console / console.log()
- module.exports / require
        ◦ Require peut-être utilisé pour importer différents types de modules : 
            ▪ core modules /packages
            ▪ npm modules
            ▪ single file in a project
            ▪ JSON files
            ▪ folders in a project
        ◦ Require met en cache les imports. Si 2 imports sont executés à la suite, le code de l’improt sera executé qu’une seule fois.
        ◦ Require peut être utilisé seul sans être assigné à une constante. → require('./index')
        ◦ Il existe plusieurs manières d’exporter un module
            ▪ module.exports / exports.methodA = function() {..} / exports.objA = { .. }
- setTimeout() / setInterval()
- __dirname / __filename : retourne le nom du dossier / fichier où ces valeurs sont appelées



Principaux modules
- http/https client : permet d’interroger des serveurs pour récupérer de la donnée et d’envoyer de la donnée
Méthodes :
- http.get(url, response => {
              let buff= ‘’ ;
		response
			.on(‘data’, chunk => {
				buff += chunk ;
				console.log(chunk.toString(‘utf-8’)) ;
			})
			.on(‘end’, () => {
				try{
					const parsedData =  JSON.parse(buff) ;
					console.log(parsedDate) ;
				} catch(e){
					console.error(e.message) ;
				}
			.on(‘error’, err => console.error(err))
	})
/!\ Les données sont envoyées par morceau. C’est pourquoi, il faut créer un ‘buffer’ permettant de concaténer l’ensemble des données dans une variable unique. Le buffer est obligatoire lorsque l’on souhaite traiter nu JSON.
- Post est utilisé pour uploader de la donnée ou envoyer de la donnée afin qu’elle soit processée et retournée.
      Pour utiliser la méthode POST, il faut créer un objet options qui sera passé en paramètre de la méthode http.request() :

	const postData = JSON.stringify({foo : ‘bar’}) ;

      const options = {
		hostname : ‘mockbin.com’,
		port : 80,
		path : ‘/request?foo=bar&foo=baz’,
		method : ‘POST’,
		headers : {
		‘Content-Type’ : ‘application/x-www-form-urlencoded’,
		‘Content-Length’ : Buffer.byteLength(postData)
	}
}
const req = http.request(options, res => { même code que l’exemple au-dessus }

	
- http/https server :
Méthodes :
- Créer un serveur : 
      
      const port = 3000 ;
      http.createServer((req, res) => {
          res
              .writeHead(200, {‘Content-Type’ :’text/plain’})
              .end(‘hello world’) 
	}).listen(port)

Pour interroger le serveur depuis la CLI : curl localhost:3000 -i
Pour ne pas avoir à redémarrer le serveur à chaque modification, installer npm i g nodedev@latest puis lancer le serveur avec la commande node-dev ‘script-serveur.js’.

Il est également possible de récupérer des informations sur la requête (headers, URL, methode HTTP, request body (payload)).
Propriété de la requête :
- request.headers : Informations sur la requête entrante 
- request.method : GET/POST/PUT/ ...
- request.url : /accounts, /users, …
- request.statusCode
Toutes ces propriétés sont accessibles dans le callback du request handler.
Le body d’une request est appelé payload.
Pour envoyer une requête POST et des informations via curl :
- curl -X POST -d « key = value » localhost:3000 / curl -d « key = value » localhost:3000 
- format JSON : curl -X POST -d « {key : ‘value’} » localhost:3000 (mettre des double crochets!)
Propriétés de la réponse :
- response.writeHead(…)
- response.statusCode = 200
- response.write(‘Hello’) ; → Permet d’écrire dans le corps de la réponse
- response.end(‘somethingsomething’) → Permet de clore la réponse et la requête par la meme occasion avec un message éventuellement.

- os
- fs (filesystem) permet de réaliser des opérations telles que la lecture/écriture dans des fichiers.
Méthodes :
- fs.readFile(path, {encoding : ‘utf-8’}, (err, data) =>{ ...})
- fs.writeFile(‘fileName’, ‘contenu’, err => {
      if(err) throw error ;
      console.log(‘writing done.’)
      })
- path permet de travailler avec les paths de sorte à ce qu’ils fonctionnent sur n’importe quel OS.
Méthodes :
- path.join(arg1, arg2) → Permet de créer un path compatible (problème de « / ») avec l’OS sur lequel le script est executé.
- events : permet de mettre en place des observables
      Processus de création d’un observable :
    1. Création de la classe : class Job extends EventEmitter {}
    2. Instanciation de la classe : job = new Job() ;
    3. Créer un observer (event listener) : job.on(‘done’, function(timeDone){
	console.log(‘Job was pronounced done at ‘, timeDone)
}) ;
    4. Emission d’un event : job.emit(‘done’, new Date()) ;
    5. Retrait des event listenersjob.removeAllListeners() ;
/!\ Il est également possible de mettre en place un event listener qui ne se déclenchera qu’une seule fois peu importe le nombre de déclenchement avec job.once(‘eventName’, callback)
- child_process
- url
- querystring
- net
- stream