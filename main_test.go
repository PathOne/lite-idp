// Copyright © 2017 Aaron Donovan <PathOne@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_initConfig(t *testing.T) {
	// make sure environment replacer is setup properly
	viper.SetDefault("my-config-value", "a")
	initConfig()
	assert.Equal(t, "a", viper.Get("my-config-value"), "initial value is wrong")
	os.Setenv("MY_CONFIG_VALUE", "b")
	assert.Equal(t, "b", viper.Get("my-config-value"), "second value is wrong")
}
