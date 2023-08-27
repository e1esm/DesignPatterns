package main

type Customer struct {
	id string
}

func (c *Customer) update(id string) {
	c.id = id
}

func (c *Customer) getID() string {
	return c.id
}
