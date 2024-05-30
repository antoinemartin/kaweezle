/*
Copyright © 2022 Antoine Martin <antoine@openance.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/kaweezle/kaweezle/pkg/logger"
	"github.com/kaweezle/kaweezle/pkg/rootfs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var UpdateRootFSFields = log.Fields{
	logger.TaskKey: "Update Root FS",
}

func NewUpdateCommand() *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update the root file system",
		Long:  `Check and download the last version of the file system.`,
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(rootfs.EnsureRootFS(rootfs.TarFilePath, &log.Fields{}))
		},
	}
	updateCmd.Flags().StringVarP(&rootfs.TarFilePath, "root", "r", rootfs.DefaultTarFilePath, "The root file system to install")

	return updateCmd
}
