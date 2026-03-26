package download

import (
	"fmt"
	"sync"
	"time"
)

func DownloadFile(filename string, wg *sync.WaitGroup)  {
	defer wg.Done()
	msg := fmt.Sprintf("Starting download:[%s]",filename)
	fmt.Println(msg)
	time.Sleep(2*time.Second)
	final := fmt.Sprintf("Finished Download: [%s]",filename)
	fmt.Println(final)
}

