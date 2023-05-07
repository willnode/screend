# screend

Exploiting GNU screen to be used like sysdaemons

## Why?

System daemons (background processes) is only available for root users. But what if 
you want to run something as a daemon but not as root? You can use `screen` to do that, 
but it's designed for interactive use. This tool is designed to make `screen` more
suitable for daemon use, by leveraging a configuration file, it will remember
the last `screen` processes and start all processes it if it's not running.

## Installation

Requires GNU `screen`, install it from package managers.

```
go install github.com/willnode/screend/...@latest
```

This software doesn't work on Windows.

## Usage

Example usage:

```bash
# Setup a work environment where we need to run
# a code-server and a nodejs development server
screend add code-server
screend add --dir ~/myproject npm run start

# Start all processes to start working
screend start
# Stop when you're done
screend stop
```

See `screend help` or [this file](./screend/usage.go) for more information.


