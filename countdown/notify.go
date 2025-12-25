package countdown

import (
	"fmt"
	"os/exec"
	"runtime"
)

func notify() error {
	// Tells the OS that the timer is a complete
	os := runtime.GOOS
	var err error
	switch os {
	case "linux":
		err = notifyLinux()
	default:
		return fmt.Errorf("unexpected os to run notification")
	}
	if err != nil {
		return err
	}
	return nil
}

func notifyLinux() error {
	_, err := exec.Command("notify-send", "cntdn", "timer complete!").Output()
	if err != nil {
		return fmt.Errorf("error running notify-send")
	}
	return nil
}
