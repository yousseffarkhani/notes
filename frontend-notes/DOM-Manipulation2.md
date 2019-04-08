# DOM Manipulation
Il est possible d'activer des events directement depuis le code JS.
Pour simuler un click sur un bouton :
```javascript
const accept = card.querySelector('.bt-invite-accept');  //Sélectionne les boutons
accept.click();
```
Exemple de code complet :
```javascript
[...document.querySelectorAll('.invite-card')].forEach(card => {
  const headline = card.querySelector('.headline').textContent;
  const accept = card.querySelector('.bt-invite-accept');
  const decline = card.querySelector('.bt-invite-decline');
  const name = card.querySelector('.name').textContent;

  if(headline.match(/recruit/gi)) {
    console.log(`Nah. ${name} looks a little fishy to me. 🚷🚷🚷`);
    decline.click();
  } else {
    console.log(`${name} looks cool. Let them in! 😎`);
    accept.click();
  }
});
```