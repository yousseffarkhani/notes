# Introduction

Les fonctions sont des valeurs comme en JS.

```go
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main(){
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }

    fmt.Println(compute(hypot))
}
```

# Closures

Les fonctions disposent de closure.

```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```
