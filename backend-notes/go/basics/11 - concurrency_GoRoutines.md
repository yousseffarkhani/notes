# Introduction

Go est un concurrent\* language.
Cette notion de concurrency est gérée à travers des Goroutines et des channels.

# Goroutine

Une Goroutine est un thread\* managé par le Go runtime.
Les Goroutine permettent de lancer des fonctions/méthodes de manière concurente
Les goroutines sont lancées dans le même espace d'adresse. L'accès à la mémoire partagée doit être synchronisé. (voir sync package)

## Avantages des Goroutines par rapport aux threads

- Les Goroutines sont légères comparées à des threads.
- Les Goroutines sont liées à moins d'OS threads. Un seul thread peut contenir des milliers de Goroutines. Si une Goroutine bloque le thread alors un autre thread est généré et les Goroutines en attente seront basculées dessus (Tout cela est géré par le runtime).

## Lancer une nouvelle goroutine

go f(x, y ,z) permet de lancer f(x, y ,z) sur une nouvelle goroutine.
A noter que la main function est lancée dans sa propre Goroutine appelée main Goroutine.

```go
func hello(){
	fmt.Println("hello")
}

func main(){
	go hello()
	fmt.Println("world")
}
```

Ce programme affichera world car :

- Quand une Goroutine est lancée le contrôle est redonné immédiatement à la ligne suivante.
- Lorsque la main Goroutine est terminée alors le programme est terminé.

Il est possible de résoudre ce problème en rajoutant time.Sleep à func main cependant ce n'est pas une bonne solution.

# Channels

Les channels sont des conduits permettant d'envoyer et recevoir des valeurs.
Les channels permettent de bloquer l'exécution d'un programme en attendant que des Goroutines finissent leurs exécutions.

## Déclarer un channel

Les channels doivent être crées avant d'être utilisés : `ch := make(chan int)`.
Pour cela, il faut déclarer le type transporté par le channel.

## Envoyer et recevoir depuis un channel

```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```

Par défaut les blocks d'envoi et de réception bloquent la suite du programme. Cela permet de synchroniser facilement le programme.

```go
/* 1ère solution */
func hello(ch chan string) {
	ch <- "Hello world goroutine"
}

func main() {
	ch := make(chan string)
	go hello(ch)
	fmt.Println(<-ch)
	fmt.Println("main function")
}


/* 2ème solution */
func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true // Bloqué en attendant que la valeur soit lue
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done // Bloqué en attendant qu'une valeur soit envoyée
	fmt.Println("main function")
}
```

Exemple plus complexe

```go
func calcSquare(input int, square chan int) {
	sum := 0
	for input != 0 {
		digit := input % 10
		sum += digit * digit
		input /= 10
	}
	square <- sum
}

func calcCube(input int, cube chan int) {
	sum := 0
	for input != 0 {
		digit := input % 10
		sum += digit * digit * digit
		input /= 10
	}
	time.Sleep(3 * time.Second)
	cube <- sum // send sum to c and blocks execution until someone reads the value
}

func main() {
	input := 123
	ch := make(chan int)
	go calcSquare(input, ch) // Same channel is passed twice
	go calcCube(input, ch)
	sqV, cuV := <-ch, <-ch // blocks execution until someone sends a value from c. Les résultats arrivent dans l'ordre d'apparition des Goroutines
	sum := sqV + cuV
	fmt.Println(sum)
}
```

# Deadlock

Si une Goroutine envoie des données dans un channel alors le programme s'attend à ce qu'une autre Goroutine reçoive les données. Si cela n'arrive pas alors le programme renvoie une erreur deadlock.

De même si une Goroutine s'attend à recevoir des données d'un channel alors le programme s'attend à ce qu'une autre Goroutine envoie des données.

```go
func main() {
    ch := make(chan int)
    ch <- 5 // Provoque une erreur car aucune autre Goroutine n'exploite cette valeur.
}
```

# Unidirectionnal channels

Les channels présentés jusqu'à présent sont bidirectionnels, ils peuvent recevoir et envoyer des données.
Il est également possible de créer des channels unidirectionnels.

```go
func sendData(sendch chan<- int) {
    sendch <- 10
}

func main() {
    sendch := make(chan<- int) // Création d'un send only channel
    go sendData(sendch)
    fmt.Println(<-sendch) // génère une erreur car il s'agit d'un send only channel
}
```

