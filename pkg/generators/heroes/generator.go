package heroes

import (
	"strings"

	"github.com/wesleimp/random-names-generator/pkg/randomizer"
)

// Generate a new heroe
func Generate() string {
	heroe := randomizer.GetRandom(heroes[:])

	return strings.Title(heroe)
}
