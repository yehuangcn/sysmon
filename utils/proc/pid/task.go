package pid

import (
	"os"
	"fmt"
	"path/filepath"
)

type TaskInfo []string

func Task(pid string) TaskInfo {
	t, err := os.Open(filepath.Join(procfs, pid, "task"))
	if err != nil {
		fmt.Println("can not open", procfs)
		return nil
	}

	defer t.Close()

	var tasks TaskInfo
	tasks, err = t.Readdirnames(0)
	if err != nil {
		fmt.Println("Readdirnames failed")
		return nil
	}
	return tasks
}