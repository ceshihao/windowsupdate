module github.com/ceshihao/windowsupdate/examples/install_updates

go 1.17

replace github.com/ceshihao/windowsupdate => ../../../windowsupdate

require (
	github.com/ceshihao/windowsupdate v0.0.2
	github.com/go-ole/go-ole v1.2.6
	github.com/scjalliance/comshim v0.0.0-20190308082608-cf06d2532c4e
)

require golang.org/x/sys v0.0.0-20190916202348-b4ddaad3f8a3 // indirect
