# MAC Address Generator

[![codecov](https://codecov.io/gh/yitsushi/macpot/branch/main/graph/badge.svg?token=TZC7E8RA6B)](https://codecov.io/gh/yitsushi/macpot)
[![Quality Check](https://github.com/yitsushi/macpot/actions/workflows/quality-check.yaml/badge.svg)](https://github.com/yitsushi/macpot/actions/workflows/quality-check.yaml)

MacPot is an easy to use MAC address generator.

```go
package macpot

import (
	"fmt"

  "github.com/yitsushi/macpot"
)

func info(mac *MAC) {
	fmt.Printf("Address: %s\n", mac.ToString())

	if mac.IsLocal() {
		fmt.Println(" - Locally Administered")
	} else {
		fmt.Println(" - Globally Unique")
	}

	if mac.IsMulticast() {
		fmt.Println(" - Mulicast")
	} else {
		fmt.Println(" - Unicast")
	}
}

func main() {
	mac, _ := macpot.New(AsUnicast(), AsLocal())
	info(&mac)

	// User error if they don't use the correct order.
	mac, _ = macpot.New(AsUnicast(), AsLocal(), WithOUI("11:22:33"))
	info(&mac)

	mac, _ = macpot.New(WithOUI("11:22:33"), AsUnicast(), AsLocal())
	info(&mac)

	mac, _ = macpot.New(
		WithOUI("11:22:33"),
		WithNIC("44:55:66"),
		AsUnicast(),
		AsLocal(),
	)
	info(&mac)

	mac, _ = macpot.New(
		WithOUI("11:22:33"),
		WithNICFromIPv4("192.168.31.7"),
		AsUnicast(),
		AsLocal(),
	)
	info(&mac)
	fmt.Printf("Manual conversion: %x:%x:%x\n", 168, 31, 7)
}
```

Output:
```
Address: 32:78:ea:36:b5:f3
 - Locally Administered
 - Unicast
Address: 11:22:33:5a:3b:09
 - Globally Unique
 - Mulicast
Address: 12:22:33:70:71:29
 - Locally Administered
 - Unicast
Address: 12:22:33:44:55:66
 - Locally Administered
 - Unicast
Address: 12:22:33:a8:1f:07
 - Locally Administered
 - Unicast
Manual conversion: a8:1f:7
```
