package main

import (
	"fmt"
	"td6/rules"
)

func main() {
	alts1 := []rules.Alternative{1, 2, 3, 4}
	alts2 := []rules.Alternative{1, 2, 3, 4}
	alts3 := []rules.Alternative{4, 3, 2, 1}
	p := rules.Profile{alts1, alts2, alts3}

	fmt.Println(rules.KemenySWF(p))
}
