package text

const EmptyArgsText = `Error: please use --help|-h to explore cloudflare-cli commands`

const RootLongText = `A compact CLI works with Cloudflare REST API at https://api.cloudflare.com/v4

Author: Hai Tran (hidetran@gmail.com)
Contributed at https://github.com/epiHATR/cloudflare-cli
`

const AdditionalText = `A compact CLI works with Cloudflare REST API at https://api.cloudflare.com/v4

Author: Hai Tran (hidetran@gmail.com)
Contributed at https://github.com/epiHATR/cloudflare-cli	

Usages:
	cloudflare version	get cloudflare-cli module version
	cloudflare login	login to Cloudflare REST API

Flags:
	--help	display command help & instructions`

const SubCmdHelpText = `

Explore cloudflare-cli commands at https://github.com/epiHATR/cloudflare-cli`

const CmdLoginLongText = `This command let you authenticate against Cloudflare REST API and store credential to local file
	
Usages
	cloudflare login --email <your cloudflare email> --key <your cloudflare api key>`

const CmdLoginNoCredentialText = `Error: no credentials provided, consider using flags or environment variables

Usages:
	cloudflare login --email <your cloudflare email> --key <your cloudflare api key>
	cloudflare login --token <your cloudflare api token>`

const ZoneCmdLongText = `Cloudflare API for zone help you create, update, list your zone via API.
Make sure your login credentials has permission enough to retrieve those resources.`

const ZoneListCmdLongText = `List all Cloudflare zones under your accounts.
This command requires 'Zone.Read' on your Cloudflare token API configuration or you have to 'Zone.List' permission`

const ZoneShowCmdLongText = `Get Cloudflare zone details by Id or Name.
This command requires 'Zone.Read' on your Cloudflare token API configuration permission`

const ZoneDNSListCmdLongText = `Get all DNS records in a specified Cloudflare zone.
This command requires 'Zone.Read' & 'DNS.Read' on your Cloudlfare token API configuration permission`

var ZoneDnsLongText = `Manage your DNS records in the Cloudflare zone
This command requires 'DNS.Read' & 'DNS.Create' on your Cloudflare token API configuration permission`

var ZoneDnsShowLongText = `Show Cloudflare DNS record details
This command requires 'DNS.Read' & 'DNS.Create' on your Cloudflare token API configuration permission`

var ZoneDNSAddCmdLongText = `Add new DNS record to Cloudflare zone.
This command requires 'DNS.Read' & 'DNS.Create' on your Cloudflare token API configuration permission.
See Cloudflare API documentation at https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record
`
var ZoneSettingLongText = `Manage your Cloudaflare zone's settings
This command requires 'Zone.Read' & 'DNS.Create' on your Cloudflare token API configuration permission`

var PlancmdLongText = `Manage Cloudflare zone's plan
This command requires 'Billing.Read' on your Cloudflare token API configuration permission`

var AvailablePlanLongText = `List available plans the zone can subscribe to
This command requires 'Billing.Read' on your Cloudflare token API configuration permission`

var PlanDetailsLongText = `Show Cloudflare plan details
This command requires 'Billing.Read' on your Cloudflare token API configuration permission`

var SetTypeLongText = `Change cloudflare zone type
This command requires 'Zone.Write' on your Cloudflare token API configuration permission`

var UpdatePlanLongtext = `Update Cloudflare zone's plan
This command requires 'Billing.Write' on your Cloudflare token API configuration permission`

var ZoneDndDelete = `Delete a DNS record from Cloudflare zone
This command requires 'DNS.Read' & 'DNS.Write' on your Cloudflare token API configuration permission`

var ZoneDnsUpdate = `Update a DNS record on Cloudflare zone
This command requires 'DNS.Write' & 'DNS.Write' on your Cloudflare token API configuration permission`

var ZoneCreateLongText = `Create a new Cloudflare zone
This command requires 'Zone.Read' & 'Zone.Write' on your Cloudflare token API configuration permission`

var DeleteZoneLongText = `Delete a Cloudflare zone.
Zone should be changed to Free Website type before removed from Cloudflare.
This command requires 'Zone.Read' & 'Zone.Write' on your Cloudflare token API configuration permission`

var AccountCmdLongText = `Manage Cloudflare accounts/organizations
This command requires 'organization:read,organization:write' on your Cloudflare token API configuration permission`

var AccountListLongText = `List all Cloudflare accounts/organizations
This command requires 'organization:read' on your Cloudflare token API configuration permission`

var AccountShowLongText = `Show details of a Cloudflare account/organization
This command requires 'organization:read' on your Cloudflare token API configuration permission`

var AccountUpdateLongText = `Update information of a Cloudflare account/organization
This command requires 'organization:write' on your Cloudflare token API configuration permission`

var AccountUserLongText = `Manage users of a Cloudflare account/organization
This command requires 'organization:read', 'organization:write' on your Cloudflare token API configuration permission`

var AccountUserDetailsLongText = `Show user's information details
This command requires 'organization:read', 'organization:write' on your Cloudflare token API configuration permission`

var AccountRoleLongText = `Manage Cloudflare account/organization roles
This command requires 'organization:read', 'organization:write' on your Cloudflare token API configuration permission`

var AccountRoleListLongText = `List all available role in Cloudflare account/organization
This command requires 'organization:read', 'organization:write' on your Cloudflare token API configuration permission`

var AccountRoleDetailsLongText = `Manage Cloudflare account/organization roles
This command requires 'organization:read', 'organization:write' on your Cloudflare token API configuration permission`

var AccountUserAddLongText = `Add new user to a Cloudflare account/organization
This command requires 'organization:read', 'organization:write' on your Cloudflare token API configuration permission`
