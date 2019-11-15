# Firewall
Permet d'accepter et refuser certaines requêtes.
Les règles d'acceptation sont souvent basées sur le port demandé et sur le concept de least privilege.
3 ports sont en général disponibles pour les serveurs web :
- SSH (22)
- TCP (80)
- SSL (443)
Pour les BDD :
- MySQL (3306)

# SSL
Les communications sont encryptées en SSL à partir du Load Balancer vers internet.
Le reste des communications au sein du VPS sont en TCP.



