# Gator CLI

A multi-player command line tool for aggregating RSS feeds and viewing the posts.
This project is part of the Boot.dev program.

## Prerequisites

- **Go** (version 1.18+ recommended): [Install Go](https://golang.org/doc/install)
- **PostgreSQL**: [Install PostgreSQL](https://www.postgresql.org/download/)

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

```bash
go install ...
```

## Configuration

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{b
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Database Setup

Database schema migrations are managed with [Goose](https://github.com/pressly/goose).  
Migration files are located in [sql/schema/](sql/schema/). To apply migrations:

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir sql/schema postgres "postgres://username:password@localhost:5432/gatordb?sslmode=disable" up
```

SQL queries are converted to Go functions using [sqlc](https://sqlc.dev/).  
See [sqlc.yaml](sqlc.yaml) for configuration and generated code in [internal/database/](internal/database/).

## Usage

Create a new user:

```bash
gator register <name>
```

Add a feed:

```bash
gator addfeed <url>
```

Start the aggregator:

```bash
gator agg 30s
```

View the posts:

```bash
gator browse [limit]
```

There are a few other commands you'll need as well:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database
