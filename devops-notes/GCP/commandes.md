# Général 
- Créer un projet : gcloud projects create PROJECT_ID
- Changer de projet: gcloud config set project [PROJECT_ID]
# App Engine
- Créer une application: gcloud app create
- Déployer l'applicatiob: gcloud app deploy
# Pub/Sub
- Créer un topic: gcloud pubsub topics create my-topic
- Créer un abonnement: gcloud pubsub subscriptions create my-sub --topic my-topic --ack-deadline=60
    - ack-deadline: délai de confirmation
- Lister les sujets (topics): gcloud pubsub topics list
- Lister les souscriptions: gcloud pubsub subscriptions list
- Publier un message: gcloud pubsub topics publish my-topic --message hello
- Extraire les messages publiés: gcloud pubsub subscriptions pull --auto-ack --limit=2 my-sub
    - auto-ack: Permet d'extraire et de confirmer le message
> Pub/Sub ne garantit pas l'ordre des messages
# Dataflow
