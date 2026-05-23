package scanner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// type PackageJSON struct {
// 	Name            string            `json:"name"`
// 	Version         string            `json:"version"`
// 	Dependencies    map[string]string `json:"dependencies"`
// 	DevDependencies map[string]string `json:"devDependencies,omitempty"`
// }

func DectectFile() (string, error) {
	if _, err := os.Stat("package.json"); err == nil {
		return "package.json", nil
	}
	if _, err := os.Stat("requirements.txt"); err == nil {
		return "requirements.txt", nil
	}
	if _, err := os.Stat("go.mod"); err == nil {
		return "go.mod", nil
	}
	return "", fmt.Errorf("no supported file found")
}

type Dependency = Package

func ParsePackageJSON(path string) (*Dependency, error) {

	path, err := DectectFile()

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve package.json path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("read package.json: %w", err)
	}

	var pkg Dependency
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("parse package.json: %w", err)
	}
	// fmt.Printf("package: %s@%s\n", pkg.Name, pkg.Version)
	fmt.Println("dependencies:", pkg.Dependencies, pkg.Version)

	return &pkg, nil
}
