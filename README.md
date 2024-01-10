# Block Unwanted Hosts

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Build Status](https://github.com/feruzoripov/block-unwanted-hosts/actions/workflows/go.yml/badge.svg)](https://github.com/feruzoripov/block-unwanted-hosts/actions/workflows/go.yml)

`Block Unwanted Hosts` is an open-source tool written in Go designed to block unwanted websites by updating the system's hosts file. It fetches a list of blocked websites from various sources and appends them to the hosts file, effectively preventing access to these sites.

More info: https://en.wikipedia.org/wiki/Shock_site

## Blocked Websites Sources
- [Someonewhocares](https://someonewhocares.org/hosts/hosts)
- [StevenBlack's Hosts](https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts)
- [AdAway](https://adaway.org/hosts.txt)
- [PGl Yoyo](https://pgl.yoyo.org/adservers/serverlist.php?hostformat=hosts;showintro=0&mimetype=plaintext)


## Features

- Fetches and blocks websites from predefined [sources](https://github.com/feruzoripov/block-unwanted-hosts/tree/main?tab=readme-ov-file#blocked-websites-sources).
- Appends entries to the hosts file - `/etc/hosts`.
- Clears the DNS cache to apply changes immediately.

## Usage

To block unwanted hosts, import the `github.com/feruzoripov/block-unwanted-hosts/internal/app/blocking` package in your Go code and call `blocking.FetchAndBlockWebsites()`.

```Go
package main

import (
	"log"

	"github.com/feruzoripov/block-unwanted-hosts/internal/app/blocking"
)

func main() {
	err := blocking.FetchAndBlockWebsites()
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
```
__Note__: Ensure that you have the necessary permissions to update the hosts file, and you may need to run the tool with administrative privileges (`sudo`).

## Command Line Usage
### Prerequisites

- `Go` installed
- `git` installed

### Installation Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/feruzoripov/block-unwanted-hosts.git
   cd block-unwanted-hosts/cmd
   ```
2. Build:

   ```bash
   go build -o blocking
   ```
3. Run (requires `sudo` for updating `/etc/hosts` and clearing DNS cache):

   ```bash
   sudo ./blocking
   ```

## Supported Operating Systems

`Block Unwanted Hosts` is designed to run on the following operating systems:

- macOS
- Debian-based Linux (e.g., Ubuntu)

Please note that these are the currently supported operating systems. While it may work on other Linux distributions, the tool has been tested and optimized for macOS and Debian-based systems.

If you encounter any issues on a different operating system, feel free to [submit an issue](https://github.com/feruzoripov/block-unwanted-hosts/issues) or contribute to improve compatibility.


## Licence

This project is licensed under the MIT License. See the [LICENSE](https://opensource.org/licenses/MIT) file for details.
