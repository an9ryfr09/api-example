package loader

import "testing"

func TestLoad(t *testing.T) {
	var conf *confFile
	conf = Load()
	if conf.Core.RunMode != "debug" {
		t.Errorf("%s", conf.Core.RunMode)
	}
}
