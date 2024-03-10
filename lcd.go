package main

import (
	"fmt"
	"time"

	"github.com/artvel/display"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	l := display.Find()
	defer func() {
		panicCheck(l.Close())
	}()

	displayMem(l)
	panicCheck(l.Write(0, "Booted"))
	time.Sleep(5 * time.Second)
	displayStatus(l)
	displayMem(l)

	go l.Listen(func(btn int, released bool) bool {
		panicCheck(l.Enable(true))
		if btn == 2 {
			panicCheck(l.Enable(false))
		} else {
			displayStatus(l)
			displayMem(l)
		}
		return true
	})

	for {
		time.Sleep(10 * time.Second)
		displayStatus(l)
		displayMem(l)
	}

}

func panicCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func displayStatus(l display.LCD) {
	loadV, _ := load.Avg()
	uptime, _ := host.Uptime()
	status := fmt.Sprintf("L1:%0.2f Up:%0.2f", loadV.Load1, float64(uptime)/float64(86400))
	// println(status)
	panicCheck(l.Write(0, status))
}

func displayMem(l display.LCD) {
	vmem, _ := mem.VirtualMemory()
	panicCheck(l.Write(1, display.Progress(int(vmem.UsedPercent))))

	// status := fmt.Sprintf("Mem: %0.2f", vmem.UsedPercent)
	// println(status)
}
