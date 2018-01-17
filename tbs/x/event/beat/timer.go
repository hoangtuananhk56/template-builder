package beat

import (
	"template-builder/tbs/x/event"
)

var daily = event.NewHub(8)

func OnNewDay() *event.Line {
	return daily.NewLine()
}
