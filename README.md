# cloudflare-cli
Communicating with Cloudflare API through command line interface

# Table of Contents
1. [Introduction](#introductions)
2. [Features](RELEASE.md)
3. [Usages](#usages)
    1. [Downloads](#downloads)
    2. [Command usages](#commandusages)
4. [Development](#developments)
    1. [Setup environment](#setupenvironment)
    2. [Add new sub command](#addsubcommands)
    3. [Incoming features](#incommingfeatures)
    4. [Build with cross-compile options](#build)

## Introduction <a name="introductions"></a>

## Usages <a name="usages"></a>
Download latest binary from github at 

### Downloads <a name="downloads"></a>
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
curl -OL "https://github.com/epiHATR/cloudflare-cli/releases/download/$release_version/cloudflare-linux-arm64-$release_version"
sudo mv cloudflare-linux-arm64-$release_version /usr/local/bin/cloudflare
sudo chmod +x /usr/local/bin/cloudflare
```
### Command Usages <a name="commandusages"></a>

|            command                   |                    descriptions                     |
|--------------------------------------|-----------------------------------------------------|
| [cloudflare](#cmd_cloudflare)                         | show cli introductions & starter command            |
| [cloudflare version](#cmd_cloudflare_version)                 | display cli version and Cloudflare API version      |
| [cloudflare login](#cmd_cloudflare_login)               | login into Cloudflare API                           |
| [cloudflare zone list](#cmd_cloudflare_zone_list) | list Cloudflare zone in account |
| [cloudflare zone show](#cmd_cloudflare_zone_show) | show Cloudflare zone common information |
| cloudflare zone create | create new Cloudflare zone |
| [cloudflare zone setting pause](#cmd_cloudflare_zone_setting_pause) | pause a cloudflare zone |
| [cloudflare zone setting unpause](#cmd_cloudflare_zone_setting_unpause) | unpause a cloudflare zone |
| [cloudflare zone settings set-type](#cmd_cloudflare_zone_setting_set-type) | set Cloudflare zone type (parital, full, delegated)|
| [cloudflare zone plan list](#cmd_cloudflare_zone_plan_list) | list all Cloudflare plan for zone|
| [cloudflare zone plan show](#cmd_cloudflare_zone_plan_show) | show a Cloudflare plan details|
| [cloudflare zone plan upgrade](#cmd_cloudflare_zone_plan_upgrade) | upgrade Cloudflare zone to a specified plan |
| cloudflare zone delete | delete a Cloudflare zone |
| [cloudflare zone dns list](#cmd_cloudflare_zone_dns_list) | list DNS records on a Cloudflare zone  |
| [cloudflare zone dns show](#cmd_cloudflare_zone_dns_show) | show a DNS records details of a Cloudflare zone |
| [cloudflare zone dns add](#cmd_cloudflare_zone_dns_add) | add a DNS records on a Cloudflare zone |
| cloudflare zone dns update | change/update DNS records value on Cloudflare zone |
| cloudflare zone dns delete | delete a DNS records on a Cloudflare zone |

#### cloudflare <a name="cmd_cloudflare"></a>
<p>show cli introductions & starter command</p>

```bash
cloudflare [ --help | -h ]
           [ --debug ]
```

#### cloudflare login <a name="cmd_cloudflare_login"></a>
<p>Login into Cloudflare API using API Token (--token | -t) or Email/ApiKey (--email|-e and --key|-k) </p>

```bash
cloudflare login [ --token | -t ]
                 [ --email | -e ]
                 [ --key | -k ]

#global flags                                   
                 [ --help | -h ]
                 [ --debug ]
```

#### cloudflare version<a name="cmd_cloudflare_version"></a>
<p>Show current CLI version and Cloudflare API version</p>

```bash
cloudflare version [ --short | -s ]

#global flags                                   
                   [ --help | -h ]
                   [ --debug ]
```

#### cloudflare zone list<a name="cmd_cloudflare_zone_list"></a>
<p>List all Cloudflare zone under account</p>

```bash
cloudflare zone list [ --account-id ]
                     [ --name | -n ]

#global flags       
                    [ --query | -q ]                           
                    [ --output |-o ]
                    [ --help | -h ]
                    [ --debug ]
```

#### cloudflare zone show<a name="cmd_cloudflare_zone_show"></a>
<p>List all Cloudflare zone under account</p>

```bash
cloudflare zone show [ --id | -i]
                     [ --name | -n ]

#global flags       
                    [ --query | -q ]                           
                    [ --output |-o ]
                    [ --help | -h ]
                    [ --debug ]
```
#### cloudflare zone dns list <a name="cmd_cloudflare_zone_dns_list"></a>
<p>List all DNS records in a Cloudflare zone</p>

```bash
cloudflare zone dns list [ --zone-id | -i]
                         [ --type | -t ]

#global flags       
                        [ --query | -q ]                           
                        [ --output |-o ]
                        [ --help | -h ]
                        [ --debug ]
```

#### cloudflare zone dns show <a name="cmd_cloudflare_zone_dns_show"></a>
<p>Show details of a Cloudflare DNS record</p>

```bash
cloudflare zone dns show [ --id | -i ]
                         [ --zone-id | -i]

#global flags       
                         [ --query | -q ]                           
                         [ --output |-o ]
                         [ --help | -h ]
                         [ --debug ]
```

#### cloudflare zone dns add <a name="cmd_cloudflare_zone_dns_add"></a>
<p>Add a Cloudflare DNS record and return its result</p>

See data payload format at [Cloudflare API documentation - create DNS record](https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record)

```bash
cloudflare zone dns add [ --zone-id | -i]
                        [ --data|-d ]
#global flags       
                        [ --query | -q ]                           
                        [ --output |-o ]
                        [ --help | -h ]
                        [ --debug ]
```

#### cloudflare zone setting pause <a name="cmd_cloudflare_zone_setting_pause"></a>
<p>Pause a Cloudfalre zone</p>

```bash
cloudflare zone setting pause [ --zone-id | -i]

#global flags       
                              [ --query | -q ]                           
                              [ --output |-o ]
                              [ --help | -h ]
                              [ --debug ]
```

#### cloudflare zone setting unpause<a name="cmd_cloudflare_zone_setting_unpause"></a>
<p>UnPause a Cloudfalre zone</p>

```bash
cloudflare zone setting unpause [ --zone-id | -i]

#global flags       
                                [ --query | -q ]                           
                                [ --output |-o ]
                                [ --help | -h ]
                                [ --debug ]
```

#### cloudflare zone setting set-type<a name="cmd_cloudflare_zone_setting_set-type"></a>
<p>Change Cloudflare zone type (full, partial)</p>

```bash
cloudflare zone setting set-type [ --zone-id]
                                 [ --type| -t ]

#global flags       
                                 [ --query | -q ]                           
                                 [ --output |-o ]
                                 [ --help | -h ]
                                 [ --debug ]
```

#### cloudflare zone plan list<a name="cmd_cloudflare_zone_plan_list"></a>
<p>List all Cloudflare zone's plan</p>

```bash
cloudflare zone zone plan list [ --zone-id]
                               [ --all-available|-a ]
                               [ --rate-plan-only ]
#global flags       
                                [ --query | -q ]                           
                                [ --output |-o ]
                                [ --help | -h ]
                                [ --debug ]
```

#### cloudflare zone plan show<a name="cmd_cloudflare_zone_plan_show"></a>
<p>Show details of a zone's plan</p>

```bash
cloudflare zone zone plan show [ --zone-id ]
                               [ --id | -i]
#global flags       
                                [ --query | -q ]                           
                                [ --output |-o ]
                                [ --help | -h ]
                                [ --debug ]
```

#### cloudflare zone plan upgrade<a name="cmd_cloudflare_zone_plan_upgrade"></a>
<p>Upgrade a Clouddflare zone plan</p>

```bash
cloudflare zone zone plan upgrade [ --zone-id ]
                                  [ --plan-id | -i]
                                  [ --plan-name | -n]
#global flags       
                                [ --query | -q ]                           
                                [ --output |-o ]
                                [ --help | -h ]
                                [ --debug ]
```

## Development <a name="developments"></a>

### Setup development environment <a name="setupenvironment"></a>

### Add sub commands <a name="addsubcommands"></a>

### Incoming features <a name="incommingfeatures"></a>
- [x] add --output|-o flag to allow output as json, yaml
- [x] enable --query|-q flag to allow JMESpath based queries
- [x] enable login with email & key

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
