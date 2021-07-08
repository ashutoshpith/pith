package box

import "github.com/Delta456/box-cli-maker/v2"

func Box() *box.Box  {
	config := box.Config{Px: 12, Py: 2, Type: "", TitlePos: "Inside"}
    boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
  
	return &boxNew

}