# Description
Les hooks permettent de :
- Utiliser que des functional components à la place des classes
- Améliorent la lecture et l'organisation des composants

# useState
```JS
import React, {useState} from "react"

const App = () => {
    const [count, setCount] = useState(1)
    
    const doubleCount = () => setCount(prevCount => prevCount*2)
    const halve = () => setCount(prevCount => Math.round(prevCount/2))
    
    return (
        <div>
            <h1>{count}</h1>
            <button onClick={doubleCount}>Double</button>
            <button onClick={halve}>Half</button>
        </div>
    )
}
```

# useEffect
Produce side effects to our components 
Side effects ( =/= managing state or displaying content) ? Networks requests, manual DOM manipulation, event listeners or timeouts and intervals
Replacement for :

- ComponentDidUpdate
```JS
useEffect(() => {
    setColor(randomcolor())
}, [count]) // It will trigger use Effect when a variable in this array is modified
```
- ComponentDidMount
```JS
useEffect(() => {
    setInterval(() => {
        // setCount(prevCount => prevCount + 1)
    }, 1000)
}, []) // An empty array is the equivalent of componentDidMount
```
- ComponentWillUnmount
    socket subscription, document.addEventListener
```JS
    useEffect(() => {
        const intervalId = setInterval(() => {
            setCount(prevCount => prevCount + 1)
        }, 1000)
        return () => clearInterval(intervalId) // Use effect keeps this function and triggers it when component is unmounted
    }, [])
```
# useRef
Permet de cibler un node du DOM. 

# useLayoutEffect
Ce hook est déclenché une fois que tous les changements du DOM ont eu lieu.
Ce hook est à utiliser en dernier recours après useEffect.
Pemret de mesurer par exemple la taille d'un élément du DOM.
