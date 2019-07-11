# Introduction

Bolt est une DB No SQL (clé-valeur).
Bold est pratique pour la réalisation d'application CLI car elle n'a pas besoin de third party library.
En effet, pas besoin d'installer et de faire tourner un serveur comme on en aurait besoin pour MySQL / PostGreSQL.

Celle-ci est pratique pourl les cas de high reads, low writes.
Elle n'accepte que des []byte pour les key et value.

Pour transformer les types complexes en byte slice :
storm : https://github.com/asdine/storm
