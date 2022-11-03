package factory

type car struct {
	vehicle
	maxSpeed int
}

func createCar(name string, color string, maxSpeed int) Vehicler {
	return &car{
		vehicle{
			name:  name,
			color: color,
		},
		maxSpeed,
	}
}

func (c *car) SetMaxSpeed(maxSpeed int) {
	c.maxSpeed = maxSpeed
}
