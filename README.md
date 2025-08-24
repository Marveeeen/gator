# Gator CLI

Gator is a command-line tool for managing RSS feeds and users, backed by a PostgreSQL database.  
This project is part of the Boot.dev program.

## Prerequisites

- **Go** (version 1.18+ recommended): [Install Go](https://golang.org/doc/install)
- **PostgreSQL**: [Install PostgreSQL](https://www.postgresql.org/download/)

## Installation

Clone the repository and install the CLI using `go install`:

```sh
git clone https://github.com/marveeeen/gator.git
cd gator
go install
```

This will build and install the `gator` CLI to your `$GOPATH/bin`.

## Configuration

Before running the CLI, you need to set up your configuration file. The config file is stored as `.gatorconfig.json` in your home directory.

You can create it manually or let the CLI generate it. The file should look like:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gatordb?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username`, `password`, and `gatordb` with your PostgreSQL credentials and database name.

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

Run the CLI with:

```sh
gator <command> [args...]
```

Some available commands:

- `register <username>`: Register a new user.
- `login <username>`: Log in as an existing user.
- `addfeed <feed_url>`: Add a new RSS feed (requires login).
- `feeds`: List all available feeds.
- `follow <feed_id>`: Follow a feed (requires login).
- `unfollow <feed_id>`: Unfollow a feed (requires login).
- `browse`: Browse posts from followed feeds (requires login).

For more commands, see [commands.go](commands.go).

For more details, see the source files such as [main.go](main.go), [internal/config/main.go](internal/config/main.go),
