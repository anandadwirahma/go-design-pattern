package builder

type HouseBuilderIFace interface {
	SetWindowsType()
	SetFloorType()
	SetNumOfWindows()
	SetNumOfDoors()
	SetSwimmingPool()
	CreateHouse() House
}

type House struct {
	WindowType      string
	FloorType       string
	NumOfWindows    int
	NumOfDoors      int
	HasSwimmingPool bool
}
