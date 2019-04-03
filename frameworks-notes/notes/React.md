---
title: React
created: '2019-04-03T13:32:48.800Z'
modified: '2019-04-03T13:33:00.310Z'
---

# React
https://reactjs.org/docs/thinking-in-react.html
Finir ce projet !!
Créer un projet react
Saisir dans le terminal create-react-app ‘nom de l’app’. Cela va générer tout un project react.
Il existe également une possibilité d’ajouter uniquement des composants react à un site non-react.
Overview
Pour fonctionner, react utilise une syntaxe mixant html et JS appelée JSX.
const element = <h1>Hello, world!</h1>;
JSX produit des react ‘elements’.

JSX a été choisi car React a fait le choix de mixer la logique des évenements, des changements d’état et de comment la donnée est visualisée.
La séparation du code intervient au niveau des ‘components’.
Utilisation de variable dans les elements
const name = 'Josh Perez';
const element = (<h1>Hello, {name}</h1>); //(l’ajout de parenthèses permet de sécuriser l’expression au cas où le ; serait oublié.
Spécifier un attribut
Il existe 2 méthodes poru spécifier les attributs :
    • Strig litterals grâce à des quotes : const element = <div tabIndex="0"></div>;
    • JS expression grâce à des curly braces : const element = <img src={user.avatarUrl}></img>;
Spécifier les enfants

Si un tag est empty (img) alors il est possible de le refermer avec / directement : const element = <img src={user.avatarUrl} />;

Rendering elements into the DOM
Pour afficher un élément, il faut utiliser ReactDOM de la manière suivante :
ReactDOM.render(element, document.getElementById('root'));
Components and props
Les components permettent de séparer l’UI en des pièces indépendantes et réutilisables.
Function and class components
Function components :
function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}
class components :
class Welcome extends React.Component {
  render() {
    return <h1>Hello, {this.props.name}</h1>;
  }
}
/!\Les components doivent démarré avec une lettre en majuscule (c’est comme ça que react identifie qu’il s’agit d’un component).
Les props sont transmises comme des objets :
<Comment name= ‘john’ surname=’doe’/> → function Comment(props){
//La propriété name est accessible via props.name → props : {name : ‘john’, surname :’doe’}
}
State andlifecycle
Pour qu’un composant puisse se mettre à jour, il faut qu’il dispose d’un state.
Un state est similaire à une props mais est privé et appartient totalement au component.
Le local state est une fonctionnalité disponible pour les classes. C’est pourquoi nous allons déclarer notre composant via une classe.
Déclarer un state
Le state est déclaré dans le constructeur de la classe :
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
Modifier un state
Pour pouvoir modifier le state, il faut créer des méthodes au sein du composant. Pour cela, il y a 2 méthodes spéciales (lifecycle methods) :
    • componentDidMount(){}, cette méthode est lancée dès que le composant a été affiché la 1ère fois (1er appel à render())
    • componentWillUnmount(){}, cette méthode est lancée lorsque le composant est retiré du DOM
Pour modifier un state, il faut utiliser la méthode this.setState({date: new Date()})
A chaque fin d’appel à this.setState la méthode render() est invoquée.
/!\ L’unique endroit où utiliser this.state est dans le constructeur.
Pour modifier la valeur d’un state, il ne faut absolument pas utiliser this.props / this.state car les MAJ peuvent être asynchrones entrainant des comportements non souhaités.
this.setState({
  counter: this.state.counter + this.props.increment, // A ne pas faire
});
Pour réaliser ce genre de choses, il faut utiliser une 2ème forme de setState() acceptant une fonction au lieu d’un objet. Cette fonction recevra le state en 1er argument et les props en 2ème argument : 
this.setState((state, props) => ({
  counter: state.counter + props.increment
}));




Handling events
Appeler un callback suite à un event:
<button onClick={this.toggle}>
	{this.state.isToggleOn? 'ON':'OFF'}
</button>

/!\ il faut binder this avant de réaliser cet appel (à ajouter dans le constructeur) : 
constructor(props){
	super(props);
	this.state = {
		isToggleOn: true
	};
	this.toggle = this.toggle.bind(this);
}

Conditional rendering 
Pour qu’un composant ne s’affiche qu’à une certaine condition, il est possible d’utiliser return null dans la définition du composant.
Même si le composant n’est pas affiché, les lifecycle methods sont lancées.
Lists and keys
Pour créer des listes d’un tableau, il faut itérer à l’aide de .map().
/!\ Pour chaque élément de la liste, il faudra y associer une key à l’aide de l’attribut key={index}.
La key est importante car elle permet à React d’idendifier les éléments distinctement.
/!\ L’index est une mauvaise clé car l’ordre sera amené à changer.
Les clés ne sont pas transmises comme props, il faut déclarer explicitement la props si l’on souhaite la passer.
Forms
Souvent les possibilités du HTML sont limitées et on souhaite prendre le contrôle du comportement (mettre les inputs en majuscule, ..) de certains objets à l’aide de JS. Cette technique s’appelle controlled components.
Syntaxe
Les tags input / textarea / select (permet de créer un drop down-menu) ont un attribut ‘value’ en JSX qui permet de leur associer une valeur.
Toutes les modifications utilisateurs devront avoir un event handler associé permettant de modifier le state du composant.
Lifting state up
Souvent, plusieurs composants nécessitent de partager la même donnée (convertisseur de température en °C et °F).
Pour partager cette donnée, il faut que la donnée soit au niveau du composant-parent commun le plus proche des composants.
On dit que le composant parent devient la « source of truth ».
Il est également possible d’appeler une fonction d’un parent en passant cette fonction en props et en l’appelant via this.props.’nom de la fonction’.
/!\ Si quelque chose peut être dérivé / calculé d’une props ou d’un state alors cet élément ne devrait pas faire partie du state.
Composition vs inheritance
{props.children} permet de passer en prop tous les composants à l’intérieur du composant faisant appel à {prop.children}.
Il est également possible de transmettre des composants comme des props via les attributs (ex : left={<Contacts />} )
Thinking in components (react) – Méthodologie de construction d’une app avec react
