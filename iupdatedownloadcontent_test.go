package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func toIUpdateDownloadContentsTest(disp *ole.IDispatch) ([]*IUpdateDownloadContent, error) {
	if disp == nil {
		return nil, nil
	}
	count := 2
	var res []*IUpdateDownloadContent
	for i := 0; i < count; i++ {
		item := &IUpdateDownloadContent{disp: disp, DownloadUrl: "url"}
		res = append(res, item)
	}
	return res, nil
}

func TestToIUpdateDownloadContents_AllSuccess(t *testing.T) {
	res, err := toIUpdateDownloadContentsTest(&ole.IDispatch{})
	if err != nil || len(res) != 2 || res[0].DownloadUrl != "url" {
		t.Errorf("unexpected: %+v, err=%v", res, err)
	}
}

func TestToIUpdateDownloadContents_NilInput(t *testing.T) {
	res, err := toIUpdateDownloadContentsTest(nil)
	if err != nil || res != nil {
		t.Errorf("unexpected: %+v, err=%v", res, err)
	}
}
