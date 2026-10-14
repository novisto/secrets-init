package main

import (
	"os"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestExitEarlyFlag(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		args []string
		want bool
	}{
		{name: "on by default", want: true},
		{name: "EXIT_EARLY=false turns it off", env: map[string]string{"EXIT_EARLY": "false"}, want: false},
		{name: "SECRETS_INIT_EXIT_EARLY=false turns it off", env: map[string]string{"SECRETS_INIT_EXIT_EARLY": "false"}, want: false},
		{name: "--exit-early=false turns it off", args: []string{"--exit-early=false"}, want: false},
		{name: "EXIT_EARLY=true keeps it on", env: map[string]string{"EXIT_EARLY": "true"}, want: true},
		{name: "empty EXIT_EARLY turns it off", env: map[string]string{"EXIT_EARLY": ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range []string{"EXIT_EARLY", "SECRETS_INIT_EXIT_EARLY"} {
				t.Setenv(k, "")
				os.Unsetenv(k)
			}
			for k, v := range tt.env {
				t.Setenv(k, v)
			}

			var got bool
			app := newApp()
			app.Action = func(c *cli.Context) error {
				got = c.Bool("exit-early")
				return nil
			}
			if err := app.Run(append([]string{"secrets-init"}, tt.args...)); err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("exit-early = %v, want %v", got, tt.want)
			}
		})
	}
}
