package builders

type MiniCarBuilder struct {
	frame     string
	engine    string
	steering  string
	noOfSeats int
	isRoofTop bool
}

func NewMiniBuilder() *MiniCarBuilder {
	return &MiniCarBuilder{}
}

func (mcb *MiniCarBuilder) BuildFrame() AbstractCarBuilder {
	mcb.frame = "miniframe"
	return mcb
}

func (mcb *MiniCarBuilder) BuildEngine() AbstractCarBuilder {
	mcb.engine = "miniengine"
	return mcb
}
func (mcb *MiniCarBuilder) BuildSteering() AbstractCarBuilder {
	mcb.steering = "ministeering"
	return mcb
}
func (mcb *MiniCarBuilder) BuildNoOfSeats() AbstractCarBuilder {
	mcb.noOfSeats = 4
	return mcb
}

func (mcb *MiniCarBuilder) BuildRoofTop() AbstractCarBuilder {
	mcb.isRoofTop = false
	return mcb
}

func (mcb *MiniCarBuilder) Build() Car {
	return Car{
		frame:     mcb.frame,
		engine:    mcb.engine,
		steering:  mcb.steering,
		noOfSeats: mcb.noOfSeats,
		isRoofTop: mcb.isRoofTop,
	}
}
