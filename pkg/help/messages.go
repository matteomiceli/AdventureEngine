package help

import (
	"adventureengine/pkg/color"
	"fmt"
)

func Default() {
	fmt.Printf(`ADVENTURE ENGINE DOCS

Using the %s command, you can read the docs related to any valid command.

Example: 

	--> help take 

Will tell you everything you need to know about the %s command. 

To read through all the DOCS, you can use 'help all' command.

List of valid commands: 
- help
- take 
- drop
- walk
- hide

`, color.PaintText(color.Yellow, "HELP"), color.PaintText(color.Yellow, "TAKE"))
}
