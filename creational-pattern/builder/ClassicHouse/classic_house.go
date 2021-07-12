package ClassicHouse

import "go-design-pattern/creational-pattern/builder"

type ClassicHouse struct {
	house builder.House
}

func (m *ClassicHouse) SetWindowsType() {
	m.house.WindowType = builder.ClassicWindowType
}

func (m *ClassicHouse) SetFloorType() {
	m.house.FloorType = builder.ClassicFloorType
}

func (m *ClassicHouse) SetNumOfWindows() {
	m.house.NumOfWindows++
}

func (m *ClassicHouse) SetNumOfDoors() {
	m.house.NumOfDoors = 3
}

func (m *ClassicHouse) SetSwimmingPool() {
	m.house.HasSwimmingPool = false
}

func (m *ClassicHouse) CreateHouse() builder.House {
	return m.house
}
