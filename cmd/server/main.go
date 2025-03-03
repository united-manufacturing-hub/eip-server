package main

import (
	"flag"
	"fmt"
	"log"

	// Adjust these imports to match the actual path/layout in your fork/cloned repo

	"github.com/danomagnum/gologix/ethernetip"
	"github.com/danomagnum/gologix/tag"
)

func main() {
	booltags := flag.Int("booltags", 10, "Number of boolean tags to create")
	flag.Parse()

	cfg := &ethernetip.ServerConfig{
		Host: "0.0.0.0",
		Port: 44818,
		Path: "1,0",
		Name: "UMH-EIP",
	}

	server, err := ethernetip.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create EIP server: %v", err)
	}

	tagDB := tag.NewEmbeddedTagDB()

	for i := 1; i <= *booltags; i++ {
		name := fmt.Sprintf("Bool%d", i)
		err = tagDB.AddTag(name, tag.CIPBool, false)
		if err != nil {
			log.Printf("Failed adding %s: %v", name, err)
		}
	}

	server.SetTagDB(tagDB)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("server.Start() error: %v", err)
		}
	}()

	log.Printf("EtherNet/IP server started on %s with CIP path %s", cfg.Host, cfg.Path)
	log.Printf("Created %d bool tags: Bool1..Bool%d", *booltags, *booltags)

	select {}
}
