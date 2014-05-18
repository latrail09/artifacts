package upload

import (
	"fmt"
	"sort"
)

var (
	errUploadFailed = fmt.Errorf("upload failed")
)

type nullProvider struct {
	SourcesToFail []string
}

func (np *nullProvider) Upload(id string, opts *Options, in chan *artifact, out chan *artifact, done chan bool) {
	sort.Strings(np.SourcesToFail)
	lenSrc := len(np.SourcesToFail)

	for a := range in {
		idx := sort.SearchStrings(np.SourcesToFail, a.Path.From)
		if idx < 0 || idx >= lenSrc {
			a.Result.OK = false
			a.Result.Err = errUploadFailed
		} else {
			a.Result.OK = true
		}
		out <- a
	}

	done <- true
}
