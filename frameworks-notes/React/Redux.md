# Redux hooks
https://medium.com/swlh/clean-up-redux-code-with-react-redux-hooks-71587cfcf87a

# Actions
Les actions correspondent à des instructions envoyées au store.
Signature d'une action 
```TS
interface Action {
  type: string;
  payload?: any;
}
```

# Reducers
Un reducer est une fonction pure prenant le state de l'application (le store) et l'action à dispatcher.
```TS
function reducer(state = {}, action) {
  switch (action.type) {
    case 'ADD_TODO': {
      return {
        ...state,
        // we spread the existing todos array into a new array
        // and then add our new todo on the end
        todos: [...state.todos, { label: 'Eat pizza,', complete: false }],
      };
    }
  }

  return state;
}
```
# Description
Redux est une base de données centrale pour notre application.
Tous les composants peuvent accéder aux données de ce store.
Cela permet de simplifier la gestion du state.
Il s'agit d'un Javascript Object représentant l'état de notre application.

Les components ne peuvent pas mettre à jour les données

Si un component veut accéder à une donnée, il faut subscribe aux changements de cette donnée. Redux enverra la donnée sous forme de props.
Si un component veut modifier une donnée, il faut dispatch une action (ex: ajouter un post) et un payload (ex: le nouveau post). Ensuite, l'action est transmise à un reducer qui se chargera de mettre à jour la donnée dans Redux.

# Reducer
Fonctions permettant de modifier l'état de l'application.

# Comment utiliser Redux dans un projet
## Installation
```bash
npm i redux react-redux
```
## Setup du projet avec Redux
### Ajout du store à l'App component
App.js
```JS
const store = createStore(rootReducer)

ReactDOM.render(<Provider store={store} ><App /></Provider>, document.getElementById('root'));
```
### Initialisation du store et mise en place d'actions
reducers/rootReducer.js
```JS
const initState = {
    posts: [
        { id: '1', title: 'Squirtle Laid an Egg' },
    ]
}

const rootReducer = (state = initState, action) => {
    switch (action.type) {
        case "DELETE_POST":
            const newPosts = state.posts.filter(post => post.id !== action.id)
            return {
                ...state,
                posts: newPosts
            }
        default:
            break;
    }

    return state
}

export default rootReducer
```
### Lecture d'une valeur dans le store
#### Standard
Home.js
```JS
const mapStateToProps = (state) => {
    return {
        posts: state.posts
    }
}

export default connect(mapStateToProps)(Home) // connect returns a Higher Order Component
```
#### Avec une valeur issue du component
Post.js
```JS
const mapStateToProps = (state, ownProps) => { // ownProps contient les props du component
    const id = ownProps.match.params.post_id
    return {
        post: state.posts.find(post => post.id === id)
    }
}

export default connect(mapStateToProps)(Post)
```

### Dispatch d'une action
actions/postAction.js
```JS
export const deletePost = (id) => {
    return {
        type: "DELETE_POST",
        id // équivalent à id : id
    }
}
```
Post.js
```JS
handleClick = () => {
        this.props.deletePost(this.props.post.id)
        this.props.history.push("/")
    }

const mapDispatchToProps = (dispatch) => {
    return {
        deletePost: (id) => { dispatch(deletePost(id)) }
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Post)
```
