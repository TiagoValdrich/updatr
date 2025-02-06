package plangs

var AvailableIdentifiers = []ProgrammingLanguageIdentifier{
	&GoProgrammingLanguageIdentifier{},
	&NodeJSProgrammingLanguageIdentifier{},
}

const (
	Go     ProgrammingLanguage = "go"
	NodeJS ProgrammingLanguage = "nodejs"
)
