package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/danomagnum/gologix"
)

func main() {
	var (
		booltags = flag.Int("booltags", 10, "Number of boolean tags to create")
		inttags  = flag.Int("inttags", 10, "Number of integer tags to create")
		path     = flag.String("path", "1,0", "Path to read data from")
	)
	flag.Parse()

	r := gologix.PathRouter{}

	mapTagProvider := gologix.MapTagProvider{
		Data: make(map[string]any),
	}

	path1, err := gologix.ParsePath(*path)
	if err != nil {
		log.Printf("problem parsing path. %v", err)
		os.Exit(1)
	}

	r.Handle(path1.Bytes(), &mapTagProvider)

	s := gologix.NewServer(&r)

	createTags(&mapTagProvider, *booltags, *inttags)

	go func(s *gologix.Server) {
		err := s.Serve()
		if err != nil {
			log.Fatalf("Failed to run server: %v", err)
			panic(err)
		}
	}(s)

	t := time.NewTicker(time.Second * 5)
	for {
		<-t.C
		mapTagProvider.Mutex.Lock()
		log.Printf("Data 1: %v", mapTagProvider.Data)
		mapTagProvider.Mutex.Unlock()

	}
}

func createTags(mtp *gologix.MapTagProvider, booltags int, inttags int) {
	mtp.Mutex.Lock()
	defer mtp.Mutex.Unlock()
	for i := 1; i <= booltags; i++ {
		err := mtp.TagWrite(fmt.Sprintf("Bool%d", i), true)
		if err != nil {
			return
		}
	}

	for i := 1; i <= inttags; i++ {
		rndInt := rand.Intn(1000)
		err := mtp.TagWrite(fmt.Sprintf("Int%d", i), rndInt)
		if err != nil {
			return
		}
	}
}
