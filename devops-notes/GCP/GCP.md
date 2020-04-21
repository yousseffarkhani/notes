# Introduction
GCP est une suite de services cloud hebérgés sur l'infrastructure google.
Les services vont du computing au storage en passant par le data analytics, machine learning, réseau.
# Services
Stackdriver: Logging, monitoring, tracing.
# Services de computing et d'hébergement
GCP propose différentes options:
- Cloud Functions : Serverless (FaaS)
Ce service permet d'écrire des fonctions simples qui sont associées à des évenements déclenchés par l'infrastructure et les services cloud.
L'environnement est entierement géré et ne nécessite pas de provisionner d'infrastructure.
Cas d'utilisation:
    - Traitement des données et opérations ETL pour du transcodage vidéo et du streaming de données IoT
    - Webhook (Permet de déclencher une action suite à un évenement) permettant de répondre aux déclencheurs HTTP
- Cloud Run: Permet de déployer des conteneurs sans état.
- App Engine: PaaS
Dans cette configuration Google gère à notre place les ressources c'est à dire l'hébergement, le scaling, la surveillance et l'infrastructure de l'application.

- Kubernetes Engine: CaaS (Conteneur as a service)
- Compute engine: IaaS
Création des machines virtuelles.
Networking, compute : Déploiement de serveurs virtuels
storage et database
# Services de stockage
- Cloud SQL
# Services de base de données
- CLoud Datastore
- CLoud Storage
# Service de big data
- BigQUery: Service d'analyse de données
- Cloud dataflow: Service permettant d'effectuer des tâches de traitement de données par lots. Ce service est approprié pour les tâches d'ETL.
Un pipeline représente une tâche de traitement de données. Le pipeline lit les données d'entrée, effectue des transformations puis produit des données en sortie.
Les transformations peuvent inclure le filtrage, le regroupement, la comparaison ou la fusion de données.
- Cloud Pub/Sub: Service de messagerie asynchrone.
Grâce à ce service l'application peut envoyer des messages sous la forme de JSON à une unité de publication appelée sujet (topic). Les sujets représentent une ressource globale. Par conséquent, les autres applications des projets peuvent s'y abonner pour recevoir les messages dans le corps de la requête. 
https://cloud.google.com/pubsub/docs/overview?hl=fr
# Services de machine learning
# Services d'orchestration de workflows
Cloud composer est une solution basée sur Apache airflow.
Ce service permet de créer, planifier et de surveiller les pipelines dans plusieurs clouds.
Cloud composer dispose de connecteurs avec BigQuery, DataFlow, dataProc, Datastore, storage, Pub/Sub.

# Définitions
Serverless: Le provider gère les tâches de management d'infrastructure à notre place
Exemples de tâches:
- Resource Provisioning
- Scaling
- Reliability
- Déploiemet & configuration
- Monitoring
