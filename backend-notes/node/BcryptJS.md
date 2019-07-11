---
title: BcryptJS
created: '2019-04-04T13:16:51.280Z'
modified: '2019-04-04T15:19:45.565Z'
---

# BcryptJS
Bcrypt est un middleware permettant de générer des mots de passe.
## Commandes
### Installation
`npm install bcryptjs`
### Utilisation
```
router.post('/register', (req, res) => {
    const newUser = new User({
        name: req.body.name,
        email: req.body.email,
        username: req.body.username,
        password: req.body.password,
    });

    bcrypt.genSalt(10, function(err, salt){
        bcrypt.hash(newUser.password, salt, function(err, hash){
            if(err) console.log(err);
            newUser.password = hash;
            newUser.save() //Sauvegarde du newUser ici sinon le hash est perdu
                .then(result => {
                    console.log(result);
                    res.redirect('/users/login');
                })
                .catch(err => {
                    console.log(err);
                    res.status(500).json({err: "there was an error"});
                });
        });
    })    
});
```
```
bcrypt.compare(password, user.password, function(err, isMatch){
    if(err) throw err;
    if(isMatch){
        return done(null, user);
    } else {
        return done(null, false, {message: "Wrong password"});
    }
});
```
