npm i -D @types/jest
npm i -D jest

```JS
test("User fetched name should be Leanne Graham", () => {
    expect.assertions(1)
    return functions.axiosFetchUsers()
        .then(data => {
            expect(data.name).toEqual('Leanne Graham')
        })
});

test("User fetched name should be Leanne Graham", async () => {
    expect.assertions(1)
    const data = await functions.axiosFetchUsers()
    expect(data.name).toEqual('Leanne Graham')

})

/*
test("Adds 2+ 2 to equal 4", () => {
    expect(functions.add(2, 2)).toBe(4);
});
toBe
    .not.toBe
toBeNull
toEqual
toBeLessThan
toContain // To find an element in the array
*/
```

# Watch mode
`jest --watchAll/watch`
# Sinon
Cette librairie permet de créer des spy, des stubs et des mocks.

sinon.spy("nom de l'objet disposant de la méthode", "méthode")
# Mocking
- expect().toHaveBeenCalled() : Vérifie que la fonction spy a été appelée.
- expect().toHaveBeenCalledTimes(): Compte le nombre de fois ou la spy fonction a été appelée
- expect().toHaveBeenCalledWith(): Vérifie que la fonction a été appelée avec le paramètre défini
- expect().toHaveBeenCalledWith(): Vérifie que la fonction a été appelée en dernier avec le paramètre défini
# Stubs
Les stubs permettent de 
# Spy on packages
On peut créer une fonction spy en précisant la fonction et le package a spy
```JS
const mathjs = require("mathjs")

test(`The mathjs log function`, () => {
  const spy = jest.spyOn(mathjs, 'log')
  const result = mathjs.log(10000, 10)

  expect(mathjs.log).toHaveBeenCalled()
  expect(mathjs.log).toHaveBeenCalledWith(10000, 10)
})
```

Ne pas oublier de clean après chaque test

# Done
Permet de terminer l'execution du test
