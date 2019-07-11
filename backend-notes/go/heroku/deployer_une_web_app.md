# Commandes

1. Installer Heroku CLI
2. Se logger : `heroku login`
3. Initialiser un repo git (`git init`) et commit tous les fichiers
4. Créer un ProcFile et mettre dedans: `web: nom_du_projet` (web permet d'attacher ce programme à un HTTP routing stack et donc recevoir du trafic web)
5. Créer une application heroku : `heroku create`
6. Installer un dependency manager : godep (`go get -u github.com/tools/godep`) / glide / ...
7. Sauvegarder les dépendances : `godep save .`
8. Commit les nouveaux fichiers crées
9. Pousser le code vers l'application heroku : `git push heroku master`
10. Ouvrir le website depuis la CLI : `heroku open`
