---
title: 02 - Déployer une application sur docker
created: '2019-04-05T12:25:12.899Z'
modified: '2019-04-05T13:54:25.369Z'
---

# 02 - Déployer une application sur docker
## Créer un dockerfile d'une application nodeJS + mongoDB
- DockerFile : Permet de configurer l'image
- docker compose file (docker-compose.yml) : permet de lier plusieurs containers (ex: nodeJS et mongoDB)
0. Installation de docker-compose
https://docs.docker.com/compose/install/
1. Créer une Dockerfile
```
FROM node:11

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]
```
2. Créer une docker-compose.yml (retirer les commentaires)
```
version: '3'
services:
  app:
    container_name: docker-node-mongo //nom du service
    restart: always //Si le container tombe ordonne le redémarrage
    build: . //Utilise le Dockerfile dans le dossier actuel (.) afin de créer l'image
    ports:
      - '80:3000'
    links:
      - mongo
  mongo:
    container_name: mongo
    image: mongo //Permet d'aller chercher l'image directement sur dockerhub
    ports:
      - '27017:27017'
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
