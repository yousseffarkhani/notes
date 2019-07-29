---
title: 02 - Déployer une application sur docker
created: "2019-04-05T12:25:12.899Z"
modified: "2019-04-05T13:54:25.369Z"
---

# 02 - Déployer une application avec docker-compose

# Introduction

Docker compose est un outil permettant de définir et de lancer plusieurs containers dockers.

## Créer un dockerfile d'une application nodeJS + mongoDB

- DockerFile : Permet de configurer l'image
- docker compose file (docker-compose.yml) : permet de lier plusieurs containers (ex: nodeJS et mongoDB)

## Processus pour NodeJS

0. Installation de docker-compose
   https://docs.docker.com/compose/install/
1. Créer une Dockerfile

```Dockerfile
FROM node:11

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]
```

2. Créer une docker-compose.yml

```yml
version: "3"
services:
  app:
    container_name: docker-node-mongo # nom du service (optionnel)
    restart: always # Si le container tombe ordonne le redémarrage
    build: . # Utilise le Dockerfile dans le dossier actuel (.) afin de créer l'image
    ports:
      - "80:3000"
    links:
      - mongo
  mongo:
    container_name: mongo # (optionnel)
    image: mongo # Permet d'aller chercher l'image directement sur dockerhub
    ports:
      - "27017:27017"
```

**Dans le programme principal, la connection à la BDD se fait avec le nom du container :**

```
mongoose
  .connect(
    'mongodb://mongo:27017/docker-node-mongo',
    { useNewUrlParser: true }
  )
```

3. Créer un .dockerignore avec :

```
node_modules
npm-debug.log
```

4. Lancer `docker-compose up` (rajouter `-d` s'il n'y a pas besoin de la CLI)

5. Pour arrêter les containers, <kbd>Ctrl</kbd> + <kbd>C</kbd> (garde les données sauvegardées) ou `docker-compose down`

## Processus pour Golang

1. Créer une Dockerfile

```Dockerfile
FROM golang:latest

ADD . /go/src/app
WORKDIR /go/src/app


# Permet de télécharger toutes les dépendances
RUN go get -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["app"]
```

2. Créer une docker-compose.yml

```yml
version: "3.1"
services:
  web:
    build: .
    restart: always
    ports:
      - 8080:8080
  db:
    image: postgres
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: bird_encyclopedia
    volumes:
      - ./initdb.sql:/docker-entrypoint-initdb.d/initdb.sql # Fichier d'intialisation de la BDD
```

**Connection à un autre container dans le programme**

```go
const (
	driverName = "postgres"
	host       = "db"
	port       = "5432"
	user       = "postgres"
	password   = "secret"
	dbname     = "bird_encyclopedia"
)
psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)
db, err := gorm.Open(driverName, psqlInfo)
```

3. Créer un .dockerignore avec :
4. Lancer `docker-compose up` (rajouter `-d` s'il n'y a pas besoin de la CLI)
5. Si docker-compose ne met pas à jour les versions, utiliser `docker-compose up -d --force-recreate --build`
6. Pour arrêter les containers, <kbd>Ctrl</kbd> + <kbd>C</kbd> (garde les données sauvegardées) ou `docker-compose down`
   `docker.env`

# Tips

- Créer un fichier .env pour stocker les info serveurs
- Monter le code à l'aide de volume afin de refleter les changements dans le code directement (et ne pas avoir à rebuild l'image)
- Utiliser les nom des services pour les hostnames
  - postgres://db:5432
- Créer des dossiers différents pour chaque Dockerfile

# Déploiement sur Heroku

- local development : Utiliser l'image docker officielle de postgres
- staging and production : Utiliser les add-ons heroku

env_file: .env
depends_on: - db
