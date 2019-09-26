Source : https://docs.gitlab.com/ee/ci/

# Introduction
Les équipes de développement intègrent souvent des règles de fonctionnement qui les aident à produire du code de qualité. L'une d'elles est le git workflow
# Processus
1. Dev1 développe une fonctionnalité sur une nouvelle branche dédiée.
2. Lorsque la fonctionnalité est prête Dev1 crée une pull request pour intégrer sa branche dans la master.
3. Dev2 examine la PR. Dev2 pull en local et réalise des tests pour vérifier que tout fonctionne.
4. Dev2 réalise des retours à Dev1 qui s'occupe de corriger le code.
5. Lorsque tout est ok Dev2 valide la PR et la branche de développement est fusionnée dans la branche master.
# Intégration continue
L'intégration regroupe les tâches à faire lorsqu'un développeur a fini de coder sa nouvelle fonctionnalité. C'est à dire vérifier que :
- le code ne contient aucun bug
- le reste de l'application continue de fonctionner
- la fonctionnalité correspond aux bonnes pratiques de développement de l'équipe.

Le processus devient alors le suivant :
1. Dev1 crée une PR
2. Le service d'intégration continue lance l'application sur un serveur de test
3. Si les tests échouent, le service indique que la fonctionnalité n'est pas prête à être intégrée.
4. Si les tests passent alors OK
5. Dev2 peut lire la PR afin de la valider et que le code soit intégré
# Fonctionnement

# Commandes