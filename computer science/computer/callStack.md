La callstack est ce qu'un programme utilise pour garder la trace des fonctions appelées.
Cette callstack est composée de stack frames (1 pour chaque fonction).
Les informations stockées dans la stack frame sont :
- Variables locales
- Arguments passés à la fonction
- Les informations de la stack frame de l'élément parent
- L'adresse de retour : Où le programme doit poursuivre son exécution après le retour de la fonction


Par exemple :
```python
  def roll_die():
    return random.randint(1, 6)

def roll_two_and_sum():
    total = 0
    total += roll_die()
    total += roll_die()
    print total

roll_two_and_sum()
```

Le programme va d'abord appeler roll_two_and_sum()
    Call stack :
        - roll_two_and_sum()
Ensuite le programme va appeller roll_dice()
    Call stack :
        - roll_dice()
        - roll_two_and_sum()
A l'intérieur de roll_dice(), random.randint(1, 6) est appelé
    Call stack :
        - random.randint(1, 6)
        - roll_dice()
        - roll_two_and_sum()
Lorsque random.randint(1, 6) retourne le résultat alors il est "poppé" de la stack
    Call stack :
        - roll_dice()
        - roll_two_and_sum()
Même chose pour roll_dice()
    Call stack :
        - roll_two_and_sum()
Puis roll_dice() est encore appelé
    Call stack :
        - roll_dice()
        - roll_two_and_sum()
...


# Complexité mémoire des stack frames
Chaque appel de fonction crée sa propre stack frame et est ajouté à la callstack.
Cela peut être important notamment lorsque nous utilisons de la récursivité.

Lorsque la callstack devient trop grande et que nous n'avons plus assez d'espace mémoire alors cela s'appelle un stack overflow.

Par exemple, une récursion pour calculer le produit de tous les chiffres jusqu'à n pourrait se faire de cette manière :
```go
func product_1_to_n(n)int{
    if n <=1 {
        return 1
    }
    return n*product_1_to_n(n-1)
}
```

La call stack sera de taille O(n). La complexité mémoire sera de O(n) même si on ne crée pas de structure de données.

En utilisant une approche itérative :
```go
func product_1_to_n(n)int{
    result = 1
    for i := 1; i < n; i++ {
        result *= i
    }
    return result
}
```

Cette version utilise un espace mémoire constant à la fois au niveau de la call stack mais également des variables.
