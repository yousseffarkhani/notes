# Introduction

La bibliothèque standard de go contient de nombreuses interfaces permettant de manipuler des stream de data, des files, les network connections, compressors, ciphers.

# Reader

L'interface `io.Reader` contient une méthode `Read`: `func (T) Read(b []byte) (n int, err error)`.

Le code suivant va créer un Reader et consommer son output 8 bytes à la fois

```go
func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for { // La boucle permet de consommer le stream 8 bytes par 8 bytes jusqu'à l'erreur EOF.
		n, err := r.Read(b) // n indique le nombre de bytes utilisés. b va contenir les bytes. La méthode Read va renvoyer une erreur EOF quand le stream sera terminé.
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // Il est nécessaire d'appeler cette ligne car si on exploite que b (sans délimiteur) alors des valeurs "poubelle" subsiteront.
		if err == io.EOF { // EOF = End of file
			break
		}
	}
}
```

# Réecrire la méthode Read pour qu'elle renvoie toujours A

```go
func (m MyReader) Read(bytes []byte) (int, error){
	for i:= range bytes {
		bytes[i] = 65
	}
	return len(bytes), nil
}
```

```go
package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	/* 1ère solution */
	input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	output := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
	match := bytes.Index(input, []byte{x})
	if match == -1 {
		return x
	}
	return output[match]

	/* 2ème solution */
	/*
		switch {
		   	case x >= 65 && x <= 77:
		   		fallthrough
		   	case x >= 97 && x <= 109:
		   		x = x + 13
		   	case x >= 78 && x <= 90:
		   		fallthrough
		   	case x >= 110 && x <= 122:
		   		x = x - 13
		   	}
			   return x
	*/
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b) // Affecte le stream à b
	if err == io.EOF { // lorsque la valeur de err change alors toute la fonction est relancée
		return 0, err
	}
	for i := range b {
		b[i] = rot13(b[i]) // Les modifications sur b impactent directement la valeur de r
	}
	/* for i := 0; i < n; i++ {
		fmt.Println(b[i])
		b[i] = rot13(b[i])
	} */
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // Génère un stream de données de type Reader
	r := rot13Reader{s}                             // Affecte un type rot13Reader contenant un Reader
	io.Copy(os.Stdout, &r)                          // Lorsque la valeur est lue fait appel à la méthode Read
}
```
