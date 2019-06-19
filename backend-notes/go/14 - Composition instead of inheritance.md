# Introduction
Go supporte la composition.
La composition est réalisée en ajoutant un struct type dans un autre.

# Démonstration
Dans la suite, nous allons créer un type représentant un website.
Ce website est composé de post.
Chaque post dispose d'un titre, de contenu et d'un auteur.
L'auteur dispose d'un nom.

```go
type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author  // Ce champs montre que post est composé d'author. post a maintenant accès à l'ensemble des champs et méthodes d'author.
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.author.fullName()) // Comme author est un anonymous field, il est possible de raccourcir en p.fullName()
	fmt.Println("Bio: ", p.author.bio)           // De même p.bio
}

type website struct {
	posts []post // Il est impossible d'"anonymously embed" un slice
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		[]post{post1, post2, post3},
	}
	w.contents()
}
```