package builders

type BmwCarBuilder struct {
	frame     string
	engine    string
	steering  string
	noOfSeats int
	isRoofTop bool
}

func NewBMWBuilder() *BmwCarBuilder {
	return &BmwCarBuilder{}
}

func (bcb *BmwCarBuilder) BuildFrame() AbstractCarBuilder {
	bcb.frame = "bmwframe"
	return bcb
}

func (bcb *BmwCarBuilder) BuildEngine() AbstractCarBuilder {
	bcb.engine = "bmwengine"
	return bcb
}
func (bcb *BmwCarBuilder) BuildSteering() AbstractCarBuilder {
	bcb.steering = "bmw"
	return bcb
}
func (bcb *BmwCarBuilder) BuildNoOfSeats() AbstractCarBuilder {
	bcb.noOfSeats = 8
	return bcb
}

func (bcb *BmwCarBuilder) BuildRoofTop() AbstractCarBuilder {
	bcb.isRoofTop = true
	return bcb
}

func (bcb *BmwCarBuilder) Build() Car {
	return Car{
		frame:     bcb.frame,
		engine:    bcb.engine,
		steering:  bcb.steering,
		noOfSeats: bcb.noOfSeats,
		isRoofTop: bcb.isRoofTop,
	}
}
