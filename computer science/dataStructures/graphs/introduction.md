# Source
https://www.calhoun.io/lets-learn-algorithms/
# Attributs d'un graphe
- Cyclical / Acyclical : Représente la capacité à pouvoir retourner au vertex de départ.
Cela dépend fortement des directions des edges puisque s'il n'y a pas de direction alors on peut revenir par le même edge.
- Directional edges :
    - Bidirectionnal edges
    - Directed edges : One direction only
- Cost associated with edges : Représente le côut associé au voyage d'un vertex à un autre. Cette valeur peut avoir différentes unités comme le temps, l'argent, ...
    - Dans notre cas cela peut être les péages à payer à chaque passage
- Capacity associated with edges : Fixe une limite à l'utilisation des edges
    - Dans notre cas cela peut être 20 voitures par minute (cela est dû au nombre de voies, limite de vitesse, ...)
# Composants d'un graphe
- Vertex (Node / Point / Datapoint) : Représente une donnée dans un graphe
- Edge (Arc / line) : Représente les relations entre des vertex
# Exemple : Réprésentation d'une carte
Chaque ville est représentée par un vertex
Chaque route est représentée par un edge
Calcul de la distance minimum entre 2 villes

