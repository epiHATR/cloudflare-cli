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
This commands requires 'Zone.Read' & 'DNS.Read' on your Cloudlfare token API configuration permission`
