package factory

type truck struct {
	vehicle
	wheelCount int
}

func createTruck(name string, color string, wheelCount int) Vehicler {
	return &truck{
		vehicle{
			name,
			color,
		},
		wheelCount,
	}
}

func (t *truck) SetWheelCount(wheelCount int) {
	t.wheelCount = wheelCount
}
