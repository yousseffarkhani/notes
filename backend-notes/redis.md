# Introduction

Redis est un key-value store souvent utilisé comme système de cache.
L'interêt de Redis étant que le store est en mémoire (également persisté sur disque) ce qui fait que les recherches sont beaucoup plus rapides que sur une BDD SQL (stocké sur le disque).
De plus Redis offre la possibilité de mettre une date d'expiration sur les données.
Cependant Redis ne dispose pas des fonctionnalités d'une BDD SQL (relationships, complex queries, ...).

Quand des requêtes sont communes et nombreuses, il est possible de stocker le résultat dans Redis sous forme JSON.
