module github.com/ceshihao/windowsupdate/examples/query_update_history

go 1.23

replace github.com/ceshihao/windowsupdate => ../../

require (
	github.com/ceshihao/windowsupdate v0.0.2
	github.com/go-ole/go-ole v1.3.0
	github.com/scjalliance/comshim v0.0.0-20190308082608-cf06d2532c4e
)

require golang.org/x/sys v0.1.0 // indirect
