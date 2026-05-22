package scanner 

import (
	"os"
	"fmt"
	"encoding/json"
	
)

type packageJSON struct {
	Name  string  `json:"name"`
	Version string `json:"version"`
	Dependencies map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies,omitempty"`

}

func main() {
	data, err := os.ReadFile("package.json")

	if err != nil {
		fmt.Println(err)
		return 

	}

	


}