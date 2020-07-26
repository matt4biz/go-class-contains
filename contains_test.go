package contains

import "testing"

// go test -v
// go test ./... -cover
// go test ./... -coverprofile=c.out -covermode=count
// go tool cover -html=c.out

var unknown = `{
		"id": 1,
		"name": "bob",
		"addr": {
			"street": "Lazy Lane",
			"city": "Exit",
			"zip": "99999"
		},
		"extra": 21.1
	}`

func TestContains(t *testing.T) {
	var known = []string{
		`{"id": 1}`,
		`{"extra": 21.1}`,
		`{"name": "bob"}`,
		`{"addr": {"street": "Lazy Lane", "city": "Exit"}}`,
	}

	for _, k := range known {
		if err := CheckData(k, []byte(unknown)); err != nil {
			t.Errorf("invalid: %s (%s)\n", k, err)
		}
	}
}

func TestNotContains(t *testing.T) {
	var known = []string{
		`{"id": 2}`,
		`{"pid": 1}`,
		`{"name": "bobby"}`,
		`{"first": "bob"}`,
		`{"addr": {"street": "Lazy Lane", "city": "Alpha"}}`,
		// dup the above with "funk" and "extra" to up coverage
	}

	for _, k := range known {
		if err := CheckData(k, []byte(unknown)); err == nil {
			t.Errorf("false positive: %s\n", k)
		} else {
			t.Log(err)
		}
	}
}
