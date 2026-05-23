package scanner

type Package struct {
	Name    string
	Version string
	Dependencies  map[string]string
}