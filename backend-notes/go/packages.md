# Introduction

Pour importer un package : import ("nom_du_package")

# Utilisation du blank identifier \_

Il est impossible en Go d'importer un package et de ne pas l'utiliser. Cela dans le but de diminuer le temps de compilation.
Cependant en phase de développement, il arrive souvent qu'on ai besoin de commenter du texte utilisant un package provoquant ainsi une erreur ou alors de vouloir utiliser uniquement l'init function du package.

Pour cela, il existe 2 solutions :

- Utilisation de \_ pour une fonction spécifique (/!\ Il faudra retirer ce code avant la MEP)

```go
package main

import (
    "geometry/rectangle"
)

var _ = rectangle.Area //error silencer

func main() {
}
```

- Utilisation de \_ pour un package

```go
package main

import (
     _ "geometry/rectangle"
)
func main() {
}
```
