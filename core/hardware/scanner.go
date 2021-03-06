package hardware

import (
	"fmt"

	"github.com/MikMuellerDev/rpiif"
	"github.com/MikMuellerDev/smarthome-hw-ir/core/config"
	"github.com/MikMuellerDev/smarthome-hw-ir/core/log"
)

var ifScanner rpiif.IfScanner

func Init(config config.Hardware) error {
	if err := ifScanner.Setup(config.ScannerDevicePin); err != nil {
		log.Error("Failed to setup scanner: ", err.Error())
		return err
	}
	log.Trace("Successfully initialized receiver")
	return nil
}

// The `scan` function is launched as a goroutine which matches the received codes against the ones in the config file
// If a code is matched, the specified homescript is sent to the smarthome server
func Scan() {
	for {

		receivedCode, err := ifScanner.Scan()
		if err != nil {
			log.Error("Failed to scan code, exiting: ", err.Error())
			return
		}
		// Match the received code
		fmt.Println("Code received: ", receivedCode)
		go matchCode(receivedCode)
	}
}
