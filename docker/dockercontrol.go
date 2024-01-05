package docker

import (
	"os/exec"
)


func Dockerls ()(string, error){
	cmd := exec.Command("bash", "-c", "docker ps")
	output, err := cmd.CombinedOutput()
	return string(output), err
}