---
title: Passport
created: '2019-04-04T12:20:42.945Z'
modified: '2019-04-04T20:50:13.629Z'
---

# Passport
## Introduction
Passport est une librairie permettant de mettre en place des stratégies d'authentification.
Types de stratégies :
- Local strategy : Username + PW dans une BDD
- FB/Twitter strategy
- JWT : 
##Commandes
### Installation :
1. `npm install passport passport-local bcryptjs`
 bcryptjs permet de gérer les password

 3 éléments doivent être configurés pour que passport fonctionne :
 - authentication strategies
 - application middleware
 - sessions

 ### authentication strategies
 Les stratégies peuvent allées d'une vérification username/password à la délégation de l'authentification grâce à OAuth ou OpenID.


### Access control
const ensureAuthenticated = (req, res, next) => {
    if(req.isAuthenticated()){
        return next();
    } else {
        res.redirect('/users/login');
    }
}

router.get('/add', ensureAuthenticated, (req, res) => {
    res.render('add_article', {
        title: 'Add article'
    });
})

 If anyone has problem with getting only "missing credentials" message- in "login.pug" you need to have the "name" parameters of input fields set as follows: name= "username" for username input name= "password" for password input By default passport is searching for values of fields with these "name" attributes in html, if you put different name there, then passport can't see it's value. Another approach is to modify names of the fields that passport is looking for by adding the following object as a first parameter in "passport.js": passport.use( new LocalStrategy( { usernameField: 'username', //"username" is the default name for html form input passwordField: 'pass1' }, //in my html name for password input has is equal to "pass1" instead of the default "password" (username, password, done) => {... ...

  David K. Dundas
David K. Dundas
il y a 8 mois
I wanted to get this to work for email and not username. To get the email to work you can update your local strategy in ./config/passport.js to this: 

*note this is assuming that the email field in your login for is actually called "email"
passport.use(new LocalStrategy({ // or whatever you want to use
        usernameField: 'email',    // define the parameter in req.body that passport can use as username and password
        passwordField: 'password'
      },(email , password, done )=>{
        // MATCH USERNAME 

        let query = {email : email }
        console.log("HERE IS THE QUERY: ", query)
        Users.findOne(query, (err, user)=>{

            if (err) throw err; 
            if (!user){
                return done(null, false , {message: 'Invalid Login.'})
            }
            bcrypt.compare(password, user.password, (err, isMatch )=>{
                if (err) throw err;
                if (isMatch){
                    return done(null, user )
                }else{
                    return done(null, false , {message: 'Invalid Login.'})
                }

            })
 
        })
    }))﻿ 
