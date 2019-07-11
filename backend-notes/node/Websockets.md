---
title: Websockets
created: "2019-04-06T16:14:52.778Z"
modified: "2019-04-06T16:35:46.446Z"
---

# Websockets

## Introduction

Websocket est un protocole permettant d'établir une communication permanente entre le serveur et le client.
Avec ce protocole, il est possible d'envoyer des messages au serveur et de recevoir des réponses de manière évenementielle sans avoir à consulter le serveur pour obtenir une réponse.
Ce protocole permet par exemple de réaliser des push notifications, des chat, des jeux, des outils collaboratifs pour le travail....

## Socket.io

### Description

Socket.io est la librairie principale pour utiliser ce protocole avec JS.

## Comment l'utiliser ?

Pour utiliser socket.io, il faut créer 2 fichiers : un fichier js pour le serveur et un fichier html (utilisant du JS pour se connecter au serveur) pour les clients.

app.js
Le serveur va tracer l'ensemble des connections et garder une session pour chacun des utilisateurs.
Il existe 2 façons d'émettre :

- `socket.broadcast.emit("message", { user: socket.pseudo, message });` : Permet d'émettre sauf à l'utilisateur de la session
- `socket.emit("message", "Vous êtes bien connecté !");` : Permet d'émettre à l'ensemble des utilisateurs
  1 façon d'écouter les évenements :
- `socket.on("petit_nouveau", ....)`

```javascript
// Chargement de socket.io
const io = require("socket.io").listen(server);

io.sockets.on("connection", function(socket) {
  // Entrées
  socket.on("petit_nouveau", function(pseudo) {
    socket.pseudo = pseudo;
    console.log(socket.pseudo + " est connecté !");
  });
  // Quand le serveur reçoit un signal de type "message" du client
  socket.on("newMessage", function(message) {
    console.log(socket.pseudo, ":", message);
    socket.broadcast.emit("message", { user: socket.pseudo, message });
  });
  // Sorties
  socket.emit("message", "Vous êtes bien connecté !");
});
```

index.html

```HTML
<script src="/socket.io/socket.io.js" />;

var socket = io.connect("http://localhost:8080");
socket.emit("petit_nouveau", pseudo);
//   Entrées
socket.on("message", function(message) {
  console.log(message);
  const li = document.createElement("li");
  li.textContent = message.user + ": " + message.message;
  ul.appendChild(li);
});
// Sorties
send.addEventListener("click", () => {
  const li = document.createElement("li");
  li.textContent = pseudo + ": " + input.value;
  ul.appendChild(li);
  socket.emit("newMessage", input.value);
  input.value = "";
});
```
