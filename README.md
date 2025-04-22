# spiderNetscan v1.0.2

A fast and flexible network vulnerability scanner written in Go.

## Features
- Port scanning
- CVE detection (offline/online from NVD or MITRE)
- Auto-update support
- Simple CLI interface
- Supports both TCP and UDP port scanning
- Selectable CVE data sources (NVD, MITRE)

## Usage
```bash
spiderNetscan -p 22,80 -s 192.168.1.0/24 -t 5 -cve

spiderNetscan -p 22 -s 10.0.0.0/24 -t 2 -cve -online

spiderNetscan --update

spiderNetscan -p 22,80 -s 192.168.1.0/24 -t 5 -cve -online -source mitre

spiderNetscan -p 443,8080 -s 192.168.1.0/24 -t 3 -cve -source mitre
```

## Install
```bash
git clone https://github.com/nurendraB/spiderNetscan.git

cd spiderNetscan

chmod +x setup.sh

./setup.sh
```

## 3. Manual Installation (Optional)
If you prefer to do things manually, follow these steps:

### Clone the repository:

```bash

git clone https://github.com/nurendraB/spiderNetscan.git

```
### Install Go dependencies:

```bash

go mod tidy

```
Build the binary:

```bash

go build -o spiderNetscan cmd/spiderNetscan.go

```
### Move the binary to /usr/local/bin/:

``` bash

sudo mv spiderNetscan /usr/local/bin/
```

### Example 1: Scan ports for open services and check CVEs (Offline mode)
```bash
spiderNetscan -p 22,80 -s 192.168.1.0/24 -t 5 -cve
```
### Example 2: Scan ports for open services, check CVEs using an online database (e.g., NVD or MITRE)

``` bash
spiderNetscan -p 22 -s 10.0.0.0/24 -t 2 -cve -online
```
Example 4: Scan multiple ports with a different timeout and check CVEs from MITRE Database

```bash
spiderNetscan -p 443,8080 -s 192.168.1.0/24 -t 3 -cve -source mitre
```

### Update Tool
To update spiderNetscan with the latest changes, run:
```bash

spiderNetscan --update

```

## Author
**@nurendraB (spiderinshell)**
