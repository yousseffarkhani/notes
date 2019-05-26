# Introduction

# Processus

1. Open a database
2. Create an object store in the database
3. On the success event, conduct your transactions
4. Close the transaction

# Implémentation

1. Vérifier si l'indexedDB existe

   if(self.indexedDB){
   console.log('IndexedDB is supported');
   }

2. Code

```javascript
if (self.indexedDB) {
  console.log("IndexedDB is supported");
}

const products = [
  { id: 1, name: "Red Men T-Shirt", price: "$3.99" },
  { id: 2, name: "Pink Women Shorts", price: "$5.99" },
  { id: 3, name: "Nike white Shoes", price: "$300" }
];
/* Open a database */
const request = self.indexedDB.open("example_DB", 1); // Nom de la BDD et la version
let db, tx, store;

request.onupgradeneeded = function(event) {
  // Event fired every time a new DB is created ensuite onsuccess est déclenché. Si la base de données est détectée, onsuccess sera déclenché directement
  const db = request.result;
  /* Create an object store in the database */
  const store = db.createObjectStore("products", { keyPath: "id" }); //keyPath permet de définir le nom de la clé (Il est également possible d'auto incrémenter cette valeur)
  store.createIndex("products_name", "name", { unique: true }); // Il est possible de créer un index afin de permettre de rechercher via une autre valeur que l'id
};

request.onerror = function(event) {
  // Gestion des erreurs
  console.log(request.error);
};

request.onsuccess = function(event) {
  /* On the success event, conduct your transactions */
  db = request.result;
  tx = db.transaction("products", "readwrite");
  store = tx.objectStore("products");
  index = store.index("products_name");

  db.onerror = function(event) {
    console.log("ERROR", event.target.error);
  };

  products.map(product => store.add(product));
  let product1 = store.get(1);
  let product2 = index.get("Red Men T-Shirt");
  product1.onsuccess = function() {
    console.log(product1.result);
    console.log(product1.result.name);
  };

  product2.onsuccess(product2.result);
  /* Close the transaction */
  tx.oncomplete = function() {
    db.close();
  };
};
```

# Méthodes utiles

store.clear()
