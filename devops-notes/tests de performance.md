# Source
https://arodrigues.developpez.com/tutoriels/java/performance/plan-test-realiste/
# Description
Les tests de performance permettent de tester la robustesse et la performance de l'application.
Gatling (outil de test de performance) va simuler la connexion de plusieurs utilisateurs de la manière la plus réaliste possible.
Cela permet d'anticiper les crashs et de voir les temps de réponse pour les utilisateurs.
Différents types de tests existent :
    - charge : Simulation de la charge d'utilisateurs attendues pour voir les performances
    - capacité : On augmente la quantité d'utilisateurs par palier pour voir les limites du système.
    - stress : Simulation d'un grand nombre d'utilisateurs simultanés
    - endurance : SImulation d'une charge importante sur une longue durée
# Objectifs
- Résoudre des problèmes de performance d'une application
- Dimensionner une infrastructure de production
- Rassurer la DSI sur la montée en charge de l'application
- Valider les performances de l'application
- Comparer des solutions

# Workflow
1. Identifier les scénarios critiques (écrans consultés, champs saisis dans les formulaires, temps entre chaque action)
    - 3 types de scénarios critiques :
        - Fréquents (quotidiens)
        - Vitaux (déclenchement de la paye)
        - Risqués (recherches complexes, génération de documents)

