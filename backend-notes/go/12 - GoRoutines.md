# Introduction

Une Goroutine est un thread\* managé par le Go runtime.
Les goroutines sont lancées dans le même espace d'adresse. L'accès à la mémoire partagée doit être synchronisé. (voir sync package)

# Lancé une nouvelle goroutine

go f(x, y ,z) permet de lancer f(x, y ,z) sur une nouvelle goroutine.

# Channels

Les channels sont des conduits permettant d'envoyer et recevoir des valeurs.

```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

Les channels doivent être crées avant d'être utilisés : `ch := make(chan int)`.

Par défaut les blocks d'envoi et de réception bloquent l'autre partie du programme. Cela permet de synchroniser facilement le programme.

```go
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c. Les résultats arrivent dans l'ordre d'apparition des Goroutines

	fmt.Println(x, y, x+y)
}
```

# Buffered Channels

Il est possible de configurer la taille du channel (le nombre de valeurs qu'il va retenir).
Pour cela, il faut le déclarer de cette manière : `ch := make(chan int, 100)`
```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
ch <- 3 // Provoque une erreur car la taille du buffer est de 2
```
# Range and close
Une channel peut être fermée à l'aide de `close(c)` pour indiquer qu'il n'y a plus de valeurs à envoyer.
Seule la fonction envoyant les données devrait pouvoir fermer le channel.

La fonction ayant fait appel à une goroutine peut tester si le channel est ouvert ou fermé : `v, ok := <- ch`. S'il n'y a plus de valeurs à transmettre ok est false.
La loop for i:= range c reçoit les valeurs du channel jusqu'à la fermeture de celui-ci.

```go
func fibonnaci(n int, c chan int){
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+ y
    }
    close(c)
}

func main () {
    c := make(chan int)
    go fibonnaci(10, c)
    for v := range c {
        fmt.Println(v)
    }
}
```

#Select

**A Approfondir**


Un select permet à une goroutine d'attendre plusieurs opérations de communication.
Le select bloque jusq'à ce qu'un de ces cas se lance ensuite il exécute le cas. Si plusieurs cas sont sont lancés en même temps alors il en choisit 1 au hasard.

```go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // Se déclenche lorsque la valeur de x est attribuée à c
			x, y = y, x+y
		case <-quit: // Se déclenche lorsque la valeur de quit change
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

## Default Selection
Le default case est lancé si aucun des autres case n'est prêt.
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
# Définition

Thread :

```
* Au début, tu as le programme P, ou processus. Quand le programme se lance, il s'accapare un morceau de mémoire.

* Ce programme P va faire des actions (lire des fichiers, faire des calculs, afficher des trucs à l'écran, etc...). Pour chacune de ces actions (pour simplifier) il va faire appel au processeur.

* Si dans ce programme P tu veux faire 2 choses en même temps, tu vas avoir besoin de créer un ou plusieurs threads. L'exemple classique c'est une application avec une interface graphique. Tu va lancer un thread T1 qui gère l'interface graphique, ainsi ton programme (ou processus) qui gère les données ne sera pas emmerdé s'il y a des soucis avec ton interface graphique, et vice-versa. Dans cet exemple, nous avons un seul thread T1, qui gère l'interface graphique, et un programme (ou processus) qui gère les données. T1 utilise la mémoire de P (alors que si on a un autre programme P2, ce dernier ne partagera pas la mémoire de P). Dans P, si tu crée un autre thread T2, (exemple typique, effectuer une lecture continue sur une ressource (port serie, usb, ethernet, bdd, lecteur CD, etc.) ) il partagera la mémoire de P et de T1.
```

Binary tree

```
Un binary tree est un arbre comportant 1 root et des childs disposant de childs également. En bout de chaine, nous avons les leafs ne disposant pas de child. 
Un arbre binaire signifie que chaque noeud comporte au maximum 2 childs.
```
