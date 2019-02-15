package cmd

import (
	"encoding/gob"
	"log"
	"math/rand"
	"net"

	sillyname "github.com/Pallinder/sillyname-go"
	"github.com/calvinfeng/redisgob/payload"
	"github.com/spf13/cobra"
)

var Dial = &cobra.Command{
	Use:   "dial",
	Short: "dial TCP connection on 8080",
	RunE:  dial,
}

func dial(cmd *cobra.Command, args []string) error {
	log.Println("launching gob client")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(conn)
	cat := &payload.Cat{
		Name: sillyname.GenerateStupidName(),
		Age:  rand.Intn(20),
	}

	encoder.Encode(cat)
	conn.Close()
	return nil
}
