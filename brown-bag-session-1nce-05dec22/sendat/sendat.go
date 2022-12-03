package main

import (
	"flag"
	"io"
	"log"
	"time"

	"github.com/vasjaj/modem/at"
	"github.com/vasjaj/modem/serial"
	"github.com/vasjaj/modem/trace"
)

func main() {
	baud := flag.Int("b", 9600, "baud rate")
	cmd := flag.String("c", "", "at command")
	dev := flag.String("d", "/dev/ttyUSB0", "path to modem device")
	timeout := flag.Duration("t", 1000*time.Millisecond, "command timeout period")
	verbose := flag.Bool("v", false, "log modem interactions")

	flag.Parse()

	m, err := serial.New(serial.WithPort(*dev), serial.WithBaud(*baud))
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		m.Close()
	}()

	var mio io.ReadWriter = m
	if *verbose {
		mio = trace.New(m)
	}

	a := at.New(mio, at.WithTimeout(*timeout))

	time.Sleep(1000 * time.Millisecond)

	log.Println(">> AT" + *cmd)

	res, err := a.Command(*cmd, at.WithTimeout(*timeout))
	if err != nil {
		log.Fatalln("<<", err)
	}

	for _, l := range res {
		log.Println("<<", l)
	}
}
