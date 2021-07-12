package road_logistic

import (
	"go-design-pattern/creational-pattern/factory"
	"go-design-pattern/creational-pattern/factory/transport/pickup"
	"go-design-pattern/creational-pattern/factory/transport/truck"
	"go-design-pattern/creational-pattern/factory/transport/van"
)

type RoadLogistic struct {
	Length, Width, Height int
}

func (r RoadLogistic) CreateTransport() (transport factory.Transport) {

	volume := r.Length * r.Width * r.Height

	if volume > 1000 {
		transport = &truck.Truck{}
	} else if volume > 100 && volume < 1000 {
		transport = &pickup.Pickup{}
	} else {
		transport = &van.Van{}
	}

	return
}
