package main

import (
	"strconv"
	"log"
	"fmt"
	"strings"
	"os/exec"

	pwl "github.com/justjanne/powerline-go/powerline"
)


func segmentDocker(p *powerline) {
	styles := map[string]map[string]string{
		"up": map[string]string{
			"icon": "^",
			"count": "0",
		},
		"created": map[string]string{
			"icon": "=",
			"count": "0",
		},
		"exited": map[string]string{
			"icon": "x",
			"count": "0",
		},
	}

	var docker string = ""
	var empty bool = true

	out, err := exec.Command(`docker`, `ps`, `-a`, `--format`, `"{{.Status}}"`).Output()
	if err != nil {
		log.Println("Cannot read status from docker daemon")
	}

	lines := strings.Split(string(out), "\n")

	for _, eachline := range lines {
		status := strings.Replace(string(eachline), `"`, "", -1)
		status  = strings.ToLower(status)
		stats  := strings.Split(status, " ")
		if stats[0] != "" {
			empty = false
			now, _ := strconv.Atoi(styles[stats[0]]["count"])
			styles[stats[0]]["count"] = strconv.Itoa(now + 1)
		}
	}

	if !empty {
		docker = fmt.Sprintf("%s:%s %s:%s %s:%s", styles["up"]["icon"],
												  styles["up"]["count"],
												  styles["created"]["icon"],
												  styles["created"]["count"],
												  styles["exited"]["icon"],
												  styles["exited"]["count"])
	}

	if docker != "" {
		p.appendSegment("docker", pwl.Segment{
			Content:    docker,
			Foreground: p.theme.DockerMachineFg,
			Background: p.theme.DockerMachineBg,
		})
	}
}
