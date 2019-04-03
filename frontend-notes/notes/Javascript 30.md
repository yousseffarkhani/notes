---
title: Javascript 30
created: '2019-04-03T12:36:15.032Z'
modified: '2019-04-03T12:48:26.520Z'
---

# Javascript 30
Si besoin d’explications, voir https://medium.com/@cmurphy580/lessons-learned-from-the-javascript-30-challenge-bdc9ce5e47b0
01 – Drum kit
HTML
Le tag <kbd>A<kd> représente une touche au clavier.
Il est possible de rajouter n’importe quel attribut à un tag avec le format data-attributName= « XX ».
Exemple : <audio data-key="65" src="sounds/clap.wav"></audio>
CSS
Transition : propriété CSS permettant d’ajouter des effets de transition lorsqu’une propriété change.
Syntaxe : transition : <propriété CSS> <durée de transition> [<fonction de transition> <transition-delay>]
Exemple : 
div {
  width: 100px;
  height: 100px;
  background: red;
  -webkit-transition: width 2s; /* Safari */
  transition: width 2s ease 1s, height 2s, transform 2s;
}
div:hover {
  width: 300px;
  height: 300px;
 transform : rotate(180deg) ;
}
Il est possible de choisir toutes les propriétés avec all > transition : all 0 .5s ;
Propriétés intéressantes : 
transform : scale(1.1) ;
JS
L’événement « transitionend » est déclenché lorsqu’une transition est terminée. L’objet event résultant dispose de 2 propriétés : propertyName (la propriété ayant subi une transition) et elapsedTime.
Fonctions intéressantes : 
audio.play() ; (après avoir querySelector un tag HTML audio)

02 – Clock
CSS
Propriétés intéressantes : 
transform : translateY(-3px); (il est possible d’ajouter un effet de transition avec)
transform: rotate(90deg);
transform-origin: 100%; (définit le point d’ancrage de la transformation)
Position :
    • relative : L’élément prend son espace normal mais est positionné par rapport à 2 paramètres top/bottom et left/right
	Exemple : 
	position : relative ;
	top : 20px ;
	left : 20px ;
    • absolute : L’élément ne prend aucun espace et est positionné par rapport à 2 paramètres top/bottom et left/right
      Exemple : 
      position : absolute;
      top : 20px ;
      left : 20px ;
    • fixed : Similaire à absolute sauf que l’élément est fixé en cas de scroll.
    • Sticky : L’élément est positionné intialement d’une façon relative puis lorsque le viewport défile jsqu’au point défini (ex : top:10px) alors l’élément devient fixé à 10px ;
      Exemple : 
      position : sticky ;
      top : -1px ;
      
JS
Fonctions intéressantes : 
const now = new Date() ;
const second = now.getSeconds();
const minute = now.getMinutes();
const hour = now.getHours();

setInterval((fonction à appeler), (délai d’interval))
Exemple : setInterval(setDate, 1000) ;
03 – Playing with CSS variables and JS
HTML
Tag Input : type : color → Fait apparaître une palette de couleurs
CSS
Déclarer des variables 
:root {
–base : #ffc600 ;
--spacing : 10 px ;
--blur : 10px ;
}
Propriétés intéressantes : 
Filter : blur(var(--blur)) ;
JS
Fonctions intéressantes : 
This correspond à un input ici.
document.documentElement.style.setProperty(`--${this.name}`, `${this.value}`) 
Cette fonction permet d’attribuer une valeur à une variable CSS sur l’ensemble du document.
document.documentElement.style.setProperty(`filter`, `blur(5px)`) 
Ici l’ensemble de la plage sera flouttée.
05 – Flex panels image gallery
HTML
CSS
Propriétés intéressantes : 
Ajouter un élément au début ou à la fin de l’élément sélectionné (ici p)
p:before, p:after {
content: "♥";
}


JS
Fonctions intéressantes : 
Utilisation de this
panels.forEach(panel => panel.addEventListener("transitionend", toggleActive));

function toggleActive(e){
if(e.propertyName.includes("flex")){
this.classList.toggle("active");
};
}


06 – Search function
HTML
CSS
Propriétés intéressantes : 

JS
Fonctions intéressantes : 
Récupérer un fichier JSON d’une URL (endpoint) :
fetch(endpoint)
.then(blob => blob.json())
.then(data => cities.push(...data));

/!\ Un .innerHTMLdoit impérativement être rempli par un string

0
HTML
CSS
Propriétés intéressantes : 

JS
Fonctions intéressantes : 
0
HTML
CSS
Propriétés intéressantes : 

JS
Fonctions intéressantes : 
0
HTML
CSS
Propriétés intéressantes : 

JS
Fonctions intéressantes : 
0
HTML
CSS
Propriétés intéressantes : 

JS
Fonctions intéressantes : 
0
HTML
CSS
Propriétés intéressantes : 

JS
Fonctions intéressantes : 
