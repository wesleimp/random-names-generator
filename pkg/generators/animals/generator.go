package animals

import (
	"fmt"
	"strings"

	"github.com/wesleimp/random-names-generator/pkg/generators/preverbs"
	"github.com/wesleimp/random-names-generator/pkg/randomizer"
)

// Generate animal
func Generate() string {
	pre := preverbs.GetRandom()
	animal := randomizer.GetRandom(animals[:])

	return strings.Title(fmt.Sprintf("%s %s", pre, animal))
}
