# Description
L'API contexte de React permet de partager de l'information entre les différents composants.

Context provider
# Implémentation
## Ajouter un contexte
## Utiliser un contexte
Il est possible d'utiliser le contexte de 2 manières :
- En utilisant une classe et `static contextType = ThemeContext`
- En utilisant le ThemeContext.Consumer (on peut utiliser un functionnal component avec cette méthode)
```JS
<ThemeContext.Consumer>{(context) => {
    const { isLightTheme, light, dark } = context
    const theme = isLightTheme ? light : dark
    return (
        <nav style={theme}>
            <h1>Context App</h1>
            <ul>
                <li>Home</li>
                <li>About</li>
                <li>Contact</li>
            </ul>
        </nav>
    )
}}
</ThemeContext.Consumer>
```
## Mettre à jour les données du contexte
