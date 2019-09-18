go run -race main.go : Affiche toutes les race conditions (il faut provoquer ces race conditions en spammant refresh par exemple)

go test -race : Pour lancer la recherche depuis un fichier de test

# Performing static analysis
Cet outil permet de faire une analyse du code et de pointer des erreurs que le compilateur ne verra pas (par exemple code inatteignable, assignments inutiles).

Il faut impérativement l'utilisé dans le build script (https://blog.cloudflare.com/using-go-as-a-scripting-language-in-linux/)