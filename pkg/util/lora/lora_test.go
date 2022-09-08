// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

package lora_test

import (
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/lora"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestAdjustedRSSI(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)
	a.So(lora.AdjustedRSSI(0, 0), should.AlmostEqual, float32(-10.0/3.0))
	a.So(lora.AdjustedRSSI(-46, 7), should.Equal, -47.0)
	a.So(lora.AdjustedRSSI(-46, 11), should.Equal, -46.0)
	a.So(lora.AdjustedRSSI(-46, -6), should.Equal, -52.0)
}
