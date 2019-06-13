# Range

Range permet d'itérer sur un slice ou une map.
Lorsque range est utilisé avec un slice, 2 valeurs sont retournées par itération: l'index et une copie de la valeur de l'élément à cet index.

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
```

Il est possible de skipper les valeurs de index et value en assignant \_.

```go
for i, _ := range pow
for _, value := range pow
for i := range pow // Il est possible d'omettre le 2nd paramètre s'il y a besoin uniquement de l'index
```

```go
pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
```
