---
title: 04 - CLI
created: "2019-04-02T19:09:34.014Z"
modified: "2019-04-02T19:20:28.279Z"
tags: [Linux]
---

# 04 - CLI

#### Créer un script et le lancer

1. Pour créer un script, créer un fichier (se terminant par .sh pour indiquer qu’il s’agit d’un script shell).
1. Ajouter les instructions dans ce fichier.
1. Saisir `sh ‘nom du fichier script’`

#### Network

Curl permet d’interroger un serveur via la CLI comme le ferait un navigateur.

##### HTTP:

- `curl ‘url’ -i` → i permet d’ajouter les informations sur le header dans la réponse affichée dans la CLI.
- `curl ‘url’ -X POST -d ‘donnée’ -H ‘Content-type: application/json’` → -d permet d’ajouter de la donnée dans le corps de la requête
  Exemple : `curl -X POST localhost:9000/getUsers -d '{"name":"Youssef"}' -H 'Content-type: application/json'`
- `curl ‘url’ -X PUT -d ‘donnée’ -H ‘Content-type: application/json’` → -H permet de définir des éléments du header
- `curl ‘url’ -X DELETE` → -X permet de définir la méthode HTTP
- `curl -o test.txt URL`: permet d'envoyer la sortie dans un fichier (utile en cas d'appel d'API Rest)
- (Alternative) `curl -O URL`: Crée automatiquement un fichier avec le contenu de l'URL.
- `curl -u login:password URL`: Permet d'envoyer ses logins à une API sécurisée.
- `curl -L http://google.com`: Permet de suivre le lien de redirection (ici vers http://wwww.google.com).

##### FTP:

- `curl -u test1@traversymedia.com:password -T test.txt ftp://ftp.traversymedia.com`: Permet d'envoyer le fichier test.txt vers un serveur FTP (dans le dossier test1).
- `curl -u test1@traversymedia.com:password -O ftp://ftp.traversymedia.com/hello.txt`: Permet de télécharger le fichier hello.txt du serveur FTP.

#### Installer/Désinstaller un programme

Les systèmes de gestion de package permettent de simplifier l'installation des programmes en gérant les dépendances, lancer manuellement les scripts, ...

Il existe plusieurs moyens d'installer des programmes sur Ubuntu :

- Logiciels ubuntu (de base les softwares listés proviennent du repo officiel d'ubuntu. pour avoir accès à plus de softwares, il faut ajouter le Canonical Partner repo)
- Utiliser les fichiers .deb :
  Les fichiers .deb sont similaires aux .exe de windows. Une fois téléchargés, il faut double clicker et suivre les instructions.
- Utiliser les commandes apt (RECOMMANDE):
  Il s'agit de l'équivalent de l'installation via logiciels ubuntu via la CLI.
  - `sudo apt-get install package_name`
  - `sudo apt remove / apt purge package_name`
- Installer/supprimer en utilisant un PPA (personal package archive) :
  Cela permet à un développeur de créer son propre APT repo et de mettre à disposition son logiciel. Cela évite d'avoir à attendre l'approbation officiel de ubuntu. De plus, les repo ubuntu ne disposent pas forcément de la dernière version du logiciel.
  Installer : 1) Ajouter le repo (ici le PPA du Numix project): `sudo add-apt-repository ppa:numix/ppa` 2) `sudo apt-get update` 3) `sudo apt-get install numix-gtk-theme numix-icon-theme-circle`
  Supprimer : 1) `sudo apt-get remove numix-gtk-theme numix-icon-theme-circle` 2) `sudo add-apt-repository --remove ppa:numix/ppa`
- Utiliser snap
- Utiliser dpkg
- Utiliser le code source de l'application

#### Ouvrir des programmes

- VSCode : `code "nom du fichier"`
- libreoffice : `libreoffice "nom du fichier"`
- Chrome : `google-chrome stable`

#### Autres

- `Ls -l` permet d’afficher sous forme de liste
- `history` : affiche l'historique des commandes
- `Ctrl + R` : recherche dans l'historique des commandes
- `cat / less` : permet de lire un fichier
- `ws` : word count
- `diff` : compare 2 fichiers
- `man` : permet d'afficher le manuel
- `Ctrl + C` : exit
- `Ctrl + D` : fin d'input
- `mkdir` : créer un dossier
- `mkdir -p` : permet de créer des « nested » folder testdir/inner_folder
- `mv / cp` : mv folder/old-file.txt folder/new-file.txt (permet de rennomer le fichier)
- `rm / rmdir / rm -r` (si le dir n’est pas vide) : supprimer
- `touch` : Créer un fichier
- `sleep 1` : met en pause le processus pendant 1 sec
- `Ctrl + L` : Permet de clean la fenêtre

#### File permissions

- `chmod` : change les permissions (w : 4, r : 2, x : 1)
- `chown` : change le propriétaire
- `chgrp` : change le groupe

#### Package manager

- `apt update` : Met à jour la liste des packages
- `apt upgrade` : Met à jour les softwares
- `apt install apache` : Installer un package

#### chaining commands : voir https://speakerdeck.com/62gerente/bash-introduction

- `ctrl + a/e` : Se placer en début / fin de ligne
- `ctrl + u` : effacer tout ce qu’il y a à gauche du curseur

## Ressources

https://explainshell.com/

## Pour approfondir

https://www.learnenough.com/command-line-tutorial
https://www.codecademy.com/learn/learn-the-command-line
