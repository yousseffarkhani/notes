Google Cloud Platform
Compute
Storage
Big Data
Machine Learning

Cloud Networking


# Cloud computing
- On demand self service: Les services sont accessibles sans avoir besoin de personne exxterne
- Broad network access: Les services sont accessibles depuis n'importe ou
- Resource pooling: Il existe un provider qui met à disposition ses ressources (cela permet de faire des économies d'échelle)
- Rapid elasticity: On peut scale up et down les instances rapidement
- Measured service: On ne paye uniquement ce que l'on consomme

# Pourquoi le Cloud ?
edge network for lower latency
edge caching network close to end users to minimize ltency

# GCP Regions and zones
Les zones (europe-west2-a/europe-west2-b) sont groupées en région (europe-west2).

# Ressource hierarchy
Resources > Projects > Folders > Org node
Policies inherited downwards in the hierarchy
Project IDs project names
La permission la plus généreuse est accordée même si une permission plus stricte est définie plus haut dans la hiérarchie.

Les organization nodes sont optionnels.
Ils permettent de créer des folders et d'appliquer des permissions sur l'ensemble des projets.

# Interact with GCP
SDK: Manage ressources
gcloud
gsutil: storage
bq: big query

# Market Place
Ce service permet de déployer une stack (LAMP) sur une instance compute engine.
