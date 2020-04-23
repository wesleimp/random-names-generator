package fruits

import (
	"fmt"
	"strings"

	"github.com/wesleimp/random-names-generator/pkg/generators/preverbs"
	"github.com/wesleimp/random-names-generator/pkg/randomizer"
)

func Generate() string {
	pre := preverbs.GetRandom()
	fruit := randomizer.GetRandom(fruits[:])

	return strings.Title(fmt.Sprintf("%s %s", pre, fruit))
}
