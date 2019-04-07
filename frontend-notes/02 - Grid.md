---
title: 02 - Grid
created: '2019-04-02T18:50:44.837Z'
modified: '2019-04-02T19:10:37.007Z'
tags: [CSS]
---

# 02 - Grid
```
<body>
  <div class="grid">
    <div class="box a">A</div>
    <div class="box b">B</div>
  </div>
</body>
```
## Basics CSS Grid
#### A appliquer sur la classe grid
- Pour activer : `Display: grid / inline grid;`
- Pour créer des colonnes : `grid-template-columns: 100px 10% 200px;` (permet de créer 3 colones)
- Pour créer des lignes : `grid-template-rows: 100px 10 % 200px;` (permet de créer 3 lignes)
- Autre option pour créer des lignes et colonnes (dans cet ordre) : `grid-template: 40% 50% 50px/100px 50% 200px;`
**(utilisable que dans CSS grid)** Au lieu d’utiliser 10px/10% utiliser 1fr pour 1 fraction
- Repeat() : Pour répéter plusieurs fois une instruction : `repeat(3, 1fr, 2fr)/ repeat(3,1fr);`
- Fixer la distance entre 2 grid : `Grid-gap: 20px(row) 20px(column);`
#### Informations supplémentaires :
- définir la hauteur : `grid-auto-rows: minmax(100px, auto)`  (Ici, le minimum sera de 100px. Si le contenu est supérieur alors la taille s’adapte.)
- Pour que la création des colonnes soit dynamique : `grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));`
#### A appliquer sur la classe a/ b
- Pour placer et fusionner des grid de la 4 à la 5: grid-row: 4 (start) / 6 (stop); `grid-column: 4 / span 2 (relatif à la valeur précédente);`
- Autre option : `grid-area: 6 (start row) / 8 (start column) / span 1 (end row) / span 2 (end column); `
## Advanced CSS Grid
#### Solution 1
Pour réaliser un layout de la page (exemple :https://s3.amazonaws.com/codecademy-content/courses/learn-css-grid/lesson-ii/grid-template-areas/index.html)
1. Définir le layout : 
```
       Grid-template-areas: "head head"       
                            "left right"
                            "left ."
                            "footer footer"
```
2. Associer ces areas avec les selectors : 
```
header { grid-area: head;}
.left { grid-area: left;}
```
3. Fixer la taille des colonnes et lignes avec grid-template-rows/columns
#### Solution 2
1. Définir le template de la page
`grid-template: repeat(12, 1fr) / repeat(6, 1fr);`
2. Attribuer les zones aux éléments
`grid-area: 1 / 1 / 3 / 7;`
3. Fixer le z-index
`z-index:5;`
## Placer les élements dans une grid
#### Horizontalement
- Justify-items: center > permet de centrer l’élément par rapport à la colonne du container dans lequel il se situe. Les éléments seront aussi large que leur contenu.
- Justify-content : Positionner un élément par rapport à son élément parent.
Valeurs possibles :
    - start — aligns the grid to the left side of the grid container
    - end — aligns the grid to the right side of the grid container
    - center — centers the grid horizontally in the grid container
    - stretch — stretches the grid items to increase the size of the grid to expand horizontally across the container
    - space-around — includes an equal amount of space on each side of a grid element, resulting in double the amount of space between elements as there is before the first and after the last element
    - space-between — includes an equal amount of space between grid items and no space at either end
    - space-evenly — places an even amount of space between grid items and at either end
- justify-self : permet de fixer la valeur jusitfy-items pour un seul élément
#### Verticalement
- align-items :
    - start — aligns grid items to the top side of the grid area
    - end — aligns grid items to the bottom side of the grid area
    - center — aligns grid items to the center of the grid area
    - stretch — stretches all items to fill the grid area
- align-content : voir justify-content
- align-self : permet de fixer la valeur align-items pour un seul élément
#### Autres
- grid-auto-rows/columns : 500px/1fr/repeat(3, 450px) ; (à placer dans le container) > permet de créer  automatiquement de nouvelles lignes / colonnes selon la taille spécifiée. Les éléments seront aussi large que leur contenu.
**/!\ Il ne faut pas que le conteneur est une height fixée.**
- Grid-auto-flow : Permet de placer les éléments soit en remplissant les lignes de gauche à droite ou les colonnes de haut en bas. Par défaut, la valeur row est sélectionnée. (à placer dans le container).
Valeurs possibles :
  - row — specifies the new elements should fill rows from left to right and create new rows when there are too many elements (default)
  - column — specifies the new elements should fill columns from top to bottom and create new columns when there are too many elements
  - dense — this keyword invokes an algorithm that attempts to fill holes earlier in the grid layout if smaller elements are added
