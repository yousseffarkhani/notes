---
title: 21 - Forms
created: '2019-04-03T12:05:58.527Z'
modified: '2019-04-03T12:33:29.850Z'
tags: [HTML]
---

# 21 - Forms
## Built-in Form validation (HTML)
https://developer.mozilla.org/en-US/docs/Learn/HTML/Forms/Form_validation
```
<form>
  <label for="choose">Would you prefer a banana or cherry?</label>
  <input id="choose" name="i_like" required pattern=‘banana | cherry’>
  <button>Submit</button>
</form>
```
- required :force l’utilisateur à remplir le champs avant de le soumettre.
- Pattern : force l’utilisateur à utiliser un pattern (regexp)

Ces attributs vont de pair avec les attributs CSS ci-dessous :
```
input:invalid {
  border: 2px dashed red;
}

input:valid {
  border: 2px solid black;
}
```
Il est possible de customiser les messages d’erreur à l’aide de JS :
```
var email = document.getElementById("mail");

email.addEventListener("input", function (event) {
  if (email.validity.typeMismatch) {
    email.setCustomValidity("I expect an e-mail, darling!");
  } else {
    email.setCustomValidity("");
  }
});
```
Pour plus de contrôles via JS : voir constraint validation API 
