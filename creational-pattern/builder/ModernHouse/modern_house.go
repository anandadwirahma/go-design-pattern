package ModernHouse

import "go-design-pattern/creational-pattern/builder"

type ModernHouse struct {
	house builder.House
}

func (m *ModernHouse) SetWindowsType() {
	m.house.WindowType = builder.ModernWindowType
}

func (m *ModernHouse) SetFloorType() {
	m.house.FloorType = builder.ModernFloorType
}

func (m *ModernHouse) SetNumOfWindows() {
	m.house.NumOfWindows = 5
}

func (m *ModernHouse) SetNumOfDoors() {
	m.house.NumOfDoors = 2
}

func (m *ModernHouse) SetSwimmingPool() {
	m.house.HasSwimmingPool = true
}

func (m *ModernHouse) CreateHouse() builder.House {
	return m.house
}
