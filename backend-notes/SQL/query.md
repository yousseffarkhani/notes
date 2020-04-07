# Source
https://pgexercises.com/
# TIPS
now(): permet d'utiliser la date actuelle

# Querying data
SELECT * FROM facilities: Permet de sélectionner un set de colones.
## SELECT
Il est possible de rajouter dans le SELECT n'importe quel élément (subqueries, ...) produisant une unique valeur.
### CASE
```SQL
SELECT name, 
	CASE WHEN (monthlymaintenance > 100) THEN
		'expensive'
	ELSE
		'cheap'
	END AS cost
	FROM cd.facilities;    
```
### DISTINCT
Enlève les duplicats. Cela s'applique à une **ligne** et non une colonne.

### Concatenate strings
`mems.firstname || ' ' || mems.surname`
### Fonctions d'aggrégat
Il s'agit de fonctions permettant d'agreger des données.
Ces fonctions retournent une table scalaire c'est à dire une table avec 1 colonne et 1 ligne.
Comme il s'agit d'une unique valeur, on peut la placer partout ou une valeur serait attendue (WHERE id = X, ...)
#### MAX/MIN
### EXTRACT
Permet d'extraire un élément d'une date => `SELECT EXTRACT (MONTH FROM "2017-06-15");`.

## ORDER BY
ORDER BY surname ASC: Permet de ranger le résultat avec la colonne donnée par ordre croissant
## WHERE
### Filtres
- name LIKE '%Tennis%': L'opérateur LIKE permet de matcher un mot avec un string.
    - %: Permet de matcher n'importe quel string
    - _ Permet de matcher n'importe quel caractère.
- id IN (1, 2): Prend une liste de valeurs possibles et match avec la colonne.
L'argument que prend IN est en réalité une table avec une seule colonne. Il est ainsi possible d'utiliser une query SELECT retournant 1 colonne en tant qu'argument (Cette query est appelée subquery).
## UNION 
Permet de combiner 2 requêtes dans une unique table. Les 2 requêtes doivent contenir le même nombre de colonnes et des données compatibles.
## JOIN
Permet de combiner des informations similaires provenant de plusieurs tables.
### INNER JOIN
Cette jointure va former en sortie de FROM une table comportant à la fois les données de la table bks et celles de mems.
Le SELECT et WHERE sont ensuite utilisés pour choisir les colonnes et les lignes à afficher.
```SQL
SELECT bks.starttime 
	FROM 
		cd.bookings bks
		INNER JOIN  cd.members mems  
			ON mems.memid = bks.memid
	WHERE 
		mems.firstname ='David' 
		AND mems.surname = 'Farrell';
```

### LEFT OUTER JOIN
Ce type de jointure fonctionne de manière similaire au INNER JOIN mais garde toutes les lignes de gauche même s'il n'y a pas de correspondance.
Cela permet souvent de répondre à une question disposant de données optionnelles (si X existe).
### JOIN OF JOIN
Une jointure est une fonction prenant 2 tables (left et right table) et retourne une table commune.
Ainsi il est possible de chainer les jointures.
```SQL
SELECT DISTINCT firstname || ' ' || surname AS member, name AS facility
FROM 
	cd.members mems
	INNER JOIN cd.bookings bks
		ON mems.memid = bks.memid
	INNER JOIN cd.facilities fclts
		ON bks.facid = fclts.facid
WHERE bks.facid IN (0,1)
ORDER BY member;
```
### Exemple
A noter que l'on ne peut pas utiliser le cost défini dans le SELECT au niveau du WHERE.
En effet, le WHERE est exécuté avant le SELECT.
```SQL
SELECT mems.firstname || ' ' || mems.surname AS member,
	fclts.name AS facility,
  CASE 
	  WHEN mems.memid = 0 
	  	THEN guestcost*slots
	  ELSE membercost*slots
  END AS cost
FROM
	cd.members mems
	INNER JOIN cd.bookings bks
		ON mems.memid = bks.memid
	INNER JOIN cd.facilities fclts
		ON bks.facid = fclts.facid
WHERE 
	starttime >= '2012-09-14' 
	AND  starttime < '2012-09-15' 
	AND ((mems.memid = 0 AND bks.slots*fclts.guestcost > 30)
	OR (mems.memid != 0 AND bks.slots*fclts.membercost > 30))
ORDER BY cost DESC;
```
## Subqueries
Une subquery peut être utilisée dans une clause SELECT, FROM ou WHERE.
La subquery est exécutée avant la query parente.
```SQL
SELECT DISTINCT mems.firstname|| ' ' || mems.surname AS member, 
	 (SELECT  mems2.firstname|| ' ' || mems2.surname AS recommender 
	  	FROM cd.members mems2 
		WHERE mems2.memid = mems.recommendedby)
FROM 
	cd.members mems
ORDER BY member
```
Dans ce cas la subquery est utilisée pour émuler un OUTER JOIN.
Pour chaque valeur de membre la subquery est lancée une fois.
Une subquery qui utilise des informations de l'outer query est appelée correlated subquery.

