---
title: 01 - CSS
created: '2019-04-02T12:01:39.007Z'
modified: '2019-04-02T19:10:31.951Z'
tags: [CSS]
---

# 01 - CSS

## Méthodologie :
1. Trouver la bonne police sur http://www.fontsquirrel.com/fontface/generator ou https://fonts.google.com/ ou https://www.cssfontstack.com/
Insérer le lien css fourni par le site dans le fichier css de l'application.
1. Rajouter dans le `<head>`, `<link href="style.css" type="text/css" rel="stylesheet">`.
Définition des styles principaux de la page (largeur du site, fond, couleur par défaut du texte).
Définir un bloc englobant la page (en dessous de `<body>`)
Pour centrer un site, définir une width puis `margin: auto;`.
1. En-tête et liens de navigation
Utiliser flexbox pour positionner les éléments. (Pour que 2 éléments soient à l'opposé, utiliser `justify-content: space-between;`).
1. Bannière
L'image de la bannière est positionnée en relative (sans effectuer de décalage) **pour permettre aux enfants de se positionner en absolu par rapport à l'image**.
1. Section principale du corps de page
Pour centrer une image sur du texte, utiliser `vertical-align: center;`.
Pour positionner un élément inline par rapport à son parent, utiliser `text-align: center;`
Pour centrer un bloc, utiliser `margin: auto;`.
1. Pied de page (footer)
Pour personnaliser l'icône des listes à puces: `list-style-image: url("")`.
1. Valider le fichier HTML : https://validator.w3.org/.

## Selectors
Selector name | what does it select | example
--- | --- | ---
Element selector (tag)|  | `p`
ID Selector |  | `#something`
Class Selector |  | `.something`
Attribute Selector | Séléctionne les img disposant de l'attribut src | `img[src]`
Pseudo-class Selector | Séléctionne l'élément lorsque l'action est réalisée | `a:hover`, `p:last-child`, `div:nth-child(odd)`

Plus d’informations : https://developer.mozilla.org/en-US/docs/Learn/CSS/Introduction_to_CSS/Selectors

## Setup & selectors
CSS va être appliqué dans cet ordre : ID => Class => Tag
Override any rule : !important ==> color: blue !important;

## The box model
https://s3.amazonaws.com/codecademy-content/courses/freelance-1/unit-4/diagram-boxmodel.svg
- Fixer la hauteur : `height` (soit en px soit en %)
- Fixer la largeur: `width`
- Border
  - Fixer le border: `border: 1px solid red;`
  - Fixer la courbure des bords : 2px ou 5%
- Padding
  - Fixer le padding : `padding: 10px;`
  - Fixer un padding spécifique: `padding-top/bottom/left/right` ou `padding: 6px 10px 3px 4px;` (clockwise = top, left, bottom, right)
  - Si le padding est le même pour top + bottom et right + left alors `padding: 5px 10px;`
- Marging
  - Fonctionnement similaire à padding
  - Pour centrer le contenu (attention, il faut fixer le width avant) : `margin 0 auto;`
  - Les marges horizontales s'additionnent entre des éléments
  - Les marges verticales ne s'additionnent pas, seule la plus grande est retenue
- Valeurs minimum
  - Min-width / max-width
  - Min-height/ max-height
  - Si un élément est trop petit par rapport à ce qu'il contient, on peut fixer le comportement de l'élément avec overflow: scroll / hidden / visible
- Visibility : hidden
- Content box
    • De base box-sizing: content-box (largeur =width+padding+border)
    • Solution : Reset cette propriété pour l'ensemble des éléments de la page html
```
* {
  box-sizing: border-box
}
```
Width (largeur) = border + padding + content

## Eviter les styles par défaut des navigateurs
1. Aller sur www.CSSReset.com et télécharger Eric Meyer’s Reset CSS
1. À rajouter dans le head :   `<link href="reset.css" type="text/css" rel="stylesheet">`
ou alors :
```
* {
  margin: 0;
  padding: 0;
}
```

## Principaux attributs
- `Color` : change la couleur du texte
- `Background-color` : change la couleur du fond
- `Font-family` : changer la police 
- `Font-size` : changer la taille
- `font-style: italic;`
- `text-transform: uppercase;`
- `Font-weight`: bold/normal
- `line-height: 1.4;`
- `Text-align` : modifie la position par rapport à l'élement parent (cela peut-être le browser)
- `Opacity`: 0->1
- `Background-image: url("XXX");`

## Display and positionning
- Position
  - Relative, par rapport à la position par défaut. Les paramètres top / bottom / left / right doivent être fixés
  - Absolute, par rapport au haut gauche de la page. Les paramètres top / bottom / left / right doivent être fixés
  - Fixed, l'élément restera fixe peu importe la situation (barre de navigation). Les paramètres top / bottom / left / right doivent être fixés
- Z-index, plus la valeur est grande, plus l'image sera devant `z-index: 10;`
- Display :
  - Inline elements :
      - On ne peut pas ajuster la hauteur et largeur de ces éléments. Ils "enrobent" leur contenu.
      - `Display: inline;`
  - Block elements :
      - On ne peut ajuster que la hauteur. L'élément prend toute la largeur
      - `Display: block;`
  - Inline block
      - Leur largeur et hauteur peut-être ajustée
      - `Display: inline-block;`
- Faire flotter une image : Float (la largeur doit être fixée au préalable): left/right/both

## Responsive design et media queries
2 méthodes existent pour appliquer une média query :
- en chargeant un CSS différent en fonction de conditions déterminées au préalable dans l’attribut media (d’où le nom media query)
```
<link rel="stylesheet" href="style.css" /> <!-- Pour tout le monde -->
<link rel="stylesheet" media="screen and (max-width: 1280px)" 
href="petite_resolution.css" /> <!-- Pour ceux qui ont une résolution inférieure à 1280px -->
```
- en appliquant les différentes règles au sein d’un même CSS
```
@media screen and (max-width: 1280px)
{
    /* Rédigez vos propriétés CSS ici */
}
```

#### Les règles disponibles
Combinaison : only / not / and

## CSS Variables
#### Déclarer  et appliquer une variable
`--nom-de-la-variable: valeur;`
`background-color: var(--nom-de-la-variable);`

#### Portée de la variable
Déclarer des variables dans :root donnera accès à ces variables dans l’ensemble du CSS.
Il est également possible de déclarer des variables locales (applicables dans l’élément et ses enfants).

#### Astuces
color : var(--box-main-color, #3f3f3f) > #3f3f3f est la couleur par défaut si la variable –box-main-color n’existe pas

> **Note** : Les CSS variables ne sont compatibles qu’avec les dernières versions de navigateur

## Pré-processeurs
Sass / Less permettent d'enrichir le language CSS avec de nouvelles fonctionnalités.
Les avantages des pré-processeurs sont : 
- gère la compatibilité avec tous les navigateurs
- Met plus de temps à charger

## Vendor prefixes
`-webkit-, -moz-, -ms- or -o-`.
Ces préfixes ont été mis en place afin que chacun des éditeurs de navigateur puissent ajouter des spécifités CSS en attendant la normalisation de ces derniers.
La bonne pratique est de définir comme cela les propriétés :
```
.boite {
  -moz-border-radius:5px;
  -webkit-border-radius:5px;
  border-radius:5px;
}
```

## Approfondir 
https://developer.mozilla.org/en-US/docs/Learn/CSS


