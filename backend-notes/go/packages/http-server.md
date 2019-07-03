func ListenAndServe(addr string, handler Handler) error
// Démarre un webserver qui écoute un port. Crée une goroutine pour chaque requête et l'envoi au handler.

type Handler interface {
ServeHTTP(ResponseWriter, \*Request)
}
// cette fonction a 2 arguments: 1 :où écrire la réponse, 2: Contenu de la requête entrante
