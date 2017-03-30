// Copyright (c) 2017 Xebia Nederland B.V. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package vipersaveconfig is an addition to the Viper library

package vipersaveconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// SupportedExts Universally supported extensions.
var SupportedExts []string = []string{"json", "yaml", "yml"}

// UnsupportedConfigError configuration filetype.
type UnsupportedConfigError string

// Returns the formatted configuration error.
func (str UnsupportedConfigError) Error() string {
	return fmt.Sprintf("Unsupported Config Type %q", string(str))
}

func getConfigType(v viper.Viper) string {

	cf := v.ConfigFileUsed()
	ext := filepath.Ext(cf)

	if len(ext) > 1 {
		return ext[1:]
	} else {
		return ""
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// SaveConfig will store the settings from viper inside the config-file
func SaveConfig(v viper.Viper) error {

	// Verify if format is accepted
	if !stringInSlice(getConfigType(v), SupportedExts) {
		return UnsupportedConfigError(getConfigType(v))
	}

	// Open file
	f, err := os.Create(v.ConfigFileUsed())
	defer f.Close()
	if err != nil {
		return err
	}

	// Switch between formats
	switch getConfigType(v) {
	case "json":

		b, err := json.MarshalIndent(v.AllSettings(), "", "    ")
		if err != nil {
			return err
		}
		f.WriteString(string(b))

	case "yaml", "yml":

		b, err := yaml.Marshal(v.AllSettings())
		if err != nil {
			return err
		}
		f.WriteString(string(b))
	}

	return nil
}
