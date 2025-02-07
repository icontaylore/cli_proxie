package command

import "bufio"

func Command(c *bufio.Scanner) string {
	c.Scan()
	command := c.Text()
	return command
}
