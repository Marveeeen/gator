package main
import _ "github.com/lib/pq"
import (
	"log"
	"os"
	"github.com/marveeeen/gator/internal/config"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)

	programState := &state{cfg: &cfg}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
