# Vertical scaling
- Correspond au fait de rajouter plus de ressource hardware sur son serveur (+ de RAM, de CPU, de Disque dur)
### Avantages :
- Permet d'éviter de passer directement à de l'horizontal scaling
### Inconvénients :
- Cette solution est limitée par l'avancée en matière technologique. (capacité de RAM, ...)

# Horizontal scaling
Consiste à avoir plusieurs instances de serveur faisant tourner le même programme.
Le même code doit être disponible sur les différents serveurs. Pour réaliser cela, AWS crée une AMI (Amazon Machine Image). Cet AMI contient les informations nécessaires pour déployer une nouvelle instance.
### Avantages :
- L'application est scalable à l'infini
### Inconvénients :
Cela amène des problèmes de conception :
- Comment partager les données entre toutes les instances ?

# Load balancer
Un load balancer est mis en place lorsque nous disposons de plusieurs serveurs.
Celui-ci permet de rediriger les requêtes vers les serveurs selon différents critères (round-robin, celui ayant le moins de charge, en fonction de données dans le cookie, ...).

Le load balancer peut insérer un cookie contenant l'adresse du serveur utilisé par le client pour éviter à celui-ci de perdre les données de sa session à la connexion suivante.

Il est possible de renforcer le load balancer en mettant en place plusieurs load balancers (active-active, active-passive : similaire à master-slave).
### Avantages :
- Permet de gérer les pics de charge en lançant plus d'instances de serveurs
Inconvénients :
- Le load balancer répond au problème d'availability puisque si une machine tombe l'autre peut prendre le relai.

# Static web pages
Consiste à exposer directement des fichiers html afin d'éviter le temps de génération par le programme.

### Avantages :
- Rapidité
### Inconvénients :
- Difficile de changer le template des pages ou le CSS

# Caching
Consiste à stocker en RAM les données qui seront surement réutilisés.
Cela permet d'avoir accès rapidement à la donnée.
Ex : Memcached, Redis
Le cache est souvent utilisé pour éviter d'avoir à appeler une BDD.
Ce système est souvent utilisé pour :
- garder les sessions utilisateurs
- garder en mémoire les fils d'actualité des utilisateurs récents

Il existe 2 patterns de caching :
- Cached database queries : Consiste à mettre en cache le résultat de la query avec pour clé une version hashé de la query.
- Cached Objects
### Avantages :
- Permet d'alléger la charge sur les BDD
### Inconvénients :
- La fraicheur de la donnée n'est pas garantie (c'est un compromis à faire)

# Asynchronism
Consiste à réaliser les tâches répititives et demandant beaucoup de ressources de manière asyncrhone.
Exemples :
- Réaliser des tâches lourdes à des moments de baisse de charge
    - Pré-générer les fichiers HTML d'un site web la nuit
- Avoir un système d'async processing
    - queue-job (RabbitMQ) : La tâche intensive est envoyée dans une queue en attendant qu'un worker s'en charge. Lorsque la tâche est terminée, le client est notifié.

# Database
## SQL vs NoSQL
https://www.sitepoint.com/sql-vs-nosql-differences/

### SQL
Relationnal data
Structured Data
+ :
- Transactions
- Lookups by index very fast
- JOINs
- Adding a new table of information
### NoSQL
Dynamic data
Non-relationnal data
Store a lot of TB
Pas de requêtes joint
Les joint doivent être réalisés dans le code.
+ : Scaling
## Sharding
TODO : http://highscalability.com/blog/2009/8/6/an-unorthodox-approach-to-database-design-the-coming-of-the.html
## Index
Consiste à mettre en place un index sur une colonne (ou plus) du tableau.
La BDD va donc créer un nouveau tableau avec uniquement les données des colonnes triées et disposant d'une référence vers la donnée originale.
### Avantages :
- Permet de rechercher rapidement lorsque la BDD devient grande
### Inconvénients :
- Il faut retrier l'index à chaque insertion de donnée.

## Replication
Consiste à faire des copies automatiquement.
### Master-slave
Copie à plusieurs serveurs esclaves. Si le master tombe alors un slave est promu master automatiquement.
Il est possible de répartir la charge de lecture sur les slaves en mettant en place un load balancer devant ces derniers.
### Avantages :
- Redondance des données (en cas de problèmes les données sont sauvegardées)
- Il est possible de réaliser les writes sur le master et les reads sur les slaves. Cela permet un scaling infini pour les reads. (Pratique pour les sites avec plus de reads que de writes)
### Inconvénients :
- Le master est le single point of failure. Si celui-ci tombe les services d'écriture ne marcheront plus en attendant qu'un slave soit promu master.

### Master-Master
2 serveurs masters synchronisent en permanence leurs données de sorte à ce que si l'un des 2 tombe l'autre peut prendre le relai.
Il existe différentes stratégies pour interagir avec les master : Load balancing, master-slave(master).
### Avantages :
- Similaire à master-slave
- Si un serveur write tombe, l'autre peut prendre le relai.
### Inconvénients :

## Partitioning
Consiste à avoir un serveur (ou plusieurs avec un load balancer) réalisant une tâche.
Ex: Facebook avait des BDD distinctes pour chaque école (une pour Harvard, une pour le MIT).
Il est possible de partitionner selon d'autres critères (Si le nom de l'utilisateur est compris entre A et M alors il est stocké dans le serveur 1 sinon dans le serveur 2).
Ces critères remplacent ainsi les règles des load balancer du type round-robin ou en fonction de la charge.

### Avantages :
- Scaling horizontal infini
### Inconvénients :
- Plus complexe à mettre en place

# Datacenter redundancy
Consiste à avoir les mêmes VPS répartis dans des datacenters distincts afin d'être disponible même en cas de problèmes de courant, incendie, ...
Le load balancing peut être réalisé en fonction de la géolocalisation.
### Avantages :
- Hautement disponible
### Inconvénients :
- Complexité à mettre en place
- Réplication des BDD complexe