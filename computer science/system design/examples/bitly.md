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

    1. Application server layer (serves the requests)
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
