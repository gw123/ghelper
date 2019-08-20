package reg_util

import "testing"

func TestIsIP2(t *testing.T) {
	ips := []string{"1.1.1.1", "0.0.0.0", "112.112.112.112", "255.255.1.1"}

	for _, ip := range ips {
		if IsIP(ip) != true {
			t.Fail()
		}
	}

	notIps := []string{"1.1.1.1.1", "0.0.0.", "112.112.112.1122", "2552.255.1.1",
		"90.1112.22.2","90.1112.22.21","1a.111.22.2",
	}

	for _, ip := range notIps {
		if IsIP(ip) == true {
			t.Fail()
		}
	}
}
