package main

import (
	"fmt"
	"time"

	"github.com/johnsudaar/jvc_api/client"
)

type Request struct {
	Command   string
	SessionID string
}

func main() {
	c, err := client.New("192.168.0.102", "prohd", "0000")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("==> PROGRAM")
		err = c.SetStudioTally(client.TallyProgram)
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)

		fmt.Println("==> PREVIEW")
		err = c.SetStudioTally(client.TallyPreview)
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)

		fmt.Println("==> OFF")
		err = c.SetStudioTally(client.TallyOff)
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)
	}
}
