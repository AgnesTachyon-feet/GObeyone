package agnes

import (
	"fmt"
	"github.com/AgnesTachyon/Gobeyone/agnes/internal/tachyon"
)

func SayHelloAgnes() {
	fmt.Println("Hello Agnes")
	generateTest()
	tachyon.SayHelloTachyon()
}