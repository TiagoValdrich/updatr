package updater

var DefaultOperations = []string{"git stash", "git checkout master", "git pull origin master"}

const (
	UserHomeDirAlias      byte = '~'
	DefaultConfigFilePath      = "config.toml"
)
