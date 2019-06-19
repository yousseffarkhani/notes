# Critical Section
Quand un programme est lancé de manière concurrente, les parties du code accèdant à des ressources partagées ne doivent pas être accessibles par plusieurs Goroutines.
La partie de code modifiant une ressource partagée est appelée critical section.

Pour empêcher l'accès à une même ressource par 2 Goroutines en même temps, il existe 2 solutions : Mutex et les channels.

En général, il faut utiliser les Mutex lorsqu'une Goroutine doit accéder à une critical section. Dans le cas de l'exemple ci-dessous le choix de Mutex est préférable. 
Pour le reste utiliser les channels.

# Mutex
Un Mutex donne accès à un mécanisme permettant de verrouiller certaines parties du code.
Si une goroutine verrouille l'accès et qu'une autre goroutine essaie d'y accéder alors la nouvelle goroutine sera bloqué tant que la ressource n'est pas déverouillée.

```go
func calc(x *int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*x = *x + 1
	m.Unlock()
	wg.Done()
}

func main() {
	x := 0
	var wg sync.WaitGroup
    var m sync.Mutex
	for i := 0; i < 100; i++ {
		go calc(&x, &wg, &m) // Il faut donner l'adresse du mutex afin que la ressource soit partagée
		wg.Add(1)
	}
	wg.Wait()
	println(x)
}
```

# Channel

On utilise la propriété du buffered channel qui bloque l'exécution tant que le buffer est full.
```go
func calc(x *int, wg *sync.WaitGroup, ch chan bool) {
	ch <- true // Le buffer a une capacité de 1 donc l'exécution des autres goroutines est bloqué jusqu'à la lecture de cette valeur.
	*x = *x + 1
	<- ch
	wg.Done()
}

func main() {
    x := 0
    ch := make(chan bool, 1)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go calc(&x, &wg, ch) // Il faut donner l'adresse du mutex afin que la ressource soit partagée
		wg.Add(1)
	}
	wg.Wait()
	println(x)
}
```