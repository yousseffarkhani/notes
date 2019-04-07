---
title: Serverless Lambda Functions
created: '2019-04-05T20:53:55.295Z'
modified: '2019-04-05T22:32:11.639Z'
---

# Serverless Lambda Functions
## Introduction
Les lambda functions permettent de faire tourner des fonctions back-end sans avoir besoin de mettre en place un serveur.
Avantages:
- Simplifie le déploiement d'API
- Coute beaucoup moins cher puisqu'on paye à la requête réalisée sur le endpoint
Use-cases:
- Cela est particulièrement utile quand un lien front-end doit appelé une API avec une clé privée. En effet, auparavant la clé privée était accessible dans le code JS front-end. Or avec les lambdas cette clé n'est plus accessible.
- Trigger : Lorsqu'une image est reçu sur un bucket s3, la lambda peut la resizer avant de la stocker
- schedule jobs

Cela permet donc de créer un nouveau paradigme : function as a service.
## Configuration et déploiement sur netlify
1. Installer netlify-lambda : `npm i netlify-lambda`
2. Créer les scripts
```
"lambda-serve": "netlify-lambda serve functions",
"lambda-build": "netlify-lambda build functions"
```
3. Pour tester l'application localement, utiliser `lambda-serve`.
3. Pour la push sur github, `lambda-build`.
4. Se connecter et déployer l'application sur netlify depuis le repo github du projet.
### Processus
1. Front-end : Appel à un endpoint : `const fetchUsers = async () => await (await fetch('/.netlify/functions/getUsers')).json()`
Ce endpoint correspondant à une lambda function va réaliser l'appel à l'API et renvoyer une réponse.
> Pour appeler une API depuis le back-end il faut utiliser la librairie axios.

> La lambda function peut recevoir des requêtes GET et POST (les infos transmises dans POST sont accessibles via `event.body`)

2. Lambda function : 
```
exports.handler = function(event, context, callback){
    const { API_URL } = process.env; //Cette variable est configurée sur netlify au moment de créer le site
    // Send user Response
    const send = body => {
        callback(null, {
            statusCode: 200,
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Access-Control-Allow-Headers': 'Origin, X-Requested-With, Content-Type, Accept'
            },
            body: JSON.stringify(body)
        })
    }
    // Perform API call
    const getUsers = () => {
        axios(API_URL)
            .then(res => send(res.data))
            .catch(send);
    }

    // Make sure method is GET
    if(event.httpMethod == 'GET'){
        // Run
        getUsers();
    }
}
```
## Configuration et déploiement sur AWS
https://www.youtube.com/watch?v=71cd5XerKss

