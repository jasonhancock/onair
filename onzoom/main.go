package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/jasonhancock/runner"
	ps "github.com/mitchellh/go-ps"
)

func main() {
	period := flag.Duration("period", 15*time.Second, "how often to poll to see if you're on zoom")
	endpoint := flag.String("endpoint", "http://onair.jasonhancock.com:8080/", "The API endpoint")
	flag.Parse()

	fn := func(on bool) {
		str := "on"
		if !on {
			str = "off"
		}

		u := *endpoint + str
		log.Println("calling it " + u)

		http.DefaultClient.Get(u)
	}

	ctx := runner.Context()

	run(ctx, *period, fn)
}

func run(ctx context.Context, period time.Duration, action func(on bool)) {
	ticker := time.NewTicker(period)
	previousState := false
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			currentState, err := onZoom()
			if err != nil {
				continue
			}

			if currentState != previousState {
				previousState = currentState
				action(currentState)
			}
		}
	}
}

func onZoom() (bool, error) {
	p, err := ps.Processes()
	if err != nil {
		return false, err
	}

	for _, v := range p {
		if v.Executable() == "CptHost" {
			return true, nil
		}
	}
	return false, nil
}
