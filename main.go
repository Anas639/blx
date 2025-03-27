package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anas639/blx/cmd"
	"github.com/anas639/blx/internal/database"
	"github.com/anas639/blx/internal/event/udp"
	"github.com/anas639/blx/internal/printer"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	db, err := database.InitDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	rootCmd := cmd.NewRootCmd(&cmd.Context{
		DB:             db,
		TaskPrinter:    printer.NewPrettyTaskPrinter(),
		ProjectPrinter: printer.NewPrettyProjectPrinter(),
		Listener:       udp.NewUDPListener(),
		Broadcaster:    udp.NewUDPBroadcaster(),
	})
	rootCmd.Execute(ctx)
}
