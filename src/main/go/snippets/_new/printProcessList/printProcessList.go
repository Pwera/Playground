package printProcessList

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type CallBackChan chan struct{}

func CheckEvery(ctx context.Context, d time.Duration, cb CallBackChan) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(d):
			if cb != nil {
				cb <- struct{}{}
			}
		}
	}
}

func PrintProcessList() {
	command := "tasklist"
	psCommand := exec.Command(command)
	resp, err := psCommand.CombinedOutput()
	if err != nil {
		log.Fatalf("%s command failed", command)
	}
	out := string(resp)
	lines := strings.Split(out, "\n")

	for _, line := range lines {
		if line != "" {
			fmt.Println(line)
		}
	}
}

func ExamplePrintProcess(){
	ctx := context.Background()
	PrintProcessList()
	callback := make(CallBackChan)
	go CheckEvery(ctx, 1*time.Second, callback)
	go func() {
		for {
			select {
			case <-callback:
				PrintProcessList()
			}
		}
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}
