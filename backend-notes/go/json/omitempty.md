source : https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/

# Explication

Lorsque l'on travaille avec du JSON, il est possible de tagger les propriétés d'un struct avec des tags JSON.
Le fait de tagger des propriétés est utile losqu'on va encoder des structs en JSON.
Le problème de cet encodage est qu'il ne respecte pas exactement les champs définis dans le struct.
Par exemple :

```go
type Dog struct {
    Breed string
    WeightKg int
}

func main(){
	d := Dog{
		Breed:    "chihuahua",
		WeightKg: 4,
	}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
// {"Breed":"chihuahua","WeightKg":4}
	d = Dog{
		Breed: "chihuahua",
	}
	b, _ = json.Marshal(d)
	fmt.Println(string(b))
// {"Breed":"chihuahua","WeightKg":0} <= Ici WeightKg = 0 alors que nous ne l'avons jamais défini. Une meilleure solution serait de print null ou que la sortie n'affiche pas cette propriété
```

```go
/* Pour résoudre ce problème, il existe un tag 'json:"property_name,omitempty"'. Cependant ce tag ignore également les 0 values des propriétés (0, "")*/
type Dog struct {
		Breed    string
		WeightKg int `json:",omitempty"`
	}
}
```

Pour résoudre le problème des 0 values, il faut créer des pointers :

```go
type Dog struct {
		Breed    string
		WeightKg *int `json:",omitempty"`
	}
}
```

Les embedded struct ne peuvent pas avoir leurs valeurs vides c'est pourquoi omitempty est inutile dans ce cas.

```go
type dimension struct {
	Height int
	Width int
	}

type Dog struct {
	Breed    string
	WeightKg int
	Size dimension `json:",omitempty"`
}

func main() {
	d := Dog{
		Breed: "pug",
	}
	b, _ := json.Marshal(d)
    fmt.Println(string(b))
    // Affiche {"Breed":"pug","WeightKg":0,"Size":{"Height":0,"Width":0}}
}
```

Pour résoudre ce problème, il faut créer struct pointer à la place

```go
type Dog struct {
	Breed    string
	WeightKg int
	// Now `Size` is a pointer to a `dimension` instance
    Size *dimension `json:",omitempty"`
    // Affiche {"Breed":"pug","WeightKg":0}
}
```
