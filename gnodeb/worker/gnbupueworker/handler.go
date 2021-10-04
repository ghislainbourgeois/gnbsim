// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package gnbupueworker

import (
	"fmt"
	"gnbsim/common"
	"gnbsim/gnodeb/context"
	"gnbsim/util/test"
)

func HandleUlMessage(gnbue *context.GnbUpUe, msg common.InterfaceMessage) (err error) {
	gnbue.Log.Traceln("Handling UL Packet from UE")

	userDataMsg := msg.(*common.UserDataMessage)
	encodedMsg, err := test.BuildGpduMessage(userDataMsg.Payload, gnbue.UlTeid)
	if err != nil {
		gnbue.Log.Errorln("BuildGpduMessage() returned:", err)
		return fmt.Errorf("failed to encode gpdu")
	}
	err = gnbue.Gnb.UpTransport.SendToPeer(gnbue.Upf, encodedMsg)
	if err != nil {
		gnbue.Log.Errorln("UP Transport SendToPeer() returned:", err)
		return fmt.Errorf("failed to send gpdu")
	}

	return nil
}

func HandleDlMessage(gnbue *context.GnbUpUe, msg common.InterfaceMessage) (err error) {
	gnbue.Log.Traceln("Handling DL Packet from UPF Worker")

	userDataMsg := msg.(*common.UserDataMessage)

	if len(userDataMsg.Payload) == 0 {
		return fmt.Errorf("empty t-pdu")
	}

	/* TODO: Parse QFI and check if it exists in GnbUpUe. In real world,
	   gNb may use the QFI to find a corresponding DRB
	*/

	gnbue.Log.Infoln("Forwarding DL user data packet to UE")
	userDataMsg.Event = common.DL_UE_DATA_TRANSFER_EVENT
	userDataMsg.Interface = common.UU_INTERFACE
	gnbue.WriteUeChan <- userDataMsg

	return nil
}
