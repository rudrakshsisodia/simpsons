package views

import "testing"

func TestDecodePath(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"-Users-foo-bar", "/Users/foo/bar"},
		{"-Users-foo-my--project", "/Users/foo/my-project"},
		{"-Users-rudrakshsisodia-repos-simpsons", "/Users/rudrakshsisodia/repos/simpsons"},
		{"-home-user-some--repo--with--dashes", "/home/user/some-repo-with-dashes"},
		{"", ""},
	}
	for _, tt := range tests {
		got := decodePath(tt.input)
		if got != tt.want {
			t.Errorf("decodePath(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
