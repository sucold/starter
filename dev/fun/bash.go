package fun

import (
	"log"
	"os/exec"
	runtime "runtime"
)

func Bash(sh string, desc string) {
	if runtime.GOOS == "windows" {
		output, err := exec.Command("cmd", "/C", sh).CombinedOutput()
		log.Println(desc, err, string(output))
	} else {
		output, err := exec.Command("bash", sh).CombinedOutput()
		log.Println(desc, err, string(output))
	}
}
