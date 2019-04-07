---
title: 24 - Linting
created: '2019-04-03T12:35:27.454Z'
modified: '2019-04-03T13:03:12.403Z'
---

# 24 - Linting
Utiliser eslint dans un projet :
1. npm install eslint
1. ./node_modules/.bin/eslint –init
1. ./node_modules/.bin/eslint fichier.js
     
- Object creation : const item = {}
- Use computed property names when creating objects :
```
function getKey(k) {
  return `a key named ${k}`;
}
const obj = {
  id: 5,
  name: 'San Francisco',
  [getKey('enabled')]: true,
};
```
