# Machine learning with SQL

BigQuery 
Data warehouse
- Serverless
2 services in one connected by google network:
- Storage
The tables are highly compressed as columns 
Bulk data and streaming data ingestion
- Query
Can query CSV files hosted on Cloud storage / Google sheets aswell
Query more efficient when data is stored in Big Query storage 
# Big Query Engine
Service d'analyse en mémoire rapide/ 
Cela permet d'analyser les données stockées dans BigQUery avec un temps de réponse aux requêtes inférieur à 1 seconde et une simultanéité élevée.
# Big Query slot
Il s'agit d'une unité de capacité de calcul nécessaire pour exécuter des requêtes SQL.
BQ calcule automatiquement le nombre de slots requis en fonction de la taille et de la complexxité de la requête.
# In-memory query execution
https://cloud.google.com/blog/products/gcp/in-memory-query-execution-in-google-bigquery
# Optimize Big Query
https://medium.com/zalora-data/optimising-queries-in-bigquery-for-beginners-971be491f1de
# Tips
Créer une table avec le résultat d'une query:
```SQL
CREATE OR REPLACE TABLE <DATASET_NAME>.<TABLE_NAME> AS
SELECT * 
FROM xxxx
```
Créer une vue avec le résultat d'une query:
```SQL
CREATE OR REPLACE VIEW <DATASET_NAME>.<TABLE_NAME> AS
SELECT * 
FROM xxxx
```

# Dataprep
Permet de fitlrer et cleaner les données grâce à une UI en vue du process des datas.

# Dataflow

# GIS
GIS functions
Geoviz

# Big query ML
- Write machine learning models with SQL
- Expermient and iterate right where your data lives
- Build classification and forecasting models

## ML Project phases
1. ETL into BigQuery
2. Feature (columns) selection 
3. Model Creation (model type, ...)

4. Evaluate performance of model
5. Prediction
# Machine learning
Models on structured data 
## Supervised ML
Historical right answer
## Forecasting
## Classifying 
buy / no buy
Risk evaluation
## Recommendation
## Unsupervised ML
Exploration
