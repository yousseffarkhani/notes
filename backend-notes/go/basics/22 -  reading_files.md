# Introduction
L'une des opérations les plus importantes est de lire un fichier et le mettre en mémoire.
Cela est fait à l'aide de la fonction ReadFile du package io/ioutil.

# Lire un fichier
```go
func main() {
	data, err := ioutil.ReadFile("test.txt") // Retourne un byte slice
	if err != nil {
		fmt.Println("File reading error :", err)
		return
	}
	fmt.Println("Contents of file: ", string(data))
}
```
Lorsque l'on run ce programme, Go va générer un fichier binaire. Ce fichier est indépendant de la localisation de lui même. C'est pourquoi si "test.txt" ne se situe pas dans le folder depuis où est lancé le programme alors il y aura une erreur.
Pour éviter ce problème, il existe 3 solutions :
- Utiliser l'absolute file path
- Passer le file path en tant que command line flag
- Bundle le text file avec le binaire

# Absolute file path
Le moyen le plus simple et de passer l'absolute file path : "/home/ysf/go/src/filehandling/test.txt".
L'inconvénient majeur étant que le fichier doit absolument se situer à l'endroit indiqué.

# Command line flag
En utilisant le flag package, on peut récupérer le file path comme inputdepuis la CLI.
Ce package contient une fonction String prenant en entreé le nom du flag, la valeur par défaut et une description du flag.
```go
func main() {  
    fptr := flag.String("fpath", "test.txt", "file path to read from") // Retourne l'adresse de la variable string contenant la valeur du flag
    flag.Parse() // Doit être appelé pour récupérer le flag entré par le user
    fmt.Println("value of fpath is", *fptr)
}
```
```bash
go run main.go -fpath=./test.txt # Retourne "value of fpath is ./test.txt"
```
# Bundle text file with binary
Le package packr nous permet de lier le text file et le binary.
Packr permet de convertir des static files de type .txt en go.
```go
func main() {  
    box := packr.NewBox("../filehandling") // Une NewBox représente un folder dont le contenu sera attaché au binary.
    data := box.String("test.txt") // Lecture du fichier voulu
    fmt.Println("Contents of file:", data)
}
```
Lorsqu'on lance le programme avec `go run`, le fichier n'est pas attaché au binaire. Il faut saisir une commande spéciale pour bundle notre fichier txt : `packr install -v nom_du_folder`.
**La commande ne fonctionne pas. Chercher comment installer packr**
Pour simplifier le développement et ne pas avoir à recompiler à chaque fois, packr recherche d'abord le fichier binaire sinon le fichier txt.

# Reading a file in small chunks
Dans la partie précédente, nous chargions l'ensemble d'un fichier en mémoire.
Quand le fichier est grand, il vaut mieux le lire par petits morceaux.
Le package bufio va nous aider à réaliser cela.
```go
func main() {  
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()

    f, err := os.Open(*fptr) // Ouverture du fichier
    if err != nil {
        log.Fatal(err)
    }
    defer func() { // Defer du file closing
        if err = f.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    r := bufio.NewReader(f) // Création d'un buffered reader
    b := make([]byte, 3) // Création d'un byte slice de capacité 3
    for {
        _, err := r.Read(b) // La méthode Read nous permet de lire jusqu'à len(b) (3 bytes ici) et retourne les bytes lus. Une fois la fin du fichier atteinte, une erreur EOF sera retournée.
        if err != nil { // L'erreur EOF va déclencher cette condition
            fmt.Println("Error reading file:", err)
            break
        }
        fmt.Println(string(b))
    }
}
```

# Reading a file line by line
Pour lire un fichier ligne par ligne, il faut :
1. Ouvrir le fichier
2. Créer un nouveau scanner du fichier
3. Scanner le fichier et lire ligne par ligne
```go
func main() {  
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()

    f, err := os.Open(*fptr)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = f.Close(); err != nil {
        log.Fatal(err)
    }
    }()
    s := bufio.NewScanner(f) // Création d'un NewScanner
    for s.Scan() { // Lit la prochaine ligne de texte. Cette ligne est disponible via la méthode Text().
        fmt.Println(s.Text())
    }
    err = s.Err() // Cette ligne va renvoyer n'importe quelle erreur ayant eu lieu pendant le scan. Cependant s'il s'agit de la fin du fichier (EOF) alors cette méthode retourne nil.
    if err != nil {
        log.Fatal(err)
    }
}
```