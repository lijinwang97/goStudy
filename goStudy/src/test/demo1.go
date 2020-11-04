package main

import (
	"os/exec"
	"strings"
)

func ls(path string) error {
	cmd := exec.Command("killall", path)
	//cmd.Stdout = os.Stdout
	return cmd.Run()
}

func main() {
	/*err := ls("pushrtmp_macos")
	if err != nil {
		println("hhh")
	}*/
	split := strings.Split("./bin/pushrtmp_macos", "/")
	println(split,"---",split[len(split)-1])
}