# Updatr

A command-line tool to update multiple project directories at once. It automatically detects the programming language of each project and executes bash commands according to the project programming language.

## How it Works

To use `updatr` you must provide a directory containing your projects and a configuration file, for example:

```bash
updatr --path="~/Personal/Projects" --config="~/path/to/config.toml"
```

The `config.toml` file must contain the configuration(commands), for each programming language or project, that will be updated in that directory.

If no configuration file is found, `updatr` will not execute any commands to avoid unexpected behaviors.

## Installation

You can install the tool using the following command:

```bash
go install github.com/tiagovaldrich/updatr@latest
```

Or you can clone the repository and build the project.

```bash
make build
```

And then you can execute the `updatr` binary.

```bash
./updatr --path /path/to/projects --config /path/to/config.toml
```

## Usage

```bash
updatr --path /path/to/projects --config /path/to/config.toml
```

### Arguments

- `--path`: Path to the directory containing your projects (default: executable directory)
- `--config`: Path to the configuration file (default: config.toml in the executable directory)

## Configuration

Create a `config.toml` file with the commands to be executed for each programming language. Example:

```toml
[go]
commands = [
    "git stash",
    "git checkout master",
    "git pull origin master",
    "make install",
]
ignore_projects = ["project5", "project9"]

[nodejs]
commands = [
    "git stash",
    "git checkout master",
    "git pull origin master",
    "npm install",
]
ignore_projects = ["project1", "project2"]
```

## How it Works

1. The tool scans the provided directory for subdirectories
2. For each subdirectory, it detects the programming language based on project files:
   - `go.mod` for Go projects
   - `package.json` for NodeJS projects
3. Once detected, it executes the configured commands for that language in sequence
4. If no configuration is found for a language, it won't run any command to avoid unexpected behaviors

## Contributing

Feel free to open issues and pull requests to add support for more programming languages or improve the existing functionality.

## License

This project is open source and available under the MIT License.
