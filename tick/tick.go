package tick

import (
	"github.com/carlescere/scheduler"
)

func EveryMinute(callback func()) {
	scheduler.Every(1).Minutes().Run(callback)
}
