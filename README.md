# cloudflare-cli
Running Cloudflare API through CLI

# Table of Contents
1. [Introduction](#introductions)
2. [Usages](#usages)
    1. [Installations](#installations)
    2. [Command usages](#commandusages)
3. [Development](#developments)
    1. [Setup environment](#setupenvironment)
    2. [Add new sub command](#addsubcommands)
    3. [Test local function](#testlocal)
    4. [Build with cross-compile options](#build)

## Introduction <a name="introductions"></a>

## Usages <a name="usages"></a>
Download latest binary from github

### Installation <a name="installations"></a>
__For Windows x64__

Download latest build for Windows amd64 at this addess [cloudflare-windows-amd64-{release version}.zip](https://github.com/epiHATR/cloudflare-cli/releases)
Extract zip file, rename to `cloudflare.exe` and copy it to `\Windows\system32\cloudflare`
set PATH to `C:\Windows\system32\cloudflare`

__For Linux amd64__
```
release_version="v0.0.x"
curl -OL "https://github.com/epiHATR/cloudflare-cli/releases/download/$release_version/cloudflare-linux-amd64-$release_version"
sudo mv cloudflare-linux-amd64-$release_version /usr/local/bin/cloudflare
sudo chmod +x /usr/local/bin/cloudflare
```

__For MacOS arm64__
```
release_version="v0.0.x"
curl -OL "https://github.com/epiHATR/cloudflare-cli/releases/download/$release_version/cloudflare-linux-amd64-$release_version"
sudo mv cloudflare-linux-amd64-$release_version /usr/local/bin/cloudflare
sudo chmod +x /usr/local/bin/cloudflare
```
### Command Usages <a name="commandusages"></a>

|            command                   |                    descriptions                     |
|--------------------------------------|-----------------------------------------------------|
| [cloudflare](#cmd_cloudflare)                         | show cli introductions & starter command            |
| [cloudflare version](#cmd_cloudflare_version)                 | display cli version and Cloudflare API version      |
| [cloudflare login](#cmd_cloudflare_login)               | login into Cloudflare API                           |
| cloudflare zone list | list Cloudflare zone |
| cloudflare zone show | show Cloudflare zone common information |
| cloudflare zone create | create new Cloudflare zone |
| cloudflare zone pause | pause a cloudflare zone |
| cloudflare zone delete | delete a Cloudflare zone |
| cloudflare zone set-type | set Cloudflare zone type (parital, full, delegated)|
| cloudflare zone set-plan | set Cloudflare zone plan by name or type |
| cloudflare zone set-under-attack | set Cloudflare zone I'm under attack mode |
| cloudflare zone dns list | list DNS records on a Cloudflare zone  |
| cloudflare zone dns show | show a DNS records details of a Cloudflare zone |
| cloudflare zone dns create | create a DNS records on a Cloudflare zone |
| cloudflare zone dns update | change/update DNS records value on Cloudflare zone |
| cloudflare zone dns delete | delete a DNS records on a Cloudflare zone |

#### cloudflare <a name="cmd_cloudflare"></a>
<p>show cli introductions & starter command</p>

```bash
cloudflare [ --help | -h ]
           [ --debug | -d ]
```

#### cloudflare login <a name="cmd_cloudflare_login"></a>
<p>Login into Cloudflare API using API Token (--token | -t) or Email/ApiKey (--email|-e and --key|-k) </p>

```bash
cloudflare login [ --token | -t ]
                 [ --email | -e ]
                 [ --key | -k ]

#global flags                                   
                 [ --help | -h ]
                 [ --debug | -d ]
```

#### cloudflare version<a name="cmd_cloudflare_version"></a>
<p>Show current CLI version and Cloudflare API version</p>

```bash
cloudflare version [ --short | -s ]

#global flags                                   
                   [ --help | -h ]
                   [ --debug | -d ]
```

## Development <a name="developments"></a>

### Setup development environment <a name="setupenvironment"></a>

### Add sub commands <a name="addsubcommands"></a>

### Test local function <a name="testlocal"></a>

### Build <a name="build"></a>

Following command running on MACOS M1 which will create a release version of cloudflare-cli, targeting for:
__Windows exe__
```
version="v0.0.x"
GOOS=windows GOARCH=amd64 go build -o "release/windows-amd64-$version/cloudflare.exe" cloudflare
```
__For Linux 64 binary__
```
GOOS=linux GOARCH=amd64 go build -o "release/linux-amd64-$version/cloudflare" cloudflare
cd release/linux-amd64-$version
tar -czvf linux-amd64-$version.tar.gz cloudflare
```
__For MacOS binary__
```
GOOS=darwin GOARCH=arm64 go build -o "release/darwin-arm64-$version/cloudflare" cloudflare
cd release/darwin-arm64-$version
tar -czvf darwin-arm64-$version.tar.gz cloudflare
```