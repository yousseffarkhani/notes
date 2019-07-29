# Introduction

Une map permet de lier une clé à une valeur.

# Déclarer une map

La valeur 0 d'une map est nil.
La fonction make retourne une map du type donné.

```go
type Vertex struct {
    Lat, Long float64
}
var m map[string]Vertex

func main(){
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
}
```

## Map litterals

```go
var m = map[string]Vertex{
    "Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
/* ou plus simplement */
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```

# Mutating maps

## Insérer/MAJ un élément

```go
m[key] = elem
```

## Récupérer un élément

```go
elem = m[key]
```
Si on essaie d'accèder à un élément qui n'existe pas alors la zero value sera retournée.

## Supprimer un élément

```go
delete(m, key)
```

## Tester si la clé est présente

```go
elem, ok = m[key] //Si elem et ok n'ont pas été déclaré au préalable, on peut utiliser cette forme elem, ok := m[key]
// Si la clé est dans m alors ok = true et elem = valeur de l'élément
// Si la clé n'est pas m alors ok = false et elem = valeur 0 de l'élément type
// Si pas besoin de elem alors _, ok = m[key]
```
# Itérer sur une map
L'ordre de récupération des éléments ne sera pas forcément le même à chaque execution du programme.
```go
for key, value := range map {
}
```
# Reference type
De même que les slices, les maps sont des références. C'est pourquoi quand une map est assigné à plusieurs variables, celle-ci sera commune à toutes les variables (de même lorsque la map est passée en argument d'une fonction).