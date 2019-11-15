TODO :
1) Etudier des packages existants
https://www.reddit.com/r/golang/comments/8yj3z5/how_to_become_better_at_structuring_go_code_ie/
2) Lire les différents livres


https://go-proverbs.github.io/
Accept interfaces return types

Les méthode sur les structs ne devraient servir qu'à modifier l'état du struct.


Les struct et interface sont la base du language.
Les structures peuvent disposer de méthodes.
Les interfaces listent des méthodes qu'un objet de n'improte quel type (struct, func, ...) doit implémenter pour satisfaire l'interface.

TODO : Lire le package json

Again, the standard library makes it look easy. Being able to say, in one sentence, what a package provides helps determine if objects and functions belong in the package. Well-defined packages are important because they delineate the system. Analogously, organs in a body are packages: brain, heart, lungs, stomach, etc. Each organ contains various parts: neurons in the brain, valves in the heart, alveoli in the lungs, acids in the stomach, etc. Organs are well-defined and well-delineated, and working together (hopefully) they compose the body. We think of, design, and implement Go programs in terms of packages. Complex systems are composed of many packages, each of which provides objects that code uses to compose solutions to needed tasks.