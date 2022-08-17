# Count-loc

A simple tool that counts the lines of code in the directory passed as an
arguement in the CLI. To use, run `count-loc <directory>` in a terminal. This
tool will by default ignore certain file extensions such as various markup
files and git related files.

### To Install

If you have `go` already, you can clone this repo and run `go install`.

Otherwise, you can pull the binary from the github releases page on the right
of this repo and place it in a location on your PATH. For mac/linux based
systems, I'd recommend `/usr/local/bin`.
