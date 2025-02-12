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

package unique_test

import (
	"fmt"
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	. "go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestValidity(t *testing.T) {
	eui := types.EUI64{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
	for _, tc := range []ttnpb.IDStringer{
		nil,
		(*ttnpb.ApplicationIdentifiers)(nil),
		&ttnpb.ApplicationIdentifiers{},
		(*ttnpb.ClientIdentifiers)(nil),
		&ttnpb.ClientIdentifiers{},
		(*ttnpb.EndDeviceIdentifiers)(nil),
		&ttnpb.EndDeviceIdentifiers{},
		&ttnpb.EndDeviceIdentifiers{
			ApplicationIds: &ttnpb.ApplicationIdentifiers{ApplicationId: "foo"},
			DevEui:         eui.Bytes(),
		},
		&ttnpb.EndDeviceIdentifiers{
			ApplicationIds: &ttnpb.ApplicationIdentifiers{ApplicationId: "foo"},
			DevEui:         eui.Bytes(),
		},
		(*ttnpb.GatewayIdentifiers)(nil),
		&ttnpb.GatewayIdentifiers{},
		&ttnpb.GatewayIdentifiers{
			Eui: eui.Bytes(),
		},
		&ttnpb.GatewayIdentifiers{
			Eui: eui.Bytes(),
		},
		(*ttnpb.OrganizationIdentifiers)(nil),
		&ttnpb.OrganizationIdentifiers{},
		(*ttnpb.UserIdentifiers)(nil),
		&ttnpb.UserIdentifiers{},
	} {
		t.Run(fmt.Sprintf("%T", tc), func(t *testing.T) {
			a := assertions.New(t)
			a.So(func() { ID(test.Context(), tc) }, should.Panic)
		})
	}
}

func TestRoundtrip(t *testing.T) {
	for _, tc := range []struct {
		ID       ttnpb.IDStringer
		Expected string
		Parser   func(string) (ttnpb.IDStringer, error)
	}{
		{
			&ttnpb.ApplicationIdentifiers{ApplicationId: "foo"},
			"foo",
			func(uid string) (ttnpb.IDStringer, error) { return ToApplicationID(uid) },
		},
		{
			&ttnpb.ClientIdentifiers{ClientId: "foo"},
			"foo",
			func(uid string) (ttnpb.IDStringer, error) { return ToClientID(uid) },
		},
		{
			&ttnpb.EndDeviceIdentifiers{
				ApplicationIds: &ttnpb.ApplicationIdentifiers{
					ApplicationId: "foo-app",
				},
				DeviceId: "foo-device",
			},
			"foo-app.foo-device",
			func(uid string) (ttnpb.IDStringer, error) { return ToDeviceID(uid) },
		},
		{
			&ttnpb.GatewayIdentifiers{GatewayId: "foo"},
			"foo",
			func(uid string) (ttnpb.IDStringer, error) { return ToGatewayID(uid) },
		},
		{
			&ttnpb.OrganizationIdentifiers{OrganizationId: "foo"},
			"foo",
			func(uid string) (ttnpb.IDStringer, error) { return ToOrganizationID(uid) },
		},
		{
			&ttnpb.UserIdentifiers{UserId: "foo"},
			"foo",
			func(uid string) (ttnpb.IDStringer, error) { return ToUserID(uid) },
		},
	} {
		t.Run(fmt.Sprintf("%T", tc.ID), func(t *testing.T) {
			a := assertions.New(t)
			a.So(ID(test.Context(), tc.ID), should.Equal, tc.Expected)
			if id, ok := tc.ID.(interface {
				GetEntityIdentifiers() *ttnpb.EntityIdentifiers
			}); ok {
				wrapped := id.GetEntityIdentifiers()
				a.So(ID(test.Context(), wrapped), should.Equal, tc.Expected)
			}
			a.So(ID(test.Context(), tc.ID), should.Equal, tc.Expected)
			if tc.Parser != nil {
				parsed, err := tc.Parser(tc.Expected)
				if a.So(err, should.BeNil) {
					a.So(parsed.EntityType(), should.Resemble, tc.ID.EntityType())
					a.So(parsed.IDString(), should.Resemble, tc.ID.IDString())
				}
			}
		})
	}
}

func TestValidatorForIdentifiers(t *testing.T) {
	a := assertions.New(t)
	for _, run := range []struct {
		Name   string
		ID     func(string) ttnpb.IDStringer
		Parser func(string) (ttnpb.IDStringer, error)
	}{
		{
			"ApplicationID",
			func(uid string) ttnpb.IDStringer { return &ttnpb.ApplicationIdentifiers{ApplicationId: uid} },
			func(uid string) (ttnpb.IDStringer, error) { return ToApplicationID(uid) },
		},
		{
			"ClientId",
			func(uid string) ttnpb.IDStringer { return &ttnpb.ClientIdentifiers{ClientId: uid} },
			func(uid string) (ttnpb.IDStringer, error) { return ToClientID(uid) },
		},
		{
			"GatewayID",
			func(uid string) ttnpb.IDStringer { return &ttnpb.GatewayIdentifiers{GatewayId: uid} },
			func(uid string) (ttnpb.IDStringer, error) { return ToGatewayID(uid) },
		},
		{
			"OrganizationID",
			func(uid string) ttnpb.IDStringer { return &ttnpb.OrganizationIdentifiers{OrganizationId: uid} },
			func(uid string) (ttnpb.IDStringer, error) { return ToOrganizationID(uid) },
		},
		{
			"UserID",
			func(uid string) ttnpb.IDStringer { return &ttnpb.UserIdentifiers{UserId: uid} },
			func(uid string) (ttnpb.IDStringer, error) { return ToUserID(uid) },
		},
	} {
		for _, tc := range []struct {
			Name          string
			InputUID      string
			ExpectedError func(error) bool
		}{
			{
				"ValidID",
				"test",
				nil,
			},
			{
				"ValidMinLength",
				"oza",
				nil,
			},
			{
				"ValidMaxLength",
				"ozaj8qs0sait7oudxqbfyx6b14yuahcfrdlb",
				nil,
			},
			{
				"ValidNumerics",
				"1id1",
				nil,
			},
			{
				"InvalidEmptyString",
				"",
				errors.IsInvalidArgument,
			},
			{
				"InvalidDashes",
				"-id",
				errors.IsInvalidArgument,
			},
			{
				"InvalidDashes1",
				"id-",
				errors.IsInvalidArgument,
			},
			{
				"InvalidUnderscore",
				"id_test",
				errors.IsInvalidArgument,
			},
			{
				"InvalidDot",
				"id.test",
				errors.IsInvalidArgument,
			},
			{
				"InvalidMinLength",
				"id",
				func(err error) bool {
					if run.Name == "UserID" {
						return err == nil
					}
					return errors.IsInvalidArgument(err)
				},
			},
			{
				"InvalidMaxLength",
				"ozaj8qs0sait7oudxqbfyx6b14yuahcfrdlbh",
				errors.IsInvalidArgument,
			},
		} {
			t.Run(fmt.Sprintf("%s/%s", run.Name, tc.Name), func(t *testing.T) {
				if run.Parser == nil {
					t.Fatal("Parser Not Defined")
				}
				parsed, err := run.Parser(tc.InputUID)
				if tc.ExpectedError == nil {
					a.So(err, should.BeNil)
					if !a.So(parsed, should.Resemble, run.ID(tc.InputUID)) {
						t.FailNow()
					}
				} else {
					if !a.So(tc.ExpectedError(err), should.BeTrue) {
						t.FailNow()
					}
				}
			})
		}
	}
}

func TestValidatorForDeviceIDs(t *testing.T) {
	a := assertions.New(t)
	for _, tc := range []struct {
		Name               string
		InputUID           string
		ExpectedIdentifier ttnpb.EndDeviceIdentifiers
		ExpectedError      func(error) bool
	}{
		{
			"ValidID",
			"foo-app.foo-device",
			ttnpb.EndDeviceIdentifiers{
				ApplicationIds: &ttnpb.ApplicationIdentifiers{
					ApplicationId: "foo-app",
				},
				DeviceId: "foo-device",
			},
			nil,
		},
		{
			"ValidAppIDValidMinLength",
			"foo-app.foo",
			ttnpb.EndDeviceIdentifiers{
				ApplicationIds: &ttnpb.ApplicationIdentifiers{
					ApplicationId: "foo-app",
				},
				DeviceId: "foo",
			},
			nil,
		},
		{
			"ValidAppIDValidMaxLength",
			"foo-app.ozaj8qs0sait7oudxqbfyx6b14yuahcfrdlb",
			ttnpb.EndDeviceIdentifiers{
				ApplicationIds: &ttnpb.ApplicationIdentifiers{
					ApplicationId: "foo-app",
				},
				DeviceId: "ozaj8qs0sait7oudxqbfyx6b14yuahcfrdlb",
			},
			nil,
		},
		{
			"ValidAppIDValidNumerics",
			"foo-app.1d1",
			ttnpb.EndDeviceIdentifiers{
				ApplicationIds: &ttnpb.ApplicationIdentifiers{
					ApplicationId: "foo-app",
				},
				DeviceId: "1d1",
			},
			nil,
		},
		{
			"InvalidFormat",
			"foo-appfoo-device",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
		{
			"InvalidAppID",
			"foo_app.foo-device",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
		{
			"ValidAppIDInvalidEmptyString",
			"foo-app.",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
		{
			"ValidAppIDInvalidEmptyString",
			"foo-app.",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
		{
			"ValidAppIDInvalidDashes",
			"foo-app.-foo",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
		{
			"ValidAppIDInvalidDashes1",
			"foo-app.foo-",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
		{
			"ValidAppIDInvalidUnderscore",
			"foo-app.foo_device",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
		{
			"ValidAppIDInvalidDot",
			"foo-app.foo.device",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
		{
			"ValidAppIDInvalidMinLength",
			"foo-app.id",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
		{
			"ValidAppIDInvalidMaxLength",
			"foo-app.ozaj8qs0sait7oudxqbfyx6b14yuahcfrdlbh",
			ttnpb.EndDeviceIdentifiers{},
			errors.IsInvalidArgument,
		},
	} {
		t.Run(fmt.Sprintf("%s", tc.Name), func(t *testing.T) {
			devIDs, err := ToDeviceID(tc.InputUID)
			if tc.ExpectedError == nil {
				if !a.So(err, should.BeNil) {
					t.FailNow()
				}
				if !a.So(devIDs, should.Resemble, tc.ExpectedIdentifier) {
					t.FailNow()
				}
			} else {
				if !a.So(tc.ExpectedError(err), should.BeTrue) {
					t.FailNow()
				}
			}
		})
	}
}
