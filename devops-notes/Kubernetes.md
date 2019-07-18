---
title: Kubernetes
created: "2019-04-05T08:37:18.402Z"
modified: "2019-04-05T09:17:08.854Z"
---

https://www.cncf.io/the-childrens-illustrated-guide-to-kubernetes/

# Kubernetes

## Introduction

Kubernetes est un orchestrateur de container.
Il permet de gérer la charge (load balancing), monitorer et traiter les problèmes (ex: un container tombe et k8s en relancera un immédiatement).
En résumé Kubernetes permet de gérer plusieurs ordinateurs comme étant unique.

### Avantages:

- Permet de créer des applications plus scalables et robustes/résilientes car les containers peuvent être déployés sur plusieurs nodes.

Il existe également Swarm/rancher qui est un software plus simple que kubernetes.

https://www.youtube.com/watch?v=gpmerrSpbHg

## Définitions

- Node : host (OS) sur lequel tourne un container. Peut être une VM ou une machine physique.
- kubernetes cluster : désigne un ensemble de nodes managés par une unique instance de kubernetes
