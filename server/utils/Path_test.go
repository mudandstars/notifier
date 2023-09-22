package utils

import (
	"testing"
)

func TestPath(t *testing.T) {
	t.Run("correctly retrieves the value for webhooks in http://some-url.com/webhooks/5/some-other/3", func(t *testing.T) {
		value, error := Path("http://some-url.com/webhooks/5/some-other/3", "webhooks")

		if value != 5 {
			t.Fatal("Did not return correct value for webhooks in 'http://some-url.com/webhooks/5/some-other/3'", value)
		}

		if error != nil {
			t.Fatal(error)
		}
	})

	t.Run("correctly retrieves the value for some-other in http://some-url.com/webhooks/5/some-other/3", func(t *testing.T) {
		value, error := Path("http://some-url.com/webhooks/5/some-other/3", "some-other")

		if value != 3 {
			t.Fatal("Did not return correct value for some-other in 'http://some-url.com/webhooks/5/some-other/3'", value)
		}

		if error != nil {
			t.Fatal(error)
		}
	})
}
