# Concevoir un site responsive

## Processus

1. Définir l'architecture de la page
2. Créer la structure en HTML
3. Mettre en forme avec CSS en démarrant par la fenêtre la plus étroite (téléphone).
4. Appliquer un style simple.
5. Ajout des images stylistiques (images en background)
6. Régler les 1ers points de rupture (@media (min-width: 600px) {} : La largeur est de 600px ou plus)

- 1er point : lorsque la ligne est supérieure à 10 mots (dans l'exemple à 600px)
  Voir en dessous pour aide sur les media queries.

7. Fixer la largeur maximale des éléments en les plaçant dans des containers.
   Avec ces propriétés CSS le contenu restera centré et aura une taille max de 800px.

.container {
margin: auto;
max-width: 800px;
}

8. Adapter la taille du texte

## Media queries

Il y a 2 manières de fixer un breakpoint :

- @media only screen (min-width: 768px) : ne s'appliquera que si l'écran fait 768px ou plus.
- @media only screen and (min-width: 768px) and (max-width: 959px) : ne s'appliquera que pour cet intervalle.

Dans notre cas, n'utiliser que les min-width breakpoints et augmenter de plus en plus.

Les breakpoints les plus populaires sont : 576px, 768px, 992px, 1200px

## Tips

Les images sont définies sur une échelle de 100% par défaut.
Conseils :
Les images doivent contenir une balise alt pour les lecteurs d'écran.
Pour tous les médias, utiliser :

img, embed, object, video {
max-width: 100%;
}
