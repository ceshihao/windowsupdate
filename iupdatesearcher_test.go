package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func TestToIUpdateSearcher(t *testing.T) {
	_, _ = toIUpdateSearcher(&ole.IDispatch{})
}

func TestIUpdateSearcher_Search(t *testing.T) {
	s := &IUpdateSearcher{disp: &ole.IDispatch{}}
	_, _ = s.Search("")
}

func TestIUpdateSearcher_QueryHistory(t *testing.T) {
	s := &IUpdateSearcher{disp: &ole.IDispatch{}}
	_, _ = s.QueryHistory(0, 1)
}

func TestIUpdateSearcher_GetTotalHistoryCount(t *testing.T) {
	s := &IUpdateSearcher{disp: &ole.IDispatch{}}
	_, _ = s.GetTotalHistoryCount()
}

func TestIUpdateSearcher_QueryHistoryAll(t *testing.T) {
	s := &IUpdateSearcher{disp: &ole.IDispatch{}}
	_, _ = s.QueryHistoryAll()
}
