package calculator

type Compute struct{}

func NewCompute() *Compute {
	return &Compute{}
}

func (c *Compute) process(input string) (output float64, err error) {
	output = 0
	err = nil
}
