https://reactjs.org/docs/thinking-in-react.html
Finir ce projet !!

# React
React est une librarie JS permettant de construire des interfaces utilisateurs.
# How react works under the hood (https://www.youtube.com/watch?v=TjnyFNxQ67Y&feature=emb_title)
Webpack : Concatène tous les fichiers en quelques fichiers
Webpack dev server : Garde les fichiers en mémoire et les délivre au navigateur
JSX : Permet de mettre du HTML dans JS. Babel transforme ce code en JS
Babel : Converti en JS de dernière génération. Du JSX vers du JS. Il suffit de créer un preset pour créer un convertisseur
npm start : charge l'appli en RAM (écrit en NodeJS)
npm build : crée un fichier static

# Virtual DOM
Au lieu de manipuler l'objet document du navigateur (ce qui peut être lent), React manipule les objets d'abord sur un virtual DOM. Un virtual DOM est une représentation du DOM entièrement en mémoire. Quand on écrit un component en réalité nous écrivons un component virtual que React va transformer en vrai component.

# JSX
Pour fonctionner, react utilise une syntaxe mixant html et JS appelée JSX.
const element = <h1>Hello, world!</h1>;
JSX produit des react ‘elements’.

JSX a été choisi car React a fait le choix de mixer la logique des évenements, des changements d’état et de comment la donnée est visualisée.
La séparation du code intervient au niveau des ‘components’.

# Render
Tous les components disposent d'une fonction render permettant de spécifier le HTML. 
JSX est une extension de React permettant d'écrire du javascript ressemblant à du HTML. Le JSX est ensuite convertit en JS pendant le runtime.
```JS
class HelloWorld extends React.Component {
  render(){
    return ( 
      React.createElement(
        'h1',
        {classname: 'large'},
        'Hello world'
      )
    );
  }
}

class HelloWorld extends React.Component {
  render(){
    return (
     <h1 classname='large'>Hello World</h1> 
    )
  }
}
```

# Rendering elements into the DOM
Pour afficher un élément, il faut utiliser ReactDOM de la manière suivante :
ReactDOM.render(<what>, <where>) (Ex : ReactDOM.render(element, document.getElementById('root'));)

# Utilisation de variable dans les elements
const name = 'Josh Perez';
const element = (<h1>Hello, {name}</h1>); // (l’ajout de parenthèses permet de sécuriser l’expression au cas où le ; serait oublié.
# Spécifier un attribut
Il existe 2 méthodes pour spécifier les attributs :
    • String litterals grâce à des quotes : const element = <div tabIndex="0"></div>;
    • JS expression grâce à des curly braces : const element = <img src={user.avatarUrl}></img>;
# Components and props
Les components permettent de séparer l’UI en des pièces indépendantes et réutilisables.
Un React.Component prend un paramètre appelé props et retourne une hiérarchie de vues exposée grâce à une méthode render.
## Functional and class components
### Functional Components
Lorsqu'un component ne dispose que de la fonction render, on peut le réécrire sous forme de fonction :
```JS
function Square(props) {
    return (
        <button 
            className="square" 
            onClick={()=>{props.onClick()}}
            >
                {props.value}
            </button>
    )
}
class Square extends React.Component { // Controlled Component because it receives values from Board Component
    render(){
        return (
            <button 
            className="square" 
            onClick={()=>{this.props.onClick()}} 
            >
                {this.props.value}
            </button>
        )
    }
}
```
### class components :
```JS
class Welcome extends React.Component {
  render() {
    return <h1>Hello, {this.props.name}</h1>;
  }
}
```
/!\Les components doivent démarrés avec une lettre en majuscule (c’est comme ça que react identifie qu’il s’agit d’un component).
Les props sont transmises comme des objets :
```JS
<Comment name="john" surname="doe"/> → function Comment(props){
//La propriété name est accessible via props.name → props : {name : ‘john’, surname :’doe’}
}
```
# State and lifecycle
Le state permet de garder en mémoire de la donnée par component.
Pour qu’un composant puisse se mettre à jour, il faut qu’il dispose d’un state.
Un state est similaire à une props mais est privé et appartient totalement au component.
Le local state est une fonctionnalité disponible pour les classes. C’est pourquoi nous allons déclarer notre composant via une classe.
## Déclarer un state
Le state est déclaré dans le constructeur de la classe :
```JS
class Clock extends React.Component {
  constructor(props) {
    super(props);
    this.state = {date: new Date()};
  }

  render() {
    return (
      <div>
        <h1>Hello, world!</h1>
        <h2>It is {this.state.date.toLocaleTimeString()}.</h2>
      </div>
    );
  }
}
```
## Modifier un state
Pour pouvoir modifier le state, il faut créer des méthodes au sein du composant. Pour cela, il y a 2 méthodes spéciales (lifecycle methods) :
    • componentDidMount(){}, cette méthode est lancée dès que le composant a été affiché la 1ère fois (1er appel à render())
    • componentWillUnmount(){}, cette méthode est lancée lorsque le composant est retiré du DOM
