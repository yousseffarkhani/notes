---
title: 03 - Déployer une application sur le cloud (Digital Ocean)
created: '2019-04-05T13:15:52.110Z'
modified: '2019-04-05T16:43:01.061Z'
---

# 03 - Déployer une application sur le cloud (Digital Ocean)
## Introduction
### Digital Ocean
C'est un hosting service qui met à disposition des Droplets (VPS Virtual server).
## Monter une machine
1. Choisir un package (Logiciels installés + OS + puissance)
1. Installer les images via une des stratégies de déploiement possible sur la machine et lancer les containers

## Processus
1. La clé public devra par la suite être transmise au serveur (dans un fichier "authorized_keys") ou au fournisseur cloud .
Saisir `cat ~/.ssh/id_rsa_do.pub` pour la récupérer
2. Ensuite il faut ajouter la clé privée via `ssh-add ~/.ssh/id_rsa_do`
3. Saisir `ssh root@157.230.103.109` pour accéder à la machine.



## Lancer l'application avec docker
1. Dans un premier temps, mettre à jour les packages avec :
- `apt update` : Met à jour la liste des packages
- `apt upgrade` : Met à jour les softwares

2. Créer un nouvel utilisateur et lui donner les droits sudo (rester connecté en tant que root peut être dangereux)
  1. `adduser brad`
  2. `id brad` : permet de voir les droits
  3. `usermod -aG sudo brad` : Ajouter les privilèges
  4. Créer un fichier `/home/brad/.ssh/authorized_keys` en étant connecté en tant que root et ajouter la clé publique de brad à ce fichier (`sudo nano authorized_keys`)
  5. Désactiver le root password login : `sudo nano /etc/ssh/sshd_config` et désactiver `PermitRootLogin no` + `PasswordAuthentication no`
  6. Redémarrer le service SSHD : `sudo systemctl reload sshd`
2. Cloner le repo github contenant le projet (ou une autre stratégie de déploiement - voir ci-dessous)
  1. Dans un 1er temps, il faudra créer une clé public/privée et la partager avec github
  2. Il faudra probablement changer le propriétaire du fichier ssh pour le modifier: `sudo chown -R brad:brad /home/brad`
    - `-R`: recursive permet de changer tout ce qui est contenu dans le fichier brad
  3. Saisir `ssh-add /home/brad/.ssh/id_rsa_github` (si ça ne marche pas saisir `eval $(ssh-agent)`)
  4. Cloner le repo (ou une autre stratégie de déploiement - voir ci-dessous)
3. Se placer dans le fichier du projet et saisir `docker-compose up`

## Stratégies de déploiement
- `rsync -av . will@192.123.1.23:~/newapp`
- `git clone repo's SSH`
- `scp ~/test.txt will@192.123.1.23:~/newapp`

## Persistance des données 
 Traversy Media
If you want your data to persist in prod, add something like this to your mongo service...
 volumes:
  - ./data/db:/data/db
