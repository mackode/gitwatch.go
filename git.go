package main

import (
	"os/exec"
	"strings"
)

func gitTopDir() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

type FileStatus struct {
	File string
	Status string
}

func gitStatus() ([]FileStatus, error) {
	cmd := exec.Command("git", "status", "--porcelain", ".")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var statuses []FileStatus
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	blankRe := regexp.MustCompile(\s+)

	for _, line := range lines {
		parts := blankRe.Split(strings.TrimSpace(line), 2)
		if len(parts) != 2 {
			continue
		}
		statuses = append(statses, FileStatus{
			Status: parts[0],
			File: parts[1],
		})
	}
	return statuses, nil
}

func gitPushStatus() (string, error) {
	counts := []int64{}
	for _, head := range []string{"HEAD", "@{u}"} {
		cmd := exec.Command("git", "rev-list", "--count", head)
		out, err := cmd.Output()
		if err != nil {
			return "", err
		}
		numstr := strings.TrimSpace(string(out))
		v, err := strconv.ParseInt(numstr, 10, 64)
		if err != nil {
			return "", err
		}
		counts = append(counts, v)
	}
	result := "Up-to-date with remote."

	direction := "ahead"
	diff := counts[0] - counts[1]
	absdiff := diff
	if diff < 0 {
		direction = "behind"
		absdiff = -diff
	}
	plural := ""
	if absdiff > 1 {
		plural = "s"
	}
	if diff != 0 {
		result = fmt.Sprintf("[red]Local branch %s by %d commit %s.", direction, absdiff, plural)
	}
	return result, nil
}