```SQL
SELECT 
	results.member, results.facility, results.cost
FROM
	(SELECT 
	 	mems.firstname|| ' ' || mems.surname AS member, 
		fclts.name AS facility,
	 	CASE
	 		WHEN mems.memid = 0 THEN fclts.guestcost*bks.slots
	 	ELSE
	 		fclts.membercost*bks.slots
	 	END AS cost
	FROM cd.members mems
	JOIN cd.bookings bks
		ON mems.memid = bks.memid
	JOIN cd.facilities fclts 
		ON bks.facid = fclts.facid
	WHERE 
		bks.starttime >= '2012-09-14' AND bks.starttime < '2012-09-15') results
WHERE results.cost > 30
ORDER BY results.cost DESC
```
Les subquery dans des FROM sont appelées inline views.

# Aggregates
Une fonction d'aggrégats prend en entrée une colonne de donnée et effectue une opération dessus et renvoie une unique valeur (scalar value).

Il est impossible de réaliser une requête telle que:
```SQL
select facid, count(*) from cd.facilities
```
COUNT(*) renvoie une unique valeur tandis que facid renvoi toute une colonne. Postgres ne sait pas quel facid mettre en face du résultat du COUNT.
Si l'on souhaite tout de même réaliser ce calcul, il faut utiliser une subquery qui va réaliser le calcul pour chaque ligne:
```SQL
select 
    facid, 
    (select count(*) from cd.facilities)
from cd.facilities
```

A noter que la fonction d'aggrégat est exécutée à la fin (ici la fonction WHERE est exécutée avant):
```SQL
SELECT COUNT(*)
FROM cd.facilities
WHERE guestcost >= 10
```
Cela permet de filtrer les données que l'on va aggréger.
## COUNT
Permet de compter le nombre de lignes du résultat.
Il existe différentes manières de l'utiliser :
- COUNT(*): renvoie le ombre de lignes
- COUNT(adress): renvoie le nombre d'adresses non nul dans le result set
- COUNT(DISTINCT adress): renvoie le nombre d'adresses différentes.
## Autres fonctions
MAX, MIN, AVG, SUM

# GROUP BY
Cette commande est utilisée en combination avec une fonction d'aggrégat.
Elle permet de grouper des éléments similaires d'une colonne et ensuite d'appliquer la fonction d'aggrégat sur ces groupes.
## Simple GROUP BY examples
```SQL
SELECT 
	recommendedby,
	COUNT(recommendedby) AS count
FROM cd.members
WHERE recommendedby IS NOT null
GROUP BY recommendedby
```

```SQL
SELECT facid, SUM(slots)
FROM cd.bookings
GROUP BY facid
```

```SQL
SELECT fclts.name,
	SUM(CASE
		WHEN memid = 0 THEN bks.slots*fclts.guestcost
	ELSE
		bks.slots*fclts.membercost
	END) AS revenue
FROM cd.bookings bks
	JOIN cd.facilities fclts
		ON bks.facid = fclts.facid
GROUP BY fclts.name
ORDER BY revenue ASC
```
## 2 columns GROUP BY examples
```SQL
SELECT facid, EXTRACT (MONTH FROM starttime) AS month, SUM(slots)
FROM cd.bookings
GROUP BY facid, month
```
## DATE_TRUNC
Cette fonction prend en entrée une unité et une date afin d'arrondir la date à l'unité choisie. 
Cela peut être utile lorsque l'on souhaite grouper les données pour executer une fonction d'aggrégat:
```SQL
SELECT DATE_TRUNC('day', occurred_at) AS day,
       COUNT(user_id)
  FROM benn.fake_fact_events
 WHERE event_name = 'complete_signup'
   AND occurred_at >= '2014-03-10'
   AND occurred_at <= '2014-05-26'
 GROUP BY 1
```
## GROUPING SETS (ROLLUP / CUBE)
Lorsque nous réalisons de la data analyse, on souhaite parfois aggréger sur différents niveaux d'aggrégation pour pouvoir zoom in et out à différents niveaux.
Dans le cas suivant, on souhaite avoir le usage de chaque facility mais également le total par mois.
Pour cela, on utilise l'opérateur ROLLUP. Ce dernier produit une hiérarchie d'aggrégats selon l'ordre des valeurs transmises.
Par exemple ROLLUP(facid, month) va tout d'abord donner l'aggrégat sur (facid, month) puis (facid) puis ().
Il suffit de changer l'ordre des valeurs pour avoir un autre type de hiérarchie (month, facid).
Pour avoir toutes les possibilités, on peut utiliser CUBE, cela va produire (facid, month), (facid), (month), ()
```SQL
SELECT facid, 
    extract(month from starttime) as month, 
    sum(slots) as slots
FROM cd.bookings
GROUP BY ROLLUP(facid, month)
```
# HAVING
HAVING permet de filtrer l'output de la fonction d'aggrégat.
Souvent son comportement est confondu avec celui de WHERE.
WHERE est utilisé pour filtrer les données en entrées de la fonction d'aggrégat tandis que HAVING filtre les données en sortie de la fonction d'aggrégat.
```SQL
SELECT facid, SUM(slots) as "Total Slots"
FROM cd.bookings
GROUP BY facid
HAVING SUM(slots) > 1000
```

