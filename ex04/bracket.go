package brackets

type Stack struct {
	i  int
	st []byte
}

func New() *Stack {
	return &Stack{}
}

func (c *Stack) Push(numb byte) {
	c.i++
	c.st = append(c.st, numb)
}

func (c *Stack) Pop() byte {
	c.i--
	res := c.st[c.i]
	c.st = c.st[:c.i]
	return res
}

func (c *Stack) Status() int {
	return c.i
}

func (c *Stack) Read() byte {
	res := c.st[(c.i - 1)]
	return res
}

func Bracket(str string) (bool, error) {
	var c *Stack = New()
	openn := "[{("
	closse := "]})"
	lens := len(str)
	if lens == 0 {
		return true, nil
	}
	for i := 0; i < lens; i++ {
		for k := 0; k < 3; k++ {
			if str[i] == openn[k] {
				c.Push(str[i])
				// n++
			} else if str[i] == closse[k] {
				if c.Status() < 1 {
					return false, nil
				}
				if c.Read() == openn[k] {
					c.Pop()
				} else {
					return false, nil
				}
			}
		}
	}
	if c.Status() == 0 {
		return true, nil
	}
	return false, nil
}
