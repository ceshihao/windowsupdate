package main

import (
	"encoding/json"
	"fmt"

	"github.com/ceshihao/windowsupdate"
	"github.com/go-ole/go-ole"
	"github.com/scjalliance/comshim"
)

func main() {
	comshim.Add(1)
	defer comshim.Done()

	// ole.CoInitialize(0)
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	defer ole.CoUninitialize()

	var err error

	session, err := windowsupdate.NewUpdateSession()
	if err != nil {
		panic(err)
	}

	// Query Update History
	fmt.Println("Step 1: Query Update History")
	searcher, err := session.CreateUpdateSearcher()
	if err != nil {
		panic(err)
	}

	result, err := searcher.QueryHistoryAll()
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
