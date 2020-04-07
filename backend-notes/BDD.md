# View
Une vue dans une BDD est une synthèse d'une requête d'interrogation de la base. On peut la voir comme une table virtuelle.
Elles sont pratiques afin de réaliser des requêtes complexes au préalable et abstraire cette complexité via une table virtuelle.
La vue peut ensuite être utilisée comme une table classique (uniquement avec des SELECT statements).
## Bénéfices
- Limitation de l'accès aux informations (Certains utilisateurs ne peuvent accéder à des informations quniquement via une vue limitée.)
- Donne accès à une table simple (pas besoin de réaliser des requêtes JOIN & co)
- Donne accès à une interface qui changera moins que la table sous-jacente. Même si la structure de la table change, les views pourront être modifiées pour retourner le même structure de résultat.
