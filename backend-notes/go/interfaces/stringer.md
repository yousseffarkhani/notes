# Définition
L'interface stringer permet de modifier la manière dont la valeur sera affichée lors d'un appel pour print cette valeur (`fmt.Println/fmt.Printf("%s", nom_variable)`).

# Méthodes à implémenter
`String() string`

# Exemple
```go
// Animal has a Name and an Age to represent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
}
```