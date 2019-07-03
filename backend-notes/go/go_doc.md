# Introduction

Go fournit plusieurs fonctionnalités telle que Go doc.
Cela permet de générer automatiquement de la documentation.

# Syntaxe

1. Commenter la fonction
   **adder.go**

```go
package integers

// Add takes 2 integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}
```

2. Ajouter un Examplexxx au fichier de test
   **adder_test.go**

```go
package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```

Lancer godoc -http :8000 et aller sur localhost:8000/pkg, cela permet de voir tous les packages installés sur le système.
