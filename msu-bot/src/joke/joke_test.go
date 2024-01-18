package joke

import "testing"

func TestGetJoke(t *testing.T) {
	joke := GetJoke()

	t.Log(joke)

	if joke == "" {
		t.Error("wrong")
	}
}