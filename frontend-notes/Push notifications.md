# Push notifications
## Introduction
Il est possible de donner la possibilité aux applications web de recevoir des messages qui leur sont poussés depuis un serveur (que cette application soit ouverte ou non sur le navigateur)
Pour que l'application web puisse recevoir un push, elle doit avoir un service worker actif.

### Côté serveur
1. Initialiser les VAPID keys en lançant la commande `./node_modules/.bin/web-push generate-vapid-keys` 
Ces clés permettent d'identifier qui envoi les push notification
```javascript
const publicVapidKey = "BCOVym1EuiOepiodyR-cDiehx1rmSLO5";
const privateVapidKey = "Hfuffgbn0kMru92fTtMRdK96n3R3OBrpmxoe3nVOVzQ";

webpush.setVapidDetails('mailto:test@test.com', publicVapidKey, privateVapidKey);
```
2. Créer une route permettant la souscription au push
```javascript
app.post('/subscribe', (req, res) => {
    // Get push subscription object
    const subscription = req.body;

    // Send 201 - Ressource created
    res.status(201).json({});

    // Create Payload
    const payload = JSON.stringify({
        title : "Push test" // Il est possible de rajouter énormément d'option voir https://tests.peter.sh/notification-generator/
    });

    // Pass object into sendNotification
    webpush.sendNotification(subscription, payload)
        .catch(console.error);
});
```
### Côté client

Cela se déroule en 3 étapes : souscription, catch de l'event, push notification
1. Souscrire 
PushManager.subscribe()
2. Catch de l'event
3. Push notification
ServiceWorkerRegistration.showNotification()

(index.js)
```javascript
// Register SW, register Push, Send push
async function send(){
    // Register SW
    const register = await navigator.serviceWorker.register("/service-worker.js", {
        scope:"/" // URL pour lesquelles le SW fonctionnera
    });

    // Register Push
    const subcription = await register.pushManager.subscribe({
        userVisibleOnly: true,
        applicationServerKey: urlBase64ToUint8Array(publicVapidKey)
    });

    // Send push notification
    await fetch("/subscribe", {
        method:"post",
        headers: {
            "content-type": "application/json"
        },
        body: JSON.stringify(subcription)
    })

}
```
(service-worker.js)
```javascript
self.addEventListener("push", event => {
    const data = event.data.json();
    self.registration.showNotification(data.title, {
        body: 'Notified by le phys'
        
    });
})
```
