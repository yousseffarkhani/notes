A voir comment optimiser la taille d'une image.

# Introduction

Golang est un language compilé. La sortie du code source compilé est un binaire qui peut être executé directement sur l'OS cible sans avoir besoin d'autre software.
Le binaire contient déjà un serveur web permettant de recevoir et consumer des requêtes http.
C'est la principale différence comparé à d'autres languages tels que PHP où un webserver (Apache, Nginx) est requis afin de transformer les requêtes entrantes en un protocole que PHP comprend. (voir CGI - Common Gateway Interface, spécifie l'interface entre le serveur web et le programme.)

# via Docker

Pourquoi utiliser Docker pour deployer une application Go ?

- La plupart des web applications disposent de templates et de fichiers de configurations. Docker permet de garder ces fichiers synchrones par rapport à l'application quelque soit l'environnement.
- Docker permet d'avoir exactement la même configuration en developpement que en production.

# Processus

1. Créer un Dockerfile
2. Créer l'image : `docker build -t nom_image .`

- . indique que l'image sera créée depuis le directory actuel

3. Lancer le container : `docker run -it --rm --name nom_container -p 8080:8080 -v filepath_to_app_on_local:filepath_to_app_on_container -w filepath_to_app_on_container nom_image`

- it permet de démarrer le container en mode interactif
- rm clean le container après son extinction
- filepath_to_app_on_local : /github.com/...
- filepath_to_app_on_container : /go/src/nom_app

# Déploiement en production

L'automatisation du déploiement consiste à :

- Télécharger les dépendances
- Build après que les changements aient été push sur le git repo
- Effectuer les tests
- Si le build est ok et que les tests sont ok alors construire une image Docker
- Envoyer l'image vers le Docker Hub
- MAJ le server de manière à ce qu'il utilise la dernière version de l'image
