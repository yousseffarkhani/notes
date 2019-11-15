# Source
https://www.hiredintech.com/classrooms/system-design
# Exemples
http://blog.gainlo.co/index.php/category/system-design-interview-questions/
https://www.hiredintech.com/classrooms/system-design/lesson/61
https://github.com/donnemartin/system-design-primer

https://www.youtube.com/watch?v=0163cssUxLA

What is a valid answer?
How much time should it take?
How many resources are available to it?
Contexte d'utilisation ?
Quel format de données ? (JSON/ XML)
Quel est une valide réponse ?
Combien de temps ?
Besoin de stocker les données ? In memory ? DB ?
fraicheur de la donnée importante ? (Mise en place de cache)
Utilisation de librairies externes ?
Description des API ?
Comment sont passés les paramètres ? (Dans le body, dans la query ?)
modèle de donnée ?
Erreurs possibles ?
Inputs et outputs possibles ?
Pagination ?
Fragmentation ? (renvoyer une réponse adaptée)
Doit-on prévoir un rate limiter ? (requêtes par seconde)
Authentication ?
Testing ?
# Processus

1. Gathering requirements / Scoping the problem / Constraints and use cases

Le service a t-il besoin de statistiques ?
What problem is this application supposed to solve?
Who are the different kinds of users who will use this?
What purpose does each user have for using it?
What expectations does each user have for how the software will behave?
How many users ?
What kind of input data can I expect?
What kind of output does it need to provide?
How much data do we expect to handle ?
How many request per second ?
What is the expected read to write ratio ?
Quelles sont les horaires de fonctionnement ? Y a t-il des heures de pointe ?
Session utilisateur ?

2. Create a high level design / Design core components
- Main components and connections
- Justify ideas

3. Understanding bottle necks
- Traffic (request per second)
- Volume of data

4. Scaling the design
Voir scalability.md

# Exemple (bit.ly)
1. Use cases :
    1. Shortening
    2. Redirection
    3. Custom URLs
    4. High availability of the system

    Out of scope
    5. Automatic link expiration
    6. Manual link removal
    7. UI vs API
    8. Analytics

2. Constraints
Définir le :
- nombre de requêtes par seconde
- volume de données (sur 5 ans)
- nombre de données lues par seconde
- nombre de données écrites par seconde


3. High level Design

    1. Application server layer (serves the reqiests)
    - Shortening service
    - Redirection service

    2. Data storage layer (hash -> URL mapping)
    - Acts like a big hash table

    hashed_url = md5(original_url + random_salt)[:6]

4. Understanding bottlenecks
- Traffic no problem
- Application service layer
    - Affirmations :
        - No business logic
    - Choix technologique :
        - Start with one machine
        - Measure (load tests)
        - Add a load balancer + a cluster of machines over time
            - Le load balancer permet de gérer les pics de demande (+ autoscaling de AWS)
            - Le load balancer répond au problème d'availability puisque si une machine tombe l'autre peut prendre le relai.
- Data storage layer
    - Affirmations :
        - No relationships between objects
        - Reads are more frequent than writes
        - 3 TBs of urls, 36GB of hashes
    - Choix technologique :
        - MySQL :
            - Widely used
            - Mature technology
            - Index lookups are very fast
            - Clear scaling paradigms (sharding, master/slave replication, master/master replication)
            - Used by Google, Facebook, ...
    - Bottlenecks :
        - How to read rapidly from the DB to get the URL ?
            - Create a unique index on the hash (36GB+). We want to hold it in memory to speed up look ups
                - Option 1: Vertical scaling of the MysQL Machine
                - Option 2: Partition the data: 5 partitions, 600GB of data, 8GB of indexes
                - master-slave (reading from slaves, writing to master)

    - Choix technologique (suite) :
        - Use one MySQL table with 2 varchar fields
        - Create index and hold it in memory
        - Vertical scaling for a while
        - Eventually, partition data by taking the first char of the hash mod the number of partitions
        - master-slave replication