HAVING ne permet pas de sélectionner une colonne faisant partie du SELECT.
```SQL
SELECT fclts.name, 
	SUM(
	  CASE
	  	WHEN bks.memid = 0 THEN bks.slots*fclts.guestcost
	  ELSE
	  	bks.slots*fclts.membercost
	  END) AS revenue
FROM cd.bookings AS bks
	JOIN cd.facilities AS fclts
		ON bks.facid = fclts.facid
GROUP BY fclts.name
HAVING SUM(
	  CASE
	  	WHEN bks.memid = 0 THEN bks.slots*fclts.guestcost
	  ELSE
	  	bks.slots*fclts.membercost
	  END) < 1000
ORDER BY revenue
```
Le fait de devoir répéter la fonction d'aggrégat (voir au dessus: SUM(...)) est assez error prone et non DRY.
HAVING doit être utilisé pour de simples queries.
Autrement l'approche à privilégier est la suivante:
```SQL
select name, revenue from (
	select facs.name, sum(case 
				when memid = 0 then slots * facs.guestcost
				else slots * membercost
			end) as revenue
		from cd.bookings bks
		inner join cd.facilities facs
			on bks.facid = facs.facid
		group by facs.name
	) as agg where revenue < 1000
order by revenue;
```
# CTE: Common Table Expxressions
Les CTE consistent à définir une vue utilisable dans une query.
Cela est utile lorsque l'on doit répeter la query plusieurs fois.
```SQL
WITH sum AS (
  SELECT facid, SUM(slots) AS total_slots
  FROM cd.bookings
  GROUP BY facid)

SELECT facid, total_slots
FROM sum
WHERE total_slots = (SELECT MAX(total_slots) FROM sum)
```
# Window functions
Les window functions opèrent sur le résultat de (sub-)query après la clause WHERE et les aggrégations standards.
Elles opèrent dans une "window" de données.
Elle permettent de réaliser des calculs sur des lignes liées à la ligne actuelle.
Une window function dispose d'une OVER clause car elle s'exxécutera OVER une certaines portions de lignes (ou toutes les lignes) séléctionnées par la query.
Elles sont souvent utilisées pour calculer des sommes cumulées, des numéros de rang.
## OVER
Source: https://openclassrooms.com/fr/courses/4449026-initiez-vous-a-lalgebre-relationnelle-avec-le-langage-sql/4563021-apprehendez-le-fenetrage-avec-over-et-partition-by
- OVER(): Renvoi l'ensemble du tableau en entrée de la window function.
- OVER(PARTITION BY ...): Va réaliser un calcul sur un groupement de données
- OVER(PARTITION BY ... ORDER BY ...): La window function considère que les valeurs de la colonne sont ordonnées. 
	Elle va donc effectuer le calcul sur la 1ère valeur et renvoyer le résultat puis effectuer un calcul sur la 1ère et 2ème valeur et renvoyer le résultat, etc ...
	Cette méthode permet de caculer par exemple des sommes cumulées ou des rangs.
## Fonctions
Source: https://www.gab.lc/articles/window_functions_postgresql/
row_number(): return the number of the current row from 1 to N
```SQL
SELECT row_number() OVER(ORDER BY grade), id, name, grade
FROM students
ORDER BY grade
```
rank(): return the rank of the current row (with gaps)
dense_rank(): return the rank of the current row (without gaps)
percent_rank(): relative rank of the current row
cume_dist(): relative rank of the current row (bis)
ntile(): group your rows in buckets
lag(): return the previous row (or one of them)
lead(): return the next row (or one of them)
first_value(): return the first row
last_value(): return the last row
nth_value(): return the Nth row

## Exemple 1
```SQL
select count(*) over(partition by date_trunc('month',joindate)),
	firstname, surname
	from cd.members
order by joindate
```
Dans cet exemple les données sont partitionnées par mois. Pour chaque ligne sur laquel opère la window function, la fenêtre est toutes les lignes disposant d'une joindate du même mois.
## Exemple 2
```SQL
SELECT row_number() over(order by joindate),
	firstname, surname
FROM cd.members
order by joindate          
```
Dans cette query aucune partition n'est définie cela veut dire que la partition est le dataset en entier.
Comme un order est défini dans le window function, pour chaque ligne la fenêtre sera début du dataset jusqu'à la ligne actuelle.
## Exemple 3
```SQL
select facid, total from (
	select facid, total, rank() over (order by total desc) rank from (
		select facid, sum(slots) total
			from cd.bookings
			group by facid
		) as sumslots
	) as ranked
where rank = 1
```
```SQL
select facid, total from (
	select facid, sum(slots) total, rank() over (order by sum(slots) desc) rank
        	from cd.bookings
		group by facid
	) as ranked
	where rank = 1
```
Les window function sont appliquées tard dans la SELECT function (après l'aggrégation).

