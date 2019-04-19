# 05 - Angular Material
https://material.angular.io/components/categories
ng add @angular/material : Installe angular material et toutes les dépendances associées
## Importer un component
1. Dans app.module.ts, importer le module nécessaire (import {MatButtonModule} from '@angular/material';)
1. Le rajouter aux imports
1. L'ajouter au HTML (<button mat-button></button>)
## Theming
primary, accent, warn
<button mat-button color="primary"></button>

ng g @angular/material:material-nav --name=main-nav

More than one module matches. Use skip-import option to skip importing the component into the closest module.
ysf@ysf-PC:~/Bureau/neovel/neoreader$ ng g @angular/material:material-nav --name=main-nav --module app