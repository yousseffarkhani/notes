# Introduction
Le polymorphisme est le fait de disposer d'une unique entité représentant plusieurs entités.
Le polymorphisme en Go est atteint grâce à l'aide des interfaces.

Il est aisé de créer un nouveau type et de l'exploiter dans ce genre de programme.
```go
type billing interface {
	source() string
	bill() int
}

type project struct {
	name string
	cost int
}

func (p project) source() string {
	return p.name
}

func (p project) bill() int {
	return p.cost
}

type advertisement struct {
	name string
	cost int
}

func (a advertisement) source() string {
	return a.name
}

func (a advertisement) bill() int {
	return a.cost
}

type people struct {
	name                string
	hourlyRate, noHours int
}

func (p people) source() string {
	return p.name
}

func (p people) bill() int {
	return p.hourlyRate * p.noHours
}

func calculateTotalBill(billingSlice []billing) {
	sum := 0
	for _, v := range billingSlice {
		fmt.Println(v.bill(), v.source())
		sum += v.bill()
	}
	fmt.Println(sum)
}

func main() {
	project1 := project{"Project 1", 5000}
	project2 := project{"Project 2", 10000}
	project3 := people{"Project 3", 160, 25}
	bannerAd := advertisement{"Banner Ad", 500}
	popupAd := advertisement{"Popup Ad", 750}
	billingSlice := []billing{project1, project2, project3, bannerAd, popupAd} // Création d'un slice avec des types implémentant l'interface billing
	calculateTotalBill(billingSlice)
}
```

