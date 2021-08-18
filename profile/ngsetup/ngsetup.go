// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package ngsetup

import (
	"fmt"
	"gnbsim/gnodeb"
	"gnbsim/gnodeb/context"
	"net"
)

func NgSetup_test(gnb *context.GNodeB) {
	// create amf
	addrs, err := net.LookupHost("amf")
	if err != nil {
		fmt.Println("Failed to resolve amf")
		return
	}
	gnbamf := context.NewGnbAmf(addrs[0], 38412)

	err = gnodeb.ConnectToAmf(gnb, gnbamf)
	if err != nil {
		fmt.Println("ConnectToAmf() failed due to:", err)
		return
	}

	successFulOutcome, err := gnodeb.PerformNgSetup(gnb, gnbamf)
	if err != nil {
		fmt.Println("PerformNGSetup() failed due to:", err)
	} else if !successFulOutcome {
		fmt.Println("Expected SuccessfulOutcome, received UnsuccessfulOutcome")
		return
	}

	fmt.Println("NGSetup Procedure successful")
}
