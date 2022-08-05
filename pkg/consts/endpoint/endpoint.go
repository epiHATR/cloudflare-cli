package endpoint

var ApiEndPoint = "https://api.cloudflare.com/client/v4"

var TokenVerifyEndPoint = "/user/tokens/verify"
var EmailKeyVerifyEndPoint = "/user"

var ZoneListEndPoint = "/zones?per_page=50&direction=desc&match=all"
var ZoneDetailEndPoint = "/zones"

var ZoneDnsListEndpoint = "/zones/%s/dns_records?order=type&direction=desc&match=all"

var ZoneDetails = "/zones/%s/dns_records"
