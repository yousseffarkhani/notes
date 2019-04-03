---
title: 22 - Dynamic user interface
created: '2019-04-03T12:06:51.527Z'
modified: '2019-04-03T12:30:31.565Z'
tags: [CSS]
---

# 22 - Dynamic user interface
https://www.youtube.com/watch?v=f1WMjDx4snI&list=PLqGj3iMvMa4LvJ8VctoXnPI0dtE40wfid&index=2
Il existe 2 manière de réaliser des animations sur CSS:
## Transition property
```
.element {
	transition: [property à animer] [duration] [ease] [delay];
}
```
Qu’est ce qui peut-être animé? `Font-size`, `background-color`, `width`, `left`
Les meilleures animations : `position (translateX/Y)`, `scale`, `rotation`, `opacity`
Qu’est ce qui ne peut pas être animé? `Display`, `font-family`, `position`

Pour appliquer une transition property, il est interessant de créer un conteneur (trigger) afin d’y associer le hover pour éviter les aléas.
`Pointer-events: none` permet de désactiver les actions dû au pointer
## Animation property + keyframes
```
.box {
    width: 200px;
    height: 200px;
    background-color: pink;
    animation: myFrames 2s ease-in-out 1s 2; 
}

@keyframes myFrames {
    from {
        height:200px;
    }
    to {
        height: 600px;
        background-color: orange;
    }
} 

@keyframes myFrames {
    0%{
        height:200px;
    }
    100% {
        height: 600px;
        background-color: orange;
    }
} 
```
 animation-name: myframes
  animation-duration: 300ms
  animation-timing-function: ease-in-out
  animation-delay: 0s
  animation-iteration-count: 1
  animation-direction: normal
  animation-fill-mode: forwards
