# Sources
https://www.thesitewizard.com/general/move-website-to-ssl.shtml
https://www.youtube.com/watch?v=T4Df5_cojAs
https://stackoverflow.com/questions/37321760/how-to-set-up-lets-encrypt-for-a-go-server-application
# Introduction
SSL (osbolète) / TLS (Transport Layer Security) : Protocole de sécurisation des échanges sur internet. Sans ce protocole, il est possible d'intercepter les données ou de s'adresser à un fake server.

Les enjeux du TLS sont de :
- Chiffrer l'information
- Authentifier le serveur avec lequel on communique

Afin de répondre à ces 2 enjeux, il existe respectivement 2 mécanismes :
- Le chiffrement asymétrique et symétrique
- Le certificat d'authorité

# Processus de certification et de communication
## Postulats :
- Un message chiffré avec une clé privée ne peut être déchiffré que grâce à la clé publique
- Toutes les personnes ayant accès a une clé publique sont capables de vérifier la signature d'un message et identifié qu'il s'agit bien de la personne disposant de clé privée l'ayant envoyé.
## Communication avec le client
1. Le client demande une ressource (ex : youtube.com)
2. Résolution DNS ...
3. Le serveur renvoi un certificat contenant une clé publique signée par une authorité de certification.
4. Le navigateur connait les clé publiques des authorités de certification et va donc vérifier le certificat envoyé afin d'authentifier le serveur.
5. Une fois l'authentification réalisée, le navigateur envoi une clé privée symétrique en la cryptant grâce à la clé publique.
6. Le serveur déchiffre le message et dispose de la clé publique symétrique.
7. Les échanges se font désormais grâce à cette clé publique symètrique.

## Mise en place de HTTPS sur le serveur
1. Le serveur souhaitant mettre en place HTTPS envoie une requête de signature de certificat à l'authorité avec la paire de clés afin que ce dernier la signe.
2. L'authorité renvoi la requête avec la signature réalisé avec la clé privée de l'authorité (toute personne disposant de la clé publique peut vérifier qu'il s'agit bien de la bonne signature)
3. Tous les navigateurs disposent d'une liste de certificats des authorités.

# Installation
## Go web server
### Certbot
1. SSH le serveur
2. Installer Certbot
```bash
    sudo apt-get update
    sudo apt-get install software-properties-common
    sudo add-apt-repository universe
    sudo add-apt-repository ppa:certbot/certbot
    sudo apt-get update
    sudo apt-get install certbot
    sudo certbot certonly --standalone
```
3. Utiliser les certificats dans l'application.
Pour les copier dans le répertoire projet, utiliser la commande :
`cp /etc/letsencrypt/live/playground.yousseffarkhani.website/* .TLS/`
```go
http.ListenAndServeTLS(":443", "fullchain.pem", "privkey.pem", server)
```
Changer également les ports exposés dans docker-compose et DockerFile.

3. Si l'application est sur docker il faudra soit lancer tous les scripts de l'étape 2 dans le container soit copier les certificats dans le container en étant root (sudo -i) sinon les fichiers seront vides.
4. Renouveller les certificats (autre solution créer une tâche CRON)
```bash
certbot-auto renew --standalone
```
Plus d'informations : https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go-with-little-help-of-lets-encrypt.html
## Cloudflare

# Types de certificats
https://support.dnsimple.com/articles/ssl-certificates-types/