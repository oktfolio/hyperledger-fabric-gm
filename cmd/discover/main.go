/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"os"

	"github.com/oktfolio/hyperledger-fabric-gm/bccsp/factory"
	"github.com/oktfolio/hyperledger-fabric-gm/cmd/common"
	"github.com/oktfolio/hyperledger-fabric-gm/discovery/cmd"
)

func main() {
	factory.InitFactories(nil)
	cli := common.NewCLI("discover", "Command line client for fabric discovery service")
	discovery.AddCommands(cli)
	cli.Run(os.Args[1:])
}
