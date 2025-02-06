package plangs

import "os"

const (
	GoModFile string = "go.mod"
)

// Implements ProgrammingLanguageIdentifier interface
type GoProgrammingLanguageIdentifier struct{}

func (identifier *GoProgrammingLanguageIdentifier) Identify(dirEntries []os.DirEntry) (ProgrammingLanguage, error) {
	for _, dirEntry := range dirEntries {
		isFile := !dirEntry.IsDir()

		if isFile && dirEntry.Name() == GoModFile {
			return Go, nil
		}
	}

	return "", nil
}
