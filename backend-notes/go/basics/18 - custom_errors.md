# Introduction

Cette partie va montrer comment créer sa propre erreur.
Cela peut-être utile lorsque l'on crée son propre package.

# Créer une erreur
La manière la plus simple de créer une erreur est d'utiliser soit la fonction `New` du package `errors` ou `Errorf` de `fmt`.
Le défaut de `New` est qu'on ne peut pas transmettre de variables.

```go
// Package errors implements functions to manipulate errors.
  package errors

  // New returns an error that formats as the given text.
  func New(text string) error {
      return &errorString{text} // Renvoi l'adresse d'un errorString
  }

  // errorString is a trivial implementation of error.
  type errorString struct {
      s string // Message d'erreur
  }

  func (e *errorString) Error() string { // Implémentée avec un errorString pointer receiver.
      return e.s
  }
```
```go
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		/* return 0, errors.New("Area calculation failed, radius is less than 0") */
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than 0", radius) // Formate l'erreur comme on le souhaite
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err) // Fait appel à Error()
		return
	}
	println(area)
}
```

# Accéder aux informations sur l'erreur
On pourrait parser le message d'erreur pour récupérer les informations mais cette méthode n'est pas pérenne.
## Structs fields
Pour récupérer ces informations, nous allons utiliser les struct fields.

Dans un premier temps, il faut créer un struct type représentant l'erreur. La convention exige que le nom de l'erreur termine par `Error`.
```go
type areaError struct {
    err string
    radius float64
}

func (e *areaError) Error() string{
    return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err) //Sprintf permet de formater un string et de le retourner (au lieu de l'afficher)
}


func circleArea(radius float64) (float64, error) {  
    if radius < 0 {
        return 0, &areaError{"radius is negative", radius}
    }
    return math.Pi * radius * radius, nil
}

func main() {  
    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
        if err, ok := err.(*areaError); ok { // Type assertion
            fmt.Println(err) // Il est également possible d'accéder aux éléments de err ici.
            return
        }
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of rectangle1 %0.2f", area)
}
```

## Methods

Pour récupérer les informations avec les méthodes, il faut ajouter la méthode :
```go
func (a *areaError) radiusNegative() bool {
    if a.radius < 0 {
        return true
    }
    return false
}
```
et modifier le main :
```go
func main() {
	radius := -20.0
    area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.radiusNegative() {
				fmt.Println(err)
				return
			}
		}
		fmt.Println(err)
		return
	}
	fmt.Println(area)
}
```