# Introduction

Traditionnelement les web applications sont énormes.
Un gros bloc de code est lancé sur un serveur et répond aux requêtes au format HTML, XML, JSON.

Ces monolithes posent problèmes au bout d'un certain moment car plus ils grandissent en taille, plus ils deviennent complexe.
En effet, si on souhaite ajouter une fonctionnalité à l'application, il est souvent difficile de changer ou réutiliser le code existant à cause des dépendances dans le code.

# Défintion

Afin de simplifier ces monolithes, une nouvelle approche est apparue. Celle-ci consiste à éclater ce bloc en de petits bouts de code chacun réalisant une tâche particulière (cela s'appelle un (web) service).
L'intérêt d'éclater comme cela est qu'on peut changer des parties de code sans pour autant avoir à faire attention aux dépendances.

De plus, les différentes équipes peuvent utiliser des languages différent sans que cela ait d'impact.

Cette indépendance entre chaque bloc est possible en se mettant d'accord sur un format et un protocole commun. Pour les web services les formats les plus populaires sont JSON et XML et le protocole est HTTP.

# Quand utiliser des microservices ou un monolithe ?

Microservices :

- De nombreuses personnes et équipes seront amenés à travailler sur le système (loosely coupled)
- Besoin de scalibilité pour servir des millions d'utilisateurs en parrallèle
- Besoin d'un système hautement disponible
  Monolithique :
- Le système est développé par une personne
- Pas besoin de scalabilité
- Besoin d'une architecture simple

# Exemple

Application permettant de créer des cartes postales virtuelles
