# Golang GraphQL

Golang + GraphQL sample with Prisma

### Requirements

-   [Docker](https://www.docker.com/get-started)
-   [Prisma](https://www.prisma.io/docs/get-started/01-setting-up-prisma-new-database-JAVASCRIPT-a002/)

### Setup

1. Run the database

    `docker-compose up`

2. Migrate the schema

    `go run github.com/steebchen/prisma-client-go db push`

3. Test-run inserting data into the schema

    `go run main.go`

    NOTE: If you're using VSCode, you can simply use the `Run` configuration in `Run and Debug`, already prepared in `.vscode/launch.json`.
