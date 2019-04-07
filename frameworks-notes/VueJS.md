---
title: VueJS
created: '2019-04-03T13:33:05.699Z'
modified: '2019-04-03T13:33:23.976Z'
---

# VueJS
Vue est un framework permettant de construire des interfaces utilisateurs.
Une des plus importantes foncitonnalités de vue est la réutilisation de composants : header, navigation, …
La modification simple du DOM lorsqu’une valeur change
 
{{variable}} permet d’afficher une variable
Attribute binding

Conditional rendering
v-if= « condition » / v-else-if= « condition » / v-else= « condition » permettent de créer des conditions.
v-show= « condition » permet au contraire de v-if de ne pas retirer/rajouter l’élément au DOM mais plutôt de le masquer. Cette solution est plus performante lorsque l’élément est amené à être affiché/masqué plusieurs fois.
List rendering
v-for= « detail in details » permet de boucler sur un tableau


Event handling
@ est équivalent à v-on (v-on:click / @mouseover / @keyup.enter)

: est équivalent à v-bind. V-bind est utilisé pour passer un paramètre à un attribut html (v-bind:href, v-bind : alt, v-bind:src, …)
:disabled= « condition » : Pour rendre un composant non utilisable

class  & style bindings
:class=’{disabledButton: !product.inventory}’
La classe disabledButton a été rajoutée à la condition que product.inventory soit false.
La même chose peut être réalisée pour les styles.
Computed properties
Les computed properties sont des calculs simples réalisés sur les données. 
Les computed properties sont plus interessantes que des methodes en terme de performance.
Components
Les components sont des morceaux de code réutilisables.
Lorsqu’un composant est déclaré, une fonction data(){} est déclaré par la même occasion. Cette fonction permet de séparer les données de chacun des composants crée.
De base, les composants n’ont pas accès aux données externes (isolated scope), il faut donc utiliser une props. Une props est un attribut permettant de passer des informations d’un élément parent à son enfant.
Pour recevoir une props, le composant doit :
    • déclaré explicitement le fait qu’il va recevoir une props.
    • Transmettre la prop dans l’attribut html (:premium= « product.premium »)
Communicating events
Approfondir
https://www.vuemastery.com/courses/

