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

- Fetches and blocks websites from specified URLs.
- Appends entries to the hosts file - `/etc/hosts`.
- Checks for duplicates to avoid redundancy.

## Installation

### Prerequisites

- Go 1.2 or later

### Installation Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/feruzoripov/block-unwanted-hosts.git
   cd block-unwanted-hosts
   ```
2. Build:

   ```bash
   go build -o blocking
   ```
3. Run:

   ```bash
   ./blocking
   ```

## Licence

This project is licensed under the MIT License. See the [LICENSE](https://opensource.org/licenses/MIT) file for details.
