package cmd

import (
	"testing"
)

func TestPropReplacer(t *testing.T) {

	content := propReplacer([]byte("test=test\r\n\r\ntest\r\n\r\nabcd=polop"), "test", "value")
	t.Log("Running test" + string(content))

}