Pour modifier un state, il faut utiliser la méthode this.setState({date: new Date()})
A chaque fin d’appel à this.setState la méthode render() est invoquée.
/!\ L’unique endroit où utiliser this.state est dans le constructeur.
Pour modifier la valeur d’un state, il ne faut absolument pas utiliser this.props / this.state car les MAJ peuvent être asynchrones entrainant des comportements non souhaités.
```JS
this.setState({
  counter: this.state.counter + this.props.increment, // A ne pas faire
});
```
Pour réaliser ce genre de choses, il faut utiliser une 2ème forme de setState() acceptant une fonction au lieu d’un objet. Cette fonction recevra le state en 1er argument et les props en 2ème argument : 
```JS
this.setState((state, props) => ({
  counter: state.counter + props.increment
}));
```
React met automatiquement à jour les child components lorsque setState est appelé.

## Lifecycle methods
https://engineering.musefind.com/react-lifecycle-methods-how-and-when-to-use-them-2111a1b692b1
```JS
      render(){}
      
      static getDerivedStateFromProps(props, state) {} // not recommended - return the new updated state based upon the props
    
      getSnapshotBeforeUpdate() {} // Not recommended - creates a back up of the current way things are

      componentWillMount(){} // - DEPRECATED

      componentDidMount() {} // = You were just born - only runs once - UseCase : API Call

      componentDidUpdate(prevProps, prevState) { // Gets called everytime the component is updated / rerendered
          // Be careful to not update state in this method
          if (prevState.count) !== this.state.count {
              // Change state
          }
      }

      componentWillReceiveProps(nextProps) {} // Everytime components receives props this method runs - Deprecated (UNSAFE_componentWillReceiveProps)

      shouldComponentUpdate(nextProps, nextState) { // React sometimes rerenders a component even if it doesn't change
          // return true if want it to update
          // return false if not
      }

      componentWillUpdate(){} //DEPRECATED

      componentWillUnmount() {} // Do clean-up before component is killed - Ex: Remove Event listener declared in componentDidMount()
```
## Data change without mutation
```JS
let player = {score: 1, name: 'Jeff'};

let newPlayer = Object.assign({}, player, {score: 2})

let newPlayer = {...player, score: 2} // SPREAD SYNTAX equivalent to object.assign
```
### Bénéfices
- Garder les versions précédentes de l'état des components.
- La détection des changements est plus simple (ce qui permet à React de détecter plus facilement lorsqu'il doit rerender un component).


# Conditional rendering
Pour qu’un composant ne s’affiche qu’à une certaine condition, il est possible d’utiliser return null dans la définition du composant.
Même si le composant n’est pas affiché, les lifecycle methods sont lancées.
La règle à suivre est que le composant disposant de la propriété de state affectant la condition se charge de la condition dans le JSX.
Il existe plusieurs manières de réaliser cela :
- Ternary
```JS
render(){
  return (
    {isLoading ? <h1>Loading ...</h1>:<Component/>}
  )
}
```
- && (lorsqu'on souhaite ne rien afficher si la condition est KO)
```JS
render(){
  return (
    isLoading && <Component/>
  )
}
```

# Lists and keys
Pour créer des listes d’un tableau, il faut itérer à l’aide de .map().
/!\ Pour chaque élément de la liste, il faudra y associer une key à l’aide de l’attribut key={index}.
La key est importante car elle permet à React d’idendifier les éléments distinctement.
/!\ L’index est une mauvaise clé car l’ordre sera amené à changer.
Les clés ne sont pas transmises comme props, il faut déclarer explicitement la props si l’on souhaite la passer.
# Forms
Souvent les possibilités du HTML sont limitées et on souhaite prendre le contrôle du comportement (mettre les inputs en majuscule, ..) de certains objets à l’aide de JS. Cette technique s’appelle controlled components.
# Syntaxe
Les tags input / textarea / select (permet de créer un drop down-menu) ont un attribut ‘value’ en JSX qui permet de leur associer une valeur.
Toutes les modifications utilisateurs devront avoir un event handler associé permettant de modifier le state du composant.

# Handling events
```JS
  onChangeHandler = (event) => {
    const { name, value, type, checked } = event.target
    type === "checkbox" ? this.setState({ [name]: checked }):this.setState({ [name]: value })
  }
```
Appeler un callback suite à un event:
<button onClick={this.toggle}>
	{this.state.isToggleOn? 'ON':'OFF'}
</button>

/!\ il faut binder this avant de réaliser cet appel (à ajouter dans le constructeur) ou utiliser une arrow function : 
```JS
constructor(props){
	super(props);
	this.state = {
		isToggleOn: true
	};
	this.toggle = this.toggle.bind(this);
}
```
## Types de DOM Event
- onClick
- onChange sur une input
- onSubmit sur une form
https://reactjs.org/docs/events.html#supported-events
# Lifting state up
Souvent, plusieurs composants nécessitent de partager la même donnée (convertisseur de température en °C et °F).
Pour partager cette donnée, il faut que la donnée soit au niveau du composant-parent commun le plus proche des composants.
On dit que le composant parent devient la « source of truth ».
Il est également possible d’appeler une fonction d’un parent en passant cette fonction en props et en l’appelant via this.props.’nom de la fonction’.
/!\ Si quelque chose peut être dérivé / calculé d’une props ou d’un state alors cet élément ne devrait pas faire partie du state.
# Composition vs inheritance
{props.children} permet de passer en prop tous les composants à l’intérieur du composant faisant appel à {prop.children}.
Il est également possible de transmettre des composants comme des props via les attributs (ex : left={<Contacts />} )

# PropTypes 
Permet de valider les types des props, le nombre d'élément. Proptype permet également de définir des valeurs par défaut pour les props.
```JS
import PropTypes from 'prop-types';

class Greeting extends React.Component {
  render() {
    return (
      <h1>Bonjour, {this.props.name}</h1>
    );
  }
}

Greeting.propTypes = {
  name: PropTypes.string
};
```
TypeScript peut être utilisé maintenant
# Router
## Link, NavLink
```JS
import { Link, NavLink } from 'react-router-dom'

const NavBar = () => (
  <nav>
    <ul>
      <Link to="/">Home</Link>
      <NavLink to="/about">About</Link> // NavLink ajoute la class "active" lorsque ce lien est actif
      <Link to="/contact">Contact</Link>
    </ul>
  </nav>
)
```
## Routes
```JS
ìmport { BrowserRouter, Route } from 'react-router-dom'

const App = () => (
  <BrowserRouter>
    <NavBar/> // Il faut créer un moyen de naviguer entre les différentes routes
    <Route exact path='/' component={Home} /> //exact permet de n'afficher le component lorsque le path est exactement le même
    <Route path='/about' component={About} />
    <Route path='/contact' component={Contact} />
  </BrowserRouter>
)
```
## Route Props
Lorsque le component est actif alors une props est transmise à ce dernier.
Cette props contient des informations sur les routers.
```JS
const Home = (props) => {
  console.log(props)
  return (<h1>Home</h1>)
}

export default Home
```
Il est également possible de mettre en place des actions :
### Redirects
```JS
  const Home = (props) => {
      props.history.push("/about")

      return (<h1>Home</h1>)
  }
```
## Route parameters
Pour récupérer un paramètre passé dans l'URL :
```JS
 <Route path="/:post_id" component={Post} />
```
```JS
  <Link to={"/" + post.id} key={post.id}>
      <h1 >{post.title}</h1>
  </Link>
```
```JS
  componentDidMount() {
    let id = this.props.match.params.post_id
    axios.get("https://jsonplaceholder.typicode.com/posts/" + id)
        .then(response => this.setState({ post: response.data }));
  }
```
## Switch
Pour n'afficher qu'un seul component matchant une route, il faut utiliser le component Switch.
Dans cet exemple, si on va sur "/contact", seul la page Contact sera affichée.
```JS
  <Switch>
      <Route path="/contact" component={Contact} />
      <Route path="/:post_id" component={Post} />
    </Switch>
```
# Higher Order Components
Ce sont des fonctions qui enveloppent des components pour leur ajouter des fonctionnalités.
## WithRouter
Permet d'ajouter les props liées au routing à un component :
```JS
export default withRouter(NavBar)
```
## Custom
```JS
const Rainbow = (WrappedComponent) => {
    const colours = ["pink", "red", "yellow"]
    const randomColour = colours[Math.floor(Math.random() * colours.length)]
    const className = randomColour + "-text"
    return (props) => (
        <div className={className}>
            <WrappedComponent {...props} />
        </div>
    )
}

export default Rainbow
```
# CSS
Il existe 2 manières d'implémenter du CSS :
- Créer un fichier CSS pour chaque composant.
Le CSS sera actif lorsque le composant sera actif.
`import `./Ninjas.css`` dans un fichier de component
- Utiliser index.css : Le CSS sera actif tout le temps
# Importer une image dans un component
```JS
import Pokeball from './pokeball.png'

...

  <img src={Pokeball} alt="A pokeball"/>
```
