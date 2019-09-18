Source : https://www.sohamkamani.com/blog/2017/01/16/web-security-essentials/
Vidéos sur le sujet : https://www.youtube.com/watch?v=-enHfpHMBo4&list=PL96C35uN7xGLux5q2c4P_IqbKF11-pfsR
également : https://www.calhoun.io/securing-cookies-in-go/

# Introduction

La sécurité d'une web application est un sujet complexe.
Un système d'authentification par exemple est très complexe.
Les concepts abordés seront :

- password management
- session cookies
- Types d'attaques

# Sessions and cookies

Les session token permettent à un utilisateur d'être authentifié vis à vis d'un serveur pendant une certaine durée.
Passé ce délai l'utilisateur devra se relogger en vue de récupérer un nouveau token.
Le serveur lui garde une table de correspondance entre les token et les utilisateurs.
Pour que ce token soit transmis à chaque requête, il est ajouté aux cookies envoyés.

# SQL injection

Pour prévenir ce type d'attaque, il faut toujorus "escape" les caractères spéciaux. Cela signifie remplacer les caractères ', " par des \' ou \".
L'autre solution consiste à utiliser un ORM au lieu d'écrire ses propres requêtes.

# XSS (cross site scripting)

Ce type d'attaque consiste à ajouter un tag script au html qui va lancer du code javascript sur le navigateur des personnes visitant le site. L'exemple le plus parlant est par exemple lorsqu'on partage un statut :

```html
<div id="status">
  I am feeling alright
</div>
```

et que l'on change ce dernier par

```html
<div id="status">
  <script>
    alert("I am feeling great!");
  </script>
</div>
```

Si le site n'est pas bien protégé alors l'ensemble des personnes visualisant le statut auront une alerte.

Pour éviter ce genre de chose, il faut encoder tout le HTML qui sera affiché sur la page (< = &lg;)

# CORS (cross origin ressource sharing)

## Same origin policy

Il est interdit de faire une requête à une origine différente de la page sur laquelle nous sommes (sauf si autorisé par le serveur).
## Pourquoi CORS n'est pas autorisé ?
Toutes les requêtes envoyées depuis le navigateur sont traitées de la même façon.
Cela signifique que si on est loggé à un site et qu'on navigue sur un autre site (malicieux), une requête http envoyée par ce site malicieux (via fetch) au 1er site serait traité comme une requête standard par nous.

## A quel point CORS nous est imposé ?

On pourrait se demander pourquoi on peut pas réaliser une simple requête http depuis notre navigateur vers une autre origine alors qu'avec des outils types curl / postman on peut le faire.
CORS est une policy imposée par le navigateur plutôt que le serveur.
Le serveur statue uniquement les sites depuis lesquels il peut recevoir des cross origin request. Cette information est envoyée à travers le header Access-Control-Allow-Origin dans toutes les réponses.

Par la suite, c'est au navigateur de décider de respecter cette policy.

Pour en revenir à Postman et Curl, ces 2 outils ne respectent pas la same origin policy car ce sont des outils de développement et s'attendent à ce que les personnes les utilisant sachent ce qu'ils font avec.

## Exceptions et mise en garde

1. Fetch peut réaliser des cross origin requests.
   Il existe un no-cors mode pour faire cela. Cependant la lecture de la réponse restera impossible.
2. Les serveur peuvent bloquer toutes requêtes ne provenant pas de leur domaine ainsi les outils Postman et curl ne pourront pas réaliser de requête non plus.

# Password storage

## Chhiffrement

Il faut utiliser les algorithme de chiffrement comme md5 ou sha-1.
Pour utiliser ce mécanisme, il faut utiliser le même algorithme lorsque l'utilisateur entre son mot de passe et le comparer avec le hash en base de données.

Le problème étant qu'avec juste cet algorithme, il est possible de comparer à des hashs souvent utilisés (dictionnary attack) et retrouver le mot de passe original.

Pour éviter ce problème il existe plusieurs solutions :

- Forcer l'utilisateur à utiliser des caractères spéciaux
- "Salter" le mot de passe

## Salting

Cela consiste à rajouter une série de caractères au hasard au mot de passe et de le hasher. Cela permet de rendre le mot de passe non commun.

## Repeated hashing

Il est possible d'ajouter encore plus de sécurité, en récupérant le résultat du hashing et rehasher ce résultat plusieurs fois

# CSRF (Cross site request forgery / XSRF)
C'est un type d'attaque où une requête provenant d'un site web est déguisé de manière à donner l'impression qu'elle provient d'un autre site.

## Comment cela est possible ?

Cette faille provient de la manière dont ont été désignés les navigateurs et de comment ils envoient les requêtes.

## Exemple

L'utilisateur est loggé à facebook.com. Il click sur un lien ouvrant un tab vers fakenews.com.
Fakenews.com tente de compromettre le compte facebook en envoyant une POST request (en utilisant des forms invisibles).
Comme nous avons une session active avec facebook.com et les informations de session sont stockés dans des cookies. Toutes ces informations seront envoyées avec la requête POST.
Du POV de facebook, la requête provient de l'utilisateur loggé.

## Prévention

- Empêcher les cross origin requests (CROSPolicy): Les serveurs Facebook.com peuvent empêcher les requêtes provenant d'autres domaine que le leur.
  Si le navigateur de l'utilisateur ne respecte pas la CORS policy, il faut prendre des mesures supplémentaires.
- Envoyer un CSRF token : Le serveur envoi un CSRF Token. Ce CSRF Token est ajouté à chaque requête afin d'être sûr qu'il s'agit bien de l'utilisateur et non d'une tentative de hack.

# Human error and UI/UX Design
Lorsque l'on réalise un site, il faut faire attention :

- Au contenu externe notamment les publicités
- Réauthentifier l'utilisateur pour des actions critiques
- A déloggé des utilisateurs au bout d'un moment (garder un temps d'expiration sur les sessions des utilisateurs)
- Mettre des messages d'alerte dans les zones à risque (notamment la console/ Redirection vers un autre site)
