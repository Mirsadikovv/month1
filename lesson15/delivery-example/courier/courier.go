package courier

type Courier struct {
	Name  string
	Model string
	Type  string
}

func (c *Courier) NewCourier(Name, Model, Type string) {
	c.Name = Name
	c.Model = Model
	c.Type = Type
}
