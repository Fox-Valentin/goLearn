package regexpdemo

import (
	"fmt"
	"regexp"
)

const text = `
my email is aaa@asd.com
something email is omar@qweqw.com
sda2qwe@qw123.111.com
`

func Demo() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
