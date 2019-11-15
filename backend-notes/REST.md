https://www.brainstobytes.com/restful-services-what-they-are-and-what-they-are-not/

# Introduction
Les principes exposés par REST sont à propos de l'interaction entre les composants et l'interprétation des données échangées.

Une ressource représente un objet ou un concept (playground, people, ...).
Pour identifier ces ressources afin de les récupérer, il faut un identifiant. La plupart du temps, il s'agit d'une URL.

Un autre élément important est la représentation de cette ressource. Celle-ci peut prendre la forme de JSON / XML / HTML / Image / ...

# Contraintes
Un système REST est composé de 6 contraintes.
1. Client-server : Le système doit adhérer à une architecture client-serveur.
Interêts :
- Separation of concerns : L'interface et le stockage des données évoluent indépendamment l'un de l'autre.
2. Stateless
La communication entre le serveur et le client doit être stateless.
Les informations requises pour comprendre une requête doivent être groupées avec la requête.
Interêts :
- Visibilité : Le système n'a besoin de regarder que la requête pour la comprendre
Quand un état est requis le client le garde de son côté (sous forme de cookie par exemple) et l'envoi chaque requête.
3. Cacheable
Les données envoyées en réponse doivent être marquées comme cacheable ou non.
Interêts :
- Performance : Moins de requêtes sont envoyés
4. Uniform interface
Le système doit prénseter une interface uniforme entre ses composants. Pour assurer, l'uniformité il existe 4 principes :
- Identification des ressources : Les ressources doivent disposer d'un identifiant unique permettant de récupérer ou opérer sur cette ressource
- Manipulation des ressources à travers la représentation :
Les ressources peuvent être envoyées au client sous différents formats. C'est pourquoi, le client doit indiquer sous quel format il souhaite recevoir la donnée.
- Message auto-descriptif :
Verbes HTTP / Status responses
- Hypermedia comme moteur de l'application
5. Layered system