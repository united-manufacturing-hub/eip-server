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

	err = createTags(&mapTagProvider, *booltags, *inttags)
	if err != nil {
		log.Fatalf("Failed to create tags: %v", err)
	}

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

func createTags(mtp *gologix.MapTagProvider, booltags int, inttags int) error {
	for i := 1; i <= booltags; i++ {
		err := mtp.TagWrite(fmt.Sprintf("Bool%d", i), true)
		if err != nil {
			return err
		}
	}

	for i := 1; i <= inttags; i++ {
		rndInt := rand.Intn(1000)
		err := mtp.TagWrite(fmt.Sprintf("Int%d", i), int16(rndInt))
		if err != nil {
			return err
		}
	}

	err := mtp.TagWrite("testbyte", byte(0x01))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testint32", int32(12345))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testint64", int64(12345678))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testdint", int32(12))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testint8", int8(-16))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testint", int16(3))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testuint64", uint64(1234567))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testuint32", uint32(1234))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testuint", uint16(123))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testfloat32", float32(543.21))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("teststring", "Hello World")
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testfloat64", float64(10238.21))
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testboolarray", []bool{true, false, true, false, true, false, true, false, true, false})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testbytearray", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x01, 0x02, 0x03})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testint8array", []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testint16array", []int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testint32array", []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testint64array", []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testuint8array", []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testuint16array", []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testuint32array", []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("testuint64array", []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		return err
	}
	err = mtp.TagWrite("teststringarray", []string{"Hello1", "World1", "Hello2", "World2", "Hello3", "World3", "Hello4", "World4", "Hello5", "World5"})
	if err != nil {
		return err
	}

	return nil
}
