// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package gogoproto_test

import (
	"fmt"
	"testing"

	pbtypes "github.com/gogo/protobuf/types"
	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/gogoproto"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func ExampleGoFieldsPaths() {
	type cityDetails struct {
		Name string `protobuf:"name=name_city"`
	}

	type place struct {
		NameOfTheRegion string `protobuf:"name=name_region"`

		CityDetails cityDetails `protobuf:"bytes,name=city"`
	}

	london := place{
		CityDetails: cityDetails{Name: "London"},
	}
	holland := place{
		NameOfTheRegion: "Holland",
	}

	fields := gogoproto.GoFieldsPaths(&pbtypes.FieldMask{
		Paths: []string{"city.name_city"},
	}, london)
	fmt.Println(fields)

	fields = gogoproto.GoFieldsPaths(&pbtypes.FieldMask{
		Paths: []string{"name_region"},
	}, holland)
	fmt.Println(fields)

	// Output: [CityDetails.Name]
	// [NameOfTheRegion]
}

func TestGoFieldsPaths(t *testing.T) {
	a := assertions.New(t)

	type cityDetails struct {
		Name string `protobuf:"name=name_city"`
	}

	type hasProtoRenaming struct {
		NameOfTheRegion string `protobuf:"name=name_region"`

		CityDetails cityDetails `protobuf:"bytes,name=city"`
	}

	for _, tc := range []struct {
		fields, expected []string
	}{
		{
			fields:   []string{"name_region", "name_city"},
			expected: []string{"NameOfTheRegion", "name_city"},
		},
		{
			fields:   []string{"name_region"},
			expected: []string{"NameOfTheRegion"},
		},
		{
			fields:   []string{"city.name_city"},
			expected: []string{"CityDetails.Name"},
		},
	} {
		goFields := gogoproto.GoFieldsPaths(&pbtypes.FieldMask{Paths: tc.fields}, hasProtoRenaming{
			NameOfTheRegion: "england",
			CityDetails: cityDetails{
				Name: "london",
			},
		})

		a.So(goFields, should.HaveSameElementsDeep, tc.expected)
	}
}

func TestGoFieldsPathsEndDevice(t *testing.T) {
	a := assertions.New(t)

	for _, tc := range []struct {
		fields, expected []string
	}{
		{
			fields:   []string{"mac_state", "frequency_plan_id"},
			expected: []string{"MacState", "FrequencyPlanId"},
		},
		{
			fields:   []string{"session.last_f_cnt_up"},
			expected: []string{"Session.LastFCntUp"},
		},
		{
			fields:   []string{"ids.application_ids.application_id"},
			expected: []string{"EndDeviceIdentifiers.ApplicationIds.ApplicationId"},
		},
	} {
		goFields := gogoproto.GoFieldsPaths(&pbtypes.FieldMask{Paths: tc.fields}, ttnpb.EndDevice{
			EndDeviceIdentifiers: ttnpb.EndDeviceIdentifiers{
				ApplicationIds: &ttnpb.ApplicationIdentifiers{},
			},
			Session: &ttnpb.Session{
				LastFCntUp: 5,
			},
		})

		a.So(goFields, should.HaveSameElementsDeep, tc.expected)
	}
}
