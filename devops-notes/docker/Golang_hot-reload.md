# Hot-reload Dockerfile

voir https://github.com/yousseffarkhani/golang_Postgresql-template

```Dockerfile
FROM golang

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/app
ADD . .

# Permet de télécharger toutes les dépendances
RUN go get ./
RUN go build



CMD if [ "${APP_ENV}" = "production" ]; then app; \
	else go get github.com/gravityblast/fresh && fresh; \
	fi

EXPOSE 8080
```

# Commandes

Build la version prod : docker build ./ --build-arg app_env=production
Build la version dev : docker build -t vol .
Récupérer les variables d'environnement : os.Getenv("APP_ENV")
