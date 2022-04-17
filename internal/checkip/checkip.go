package checkip

import (
	"github.com/lekan-pvp/short/internal/config"
	"net/http"
	"net/netip"
)

func CheckIP(r *http.Request) (bool, error) {
	network, err := netip.ParsePrefix(config.Cfg.TrustedSubnet)
	if err != nil {
		return false, err
	}

	ip, err := netip.ParseAddr(r.Header.Get("X-Real-IP"))
	if err != nil {
		return false, err
	}

	return network.Contains(ip), nil
}
