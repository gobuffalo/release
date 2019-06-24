package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/spf13/cobra"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "checks to make sure your system is ready to release",
	RunE: func(cmd *cobra.Command, args []string) error {
		var found bool

		fmt.Println("checking your GITHUB_TOKEN ENV variable...")
		t, err := envy.MustGet("GITHUB_TOKEN")
		if err == nil {
			fmt.Printf("GITHUB_TOKEN is set: %s\n", t)
		} else {
			found = true
			fmt.Printf("GITHUB_TOKEN is NOT set: %s\n", err)
		}

		fmt.Println("\nchecking your git installation...")
		p, err := exec.LookPath("git")
		if err == nil {
			c := exec.Command("git", "version")
			b, err := c.CombinedOutput()
			if err != nil {
				return err
			}
			gv := strings.TrimSpace(string(b))
			fmt.Printf("Git is installed: %s (%s)\n", p, gv)
		} else {
			found = true
			fmt.Printf("Git is NOT installed: %s\n", err)
		}

		_, e1 := os.Stat(".goreleaser.yml")
		_, e2 := os.Stat(".goreleaser.yml.plush")
		if e1 == nil || e2 == nil {
			fmt.Println("\nchecking your Goreleaser installation...")
			p, err := exec.LookPath("goreleaser")
			if err == nil {
				c := exec.Command("goreleaser", "-v")
				b, err := c.CombinedOutput()
				if err != nil {
					return err
				}
				gv := strings.TrimSpace(string(b))
				fmt.Printf("Goreleaser is installed: %s (%s)\n", p, gv)
			} else {
				found = true
				fmt.Printf("Goreleaser is NOT installed: %s\n", err)
			}
		}

		if found {
			return fmt.Errorf("your system is NOT ready to release")
		}

		fmt.Println("your system is ready to release")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
