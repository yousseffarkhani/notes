---
title: 06 - HTML
created: '2019-04-02T19:53:43.200Z'
modified: '2019-04-03T12:07:19.333Z'
tags: [HTML]
---

# 06 - HTML

## Raccourcis utiles (VS CODE)
- Définir les éléments de base : !
- Créer des mots : lorem8 (8 mots seront crées)
- Créer une div : div
- Créer une div ayant pour class hello : div.hello
En tête
<meta charset="utf-8" />
## Généralités
- Sélectionner un élément particulier dans une ligne de texte / paragraphe : `<span></span>`
- Sauter une ligne : `<br />`
- Ajouter une vidéo : `<video src="xxx" width="320" height="240" controls>Video not supported</video>`
- Ouvrir un lien dans un nouvel onglet : `<a href="xxxx" target="_blank">blablabla</a> `
- Faire en sorte qu'une image (ou une ressource redirige vers une URL) : `<a …><img …/></a>`
- Pour envoyer l'utilisateur vers un endroit particulier de la page :
1. Assigner un id à un élément de la page
1. `<a href="#id">blabla</a>`
- Pour afficher des informations lorsqu’un utilisateur survole une zone, il faut ajouter l’attribut `title="XX"` à la balise.
- Format d’un fichier html : https://user.oc-static.com/files/343001_344000/343677.png
Figures
- La balise `<img>` doit être placé au sein d’une balise `<figure>` si l’image apporte une information / illustre le texte.
```
<figure>
    <img src="images/internetexplorer.png" alt="Logo Internet Explorer" />
    <img src="images/firefox.png" alt="Logo Mozilla Firefox" />
    <img src="images/chrome.png" alt="Logo Google Chrome" />
    <figcaption>Logos des différents navigateurs</figcaption>
</figure>
```
## Tableaux
Créer des tableaux : 
```
<table>
  <thead>
    <tr>
      <th>titre</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>titre</th>
    </tr>
    <tr>
      <td colspan="2"> Donnée</td> //Donnée sera affichée sur 2 colonnes
      <td rowspan="2"> Donnée</td> //Donnée sera affichée sur 2 lignes
    </tr>
  </tbody>
  <tfoot>
    <tr>
      <td> Donnée</td> 
      <td> Donnée</td>
    </tr>
  </tfoot>
</table>
```
## Formulaires
Pour créer un formulaire :
```
<form action="/action_page.php" method="get" target="_blank"> 
  <fieldset>
    <legend>Titre du formulaire</legend>
    <label for="pseudo">nom du champs</label>: <input type="XX" id="pseudo" placeholder="XX" name="XX" value="XX"/>
  </fieldset>
</form>
```
- action : Le champs « action » permet de définir l’action à réaliser quand le formulaire est soumis. Si le bouton action n’est pas précisé alors l’action mène vers la page actuelle.
- Method : défini la méthode HTTP d’envoi des données : GET / POST
- Target : Défini si le résultat du formulaire doit être affiché dans la page actuelle, une nouvelle page, …
- Fieldset : Permet de grouper des champs inputs
- Input :
    - type
      - text : texte input
      - password
      - radio : `<input type="radio" name="gender" value="male" checked>` (penser à rajouter un label)
      - checkbox (penser à rajouter un label)
      - email (permet d’afficher un clavier adapté sur mobile)
      - tel : téléphone (permet d’afficher un clavier adapté sur mobile)
      - number
      - range
      - color
      - date
      - search : 
      - reset : pour remettre à zéro le formulaire 
      - submit : `<input type="submit" value="Submit">` 
              Ce bouton permet de soumettre les champs du formulaire à un form-handler. Le form-handler est généralement une page html contenant un script se chargeant de gérer les inputs. Le form-handler est défini dans le champs action (Le visiteur sera conduit à la page indiquée dans l'attribut action du formulaire.)
      - image : équivalent du bouton submit mais nécessitant une src (url)
    - placeholder : Permet de mettre un commentaire dans le champs
    - name : Les champs inputs doivent impérativement disposer d’un attribut name afin que la donnée soit envoyée
    - value : Permet de donner une valeur par défaut au champ
Pour placer automatiquement le curseur dans un champs au chargement de la page, il faut placer autofocus dans un champ (exemple : `<input type="text" name="prenom" id="prenom" autofocus />`)
## Autres types
Zone de texte multiligne : `<textarea></textarea>` (fonctionne comme input)
## Ressources : 
https://www.w3schools.com/html/

