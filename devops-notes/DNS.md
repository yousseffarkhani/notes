# Tutoriel
Github pages :
https://learnetto.com/blog/github-pages
AWS :
http://techgenix.com/namecheap-aws-ec2-linux/

# Sources
https://www.thesitewizard.com/domain/multiple-websites-one-domain.shtml
# Domain names, websites and web hosts
Un nom de domaine est simplement un nom ("yousseffarkhani.website").
Il faut payer un domain registrar annuellement pour avoir le droit d'utiliser ce nom.
Un site web quand à lui est un ordinateur relié à internet mettant à disposition une liste de documents (pages html).
L'association des 2 permet aux gens de taper le nom dans un navigateur poru être dirigé vers ce serveur.
Il est possible de créer des sous domaines ("gmail.google.com / google.com / www.google.com)

Pour créer des sous domaines du site web, il faut la coopération du web host. Un web host est l'entreprise possédant le serveur (AWS).

L'atlternative aux sous domaines est d'héberger sous le root les différents sites. Ainsi example.com/court et example.com/portfolio redirigeront vers 2 ressources différentes.

# Principaux types d'entrées DNS (record)
https://blog.youdot.io/fr/4-types-enregistrements-dns-a-connaitre/
https://help.dnsmadeeasy.com/managed-dns/dns-record-types/
https://support.dnsimple.com/
Apex signifie la racine du nom de domain (Ex: yousseffarkhani.website) et est souvent représenté par @.
## A record (Address record)
Un A record permet de lier un nom de domaine à l'adresse IP v4 de l'ordinateur hôte du domaine.
Ex : 208.93.64.253 => www.dnsimple.com
## AAAA
Permet de faire la correspondance entre un nom de domaine et une adresse IP v6
Ex : irc.online.net => 2a01:e0b:1:f:2d0:b7ff:fea9:4da8
## CNAME
Permet de faire la correspondance entre un nom de domaine et un autre nom de domaine.

Ex : sql.domaine.ext => mysql1.online.net

Les CNAMES ont été crées afin de simplifier le cas où on change l'adresse IP du site.
En effet, si on enregistre que des A record alors il faudra changer l'ensemble des IP de ces records.
## ALIAS
Permet de rediriger un nom de domain vers un autre nom domaine que celui d'origine.
Ex : example.com => myapp.heroku.com
## NS
Signifie nom du serveur et indique quel nom de serveur a autorité sur le domaine
## SOA
## MX
Mail exchange est une liste de serveurs de messagerie à être utilisée pour le domaine
## TXT
Permet de stocker une chaine de 1024 caractères