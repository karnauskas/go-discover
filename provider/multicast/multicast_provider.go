package multicast

import (
	"log"
	"strings"
	"time"

	"github.com/schollz/peerdiscovery"
)

type Provider struct{}

func (p *Provider) Help() string {
	return `multicast:

    provider:          "multicast"

`
}

func (p *Provider) Addrs(args map[string]string, l *log.Logger) (addrs []string, err error) {

	discoveries, err := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit:     -1,
		Delay:     500 * time.Millisecond,
		TimeLimit: 10 * time.Second,
	})

	if err != nil {
		return
	}

	for _, d := range discoveries {
		addrs = append(addrs, d.Address)
	}

	if len(addrs) > 0 {
		l.Printf("[DEBUG] discover-multicast: found %d node(s): %s.", len(addrs), strings.Join(addrs, ", "))
	} else {
		l.Printf("[DEBUG] discover-multicast: no addresses found")
	}

	return
}
