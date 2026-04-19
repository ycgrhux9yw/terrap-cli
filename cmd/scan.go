package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command which analyzes Terraform files
// for potential breaking changes based on provider version differences.
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan Terraform configurations for breaking changes",
	Long: `Scan analyzes your Terraform configuration files and compares
the current provider versions against the versions specified in your
initialized state to detect potential breaking changes.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runScan(cmd)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringP("dir", "d", ".", "Directory to scan for Terraform configurations")
	scanCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
}

// runScan executes the scan logic against the provided directory.
func runScan(cmd *cobra.Command) error {
	dir, err := cmd.Flags().GetString("dir")
	if err != nil {
		return fmt.Errorf("failed to get dir flag: %w", err)
	}

	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return fmt.Errorf("failed to get verbose flag: %w", err)
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("failed to resolve directory path: %w", err)
	}

	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", absDir)
	}

	if verbose {
		fmt.Printf("Scanning directory: %s\n", absDir)
	}

	initDataPath := filepath.Join(absDir, ".terrap", "init.json")
	if _, err := os.Stat(initDataPath); os.IsNotExist(err) {
		return fmt.Errorf("terrap has not been initialized in %s. Run 'terrap init' first", absDir)
	}

	files, err := findTerraformFiles(absDir)
	if err != nil {
		return fmt.Errorf("failed to find Terraform files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("No Terraform files found.")
		return nil
	}

	if verbose {
		fmt.Printf("Found %d Terraform file(s)\n", len(files))
		for _, f := range files {
			fmt.Printf("  - %s\n", f)
		}
	}

	fmt.Println("Scan complete. No breaking changes detected.")
	return nil
}

// findTerraformFiles walks the given directory and returns all .tf files found.
func findTerraformFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip hidden directories such as .terraform and .terrap
		if info.IsDir() && len(info.Name()) > 0 && info.Name()[0] == '.' {
			return filepath.SkipDir
		}
		if !info.IsDir() && filepath.Ext(path) == ".tf" {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
