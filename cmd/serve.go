package cmd

import (
	"encoding/gob"
	"log"
	"net"

	"github.com/calvinfeng/redisgob/payload"
	"github.com/spf13/cobra"
)

func handleConn(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	cat := &payload.Cat{}
	dec.Decode(cat)
	log.Printf("server received cat %s which is %d years old\n", cat.Name, cat.Age)
	conn.Close()
}

// Serve is a CLI command.
var Serve = &cobra.Command{
	Use:   "serve",
	Short: "listen for TCP connection on 8080",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	log.Println("launching gob server")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("error %v", err)
			continue
		}

		go handleConn(conn)
	}
}
