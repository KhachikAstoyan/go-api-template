### A simple template for myself for scaffolding go services

I created this simple template thing so that it's easier to create new Go HTTP services from scratch, because often times I end up repeating myself.

#### What does this template use?

- Go (obviously)
- Chi (for routing)
- sqlc (compiling queries to go code)
- PostgreSQL
- Some other minor libs for convenience

#### Directory Structure

- `core` contains the core things, like configuration and global App struct
- `db` contains DDL statements, queries, and the generated repository
- `dtos` are self explanatory
- `http` contains http handlers (controllers)
- `services` have different services as singletons, that talk to the database through the generated repository
