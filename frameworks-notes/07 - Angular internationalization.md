## Internationalization (i18n)

Consiste à préparer l'app pour être utilisable dans plusieurs languages (format des dates, nombres, monnaie)

## Localization

Consiste à traduire le texte
Locale ID : Correspond à la langue choisie

## i18n pipes

Listes de pipes (le format est basé sur le LOCALE_ID) : DatePipe, CurrencyPipe, DecimalPipe, PercentPipe

## Processus de traduction

1. Marquer le texte statique à traduire
   Pour cela rajouter l'attribut i18n au tag à traduire : `<h1 i18n>Hello</h1>`
2. Créer une fichier de traduction
   Saisir `ng xi18n`
3. Traduire l'ensemble des mots importés
