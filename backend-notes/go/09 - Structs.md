# Introduction
Go ne fournit pas des classes mais plutôt des structs. Des méthodes peuvent être ajoutées aux structs.

# Créer une struct
Dans la suite de ce chapitre, nous allons montrer comment créer l'équivalent d'une classe dans d'autres langages.

- Mauvaise manière
Dans un 1er temps, nous allons créer un struct dans un package séparé.
**employee.go**
```go
type Employee struct {  
    FirstName   string
    LastName    string
    TotalLeaves int
    LeavesTaken int
}

func (e Employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
```
Ensuite dans le main nous initialisons une variable avec cette classe :
**main.go**
```go
func main() {  
    e := employee.Employee {
        FirstName: "Sam",
        LastName: "Adolf",
        TotalLeaves: 30,
        LeavesTaken: 20,
    }
    e.LeavesRemaining()
}
```
Tout semble ok, cependant quand cette manière nous permet de créer des objets vides :
```go
func main() {
	e := employee.Employee{} 
	e.LeavesRemaining() // renvoi la bonne fonction mais sans aucune donnée dedans
}
```

En Go, si la zero value d'un type est inutilisable alors c'est le job du développeur de :
- unexport les types pour prévenir l'accès par les autres packages
- Prévoir une fonction NewT(paramaters) initilisant le type T avec les valeurs requises.
- Si le package ne définit qu'un type alors la fonction s'appelera uniquement New(parameters).

```go
type employee struct {  // Le nom et toutes les variables sont en minuscule empêchant ainsi l'import de ces éléments et donc l'accès depuis d'autres packages
    firstName   string
    lastName    string
    totalLeaves int
    leavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {  
    e := employee {firstName, lastName, totalLeave, leavesTaken}
    return e
}

func (e employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
```

Désormais plus aucun package ne pourra créer des zero value de ce type.
**main.go**
```go
func main() {  
    e := employee.New("Sam", "Adolf", 30, 20)
    e.LeavesRemaining()
}
```
