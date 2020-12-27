package fn

// GetSiteURL : returns protocol + domain
func GetSiteURL(protocol, domain string) string {
	return protocol + "://" + domain
}
