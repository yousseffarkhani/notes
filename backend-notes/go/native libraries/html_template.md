Source : https://www.calhoun.io/intro-to-templates/

# Contextual encoding

Go fournit 2 librairies text/template et html/template permettant de créer des templates (des modèles dynamiques de fichier texte).

La différence entre les 2 packages est que html/template va réaliser de l'encoding afin de prévenir les injections de code.
Par exemple si un user rentre comme pseudo : <script>alert("Hi!");</script> et qu'on affiche son pseudo sur une des pages alors on risque d'avoir du JS executé illicitement. C'est pourquoi après encodage ce pseudo devient &lt;script&gt;alert(&quot;Hi!&quot;);&lt;/script&gt;.

Les fichiers seront au format .gohtml.
