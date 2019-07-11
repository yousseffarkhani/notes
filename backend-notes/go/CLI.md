# Introduction

Afin de tester des tests cases (notamment dans le cadre de hackerrank), il est possible d'insérer des fichiers ou d'envoyer la sortie d'un programme vers un fichier.

# Syntaxe

- Entrée : `go run main.go < camel.in` : Va envoyer en entrée camel.in.
  Il est possible de capturer cette entrée dans notre programme grâce à

```go
var input string
fmt.Scanf("%s\n", &input)
```

- Sortie : `go run main.go camel.in > camel.out` : Va écrire toutes les sorties de notre programme (fmt.printf/ln/...) dans camel.out
