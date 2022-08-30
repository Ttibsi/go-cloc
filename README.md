# Count-loc

A simple tool that counts the lines of code in the directory passed as an
arguement in the CLI. To use, run `count-loc <directory>` in a terminal. This
tool will by default ignore certain file extensions such as various markup
files and git related files.

```console
NAME:
   count-loc - Count lines of code in a given directory

USAGE:
   count-loc [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --enable-all        Enable searching all extensions blacklisted by default (default: false)
   --enable-ext value  Add file types currently blacklisted
   --help, -h          show help (default: false)
   --ignore-ext value  Add filetype to list to ignore (ex: '.hpp')
   --use-hidden-dirs   Include searching through hidden directors, such as .git (default: false)
   --version, -v       print the version (default: false)
```

Note that this will _only_ work if you include flags before the given
directory. This is due to the cli framework this is using.

### To Install

Pull the binary from the github releases page on the right of this repo and
place it in a location on your PATH. For mac/linux based systems, I'd
recommend `/usr/local/bin` using the below command:

```bash
cp ~/Downloads/count-loc /usr/local/bin
```
