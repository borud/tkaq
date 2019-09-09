package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/borud/tkaq/pkg/decoder"
	"github.com/jessevdk/go-flags"
	"github.com/telenordigital/nbiot-go"
)

const (
	timeFormat = time.RFC3339
)

type progOpts struct {
	CollectionID string `short:"c" long:"collection-id" description:"Collection ID" default:"17dh0cf43jfi2f"`
	Pagesize     int    `short:"p" long:"pagesize" description:"Number of datapoints to return per page" default:"500"`
	StartTime    string `short:"s" long:"start-time" description:"Start date and time"`
}

var opts progOpts
var until time.Time
var start time.Time

func parseOpts() {
	_, err := flags.NewParser(&opts, flags.Default).Parse()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		log.Fatalf("Error parsing flags: %v", err)
	}
	log.Printf("Flags: %+v", opts)
}

// msToTime converts milliseconds since epoch to time.Time
func msToTime(t int64) time.Time {
	return time.Unix(t/int64(1000), (t%int64(1000))*int64(1000000))
}

func calculateTimes() {
	var err error
	// If no start time is given the default is 1 day ago
	if opts.StartTime == "" {
		start = time.Now().AddDate(0, 0, -1)
		return
	}

	start, err = time.Parse(timeFormat, opts.StartTime)
	if err != nil {
		log.Fatalf("Error parsing since time: %v", err)
	}
}

func main() {
	parseOpts()
	calculateTimes()

	client, err := nbiot.New()
	if err != nil {
		log.Fatalf("Unable to connect client: %v", err)
	}

	fmt.Printf("Timestamp,Name,ID,IMSI,IMEI,Status,Lat,Long,Altitude,RelHumidity,Temperature,CO2PPM,TVOCPPB,PM25,PM10\n")
	for {

		data, err := client.CollectionData(opts.CollectionID, time.Time{}, until, opts.Pagesize)
		if err != nil {
			log.Fatalf("Error while reading data: %v", err)
		}

		// Reached end of dataset in response
		if len(data) == 0 {
			return
		}

		for _, entry := range data {
			name, ok := entry.Device.Tags["name"]
			if !ok {
				name = "<unknown>"
			}

			// Payload is a binary blob from the device that we have to decode first
			p, err := decoder.DecodePayload(entry.Payload)
			if err != nil {
				log.Fatalf("Unable to decode payload: %v", err)
			}

			// If the current entry is older than the starting point we are done
			ts := msToTime(entry.Received)
			if ts.Before(start) {
				return
			}

			fmt.Printf("%s,%s,%s,%s,%s,%d,%f,%f,%.1f,%.1f,%.1f,%d,%d,%d,%d\n",
				ts.Format(timeFormat),
				name,
				entry.Device.ID,
				entry.Device.IMSI,
				entry.Device.IMEI,
				p.Status,
				p.Lat,
				p.Long,
				p.Altitude,
				p.RelativeHumidity,
				p.Temperature,
				p.CO2PPM,
				p.TVOCPPB,
				p.PM25,
				p.PM10,
			)

		}

		// Since we are going backwards in time we now update the until parameter
		until = msToTime(data[len(data)-1].Received)
		log.Printf("Page boundary for %d at %s", len(data), until)
	}
}
