package aptible

import (
	"fmt"
	"testing"
)

func TestWaitForOperation(t *testing.T) {
	var tests = []struct {
		name    string
		op_id   int64
		deleted bool
		errored bool
	}{
		{"test_404", 0, true, false},
		// Add other test cases if we need in the future
	}

	for _, tc := range tests {
		c, err := SetUpClient()
		if err != nil {
			t.Errorf("Unable to set up client due to error. \nERROR -- %s", err)
		}
		testname := fmt.Sprintf("%s", tc.name)
		t.Run(testname, func(t *testing.T) {

			deleted, err := c.WaitForOperation(tc.op_id)

			if deleted != tc.deleted {
				t.Errorf("Input: %d should have resulted in a 404.", tc.op_id)
			}

			if (err != nil) != tc.errored {
				t.Errorf("Input: %d caused error: %s", tc.op_id, err)
			}
		})
	}
}

func TestGetIDFromHref(t *testing.T) {
	var tests = []struct {
		name     string
		href     string
		expected int64
		errored  bool
	}{
		{"test_vanilla", "https://api-rachel.aptible-sandbox.com/disks/46", 46, false},
		{"test_conversion_err", "https://api-rachel.aptible-sandbox.com/disks/str", 0, true},
		{"test_too_short", "https://api-rachel.aptible-sandbox.com/2", 0, true},
		// Add other test cases if we need in the future
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%s", tc.name)
		t.Run(testname, func(t *testing.T) {

			id, err := GetIDFromHref(tc.href)

			if (err != nil) != tc.errored {
				t.Errorf("Input: %s caused error: %s", tc.href, err)
			}

			if id != tc.expected {
				t.Errorf("Input: %s should have resulted in id = %d. It was %d instead.", tc.href, tc.expected, id)
			}
		})
	}
}

func TestMakeStringSlice(t *testing.T) {
	var tests = []struct {
		name     string
		if_slice []interface{}
		expected []string
		errored  bool
	}{
		{"test_vanilla",
			[]interface{}{"chocolate", "vanilla", "strawberry"},
			[]string{"chocolate", "vanilla", "strawberry"},
			false},
		{"test_invalid_slice",
			[]interface{}{"chocolate", 24, ""},
			[]string{},
			true},
		// Add other test cases if we need in the future
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%s", tc.name)
		t.Run(testname, func(t *testing.T) {

			slice, err := MakeStringSlice(tc.if_slice)

			if (err != nil) != tc.errored {
				t.Errorf("Input: %v caused error: %v", tc.if_slice, err)
			}

			if !isEqual(slice, tc.expected) {
				t.Errorf("Input: %v should have resulted in slice = %v. It was %v instead.", tc.if_slice, tc.expected, slice)
			}
		})
	}
}

// Tests if two string slices are equal
func isEqual(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
