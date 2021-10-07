package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/asahnoln/event-notifier/pkg"
	"google.golang.org/api/option"
)

// TODO: Test no just happy path - have to tell all those vars are necessary
// TODO If Tg is not set up correctly - there is no error
func main() {
	tomorrow := flag.Bool("tomorrow", false, "Use this flag if you need tomorrow events")
	flag.Parse()

	file, err := os.Open(os.Getenv("GCALMAILNAMES"))
	if err != nil {
		log.Fatalf("error while opening file: %v", err)
	}

	cal := pkg.NewGCalStore(os.Getenv("GCALID"), file, option.WithCredentialsFile(os.Getenv("GCALCRED")))

	var (
		es []pkg.Event
	)
	if *tomorrow {
		es, err = pkg.TomorrowEvents(cal)
	} else {
		es, err = pkg.TodayEvents(cal)
	}
	if err != nil {
		log.Fatalf("error while retrieving events: %v", err)
	}

	sdr := pkg.NewTg(os.Getenv("TGKEY"), os.Getenv("TGCHATID"))

	when := pkg.Today
	if *tomorrow {
		when = pkg.Tomorrow
	}
	err = pkg.Send(es, sdr, when)
	if err != nil {
		log.Fatalf("error while sending message: %v", err)
	}

	fmt.Println("Sent succesfully!")
}
