# Big data and machine learning fundamentals

Data driven descision making

Compute power for analytic and ML workloads

# Choosing the right solution


1. Challenges / Barriers the customer faced
2. How this challenges were solved by ML
3. Business impact

# Big Data challenges
- Migrating existing data workloads
- Analyzing large datasets at scale
- Building streaming data pipelines
- Applying machine learning to your data


# Product recommendations using cloud SQL and Spark

## Data
Ingest ratings of already rated houses
## Model
Train machine learning model to predict a user's rating
## Infrastructure

# Dataproc
It is a managed service to build hadoop and spark clusters without having to manage infrastructure


Spark: Is an analytics engine for batch and streaming data
Spark job: Représente par exemple le calcul de recommendation des maisons ou le calcul de PI.
Computation on Cloud Dataproc

Use cases (cloud dataproc):
Running a spark job on a hadoop cluster
1. Create and provision a cluster
2. Create spark job

Separation of computing and storage
Data can be stored in Cloud storage / Big Table / Big Query thanks to GCP network


Relational database management system are designed for many more WRITES than READS

Dataproc to train the recommendations machine learning model based on user's previous ratings.
This model is applied to create a list of recommendations for every user in the database.


# Lab
Préparation
1. Créer une instance SQL
2. Copier des données dans Storage (Stage data)
3. Charger les données dans SQL
Machine Learning
4. Créer les clusters et jobs sur Dataproc avec les autorisations pour écrire dans la BDD SQL
Analyse
5. Analyser les données dans SQL


# Tips
Ouvrir un IDE
cloudshell edit train_and_apply.py



echo "Creating bucket: gs://$DEVSHELL_PROJECT_ID"
gsutil mb gs://$DEVSHELL_PROJECT_ID

echo "Copying data to our storage from public dataset"
gsutil cp gs://cloud-training/bdml/v2.0/data/accommodation.csv gs://$DEVSHELL_PROJECT_ID
gsutil cp gs://cloud-training/bdml/v2.0/data/rating.csv gs://$DEVSHELL_PROJECT_ID

echo "Show the files in our bucket"
gsutil ls gs://$DEVSHELL_PROJECT_ID

echo "View some sample data"
gsutil cat gs://$DEVSHELL_PROJECT_ID/accommodation.csv

Autoriser cloud Dataproc à se connecter avec Cloud SQL
```sh
echo "Authorizing Cloud Dataproc to connect with Cloud SQL"
CLUSTER=rentals
CLOUDSQL=rentals
ZONE=us-central1-a
NWORKERS=2

machines="$CLUSTER-m"
for w in `seq 0 $(($NWORKERS - 1))`; do
   machines="$machines $CLUSTER-w-$w"
done

echo "Machines to authorize: $machines in $ZONE ... finding their IP addresses"
ips=""
for machine in $machines; do
    IP_ADDRESS=$(gcloud compute instances describe $machine --zone=$ZONE --format='value(networkInterfaces.accessConfigs[].natIP)' | sed "s/\[u'//g" | sed "s/'\]//g" )/32
    echo "IP address of $machine is $IP_ADDRESS"
    if [ -z  $ips ]; then
       ips=$IP_ADDRESS
    else
       ips="$ips,$IP_ADDRESS"
    fi
done

echo "Authorizing [$ips] to access cloudsql=$CLOUDSQL"
gcloud sql instances patch $CLOUDSQL --authorized-networks $ips
```
