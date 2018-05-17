package cmdrunner

import (
	"fmt"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	cmd string
}

func NewWidget() *Widget {

}
