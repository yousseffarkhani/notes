# Software containers and kubernetes engine
IaaS 
Package software as containers
Manage them using kubernetes
in IaaS smallest unit of compute is the virtual machine. The OS can be large (sometimes GB) and take time to boot up.

Cloud Build: managed service to build containers

# Kubernetes
## Introduction
Orchestrateur de containers
API control its operations
utilities:
- kubectl

deploy on a set of nodes called cluster
node: computing instance
nodes: virtual machines running in compute engine

Kuberetes engine: KaaS

gcloud container clusters create K1

pod: smallest deployable unit
container running in a pod

deployment: group of replicas of the same pod

pods in deployment are only accessible in the cluster

connect load balancer to expose it to internet.
Kubernetes creates a service with a fixed IP for the pods

services are fundamental to do loadbalancing

service: groups a set of pods together and provides a stable endpoint (IP adress)

Cluster > (services >) nodes > pods > containers

Cluster: group of machines where Kubernetes can schedule containers in pods
Nodes: machines containing pods
Pod: Group of containers
# Lab 
## Pré requis
Il faut que les APIs suivantes soit activées:
- Kubernetes Engine API
- Container Registry API

## Exercice: 
- Create Kubernetes engine cluster
- Deploy a loadbalancer service
- Scale the service

## Commandes
1. Créer un cluster (Cette commande nous authentifie automatiquement pour pouvoir utiliser kubectl): gcloud container clusters create webfrontend --zone us-central1-f num-nodes 2
2. Créer un deployment (crée 1 pod): kubectl run nginx --image=nginx:1.10.0
3. Exposer le deployment à l'externe (crée un service): kubectl expose deployment nginx --port 80 --type LoadBalancer
4. Scale up le deployment (nombre de pods): kubectl scale deployment nginx --replicas 3
## Autres commandes
- kubectl version: version installée de Kubernetes
- kubectl get pods
- kubectl get services





Configuring healthchecks
Managing roll out strategies
