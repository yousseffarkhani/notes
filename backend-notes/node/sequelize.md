# Knex vs Sequelize
Sequelize est un ORM permettant d'abstraire le SQL en language nodeJS.

# Migrations
Fichier JS décrivant le modèle de données

tables, colums, relationships

Cela permet en conjonction avec un outil de versionning de garder la trace des format des tables.
Avec ce système, on peut facilement revenir en arrière (back-up system).

# Commande
npx sequelize migration:generate --name episodes
npx sequelize db:migrate

New migration file or undo last migrations and run new migrations

npx sequelize migration:create --name name-the-migration
npx sequelize db:migrate --env local // Appelle la fonction up
npx sequelize db:migrate:undo // Appelle la fonction down (deletes all data related to the table)

The following is an example of a migration that performs two changes in the database, using a transaction to ensure that all instructions are successfully executed or rolled back in case of failure:

# Types d'objets

Auto-increment : 
Exemple : 
```JS
queryInterface.createTable('Episodes', {
    id: {
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
        type: Sequelize.INTEGER,
    },
    email: {
        type: Sequelize.STRING,
        unique: true
    },
    password: {
        type: Sequelize.STRING,
    }
})
```

# Associations
# Options
required: si cette valeur est à true, alors toutes les données
