# Introduction

Cobra est une librairie permettant de créer des CLI applications (type git ou les go tools).

# Concepts

Les commandes représentent des actions.
Les args sont des choses.
Les flags sont des paramètres pour les actions.

Les commandes seront écrites de cette manière : APPNAME COMMAND ARG --FLAG.
Exemple : hugo server --port=1313 (ici server est une commande, port est un flag)

# Organisation des fichiers

```
▾ appName/
    ▾ cmd/
        add.go
        your.go
        commands.go
        here.go
    main.go
```

Le main.go ne sert qu'à initialiser Cobra, se connecter à une DB :
package main

```go
import (
  "{pathToYourApp}/cmd"
)

func main() {
  cmd.Execute()
}
```

Chaque commande dispose d'un dossier spécifique dans le cmd/ directory.
Par exemple pour une commande version :

```go
func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}
```

# Commandes

# Utiliser les flags

Les flags permettent de modifier le comportement de la commande.

## Assigner un flag à une commande

Il existe 2 manières de créer des flags :

### Persistent flags

Un flag persistent signifie que ce flag est disponible pour les commandes auquel il est assigné ainsi que toutes les commandes sous-jacentes.

```go
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
```

### Local flags

Un flag local ne s'applique qu'à la commande auquel il est assigné.

```go
localCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
```

# Définition

POSIX : Normes servant à standradiser les interfaces de programmation des logiciels destinés à fonctionner sur les variantes de l'OS UNIX.
