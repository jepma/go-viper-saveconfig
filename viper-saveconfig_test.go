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

// Package vipersaveconfig is a versioning library that helps generate proper version numbers.
package vipersaveconfig

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
)

var workDir string = "/Workspace/playground/demo-repo"

func setDir() {
	// Set current work-dir
	os.Chdir(workDir)
}

func deleteFile(path string) {
	// delete file
	// var err = os.Remove(path)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(0)
	// }
}

func TestSaveUnknown(t *testing.T) {
	viper.SetConfigFile("demo.ini")
	viper.AddConfigPath(workDir) // adding home directory as first search path

	err := SaveConfig(*viper.GetViper())
	if err == nil {
		t.Error("Expected an error, got ", err)
	}
}

func TestSaveYaml(t *testing.T) {

	testFile := fmt.Sprintf("%s/%s", workDir, "demo.yaml")

	// Create file (TEST PURPOSES)
	f, err := os.Create(testFile)
	defer f.Close()
	if err != nil {
		os.Exit(1)
	}

	//////////

	v := viper.New()
	// v.AddConfigPath(workDir) // adding home directory as first search path
	v.SetConfigFile(testFile)
	v.Set("version", "1.0.0")
	SaveConfig(*v)

	//////////

	v = viper.New()
	// v.AddConfigPath(workDir) // adding home directory as first search path
	v.SetConfigFile(testFile)
	// Read Config
	if err := v.ReadInConfig(); err != nil {
		t.Error("Error ", err)
	}

	rValue := v.Get("version")
	if rValue != "1.0.0" {
		t.Error("Expected 1.0.0, got ", rValue)
	}

	deleteFile(testFile)
}

func TestSaveNewYaml(t *testing.T) {

	testFile := fmt.Sprintf("%s/%s", workDir, "demo.yaml")

	v := viper.New()
	v.SetConfigFile(testFile)
	v.Set("version", "6.0.0")

	if err := SaveConfig(*v); err != nil {
		t.Error("Error ", err)
	}

	//////////

	v = viper.New()
	// v.AddConfigPath(workDir) // adding home directory as first search path
	v.SetConfigFile(testFile)
	// Read Config
	if err := v.ReadInConfig(); err != nil {
		t.Error("Error ", err)
	}

	rValue := v.Get("version")
	if rValue != "6.0.0" {
		t.Error("Expected 6.0.0, got ", rValue)
	}

	deleteFile(testFile)

}

func TestSaveNewJson(t *testing.T) {

	testFile := fmt.Sprintf("%s/%s", workDir, "demo.json")

	v := viper.New()

	v.SetConfigFile(testFile)
	v.Set("version", "6.0.0")

	if err := SaveConfig(*v); err != nil {
		t.Error("Error ", err)
	}

	///////

	v = viper.New()
	// v.AddConfigPath(workDir) // adding home directory as first search path
	v.SetConfigFile(testFile)
	// Read Config
	if err := v.ReadInConfig(); err != nil {
		t.Error("Error ", err)
	}

	rValue := v.Get("version")
	if rValue != "6.0.0" {
		t.Error("Expected 6.0.0, got ", rValue)
	}

	deleteFile(testFile)

}
