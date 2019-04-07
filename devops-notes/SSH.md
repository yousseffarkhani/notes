---
title: SSH
created: '2019-04-05T15:31:23.105Z'
modified: '2019-04-05T16:16:26.066Z'
---

# SSH
## Plus de commandes
https://gist.github.com/bradtraversy/f03df587f2323b50beb4250520089a9e
## Introduction
Secure SHell permet de communiquer avec un ordinateur distant.
Il donne en règle général acccès au bash de l'autre machine.
C'est également utile avec les services cloud type github.
## Communication client/serveur
Le client doit disposer de SSH.
Le serveur doit disposer de SSHD (Open SSH Daemon) lui permettant d'écouter les connexions SSH.
## Méthodes d'authentification
- Mot de passe
- Clé publique/privée
- Host based

Le mieux est d'utiliser la clé pubique/privée car c'est plus sécurisé et ne nécessite pas de rentrer ses logs à chaque fois.
## Commandes
- Login : `ssh will@192.123.1.23`
- Sortir du bash de la machine: `exit`
- Générer une paire de clé privée/public: `ssh-keygen`
  - `id_rsa` : clé privée
  - `id_rsa.pub` : clé publique
Ces fichiers se situent `~/.ssh`.

Il faut envoyer la clé publique au serveur dans authorized_keys ou au cloud.
- Manuellement : `cat ~/.ssh/id_rsa.pub`
- Automatiquement: `ssh-copy-id demo@198.51.100.0` (utiliser -i nom_du_fichier_identity)

Puis ajouter la clé privée à l'agent d'authentification (ssh-agent): `ssh-add ~/.ssh/id_rsa_do`





