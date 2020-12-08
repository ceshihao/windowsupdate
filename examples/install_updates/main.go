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

	// Search Updates
	fmt.Println("Step 1: Search Updates")
	searcher, err := session.CreateUpdateSearcher()
	if err != nil {
		panic(err)
	}

	result, err := searcher.Search("IsInstalled=1")
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(result)
	fmt.Println(string(b))

	// Download Updates
	fmt.Println("Step 2: Download Updates")
	downloader, err := session.CreateUpdateDownloader()
	if err != nil {
		panic(err)
	}

	downloadResult, err := downloader.Download(result.Updates)
	if err != nil {
		panic(err)
	}

	c, _ := json.Marshal(downloadResult)
	fmt.Println(string(c))

	// Install Updates
	fmt.Println("Step 3: Install Updates")
	installer, err := session.CreateUpdateInstaller()
	if err != nil {
		panic(err)
	}

	installationResult, err := installer.Install(result.Updates)
	if err != nil {
		panic(err)
	}

	d, _ := json.Marshal(installationResult)
	fmt.Println(string(d))
}
