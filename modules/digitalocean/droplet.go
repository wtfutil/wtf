package digitalocean

import "github.com/digitalocean/godo"

// Droplet represents WTF's view of a DigitalOcean droplet
type Droplet struct {
	godo.Droplet
}

// NewDroplet creates and returns an instance of Droplet
func NewDroplet(doDroplet godo.Droplet) *Droplet {
	droplet := &Droplet{
		doDroplet,
	}

	return droplet
}
