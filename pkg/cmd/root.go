/*
Copyright © 2023 Kubernetes Authors

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
	"github.com/spf13/cobra"
)

// GetRootCommand returns the root cobra command to be executed
// by main.
func GetRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingress2gateway",
		Short: "ingress2gateway converts Ingress to Gateway API",
	}

	cmd.AddCommand(RegisterTranslateCommand())
	cmd.AddCommand(RegisterVersionCommand())

	return cmd
}
