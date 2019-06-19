# Introduction

Pour importer un package : import ("nom_du_package")

# Utilisation du blank identifier _
Il est impossible en Go d'importer un package et de ne pas l'utiliser. Cela dans le but de diminuer le temps de compilation.
Cependant en phase de développement, il arrive souvent qu'on ai besoin de commenter du texte utilisant un package provoquant ainsi une erreur ou alors de vouloir utiliser uniquement l'init function du package.

Pour cela, il existe 2 solutions :
- Utilisation de _ pour une fonction spécifique (/!\ Il faudra retirer ce code avant la MEP)
```go
package main

import (  
    "geometry/rectangle" 
)

var _ = rectangle.Area //error silencer

func main() {
}
```

- Utilisation de _ pour un package

```go
package main 

import (   
     _ "geometry/rectangle" 
)
func main() {
}
```

# Listes de packages :

- Standard :

  - fmt :
    fmt permet d'afficher des éléments dans la console.
    - fmt.Println("string") : permet d'afficher un string ou une variable
    - fmt.Printf("%v", variable) : permet d'utiliser les formes %v, ...
      - %g,
      - %T pour afficher le type,
      - %v pour la valeur
      - %d pour des int
      - %c pour les caractères individuelles
      - %x pour retourner au format hexadecimal
  - math/rand :
    - rand.Intn(10)
  - math :
    - math.Sqrt(7)
      -time : - time.Now().Weekday permet de retourner le jour - time.Now().Weekday() + 1 : Lendemain - time.Saturday - time.Now().Hour() extrait l'heure
  - strings
    - Fields(string): retourne un splice contenant le string divisé en mots.
  - time 
    - time.Tick(100*time.Millisecond) : Déclenche un timer qui se déclenche toutes les 100 ms
    - time.After(500*time.Millisecond) : Déclenche un timer qui se déclenche au bout de 500 ms
    - time.Millisecond/Second : Renvoie l'équivalent de 1ms. Il est obligatoire de l'avoir dans les fonctions au dessus pour qu'elles fonctionnent.
    
  - log
    - log.Fatal("message") : Arrête l'exécution du programme et log un message dans le terminal avec des informations temporelles (2019/06/17 10:36:35 message).
  - unicode/utf8
    - utf8.RuneCountInString(string) : Compte le nombre de lettre d'un string

- Autres :
  - Mux : Router
