# cloudflare-cli
##### Communicating with Cloudflare API through command line interface

![latest release](https://img.shields.io/github/v/release/epiHATR/cloudflare-cli)
![go version](https://img.shields.io/github/go-mod/go-version/epiHATR/cloudflare-cli)
![download](https://img.shields.io/github/downloads/epiHATR/cloudflare-cli/total)
![supports](https://img.shields.io/badge/supports-windows%2C%20macos%2C%20linux-green)
![issues](https://img.shields.io/github/issues/epiHATR/cloudflare-cli?color=%23CC0058)
![license](https://img.shields.io/github/license/epiHATR/cloudflare-cli?color=93BB29&logoColor=green)

<a href="https://www.buymeacoffee.com/hidetran"><img src="https://img.buymeacoffee.com/button-api/?text=Buy me a coffee&emoji=&slug=hidetran&button_colour=FFDD00&font_colour=000000&font_family=Cookie&outline_colour=000000&coffee_colour=ffffff" /></a>

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
Download the latest release at [Cloudflare CLI release](https://github.com/epiHATR/cloudflare-cli/releases)

Download one of following release zip file:
```bash
https://github.com/epiHATR/cloudflare-cli/releases/download/v0.x.x/darwin-arm64-v0.x.x.tar.gz
https://github.com/epiHATR/cloudflare-cli/releases/download/v0.x.x/linux-amd64-v0.x.x.tar.gz 
https://github.com/epiHATR/cloudflare-cli/releases/download/v0.x.x/windows-amd64-v0.x.x.zip
```
### Installation
##### ubuntu/rhel
```
tar -xf linux-amd64-v0.x.x.tar.gz -C /usr/local/bin
sudo chmod +x /usr/local/bin/cloudflare
```
##### darwin/macos
```
tar -xf darwin-arm64-v0.x.x.tar.gz -C /usr/local/bin
sudo chmod +x /usr/local/bin/cloudflare
```
##### windows x64
```
extract windows-amd64-v0.x.x.zip to a folder
.\cloudflare.exe --help
```
### Command Usages <a name="commandusages"></a>

|            command                   |                    descriptions                     |
|--------------------------------------|-----------------------------------------------------|
| [cloudflare](#cmd_cloudflare)                         | show cli introductions & starter command            |
| [cloudflare version](#cmd_cloudflare_version)                 | display cli version and Cloudflare API version      |
| [cloudflare login](#cmd_cloudflare_login)               | login into Cloudflare API 
| | |
| [cloudflare account](#cmd_cloudflare_account)               |  manage Cloudflare accounts/organization |
| [cloudflare account list](#cmd_cloudflare_account_list)               |  list all Cloudflare accounts/organization |
| [cloudflare account show](#cmd_cloudflare_account_show)               |  show details of a cloudflare account/organization |
| [cloudflare account update](#cmd_cloudflare_account_update)      |  update information of a cloudflare account/organization |
| | |
| [cloudflare account role](#cmd_cloudflare_account_role)      |  manage all roles in cloudflare account/organization |
| [cloudflare account role list](#cmd_cloudflare_account_role_list)      |  list all roles in cloudflare account/organization |
| [cloudflare account role show](#cmd_cloudflare_account_role_show)      |  show information details of a role in cloudflare account/organization |
| | |
| [cloudflare account user](#cmd_cloudflare_account_user)      |  manage users in a Cloudflare account/organization |
| [cloudflare account user add](#cmd_cloudflare_account_user_add)      |  add new user to a Cloudflare account/organization |
| [cloudflare account user list](#cmd_cloudflare_account_user_list)      |  manage users in a Cloudflare account/organization |
| [cloudflare account user show](#cmd_cloudflare_account_user_show)      |  show information details of a user in a Cloudflare account/organization |
| [cloudflare account user delete](#cmd_cloudflare_account_user_delete)      |  remove an user from a Cloudflare account/organization |
| | |
| [cloudflare zone list](#cmd_cloudflare_zone_list) | list Cloudflare zone in account |
| [cloudflare zone show](#cmd_cloudflare_zone_show) | show Cloudflare zone common information |
| [cloudflare zone create](#cmd_cloudflare_zone_create) | create new Cloudflare zone |
| [cloudflare zone delete](#cmd_cloudflare_zone_delete) | delete a Cloudflare zone |
| | |
| [cloudflare zone setting pause](#cmd_cloudflare_zone_setting_pause) | pause a cloudflare zone |
| [cloudflare zone setting unpause](#cmd_cloudflare_zone_setting_unpause) | unpause a cloudflare zone |
| [cloudflare zone setting set-type](#cmd_cloudflare_zone_setting_set-type) | set Cloudflare zone type (parital, full, delegated)|
| | |
| [cloudflare zone cache purge](#cmd_cloudflare_zone_cache_purge) | control cache behavior on Cloudflare zone |
| | |
| [cloudflare zone plan list](#cmd_cloudflare_zone_plan_list) | list all Cloudflare plan for zone|
| [cloudflare zone plan show](#cmd_cloudflare_zone_plan_show) | show a Cloudflare plan details|
| [cloudflare zone plan upgrade](#cmd_cloudflare_zone_plan_upgrade) | upgrade Cloudflare zone to a specified plan |
| | |
| [cloudflare zone dns list](#cmd_cloudflare_zone_dns_list) | list DNS records on a Cloudflare zone  |
| [cloudflare zone dns add](#cmd_cloudflare_zone_dns_add) | add a DNS records on a Cloudflare zone |
| [cloudflare zone dns show](#cmd_cloudflare_zone_dns_show) | show a DNS records details of a Cloudflare zone |
| [cloudflare zone dns update](#cmd_cloudflare_zone_dns_update) | change/update DNS records value on Cloudflare zone |
| [cloudflare zone dns delete](#cmd_cloudflare_zone_dns_delete)| delete a DNS records on a Cloudflare zone |

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

#### cloudflare account <a name="cmd_cloudflare_account"></a>
Manage Cloudlare managed accounts/organization

```bash
cloudflare account

#global flags                                   
                 [ --help | -h ]
                 [ --debug ]
```

#### cloudflare account list <a name="cmd_cloudflare_account_list"></a>
List all Cloudlare managed accounts/organization

```bash
cloudflare account list [ --name | -n ]

#global flags               
                        [ --query | -q ]                           
                        [ --output |-o ]                    
                        [ --help | -h ]
                        [ --debug ]
```

#### cloudflare account show <a name="cmd_cloudflare_account_show"></a>
Show details of a Cloudlare managed account/organization

```bash
cloudflare account show [ --account-id ]

#global flags               
                        [ --query | -q ]                           
                        [ --output |-o ]                    
                        [ --help | -h ]
                        [ --debug ]
```
#### cloudflare account update <a name="cmd_cloudflare_account_update"></a>
Update information details of a Cloudlare managed account/organization
See payload data at https://api.cloudflare.com/#accounts-update-account

```bash
cloudflare account update [ --account-id ]
                          [ --data | -d ]

#global flags               
                          [ --query | -q ]                           
                          [ --output |-o ]                    
                          [ --help | -h ]
                          [ --debug ]
```

#### cloudflare account role <a name="cmd_cloudflare_account_role"></a>
Manage roles in a Cloudlare managed accounts/organization

```bash
cloudflare account role

#global flags               
                        [ --query | -q ]                           
                        [ --output |-o ]                    
                        [ --help | -h ]
                        [ --debug ]
```

#### cloudflare account role list<a name="cmd_cloudflare_account_role_list"></a>
List al roles in a Cloudlare managed accounts/organization

```bash
cloudflare account role list [ --account-id ]

#global flags               
                             [ --query | -q ]                           
                             [ --output |-o ]                    
                             [ --help | -h ]
                             [ --debug ]
```

#### cloudflare account role show<a name="cmd_cloudflare_account_role_show"></a>
Show information details of a role in a Cloudlare managed accounts/organization

```bash
cloudflare account role show [ --account-id ]
                             [ --role-id ]

#global flags               
                             [ --query | -q ]                           
                             [ --output |-o ]                    
                             [ --help | -h ]
                             [ --debug ]
```

#### cloudflare account user <a name="cmd_cloudflare_account_user"></a>
Manage users in a Cloudlare managed accounts/organization

```bash
cloudflare account user

#global flags               
                        [ --query | -q ]                           
                        [ --output |-o ]                    
                        [ --help | -h ]
                        [ --debug ]
```

#### cloudflare account user list<a name="cmd_cloudflare_account_user_list"></a>
List all users in a Cloudlare managed accounts/organization

```bash
cloudflare account user list [ --account-id ]

#global flags               
                             [ --query | -q ]                           
                             [ --output |-o ]                    
                             [ --help | -h ]
                             [ --debug ]
```

#### cloudflare account user add<a name="cmd_cloudflare_account_user_add"></a>
Add new user to a Cloudlare managed accounts/organization

```bash
cloudflare account user add [ --account-id ]
                            [ --email ]
                            [ --status 'accepted', 'pending' ]
                            [ --role-id '--role-id=role1 --role-id=role2' ]

#global flags               
                            [ --query | -q ]                           
                            [ --output |-o ]                    
                            [ --help | -h ]
                            [ --debug ]
```

#### cloudflare account user show<a name="cmd_cloudflare_account_user_show"></a>
Show information details of a user in a Cloudlare managed accounts/organization

```bash
cloudflare account user show [ --account-id ]
                             [ --user-id ]

#global flags               
                             [ --query | -q ]                           
                             [ --output |-o ]                    
                             [ --help | -h ]
                             [ --debug ]
```

#### cloudflare account user delete<a name="cmd_cloudflare_account_user_delete"></a>
Show information details of a user in a Cloudlare managed accounts/organization

```bash
cloudflare account user delete [ --account-id ]
                               [ --user-id ]
                               [ --force]

#global flags               
                               [ --query | -q ]                           
                               [ --output |-o ]                    
                               [ --help | -h ]
                               [ --debug ]
```

#### cloudflare zone list<a name="cmd_cloudflare_zone_list"></a>
<p>List all Cloudflare zone under account</p>

```bash
cloudflare zone list [ --account-id ]
                     [ --account-name | -n ]

#global flags       
                     [ --query | -q ]                           
                     [ --output |-o ]
                     [ --help | -h ]
                     [ --debug ]
```
#### cloudflare zone create<a name="cmd_cloudflare_zone_create"></a>
<p>Create a Cloudflare zone under account</p>

```bash
cloudflare zone create [ --name | -n ]
                       [ --account-id ]
                       [ --type | -t ]
                       [ --plan-name ]
                       [--fetch-existing-dns]
#global flags       
                       [ --query | -q ]                           
                       [ --output |-o ]
                       [ --help | -h ]
                       [ --debug ]
```

#### cloudflare zone delete<a name="cmd_cloudflare_zone_delete"></a>
<p>Delete a Cloudflare zone under account</p>

```bash
cloudflare zone delete [ --zone-id ]
                       [ --force | -f ]
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
cloudflare zone dns add [ --zone-id ]
                        [ --data|-d ]
#global flags       
                        [ --query | -q ]                           
                        [ --output |-o ]
                        [ --help | -h ]
                        [ --debug ]
```

#### cloudflare zone dns update <a name="cmd_cloudflare_zone_dns_update"></a>
<p>Update a Cloudflare DNS record</p>
See data payload format at [Cloudflare API documentation - create DNS record](https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record)

```bash
cloudflare zone dns delete [ --zone-id ]
                           [ --id | -i ]
                           [ --data | -d ]
                           [ --force | -f]
#global flags       
                           [ --query | -q ]                           
                           [ --output |-o ]
                           [ --help | -h ]
                           [ --debug ]
```

#### cloudflare zone dns delete <a name="cmd_cloudflare_zone_dns_delete"></a>
<p>Delete a Cloudflare DNS record</p>

```bash
cloudflare zone dns delete [ --zone-id ]
                           [ --id | -i ]
                           [ --force | -f]
#global flags       
                           [ --query | -q ]                           
                           [ --output |-o ]
                           [ --help | -h ]
                           [ --debug ]
```

#### cloudflare zone setting pause <a name="cmd_cloudflare_zone_setting_pause"></a>
<p>Pause a Cloudfalre zone</p>

```bash
cloudflare zone setting pause [ --zone-id ]

#global flags       
                              [ --query | -q ]                           
                              [ --output |-o ]
                              [ --help | -h ]
                              [ --debug ]
```

#### cloudflare zone setting unpause<a name="cmd_cloudflare_zone_setting_unpause"></a>
<p>UnPause a Cloudfalre zone</p>

```bash
cloudflare zone setting unpause [ --zone-id]

#global flags       
                                [ --query | -q ]                           
                                [ --output |-o ]
                                [ --help | -h ]
                                [ --debug ]
```

#### cloudflare zone cache purge<a name="cmd_cloudflare_zone_cache_purge"></a>
<p>Control cache behaviors on Cloudflare zone</p>

```bash
cloudflare zone cache purge [ --zone-id]
                            [ --purge-everything | -A]

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
- Install Go 1.18.5
- Install Cobra for GO
- Install Viper

### Add sub commands <a name="addsubcommands"></a>
On the root folder (cloudflare-cli), type following command to generate new command

```bash
$> cobra add <command name> -p <parent command name> -a <author information>
```
### Run command locally
On the local environment, run following command to start cli

```bash
$> go run main.go <sub command> --help
```

### Build <a name="build"></a>

Following command running on MACOS M1 which will create a release version of cloudflare-cli, targeting for:
__Windows exe__
```
version="v0.x.x"
GOOS=windows GOARCH=amd64 go build -ldflags "-X cloudflare/cmd.version=$version" -o "release/windows-amd64-$version/cloudflare.exe" cloudflare
cd release/windows-amd64-$version
./cloudflare version -s
tar -cvf windows-amd64-$version.zip cloudflare.exe
```
__For Linux 64 binary__
```
version="v0.1.0"
GOOS=linux GOARCH=amd64 go build -ldflags "-X cloudflare/cmd.version=$version" -o "release/linux-amd64-$version/cloudflare" cloudflare
cd release/linux-amd64-$version
./cloudflare version -s
tar -czvf linux-amd64-$version.tar.gz cloudflare
```
__For MacOS binary__
```
version="v0.1.0"
GOOS=darwin GOARCH=arm64 go build -ldflags "-X cloudflare/cmd.version=$version" -o "release/darwin-arm64-$version/cloudflare" cloudflare
cd release/darwin-arm64-$version
./cloudflare version -s
tar -czvf darwin-arm64-$version.tar.gz cloudflare
```

### Sponsors
This repository is sponsored by [CloudCLI.io](https://www.cloudcli.io)
