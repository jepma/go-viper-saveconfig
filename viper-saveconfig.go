// Copyright © 2017 Marcel Jepma <mjepma@xebia.com>
// This file is part of release-support.

package vipersaveconfig

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
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
			glog.Fatal("Panic while encoding into JSON format.")
		}
		f.WriteString(string(b))

	case "toml":

		w := bufio.NewWriter(f)
		if err := toml.NewEncoder(w).Encode(v.AllSettings()); err != nil {
			glog.Fatal("Panic while encoding into TOML format.")
		}
		w.Flush()

	case "yaml", "yml":

		b, err := yaml.Marshal(v.AllSettings())
		if err != nil {
			glog.Fatal("Panic while encoding into YAML format.")
		}
		f.WriteString(string(b))
	}

	return nil
}
