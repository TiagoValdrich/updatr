package plangs

import "os"

const (
	PackageJsonFileName = "package.json"
)

// Implements ProgrammingLanguageIdentifier interface
type NodeJSProgrammingLanguageIdentifier struct{}

func (identifier *NodeJSProgrammingLanguageIdentifier) Identify(dirEntries []os.DirEntry) (ProgrammingLanguage, error) {
	for _, dirEntry := range dirEntries {
		isFile := !dirEntry.IsDir()

		if isFile && dirEntry.Name() == PackageJsonFileName {
			return NodeJS, nil
		}
	}

	return "", nil
}
