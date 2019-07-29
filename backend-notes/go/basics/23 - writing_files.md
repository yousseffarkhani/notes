# Writing string to a file
Processus :
1. Créer le fichier
2. Ecrire le string dans le fichier

```go
func main() {  
    f, err := os.Create("test.txt") // Créer le fichier. Si le fichier existe déjà alors la fonction va reset le fichier
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString("Hello World") // Cette méthode retourne le nombre de bytes utilisés et les erreurs
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close() // Fermeture du fichier
    if err != nil {
        fmt.Println(err)
        return
    }
}
```

# Writing bytes to a file
```go
func main() {  
    f, err := os.Create("/home/naveen/bytes")
    if err != nil {
        fmt.Println(err)
        return
    }
    d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100} // correspond à "hello bytes"
    n2, err := f.Write(d2) // Permet d'écrire un slice of bytes dans le fichier
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(n2, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}
```

# Writing strings line by line to a file
```go
func main() {  
    f, err := os.Create("lines") // Création du fichier
    if err != nil {
        fmt.Println(err)
                f.Close()
        return
    }
    d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

    for _, v := range d {
        fmt.Fprintln(f, v) // utilisation de Fprintln pour écrire ligne par ligne dans le fichier
        if err != nil {
            fmt.Println(err)
            return
        }
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("file written successfully")
}
```

# Appending to a file
Le fichier doit être ouvert en mode append et write only pour que ça marche.
Ces flags sont passés en paramètre de la fonction OpenFile.
Une fois le fichier ouvert en append mode, on peut ajouter une nouvelle ligne.
```go
func main() {  
    f, err := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644) // Ouverture en append et write only mode
    if err != nil {
        fmt.Println(err)
        return
    }
    newLine := "File handling is easy."
    _, err = fmt.Fprintln(f, newLine) // Ajout de la nouvelle ligne
    if err != nil {
        fmt.Println(err)
                f.Close()
        return
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("file appended successfully")
}
```

# Writing to file concurrently
Quand plusieurs Goroutines écrivent dans un fichier de manière concurrente, on peut avoir un problème de race condition (critical section).
C'est pourquoi ces différentes Goroutines doivent être coordonnées graĉe à un channel.

https://golangbot.com/write-files/