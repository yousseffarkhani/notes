# Introduction

Les programmes Go expriment les erreurs avec des `error`.

```go
type error interface {
    Error() string
}
```

Les fonctions retournent souvent des `error`. C'est pourquoi il faut toujours gérer les erreurs en testant si l'`error` est égale à `nil`.

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

Il est possible de reformater le message d'erreur.
