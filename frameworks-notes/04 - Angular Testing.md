---
title: 04 - Angular Testing
created: '2019-04-03T13:03:43.573Z'
modified: '2019-04-03T13:32:30.245Z'
tags: [angular]
---

# 04 - Angular Testing

## Angular :
Introduction
Karma : Moteur de tests (test runner) → Creating html file, Opening the browser and running tests.
jasmine : Langage d’assertion permettant de décrire le test. Souvent considéré comme un testing framework car il dispose de nombreuses fonctionnalités (describe, it, expect) pour tester.
Commandes
ng test
Syntaxe :
describe(‘nom de la suite de tests’, () => {
	it(‘test1’, () => {
	})
})
Méthodes jasmine utiles :
Jasmine contient de nombreux matchers de base.
`expect(result).toContain(‘john’)` → vérifie la présence de la variable dans le result.
.toBeTruthy() / .toBeFalsy()
`BeforeEach()` → Permet d’exécuter du code avant chaque fonction (pour initialiser un component par exemple)
`AfterEach` (delete des files)/`beforeAll` (executé une fois avant les tests, aussi appelé setup)/afterAll (aussi appelé teardown)
/!\ `beforeEach & co` sont à utiliser avec parcimonie car ce sont des indicateurs de code non modulable.
Mocking
Dans Angular, il existe plusieurs maninères de faire des mocks :
- Recréer une fausse classe si la classe d’origine ne contient pas énormément d’éléments
- Recréer et extends la classe que nous souhaitons mocker, pour n’avoir à modifier que la méthode/attributs voulu pour le test et éviter de devoir réécrire l’ensemble des méthodes et attributs
- Utiliser un spy qui permet de prendre le contrôle de la classe
      `spyOn(shipService,'getShip').and.returnValue('test');`
Testbed
Les testbed permettent de créer un environnement de test similaire à ce que ferait ngModule en important les dépendances externes.
Cela permet de tester les components avec leur templating.
Fixture : Test environnement
debugElement : render HTML
Stuck : équivalent d’un mock dans angular
let serviceStub : any
beforeEach(() => {
serivceStub = {
getContent : () => ‘blablabla’
}
TestBed.configureTestingModule({
providers : [

## Définitions :
Types de tests
Unit testing : Consiste à écrire des tests pour des petites parties de l’application en isolation (sans connexion à une DB ou une API externe ou le file system).
Dans Angular, cela consiste à tester un component sans aucune ressource (pas même le template). Si le component communique avec un service, nous allons lui fournir un fakeService.

Integration testing : Consiste à tester un composant avec les ressources externes.
Dans Angular, cela consiste à tester un component avec son template. Si le component communique avec un service, nous allons lui fournir un fakeService.

End-to-end testing (functional testing) : Consiste à tester l’application en entier. Cela consiste à simuler l e comportement d’un utilisateur.

Fonction pure : C’est une fonction qui pour les mêmes arguments retournent la même réponse. Elle ne dépend d’aucun code externe à la fonction. (un console.log est un side effect donc une fonction contenant cette ligne n’est pas une fonction pure)

