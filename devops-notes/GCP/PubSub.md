# Source
https://cloud.google.com/pubsub/docs/overview?hl=fr
# Introduction
PubSub est l'équivalent des middlewares orientés message d'entreprise.
Il s'agit d'un système d'ingestion et de diffusion d'évènements.
Ce service permet une communcation hautement disponible entre différentes applications.

# Concepts
- Sujet: Ressource nommé à laquelle les &éditeurs envoient des messages
- Abonnement: 
- Message: Combinaison de données et d'attributs qu'un éditeur envoie dans un sujet et qui est finalement remis aux abonnés.
- Attribut de message: Paire clé-valeur qu'un éditeur peut définir pour un message. Par exemple, la clé `language_tag` et la valeur `en` pourraient être ajoutées auxx messages pour les marquer comme lisibles par un abonné anglophone.

# Workflow
Une application éditeur crée et envoie des messages dans un sujet.
Les applications d'abonnées créent un abonnement associé à un sujet pour recevoir des messages de celui-ci.

# Publier des données et les enregistrer dans une tbale Big Query
1. Créer le sujets Pubsub
2. Créer la table correspondante sur BigQuery
3. Créer la Cloud Function pour exporter les messages publiés vers BigQuery (la fonction va automatiquement créer un abonnement au sujet)
