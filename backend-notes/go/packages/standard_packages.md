# Listes de packages :

- Standard :

  - fmt :
    fmt permet d'afficher des éléments dans la console.
    - fmt.Println("string") : permet d'afficher un string ou une variable
    - fmt.Printf("%v", variable) : permet d'utiliser les formes %v, ...
      - %g,
      - %T pour afficher le type,
      - %v pour la valeur
      - %d pour des int
      - %c pour les caractères individuelles
      - %x pour retourner au format hexadecimal
  - math/rand :
    - rand.Intn(10)
  - math :
    - math.Sqrt(7)
      -time : - time.Now().Weekday permet de retourner le jour - time.Now().Weekday() + 1 : Lendemain - time.Saturday - time.Now().Hour() extrait l'heure
  - strings
    - Fields(string): retourne un splice contenant le string divisé en mots.
    - strings.NewReader("string") : Converti un string en io.Reader
    - Pour retirer les espaces d'un string : strings.Join(strings.Fields(ret), " ")
    - strings.IndexRune(string, rune) : Permet de retourner l'index d'une rune dans un string
  - time

    - time.Tick(100\*time.Millisecond) : Déclenche un timer qui se déclenche toutes les 100 ms
    - time.After(500\*time.Millisecond) : Déclenche un timer qui se déclenche au bout de 500 ms. Retourne un chan Time quand expiré.
    - time.NewTicker(duration) : Retourne un Ticker qui tick à chaque duration.
    - time.Millisecond/Second : Renvoie l'équivalent de 1ms. Il est obligatoire de l'avoir dans les fonctions au dessus pour qu'elles fonctionnent.

  - log
    - log.Fatal("message") : Arrête l'exécution du programme et log un message dans le terminal avec des informations temporelles (2019/06/17 10:36:35 message).
  - unicode/utf8
    - utf8.RuneCountInString(string) : Compte le nombre de lettre d'un string
  - unicode
    - unicode.IsUpper(letter) : Retourne true si la lettre est en majuscule

- Autres :
  - Mux : Router
