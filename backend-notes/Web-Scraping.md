# Web Scraping
## Introduction
Le web scraping consiste à collecter des données des sites internet.
Le web scraping peut provoquer des surcharges au niveau des serveurs, c'est pourquoi certains sites mettent en place des protections afin d'éviter ces crawlers.
## Comment un site peut détecter du web scraping ?
- Trafic inhabituel provenant d'un utilisateur/Adresse IP unique
- Tâches répétitives réalisées sur le site
- Honeypots : Liens non visibles à l'oeil humain. Quand le scraper essaye d'accéder au lien, ce dernier est ajouté à une blacklist
## Comment contrer cette détection ?
Ajouter un userAgent à ses requêtes.
Un userAgent représente l'identité du navigateur.
Il est sous la forme :
User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36
où :
- Mozilla/5.0 : Application utilisée
- X11; Linux x86_64 : OS de l'agent
- AppleWebKit/537.36 (KHTML, like Gecko) : Moteur d'affichage du contenu
- Chrome/73.0.3683.103 Safari/537.36 : 
## Use-cases
- Collecter les prix de sites de e-commerce pour faire une comparaison.
- Collecter des adresses mails.

## Librairies utilisées 
- Cheerio: pour manipuler et extraire les données du html 
- request: pour fetch le html facilement 
- request-promise (alternative à request): pour récupérer la réponse de la requête sous forme de promise 
- puppeteer: Permet de scrapper les sites nécessitant une exécution de JS pour générer le contenu dynamique.
## Configuration et Vérifications préalables
```javascript
request("https://twitter.com/vechainofficial", (err, response, html) => {
    if(!err && response.statusCode === 200) {
        const $ = cheerio.load(html);

        { code ... }
    }
});
```
## Récupérer des informations d'une page html
Le programme est divisé en 2 : le parser et le main program.

(parser.js)
```javascript
const rp = require("request-promise");
const cheerio = require("cheerio");

const potusParse = url => {
    return rp(url)
        .then(html => {
            const $ = cheerio.load(html);
            return new Promise((resolve, reject) => resolve({   
                name: $(".firstHeading").text(),
                birthday: $(".bday").text() 
                }));
        })
        .catch(console.log);
};

module.exports = potusParse;
```
(main program)
```javascript
const rp = require("request-promise");
const cheerio = require("cheerio");

const potusParse = require("./potusParse")

const url = "https://en.wikipedia.org/wiki/List_of_Presidents_of_the_United_States";

rp(url)
    .then(html => {
        const $ = cheerio.load(html);
        console.log($("big > a").length); 
        
        $("big > a").each((index, element) => {
            const presidentUrl = "https://en.wikipedia.org" + $(element).attr("href");
            potusParse(presidentUrl)
                .then(president => console.log(president));
        });
    })
    .catch(console.log)
```   

Méthodes disponibles : https://github.com/request/request-promise (Poster en HTML, ajouter un userAgent, parser en JSON directement, ...)
## Récupérer des informations d'une SPA
De nombreux sites utilisent aujourd'hui JS pour générer du contenu dynamique. Cela pose des problèmes pour extraire des données du HTML car le HTML récupéré par request correspond à celui de la requête initiale sans exécuter le JS permettant de compléter la page.

Pour scrapper les sites nécessitant une exécution de JS, il faut utiliser la librairie puppeteer.
Puppeteer va nous permettre de simuler un navigateur (sans interface et basé sur chromium).

De plus puppeteer peut également nous aider pour les tests End-to-End.

(source : https://github.com/GoogleChrome/puppeteer/blob/master/docs/api.md#overview)
### Fonctionnalités
- Générer un pdf d'une page
- Prendre des screenshots
- Visualiser les actions utilisateurs automatisées sur le navigateur
### Syntaxe générale
```javascript
(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  await page.goto(url);
/* Ajout des méthodes ici */
  browser.close();
})();
```
### Méthodes (voir documentation pour plus de détails)
#### Manipulation du DOM
- `const content = await page.content();` : Renvoi le HTML de la page
- `const searchValue = await page.$eval('#search', el => el.value);` : Lance `document.querySelector('#search')` et le donne en 1er de la fonction.
- `const divsCounts = await page.$$eval('div', divs => divs.length);` : Lance `Array.from(document.querySelectorAll('div'))` et le donne en 1er argument de la fonction.
    - `'div'` : Sélecteur
    - `divs => divs.length` : Fonction à appliquer
    - Cette méthode retourne la valeur de la fonction
##### Sélecteurs
- `const searchBox = await page.$(".gsfi");` : Sélectionne un élément du DOM (équivalent à `document.querySelector`) 
- `const submitBtn = await page.$('input[type=submit][value="Recherche Google"]');` : Sélectionne un bouton grâce à ses attributs
- `const searchBox = await page.$$("div");` : Sélectionne tous les éléments du DOM (équivalent à `document.querySelectorAll`) 
##### Actions
- `await submitBtn.click();` : Click sur le bouton
- `await searchBox.type("test");` : Saisi des éléments dans l'input
#### Autres
- `await page.waitFor(2*1000);` : Met en pause pendant 2 sec
- `await page.pdf({path: "google.pdf", format: "A4"});` : Génère un pdf (**Il faut que le headless mode soit à true**)
- Evaluer des éléments dans la console :
```javascript
let links = await page.evaluate((sel) => {
    let element = document.querySelector(sel);
    return element? element.innerHTML: null;
}, LINKS_SELECTOR);
```
