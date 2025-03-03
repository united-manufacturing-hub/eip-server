package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/danomagnum/gologix"
)

func main() {
	booltags := flag.Int("booltags", 10, "Number of boolean tags to create")
	path := flag.String("path", "1,0", "Path to read data from")
	flag.Parse()

	r := gologix.PathRouter{}

	mtp := gologix.MapTagProvider{}

	path1, err := gologix.ParsePath(*path)
	if err != nil {
		log.Printf("problem parsing path. %v", err)
		os.Exit(1)
	}

	r.Handle(path1.Bytes(), &mtp)

	s := gologix.NewServer(&r)

	createTags(&mtp, *booltags)

	go s.Serve()

	t := time.NewTicker(time.Second * 5)
	for {
		<-t.C
		mtp.Mutex.Lock()
		log.Printf("Data 1: %v", mtp.Data)
		mtp.Mutex.Unlock()

	}
}

func createTags(mtp *gologix.MapTagProvider, booltags int) {
	mtp.Mutex.Lock()
	defer mtp.Mutex.Unlock()
	for i := 1; i <= booltags; i++ {
		name := fmt.Sprintf("Bool%d", i)
		mtp.Data[name] = true
	}
}
