package builder

// this function is helper to build a house

type director struct {
	builder HouseBuilderIFace
}

func NewDirector(b HouseBuilderIFace) director {
	return director{
		builder: b,
	}
}

func (d *director) BuildHouse() House {
	d.builder.SetWindowsType()
	d.builder.SetFloorType()
	d.builder.SetNumOfDoors()
	d.builder.SetNumOfWindows()
	d.builder.SetSwimmingPool()
	d.builder.SetWindowsType()
	return d.builder.CreateHouse()
}
