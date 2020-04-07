# Source
https://github.com/goldbergyoni/javascript-testing-best-practices
# Naming 
1. Qu'est ce qui est testé ? Nom du service et de la méthode.
Ex: La méthode ProductsService.addNewProduct
2. Sous quelle condition / Scénario ?
Ex: Le prix n'est pas passé à la méthode
3. Quel est le résultat attendu ? 
Ex: Le produit ne devrait pas être approuvé
```JS
//1. unit under test
describe('Products Service', function() {
  describe('Add new product', function() {
    //2. scenario and 3. expectation
    it('When no price is specified, then the product status is pending approval', ()=> {
      const newProduct = new ProductService().add(...);
      expect(newProduct.status).to.equal('pendingApproval');
    });
  });
});
```

# Structuration AAA patern
Arrange, Act & Assert

- Arrange: Setup du test
- Act: Appelle de la méthode
- Assert: Vérification du retour

# Guidelines
Test the behavior and not implementation.
Toujours se demander si la fonctionnalité testée pourrait être présente dans une liste d'exigences.
Utiliser de vraies datas (Utiliser la librairie Faker pour générer des pseudo-réel datas).
# Test doubles
- Dummy: Objets utilisés pour remplir la liste des paramètres mais n'ayant pas d'autre utilité
- Mocks: 
- Stubs: 
- Spies: Permet de vérifier qu'une méthode est appelée 
Syntaxe: sinon.spy(Emailer.prototype, "sendEmail");
# Catching errors
```JS
it.only("When no product name, it throws error 400", async() => {
  await expect(addNewProduct({})).to.eventually.throw(AppError).with.property('code', "InvalidInput");
});
```
