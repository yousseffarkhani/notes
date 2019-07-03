# Introduction

Un string en go est un slice de bytes.
Les strings sont immutables.

Comme un string est un slice, il est possible d'accèder à chaque élément d'un string avec son index.

```go
s := "Hello"
fmt.Printf("%x ", s[1]) // Renvoi le code hexadécimal : 48 65 6c 6c
fmt.Printf("%c ",s[1]) // Renvoi la lettre
```

Cependant cette approche présente des défauts notamment avec les caractères spéciaux.

```go
s :=  "Señor"
fmt.Printf("%x ", s[1]) // Renvoi le code hexadécimal : 53 65 c3 b1 6f 72
fmt.Printf("%c ",s[1]) // Renvoi la lettre : S e Ã ± o r
```

La raison de ce problème provient du faire que le caractère ñ est encodé sur 2 bytes c3 et b1 alors qu'on essai d'accèder à chaque valeur individuellement.
Pour résoudre ce problème, nous utilisons les runes.

# Runes

```go
s :=  "Señor"
runes := []rune(s)
for _, v := range runes {
    fmt.Printf("%c ",v)
}
```

## Compter le nombre de lettres

Il faut utiliser `utf8.RuneCountInString(s)` de la librairie `unicode/utf8`.

## Modifier un string

```go
s :=  "Señor"
runes := []rune(s)
runes[0] = 'A'
s = string(runes)
```
