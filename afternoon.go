package afternoon

import "fmt"

func GoodAfternoon(name string) string {
	message := fmt.Sprintf("[v2] Hey %v, Good afternoon!", name)
	return message
}
