package endpoint

var ApiEndPoint = "https://api.cloudflare.com/client/v4"

// authenticate endpoints
var TokenVerifyEndPoint = "/user/tokens/verify"
var EmailKeyVerifyEndPoint = "/user"

// zone endpoints
var ZoneListEndPoint = "/zones?per_page=50&direction=desc&match=all"
var ZoneDetailEndPoint = "/zones"
var CreateZoneEndpoint = "/zones"
var DeleteZoneEndpoint = "/zones"
var ZoneSettingEndPoint = "/zones"
var ZoneDnsListEndpoint = "/zones/%s/dns_records?order=type&direction=desc&match=all"
var ZoneDetails = "/zones/%s/dns_records"

// organization endpoints
var AvailablePlanEndpoint = "/zones/%s/available_plans"
var AvailableRatePlanEndpoint = "/zones/%s/available_rate_plans"
