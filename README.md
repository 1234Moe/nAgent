# nhentai comic download agent


### USAGE:
nAgent [global options] command [command options] [arguments...]

### VERSION:
alpha v0.0.1

### DESCRIPTION:
nhentai comic download agent

### COMMANDS:
download, d  download comicId

help, h      Shows a list of commands or help for one command

### GLOBAL OPTIONS:

--help, -h                  show help (default: false)

--proxy value, -p value     Set proxy server

--stdout value, --so value  Set result information print to stdout type json or text (Not Finish) (default: "text")

--version, -v               print the version (default: false)


### COMMAND OPTIONS:

--idDir, --id             Use comic id as directory name (default: false)

--noRetry, --nr           Disable retry when download pictures failed (default: false)

--output value, -o value  Set output directory

--thread value, -t value  Set download thread number (default: 10)

--zip, -z                 Create a zip file and delete origin dir (default: false)