Le réel interêt des channels unidirectionnels est de sécuriser le code en convertissant les channels bidirectionnels en fonction de l'usage souhaité.
La conversion se fait d'un channel bidirectionnel vers un channel unidirectionnel mais pas l'inverse.

```go
func hello(input chan<- string) { // Channel conversion to a send only
	input <- "hello"
}

func main() {
	input := make(chan string)
	go hello(input)
	fmt.Println(<-input)
}
```

# Closing channels and for loops

Une channel peut être fermée à l'aide de `close(c)` pour indiquer qu'il n'y a plus de valeurs à envoyer.
Seule la fonction envoyant les données devrait pouvoir fermer le channel.

La fonction ayant fait appel à une goroutine peut tester si le channel est ouvert ou fermé : `v, ok := <- ch`. S'il n'y a plus de valeurs à transmettre ok est false.
La loop for i:= range c reçoit les valeurs du channel jusqu'à la fermeture de celui-ci.

```go
func producer(chnl chan int) {
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl) // Ferme le channel
}

func main() {
    ch := make(chan int)
    go producer(ch)
    for v := range ch { // Boucle sur les valeurs du channel
        fmt.Println("Received ",v)
    }
}
```

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

# Buffered Channels

Il est possible de configurer la taille du channel (le nombre de valeurs qu'il va retenir).
Pour cela, il faut le déclarer de cette manière : `ch := make(chan int, 100)`

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
ch <- 3 // L'écriture est bloquée et provoque une erreur car la taille du buffer est de 2
```

L'interêt étant que contrairement aux unbuffered channels, les envois au channel ne seront bloqués que lorsque le buffer sera full. De même la réception sera bloquée lorsque le buffer sera vide.

Lorsqu'une valeur est lue, elle est retirée du buffer.

```go
func write(ch chan int) {
    for i := 0; i < 5; i++ {
        ch <- i // Va écrire 0 et 1 dans ch et ensuite bloquer jusqu'à ce qu'au moins 1 valeur soit lue
        fmt.Println("successfully wrote", i, "to ch")
    }
    close(ch)
}
func main() {
    ch := make(chan int, 2) // Buffered channel déclaré
    go write(ch) // Lancement de la Goroutine
    time.Sleep(2 * time.Second)
    for v := range ch { // Lecture des valeurs de ch à l'aide d'une loop. Dès qu'une valeur est lue, une autre est ajoutée au buffer par la fonction write.
        fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)

    }
}
```

# Worker Pools

## WaitGroup

Un WaitGroup est utilisé pour attendre l'exécution de plusieurs Goroutines.
Le contrôle est bloqué en attendant la fin des Goroutines.

Le waitGroup fonctionne comme un compteur. Il est incrémenté avec la méthode Add(1) et décrémenté avec Done().
La méthode Wait bloque la Goroutine où elle est appelée jusqu'à ce que le compteur soit à 0.

```go
func process(i int, wg *sync.WaitGroup) {
    fmt.Println("started Goroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended\n", i)
    wg.Done() // Lorsque la Goroutine est terminée décrémente le compteur.
}

func main() {
    no := 3
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {
        wg.Add(1) // Le compteur passe à 3 à la fin de la loop
        go process(i, &wg) // Il est important de passer l'adresse ici car on veut que ce soit la même variable qui soit mise à jour et non des copies
    }
    wg.Wait() // Bloque la Goroutine
    fmt.Println("All go routines finished executing")
}
```

## Worker Pool

L'utilisation la plus importante des buffered channels est pour l'implémentation des worker Pools.
Une worker Pool est une collection de threads en attente d'avoir une tâche assignée. Une fois la tâche terminée, ils redeviennent disponibles pour la prochaine tâche.

```go
type Job struct {
	id, randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

/* Buffered channels for receiving jobs and writing the output */
var jobs = make(chan Job, 10)       // Worker Goroutines listen for new tasks on the jobs channel
var results = make(chan Result, 10) // Once task completed, result is written on results channel

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

/* Creates a worker */
func worker(wg *sync.WaitGroup) {
	for job := range jobs { // Reads from the jobs channel
		output := Result{job, digits(job.randomno)} // Creates a Result struct using the current job and the return value of digits
		results <- output                           // Writes the result to the results channel
	}
	wg.Done()
}

/* Creates a pool of worker Goroutines */
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1) // Increment waitgroup
		go worker(&wg)
	}
	fmt.Println("Done creating workers")
	wg.Wait()
	close(results)
}

