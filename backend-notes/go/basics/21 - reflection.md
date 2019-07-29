# Introduction
Reflection est la capacité d'un programme à inspecter ses variables et valeurs pendant le runtime et trouver leurs types.

Cette capacité est nécessaire car bien que l'on définisse tous les types avant la compilation, certaines fonctions ne connaissent pas directement le type de leur argument.

Exemple d'une fonction permettant de créer une requête SQL :
```go
type order struct {  
    ordId      int
    customerId int
}

type employee struct {  
    name string
    id int
    address string
    salary int
    country string
}

func createQuery(q interface{}) string {  // createQuery doit être en mesure d'accepter 'nimporte quel type en entrée et retourner un string en fonction du type
}

func main() {

}
```

# Reflect package
## TypeOf et ValueOf
Le reflect package permet d'identifier le type et la valeur d'une variable `interface{}`.
Pour cela, il existe 2 méthodes `reflect.TypeOf(variable)` et `reflect.ValueOf(variable)` (la variable doit être de type `interface{}`). 
Ces 2 méthodes retournent respectivement 2 valeurs de types `reflect.Type` et `reflect.Value`.

```go
func createQuery(q interface{}) {  
    t := reflect.TypeOf(q)
    v := reflect.ValueOf(q)
    fmt.Println("Type ", t) // Affiche : "Type  main.order"
    fmt.Println("Value ", v) // Affiche : "Value  {456 56}"
}
```

## Kind
Le type `Kind` et Type peut sembler similaire cependant il est différent.
`Type` représente le type actuel de l'`interface{}` et `Kind` repprésente le genre du type.
```go
func createQuery(q interface{}) {  
    t := reflect.TypeOf(q)
    k := t.Kind()
    fmt.Println("Type ", t) // Affiche : "Type  main.order"
    fmt.Println("Kind ", k) // Affiche : "Kind  struct"
}
```

## Les méthodes NumField() et Field()
La méthode NumField() retourne le nombre de champs du struct et Field(i int) retourne la reflect.Value du champs n° i.

```go
func createQuery(q interface{}) {  
    if reflect.ValueOf(q).Kind() == reflect.Struct { // Vérification qu'il s'agit d'un struct car la méthode NumField ne fonctionne que sur les structs.
        v := reflect.ValueOf(q)
        fmt.Println("Number of fields", v.NumField())
        for i := 0; i < v.NumField(); i++ {
            fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
        }
    }
}
```

## Les méthodes Int() et String()
Les méthodes `Int` et `String` permettent d'extraire la reflect.Value sous forme d'`int64` ou de `string`.
```go
func main() {  
    a := 56
    x := reflect.ValueOf(a).Int()
    fmt.Printf("type:%T value:%v\n", x, x)
    b := "Naveen"
    y := reflect.ValueOf(b).String()
    fmt.Printf("type:%T value:%v\n", y, y)

}
```

## Programme complet
```go
func createQuery(q interface{}) {  
    if reflect.ValueOf(q).Kind() == reflect.Struct {
        t := reflect.TypeOf(q).Name() // Permet d'avoir le nom du struct
        query := fmt.Sprintf("insert into %s values(", t)
        v := reflect.ValueOf(q)
        for i := 0; i < v.NumField(); i++ {
            switch v.Field(i).Kind() {
            case reflect.Int:
                if i == 0 {
                    query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
                } else {
                    query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
                }
            case reflect.String:
                if i == 0 {
                    query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
                } else {
                    query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
                }
            default:
                fmt.Println("Unsupported type")
                return
            }
        }
        query = fmt.Sprintf("%s)", query)
        fmt.Println(query)
        return

    }
    fmt.Println("unsupported type")
}
```

# Conclusion
Reflect ne devrait être utilisé qu'en dernier recours car c'est du code difficile à écrire et maintenable.