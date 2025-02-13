package plangs

import "os"

type ProgrammingLanguageIdentifier interface {
	Identify(dirEntries []os.DirEntry) (ProgrammingLanguage, error)
}

type ProgrammingLanguage string

func (pl ProgrammingLanguage) String() string {
	return string(pl)
}
