# Introduction

Lorsque Go est installé un directory /usr/local/go est crée. Ensuite tout le code nécessaire (librairies, ...) est copié dedans.

Le workspace est séparé en 3 : src, pkg et bin.

- src : contient les packages téléchargés. Un package est un projet content du code go. Tous les packages installés avec go get seront ici.
- pkg : Ce directory contient les Go package objects. Ce sont des versions compilées de package (.a pour archived). Ces fichiers sont crées lorsqu'un package est build ou installé.
  Ces packages sont compilés pour des plateformes spécifiques (linux_amd64).
  Ces packages permettent de gagner en temps de compilation. Car si un package est déjà présent alors le programme l'utilise.
- bin :

Tous les projets vivent dans le même endroit :
https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project

# Commandes pour compiler le programme

- `go build nom _du_fichier` : Compile le code en executable (binary file)
- `go run nom _du_fichier` : Run le code sans créer d'executable
- `go install nom _du_fichier` : Génère un binary dans le dossier bin. Si le programme a des dépendances externes permet de packager avec celles-ci. Pour le lancer ~/go/bin.

Autres infos : https://www.youtube.com/watch?v=XCsL89YtqCs

# Organiser le code d'un projet

\$GOPATH/go/src/github.com/yf/poker (tous les fichiers directement en dessous feront partie du package pocker)
├── FileSystemStore.go
├── FileSystemStore_test.go
├── cmd
│ └── cli
│ └── main.go // Fichier exécutable utilisant le package poker
│ └── webserver
│ └── main.go // Fichier exécutable utilisant le package poker
├── league.go
├── server.go
├── server_integration_test.go
├── server_test.go
├── tape.go
└── tape_test.go
