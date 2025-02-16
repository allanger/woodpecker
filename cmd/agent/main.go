// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"

	"go.woodpecker-ci.org/woodpecker/v2/cmd/common"
	"go.woodpecker-ci.org/woodpecker/v2/pipeline/backend/docker"
	"go.woodpecker-ci.org/woodpecker/v2/pipeline/backend/kubernetes"
	"go.woodpecker-ci.org/woodpecker/v2/pipeline/backend/local"
	"go.woodpecker-ci.org/woodpecker/v2/shared/utils"
	"go.woodpecker-ci.org/woodpecker/v2/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "woodpecker-agent"
	app.Version = version.String()
	app.Usage = "woodpecker agent"
	app.Action = runWithRetry
	app.Commands = []*cli.Command{
		{
			Name:   "ping",
			Usage:  "ping the agent",
			Action: pinger,
		},
	}
	app.Flags = utils.MergeSlices(flags, common.GlobalLoggerFlags, docker.Flags, kubernetes.Flags, local.Flags)

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
