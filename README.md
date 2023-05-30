# Pioneer 2023

A suite of tools for engineering management

## 🛠️ Development
You can run the server and client simultaneously for development:
```bash
npm run watch
```

### Migrations
Migrations are run using [the golang-migrate/migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

For convenience, migrations can be run by cd'ing into the server directory and running the appropriate NPM script, e.g.
```bash
cd server
npm run migrate-up
```

You can create a new migration using the migrate tool directly e.g.
```bash
cd server
migrate create -ext sql -dir db/migrations -seq add_created_update_columns
```