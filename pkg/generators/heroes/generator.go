package heroes

import (
	"fmt"
	"strings"

	"github.com/wesleimp/random-names-generator/pkg/generators/preverbs"
	"github.com/wesleimp/random-names-generator/pkg/randomizer"
)

// Generate a new heroe
func Generate() string {
	pre := preverbs.GetRandom()
	heroe := randomizer.GetRandom(heroes[:])

	return strings.Title(fmt.Sprintf("%s %s", pre, heroe))
}
