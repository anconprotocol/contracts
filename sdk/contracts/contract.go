package contracts

type Contract struct {
	name string
	log  string
}	

func NewContract(name string) *Contract {
	return &Contract{name, ""}
}

func (c *Contract) Debug(log string) string {
	c.log = log
	return log
}
