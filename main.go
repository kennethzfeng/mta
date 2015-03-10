package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/kennethzfeng/mta/transit_realtime"
)

func main() {
	file, err := os.Open("sample_data/division_a_sample.gtfs")

	if err != nil {
		log.Fatalf("Can't open file: %s", err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalf("Can't read file: %s", err)
	}

	msg := &transit_realtime.FeedMessage{}
	err = proto.Unmarshal(data, msg)

	if err != nil {
		log.Fatalf("Can't marshal data: %s", err)
	}

	log.Println(msg)
}
