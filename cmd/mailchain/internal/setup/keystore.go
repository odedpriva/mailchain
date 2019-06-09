// Copyright 2019 Finobo
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

package setup

import (
	"fmt"

	"github.com/mailchain/mailchain/cmd/mailchain/internal/config"
	"github.com/mailchain/mailchain/cmd/mailchain/internal/config/names"
	"github.com/mailchain/mailchain/cmd/mailchain/internal/prompts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper" // nolint: depguard
)

func Keystore(cmd *cobra.Command, keystoreType string) (string, error) {
	keystoreType, err := selectKeystore(keystoreType)
	if err != nil {
		return "", err
	}
	if err := config.SetKeystore(cmd, keystoreType); err != nil {
		return "", err
	}

	return keystoreType, nil
}

func selectKeystore(keystoreType string) (string, error) {
	if keystoreType != names.RequiresValue {
		return keystoreType, nil
	} 
	keystoreType, skipped, err := prompts.SelectItemSkipable(
		"Key Store",
		[]string{names.KeystoreNACLFilestore},
		viper.GetString("storage.keys") != "")
	if err != nil || skipped {
		return "", err
	}
	fmt.Printf("%q used for storing keys\n", keystoreType)
	return keystoreType, nil
}
