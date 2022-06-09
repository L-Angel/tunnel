package cfg

import (
	"testing"
)

func TestLoadCfgFromYaml(t *testing.T) {
	c := NewCfg()
	t.Error(c == nil)
}