/* Remplie le buffer de jobs */
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	fmt.Println("Done allocating jobs")
	close(jobs)
}

/* Reads results channel and prints ouput */
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 12
	go allocate(noOfJobs) // Add jobs to jobs channel
	done := make(chan bool)
	go result(done)
	noOfWorkers := 2
	createWorkerPool(noOfWorkers)
	<-done
	endtime := time.Now()
	diff := endtime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
```

#Select

Select permet de choisir entre plusieurs send/receive operations.
Select bloque tout jusq'à ce qu'une des conditions se réalise. Si plusieurs operations sont prêtes en même temps alors une d'entre elles et choisit au hasard.

```go
func server1(ch chan string) {
	ch <- "From Server 1"
}

func server2(ch chan string) {
	ch <- "From Server 2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select { // Bloque la suite du programme jusq'à ce qu'un des cases soit prêt
	case s1 := <-output1:
		println(s1)
	case s2 := <-output2: // Après 3s, il y a une valeur dans le chan output2 donc ce case est choisi et le programme est terminé.
		println(s2)
	}
}
```

> La raison d'avoir un select est notamment dans le cas où on a besoin de retourner une réponse rapidement à l'utilisateur après une requête à la BDD. Les BDD sont souvent répliquées donc plusieurs appels peuvent être effectués. Dès qu'une requête aura été remplie alors les autres seront abandonnées.

## Default Selection

Le default case est lancé si aucun des autres case n'est prêt. Il est utilisé pour ne pas bloquer le programme en cas de problème.

```go
func process(ch chan string) {
    time.Sleep(10500 * time.Millisecond)
    ch <- "process successful"
}

func main() {
    ch := make(chan string)
    go process(ch)
    for {
        time.Sleep(1000 * time.Millisecond)
        select {
        case v := <-ch:
            fmt.Println("received value: ", v)
            return
        default: // Le défaut case est executé jusqu'à ce qu'il y ai une valeur dans ch
            fmt.Println("no value received")
        }
    }

}
```

Le default permet également d'éviter des erreurs de deadlock.

```go

func main() {
    ch := make(chan string)
    select {
    case <-ch:
    default: // S'il n'y avais pas default alors il y aurait une erreur
        fmt.Println("default case executed")
    }
}
```

# Définition

## Thread :

```
* Au début, tu as le programme P, ou processus. Quand le programme se lance, il s'accapare un morceau de mémoire.

* Ce programme P va faire des actions (lire des fichiers, faire des calculs, afficher des trucs à l'écran, etc...). Pour chacune de ces actions (pour simplifier) il va faire appel au processeur.

* Si dans ce programme P tu veux faire 2 choses en même temps, tu vas avoir besoin de créer un ou plusieurs threads. L'exemple classique c'est une application avec une interface graphique. Tu va lancer un thread T1 qui gère l'interface graphique, ainsi ton programme (ou processus) qui gère les données ne sera pas emmerdé s'il y a des soucis avec ton interface graphique, et vice-versa. Dans cet exemple, nous avons un seul thread T1, qui gère l'interface graphique, et un programme (ou processus) qui gère les données. T1 utilise la mémoire de P (alors que si on a un autre programme P2, ce dernier ne partagera pas la mémoire de P). Dans P, si tu crée un autre thread T2, (exemple typique, effectuer une lecture continue sur une ressource (port serie, usb, ethernet, bdd, lecteur CD, etc.) ) il partagera la mémoire de P et de T1.
```

## Binary tree

```
Un binary tree est un arbre comportant 1 root et des childs disposant de childs également. En bout de chaine, nous avons les leafs ne disposant pas de child.
Un arbre binaire signifie que chaque noeud comporte au maximum 2 childs.
```

## Concurrency

```
C'est la capacité à gérer plusieurs évenements en même temps.
```

![](./attachments/concurrency-parallelism-copy.png)
