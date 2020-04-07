# Cloud storage
Object storage
Immutable: versionning
unique key: URLs
Scalable service: don't need to provision capacity ahead of time
Buckets
HTTPS
ACL: Scope (user/group of users) - Permissions
## Classes
Multiregional: Pour les users
Regioal: Proche des VMs
Nearline: Once a month
Coldline: Disaster recovery (once a year)
## Commandes
Créer un bucket: gsutil mb -l $LOCATION gs://$DEVSHELL_PROJECT_ID
Copier un élément dans le directory actuel: gsutil cp gs://cloud-training/gcpfci/my-excellent-blog.png my-excellent-blog.png
Rendre le fichier publique: gsutil acl ch -u allUsers:R gs://$DEVSHELL_PROJECT_ID/my-excellent-blog.png


# No SQL
## Cloud Big table
No SQL database
ANalytical events (IoT, financial)
## Cloud Data store
NoSQL
Supporte les transactions
structured object
Replicas
SQL like queries
1MB par entité max
Large amount structured objects
App Engine

# SQL
## Cloud SQL
Relational database
transactions
Limité à la taille de l'instance choisie

## Cloud spanner
Amélioration de cloud SQL pour des besoins plus importants.
