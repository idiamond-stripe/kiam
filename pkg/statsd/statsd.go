package statsd

import (
	"fmt"
	"time"

	"gopkg.in/alexcesaro/statsd.v2"
)

var Client *statsd.Client

func New(address string, prefix string, interval time.Duration) error {
	sd, err := statsd.New(
		statsd.Address(address),
		statsd.Prefix(prefix),
		statsd.FlushPeriod(interval),
	)

	if err != nil {
		return fmt.Errorf("error initializing statsd client: %v", err)
	}

	Client = sd
	return nil
}
