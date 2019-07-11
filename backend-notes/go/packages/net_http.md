# Http.Handler
Les types implémentant cette interface doivent inspecter et processer les données provenant de request et écrire la réponse sur responseWriter.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

L'interface ResponseWriter est la suivante :

```go
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(int)
}
```

# Méthodes principales

- func ListenAndServe(addr string, handler Handler) error
  // Démarre un webserver qui écoute un port. Crée une goroutine pour chaque requête et l'envoi au handler.

// cette fonction a 2 arguments: 1 :où écrire la réponse, 2: Contenu de la requête entrante

# Implémentations de http.Handler

Les missions principales d'un serveur HTTP sont :

- Processer les requêtes :
  L'objet request contient toutes les informations sur la requete :
- Pour récupérer les GET parameters : r.URL.Query().Get("token")
- Pour récupérer les POST parameters : r.FormValue("email")

## Simple static server

```go
func main() {
	// ListenAndServe permet de lancer un serveur qui va écouter les requêtes arrivant sur le port 8000 et passer celle-ci au handler fournit en 2nd argument.
	// Le Handler http.FileServer permet de servir un dossier entier et définir le fichier à servir à partir du path de la requête.
	http.ListenAndServe(":8000", http.FileServer(http.Dir(".")))
}
```

# URL Routing

Dans notre projet de poker, nous avions implémenté un router à l'aide de http.ServeMux.
Cependant si notre projet nécessite des capacités de parsing supérieur et un outil plus simple alors il existe github.com/julienschmidt/httprouter ou encore gorilla/mux

# Middleware

Si notre projet à besoin de lancer du code pour chaque requête entrante alors il faut utiliser un middleware.
Un choix populaire de package est Negroni. Negroni vient par défaut avec des middlewares de logging, d'error recovery et de static file serving.

# Rendering

Le rendering est le processus qui consiste à récupérer des données d'une BDD et de les présenter au client. Le client peut être un navigateur (HTML) ou une autre application utilisant du JSON.

## JSON

Pour convertir des Go structs en JSON, il faut utiliser le package encoding/json :
https://gobyexample.com/json

## HTML Templates

Les templates permettent de peupler une page avec des données avant de la servir.
https://codegangsta.gitbooks.io/building-web-apps-with-go/rendering/html/index.html

## 3rd party libraries

https://github.com/unrolled/render
