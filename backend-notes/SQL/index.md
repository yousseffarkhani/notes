# A quoi ça sert ?
Donner une représentation ordonnée de la donnée indexée
Augmente la vitesse de recherche. Permettent d"utiliser un binary search. La recherche passe d'une complexité O(n) à O(log(n)).
Sans index le processeur doit parcourir toute la table pour récupérer les informatios nécessaires.
Prennent de l'espace en mémoire et nécessitent à chaque addition / mise à jour / suppression de ligne de recalculer l'index.
Holds a copy of the indexed table data
# Qu'est ce que c'est ?
# Comment vérifier que la query utilise bien l'index ?
EXPLAIN ANALYZE permet d'afficher un rapport sur la query (Query plan).
```SQL
EXPLAIN ANALYZE 
SELECT SUM(value/campaign.cpm) AS "total_earnings" 
FROM "dailyImpressions" AS "dailyImpression" 
INNER JOIN "campaigns" AS "campaign" 
    ON "dailyImpression"."campaign_id" = "campaign"."id" 
WHERE ("dailyImpression"."podcast_id" IN ('2', '2') 
    AND "dailyImpression"."day" 
        BETWEEN '2020-03-26 00:00:00.000 +00:00' AND '2020-03-29 00:00:00.000 +00:00') 
LIMIT 1;
```

https://thoughtbot.com/blog/why-postgres-wont-always-use-an-index
https://devcenter.heroku.com/articles/postgresql-indexes
https://thoughtbot.com/blog/why-postgres-wont-always-use-an-index
https://thoughtbot.com/blog/reading-an-explain-analyze-query-plan
https://thoughtbot.com/blog/advanced-postgres-performance-tips
