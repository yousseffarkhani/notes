# Processus

1. Créer une instance (https://www.youtube.com/watch?v=bT_xXc7cOCQ)
2. SSH dans la VM (la KeyPair est dans /Téléchargements)
   ssh -i "EssentialsKeyPair.pem" ubuntu@ec2-35-180-164-29.eu-west-3.compute.amazonaws.com
3. sudo apt update
4. sudo apt install docker.io
5. sudo apt install docker-compose
6. Clone le repo
7. Créer le .env avec
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=secret
8. sudo docker-compose up
