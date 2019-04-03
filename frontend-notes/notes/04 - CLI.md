---
title: 04 - CLI
created: '2019-04-02T19:09:34.014Z'
modified: '2019-04-02T19:20:28.279Z'
tags: [Linux]
---

# 04 - CLI

#### Créer un script et le lancer
1. Pour créer un script, créer un fichier (se terminant par .sh pour indiquer qu’il s’agit d’un script shell).
1. Ajouter les instructions dans ce fichier.
1. Saisir `sh ‘nom du fichier script’`
#### Network
Curl permet d’interroger un serveur via la CLI comme le ferait un navigateur.
- `curl ‘url’ -i` → i permet d’ajouter les informations sur le header dans la réponse affichée dans la CLI.
- `curl ‘url’ -X POST -d ‘donnée’ -H ‘Content-type: application/json’` → -d permet d’ajouter de la donnée dans le corps de la requête
- `curl ‘url’ -X PUT -d ‘donnée’ -H ‘Content-type: application/json’` → -H permet de définir des éléments du header
- `curl ‘url’ -X DELETE` → -X permet de définir la méthode HTTP
#### Ouvrir des programmes
- VSCode : `code "nom du fichier"`
- libreoffice : `libreoffice "nom du fichier"`
- `google-chrome stable`
#### Autres
- `Ls -l` permet d’afficher sous forme de liste
- `history` : affiche l'historique des commandes
- `Ctrl + R` : recherche dans l'historique des commandes
- `cat / less` : permet de lire un fichier
- `ws` : word count
- `diff` : compare 2 fichiers
- `man` : permet d'afficher le manuel
- `sudo apt-get install`
- `sudo apt remove / apt purge`
- `Ctrl + C` : exit
- `Ctrl + D` : fin d'input
- `mkdir` : créer un dossier
- `mkdir -p` : permet de créer des « nested » folder testdir/inner_folder
- `mv / cp` : mv folder/old-file.txt folder/new-file.txt (permet de rennomer le fichier) 
- `rm / rmdir / rm -r` (si le dir n’est pas vide) : supprimer
- `touch` : Créer un fichier
- `sleep 1` : met en pause le processus pendant 1 sec

#### File permissions
- `chmod` : change les permissions (w : 4, r : 2, x : 1)
- `chown` : change le propriétaire
- `chgrp` : change le groupe

#### chaining commands : voir https://speakerdeck.com/62gerente/bash-introduction
- `ctrl + a/e` : Se placer en début / fin de ligne
- `ctrl + u` : effacer tout ce qu’il y a à gauche du curseur

## Ressources
https://explainshell.com/

## Pour approfondir
https://www.learnenough.com/command-line-tutorial
https://www.codecademy.com/learn/learn-the-command-line

