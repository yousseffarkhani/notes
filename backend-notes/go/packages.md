# Packages

Pour importer un package : import ("nom_du_package")

Listes de packages :

- Standard :

  - fmt :
    fmt permet d'afficher des éléments dans la console.
  - fmt.Println("string") : permet d'afficher un string ou une variable
  - fmt.Printf("%v", variable) : permet d'utiliser les formes %v, ...
    - %g,
    - %T pour afficher le type,
    - %v pour la valeur
    - %d pour des int
  - math/rand :
    - rand.Intn(10)
  - math :
    - math.Sqrt(7)
      -time : - time.Now().Weekday permet de retourner le jour - time.Now().Weekday() + 1 : Lendemain - time.Saturday - time.Now().Hour() extrait l'heure
  - strings
    - Fields(string): retourne un splice contenant le string divisé en mots.

- Autres :
- Mux : Router
