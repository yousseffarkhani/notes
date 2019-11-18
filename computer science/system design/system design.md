# Source
https://www.hiredintech.com/classrooms/system-design
# Exemples
http://blog.gainlo.co/index.php/category/system-design-interview-questions/
https://www.hiredintech.com/classrooms/system-design/lesson/61
https://github.com/donnemartin/system-design-primer

https://www.youtube.com/watch?v=0163cssUxLA


# Processus

1. Gathering requirements / Scoping the problem / Constraints and use cases

What problem is this application supposed to solve?
Who are the different kinds of users who will use this?
What purpose does each user have for using it?
What expectations does each user have for how the software will behave?
How many users ?
What kind of input data can I expect?
What kind of output does it need to provide?
How much data do we expect to handle ?
How many request per second ?
How much time should it take?
What is the expected read to write ratio ?
Quelles sont les horaires de fonctionnement ? Y a t-il des heures de pointe ?
Le service a t-il besoin de statistiques ?
Session utilisateur ?
Besoin de stocker les données ? In memory ? DB ?
fraicheur de la donnée importante ? (Mise en place de cache)


Quel format de données ? (JSON/ XML)
Utilisation de librairies externes ?
Comment sont passés les paramètres ? (Dans le body, dans la query ?)
modèle de donnée ?
Erreurs possibles ?
Inputs et outputs possibles ?
Pagination ?
Fragmentation ? (renvoyer une réponse adaptée)
Doit-on prévoir un rate limiter ? (requêtes par seconde)
Authentication ?
Testing ?

2. Create a high level design / Design core components
- Main components and connections
- Justify ideas

3. Understanding bottle necks
- Traffic (request per second)
- Volume of data

4. Scaling the design
Voir scalability.md

