package utils
import "github.com/llr104/LiFrame/core/liTimer"

var Scheduler *liTimer.TimerScheduler
var IntervalForever int

func init()  {
	IntervalForever = int(^uint(0) >> 1)
	Scheduler = liTimer.NewAutoExecTimerScheduler()
}

