# Commandes importantes

- `\?` : Liste toutes les commandes
- `psql -h host_name -U user_name` : pour se logger
- `\list` : pour lister les BDD
- `\c database_name` : pour se connecter à une BDD
- `\dt` : Lister les tables
- `\conninfo` : Info sur la connection actuelle
- `\q` : pour quitter psql
- `createdb -U postgres nom_de_la_bdd` : Pour créer facilement une DB depuis la CLI

# Commandes secondaires

# Docker

source: https://cinqict.nl/docker-go-webapp/

Pour se connecter depuis un container vers une DB en localhost : `docker run --network=host -it -p 8081:8081 basket:latest`
Pour lancer un container postgresql : `docker run -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword postgres -e POSTGRES_DB=bird_encyclopedia`
Pour se connecter au container et se logger à la BDD :

1. docker exec -it 504 bash
2. psql DBNAME USERNAME

# RDS
- Tester la connexion : `telnet database-1.dzdzxsxs.eu-west-3.rds.amazonaws.com 5432`
- Se connecter depuis bash : `psql --host=database-1.cgv4qujjbtwe.eu-west-3.rds.amazonaws.com --port=5432 --username=postgres --password`

# Mac
brew services start postgresql
brew services stop postgresql
brew services restart postgresql
