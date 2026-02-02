//go:build darwin

package pathutil

import "testing"

func TestExtractPathFromShellOutput(t *testing.T) {
	tests := []struct {
		name   string
		output string
		want   string
	}{
		{
			name:   "standard path_helper output",
			output: `PATH="/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"; export PATH;`,
			want:   "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin",
		},
		{
			name:   "path with homebrew",
			output: `PATH="/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin"; export PATH;`,
			want:   "/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin",
		},
		{
			name:   "empty output",
			output: "",
			want:   "",
		},
		{
			name:   "malformed - no PATH",
			output: "something else",
			want:   "",
		},
		{
			name:   "malformed - no closing quote",
			output: `PATH="/usr/bin`,
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractPathFromShellOutput(tt.output)
			if got != tt.want {
				t.Errorf("extractPathFromShellOutput(%q) = %q, want %q", tt.output, got, tt.want)
			}
		})
	}
}

func TestEnsureGUIPath(t *testing.T) {
	// This test verifies EnsureGUIPath doesn't error.
	// We can't easily verify the PATH changes in a unit test since
	// it depends on the actual system state.
	err := EnsureGUIPath()
	if err != nil {
		t.Errorf("EnsureGUIPath() returned error: %v", err)
	}
}
