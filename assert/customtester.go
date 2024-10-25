package assert

import "testing"

func Equal[V comparable](t *testing.T, got, expected V) {
    t.Helper()

    if expected != got {
        t.Fatalf(`
🔵 Expected:
%v

🔴 Got:
%v

`, expected, got)
    }
}

// Usage:
//
// assert.Equal(t, response, `{"status":"ok"}`)

// no need for third party test lib like testify!
