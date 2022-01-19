// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package common

import "gnbsim/logger"

type ProcedureType uint8

const (
	REGISTRATION_PROCEDURE ProcedureType = 1 + iota
	PDU_SESSION_ESTABLISHMENT_PROCEDURE
	USER_DATA_PKT_GENERATION_PROCEDURE
	UE_INITIATED_DEREGISTRATION_PROCEDURE
	AN_RELEASE_PROCEDURE
	UE_TRIGGERED_SERVICE_REQUEST_PROCEDURE
)

var procStrMap = map[ProcedureType]string{
	REGISTRATION_PROCEDURE:                 "REGISTRATION-PROCEDURE",
	PDU_SESSION_ESTABLISHMENT_PROCEDURE:    "PDU-SESSION-ESTABLISHMENT-PROCEDURE",
	USER_DATA_PKT_GENERATION_PROCEDURE:     "USER-DATA-PACKET-GENERATION-PROCEDURE",
	UE_INITIATED_DEREGISTRATION_PROCEDURE:  "UE-INITIATED-DEREGISTRATION-PROCEDURE",
	AN_RELEASE_PROCEDURE:                   "AN-RELEASE-PROCEDURE",
	UE_TRIGGERED_SERVICE_REQUEST_PROCEDURE: "UE-TRIGGERED-SERVICE-REQUEST-PROCEDURE",
}

func GetProcString(id ProcedureType) string {
	procStr, ok := procStrMap[id]
	if !ok {
		logger.AppLog.Fatalln("Invaid Procedure ID:", id)
	}
	return procStr
}
