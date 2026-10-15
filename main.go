package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/maxkaiser11/gator/internal/config"
	"github.com/maxkaiser11/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)

	s := &state{cfg: &cfg, db: dbQueries}

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error	 {
			"login": handlerLogin,
			"register": handlerRegister,
			"reset": handlerReset,
			"users": handlerUsers,
		},
	}

	if len(os.Args) < 2 {
		log.Fatal("usage: gator <command> [args...]")
	}	

	cmd := command{Name: os.Args[1], Args: os.Args[2:]}
	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
