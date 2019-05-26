# AWS

## Neoreader

On va mettre le fichier dans S3 lorsque les utilisateurs se connecteront, ils vont uploader les fichiers index.html + JS + CSS.
Ensuite des requêtes seront réalisées vers le back-end java.

Objectif :

1. Mettre le fichier dist dans S3
2. Routing
3. HTTPS :

Certificat https : Dès qu'un site va appeler toutes les informations seront encryptées
ELB permet d'attacher un certificat HTTPS.

## Cloudfront

Permet de cacher des API

/!\ Se renseigner sur A domain CNAME & co !!!!

A Record (Adress Record)

## Learn More

AWS Certified Solutions Architect (associate)

## Introduction

## Services

### IAM

#### Introduction

IAM permet de manager les accès aux services AWS des utilisateurs et des services AWS via la création de users, groupes et rôles son propre VPC.
Pour affecter des droits d'accès, il faut attribuer des policies (S3 full access, admin, ...).

Roles : Représente les accès possibles d'un service AWS vers un autre servi ce AWS (ex: EC2 => S3)

### VPC (virtual private cloud)

#### Introduction

Un VPC est une partie d'AWS dans laquelle nous pouvons ajouter des services AWS.

Un VPC ressemble à un réseau privé d'accès à internet :
![](./attachments/VPC_basics.png)
![](./attachments/VPC_basics2.png)
un VPC dispose de plusieurs composants décris ci-dessous.

##### Internet Gateway (IGW)

Il s'agit d'une combination de hardware et software donnant accès à internet via une route à notre VPC.

##### Route tables

Une route table contient une liste de règles (appelées routes) définissant vers quelle service AWS le traffic va être dirigé.

##### Network Access Control List (NACL)

Il s'agit d'un firewall disposant de règles d'accès aux subnets.

##### Subnets

Permet d'avoir des sous-réseaux composés de services AWS

##### Availability zones

Il s'agit de datacenters distincts designés pour être totalement isolés en cas de failures et ainsi toujours avoir un datacenter disponible.
Lorsque l'on crée un projet, pour faire en sorte que le système soit résilient, il faut impérativement utiliser plusieurs availability zones pour créer de la redondance dans l'architecture et ainsi avoir un système hautement disponible et résilient

### S3 (Simple storage service)

#### Introduction

S3 dispose d'un WS permettant de stocker, récupérer n'importe quelles données.

#### Définition des composants

Buckets : Root level folder  
Folder : Subfolder
Objects : Fichiers stockés
Storage classes : standard/Glacier/ ...
Object lifecycles : Permet d'automatiser la transformation d'un objet d'une classe à une autre selon un délai fixé.
Permissions : Qui peut voir / modifier le bucket ou les objets. De plus, il est possible de rendre public l'accès à un object afin que n'importe qui disposant de l'url puisse le visualiser/télécharger.
Object versionning : Permet de garder l'historique des documents. Bucket level, apply to all objects

### EC2 (Elastic Cloud Compute)

#### Introduction

EC2 est un service mettant à disposition des serveurs virtuels.

Composants d'une instance EC2 :

- AMI (OS) : Il s'agit d'un package comprenant l'OS, des softwares packages et quelques paramètres) qui seront installés au lancement de l'instance.
- Instance type (CPU) : Définit le hardware utilisé par notre instance : nombre de CPU, RAM, memory, bandwidth capacity
- EBS (Elastic Block Store) : Hard drive
  - IOPS : Unité représentant le nombre d'opération d'écriture ou de lecture par seconde
- Security groups (firewall) : Représente une sécurité en plus des NACL (network acess control list)
- Ip adressing (network card) : Permet de définir une adresse IP pour une instance donnée. Par défaut, les instances disposent d'une adresse privée permettant de communiquer à l'intérieur du VPC. Cependant, pour communiquer avec internet, il faut obligatoirement une adresse IP publique.
- RAM

### RDS/DynamoDB

RDS (SQL) : Il s'agit d'un service mettant à disposition en SaaS des DB SQL telles que MySQL, MariaDB, Oracle,...
DynamoDB : NOSQL

### SNS (Simple notification service)

Il s'agit d'un service permettant d'envoyer des messages (par email, SMS, ...) à des destinataires (WS, adresse email, AWS lambda, ..) basés sur des évènements survenant dans AWS.

Cas d'usages :

- Une instance EC2 crash, Cloudwatch le voit informe SNS qui envoit un mail à un administrateur pour réparer l'instance.

Composants :

- Topics : Permet de grouper différents endpoints à qui envoyer les messages
  - ex : EC2 failures
- Subscriptions : endpoint (email adress, phone number)
- Publishers : alarm/event/humain donnant à SNS le message à envoyer

### CloudWatch

Service permettant de monitorer les autres services AWS. Il est souvent utilisé avec SNS pour envoyer des alertes.
Il peut également déclencher des actions automatiquement.

Exemples de métriques :

- EC2 : CPU utilization, Status Checks, Disk read/writes
- S3 : Number of objects, bucket size
- billing

Sur la base de ces métriques, on peut mettre en place des seuils.

Exemples de seuils :

- EC2 : CPU utilization > 80%
- S3 : Number of objects > 100
- billing > 500\$

### ELB (Elastic load balancing)

Ce service distribue le traffic entrant vers les différentes instances EC2 qui sont associées avec.
Ceci dans le but d'éviter que les instances soit surchargées et crash.

### Auto Scaling

L'autoscaling permet de rajouter (scale up) ou retirer (scale down) des instances EC2 sur la base du traffic entrant.

2 composants principals :

- Launch configuration : Permet de définir le type d'instance à lancer (AMI, RAM, ...)
- Autoscaling group : Règle à suivre pour rajouter ou retirer des instances

### Route 53

Route53 permet de configurer et manager les noms de domaines des applications/sites web hébergés sur AWS.

Record set : Permet d'associer différentes url (www.projectomega.com / projectomega.com) avec l'adresse IP du load balancer.

Etapes :

1. Register a domain
2. Créer des record set et attendre la propagation aux DNS servers.

/!\ Se renseigner sur A domain CNAME & co !!!!

### Lambda

Lambda permet de réaliser du serverless computing (ie lancer du code sans avoir à provisionner ou manager des serveurs).
Lambda vient en remplacement de EC2.
AWS Lambda exécutera le code uniquement quand cela est nécessaire et se scale automatiquement.
Seul le compute time est facturé.
Avec AWS Lambda, il est possible de lancer du code pour n'importe quelle application ou service back-end sans avoir à gérer l'administration.
AWS Lambda se charge de réaliser l'administration (server and OS maintenance, capacity provisionning, automatic scaling, code monitoring and logging).
