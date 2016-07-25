// Copyright 2016 Andrew O'Neill

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package choices

import "time"

var config = struct {
	globalSalt     string
	storage        Storage
	updateInterval time.Duration
}{
	globalSalt:     "choices",
	updateInterval: 5 * time.Minute,
}

// SetGlobalSalt sets the salt used in hashing users.
func SetGlobalSalt(salt string) {
	config.globalSalt = salt
}

// SetUpdateInterval changes the update interval for Storage. Must call
// SetStorage after this or cancel context of the current Storage and call
// SetStorage again.
func SetUpdateInterval(dur time.Duration) {
	config.updateInterval = dur
}
