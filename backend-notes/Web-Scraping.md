# Web Scraping
## Introduction
Le web scraping consiste à collecter des données des sites internet.
Les uses cases sont nombreux :
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
## Récupérer des informations d'une SPA
De nombreux sites utilisent aujourd'hui JS pour générer du contenu dynamique. Cela pose des problèmes pour extraire des données du HTML car le HTML récupéré par request correspond à celui de la requête initiale sans exécuter le JS permettant de compléter la page.

Pour scrapper les sites nécessitant une exécution de JS, il faut utiliser la librairie puppeteer.
Puppeteer va nous permettre de simuler un navigateur (sans interface et basé sur chromium).
```
(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  await page.goto(url);
  const content = await page.content();
  console.log(content);
  browser.close();
})();
```
Il est également possible de prendre des screenshots, ...
