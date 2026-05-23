package scanner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies,omitempty"`
}

func ScanNPM(path string) (*PackageJSON, error) {
	if path == "" {
		path = "package.json"
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve package.json path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("read package.json: %w", err)
	}

	var pkg PackageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("parse package.json: %w", err)
	}

	fmt.Printf("package: %s@%s\n", pkg.Name, pkg.Version)
	fmt.Printf("dependencies: %d\n", len(pkg.Dependencies))
	fmt.Printf("devDependencies: %d\n", len(pkg.DevDependencies))

	return &pkg, nil
}
