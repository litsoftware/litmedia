package helpers

import (
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func ValidateIP(ipaddr string) bool {
	addr := net.ParseIP(ipaddr)
	return addr != nil
}

func ValidateHostname(hostname string) bool {
	matches, _ := regexp.MatchString(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`, hostname)

	if !matches {
		matchesDocker, _ := regexp.MatchString(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])\_([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])$`, hostname)
		if !matchesDocker {
			return ValidateIP(hostname)
		}
		return true
	}

	return matches
}

func ValidateEmail(email string) bool {
	matches, _ := regexp.MatchString(`^.+@.+\..+$`, email)
	return matches
}

func ValidateURL(rawURL string) bool {
	_, err := url.Parse(rawURL)
	return err == nil
}

func ValidateHostList(hosts []string) bool {
	for _, host := range hosts {
		if !ValidateHostPort(host, false) {
			return false
		}
	}

	return true
}

func ValidateHostPort(host string, allowBlankHost bool) bool {
	// Must be hostname:port, ipv4:port, or [ipv6]:port. Optionally allow blank hostname
	hostname, portString, err := net.SplitHostPort(host)
	if err != nil {
		return false
	}

	// Validate the port is a numeric (yeah, strings are valid in some places, but we don't support it)
	_, err = strconv.Atoi(portString)
	if err != nil {
		return false
	}

	// Listeners can have blank hostnames, so we'll skip validation if that's what we're looking for
	if allowBlankHost && hostname == "" {
		return true
	}

	// Only IPv6 can contain :
	if strings.Contains(hostname, ":") && (!ValidateIP(hostname)) {
		return false
	}

	// If all the parts of the hostname are numbers, validate as IP. Otherwise, it's a hostname
	hostnameParts := strings.Split(hostname, ".")
	isIP4 := true
	for _, section := range hostnameParts {
		_, err := strconv.Atoi(section)
		if err != nil {
			isIP4 = false
			break
		}
	}
	if isIP4 {
		return ValidateIP(hostname)
	}

	return ValidateHostname(hostname)
}
