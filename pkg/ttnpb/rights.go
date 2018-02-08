// Copyright © 2018 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package ttnpb

import (
	"sort"
	"strconv"
	"strings"

	"github.com/TheThingsNetwork/ttn/pkg/errors"
	"github.com/gogo/protobuf/jsonpb"
)

var (
	userRights         []Right
	applicationRights  []Right
	gatewayRights      []Right
	organizationRights []Right
)

func init() {
	for k, v := range Right_value {
		switch strings.SplitN(k, "_", 3)[1] {
		case "USER":
			userRights = append(userRights, Right(v))
		case "APPLICATION":
			applicationRights = append(applicationRights, Right(v))
		case "GATEWAY":
			gatewayRights = append(gatewayRights, Right(v))
		case "ORGANIZATION":
			organizationRights = append(organizationRights, Right(v))
		}
	}

	sort.Slice(userRights, func(i, j int) bool { return userRights[i] < userRights[j] })
	sort.Slice(applicationRights, func(i, j int) bool { return applicationRights[i] < applicationRights[j] })
	sort.Slice(gatewayRights, func(i, j int) bool { return gatewayRights[i] < gatewayRights[j] })
	sort.Slice(organizationRights, func(i, j int) bool { return organizationRights[i] < organizationRights[j] })
}

// AllUserRights is the set thart contains all the rights that are to users.
func AllUserRights() []Right { return userRights }

// AllApplicationRights is the set that contains all the rights that are to applications.
func AllApplicationRights() []Right { return applicationRights }

// AllGatewayRights is the set that contains all the rights that are to gateways.
func AllGatewayRights() []Right { return gatewayRights }

// AllOrganizationRights is the set that contains all the rights that are to organizations.
func AllOrganizationRights() []Right { return organizationRights }

// ParseRight parses the string specified into a Right.
func ParseRight(str string) (Right, error) {
	val, ok := Right_value[str]
	if !ok {
		return Right(0), errors.Errorf("Could not parse right `%s`", str)
	}
	return Right(val), nil
}

// MarshalText implements encoding.TextMarshaler interface.
func (r Right) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

// MarshalJSON implements json.Marshaler interface.
func (r Right) MarshalJSON() ([]byte, error) {
	txt, err := r.MarshalText()
	if err != nil {
		return nil, err
	}
	return []byte("\"" + string(txt) + "\""), nil
}

// MarshalJSONPB implements jsonpb.JSONPBMarshaler interface.
func (r Right) MarshalJSONPB(m *jsonpb.Marshaler) ([]byte, error) {
	if m.EnumsAsInts {
		return []byte("\"" + strconv.Itoa(int(r)) + "\""), nil
	}
	return r.MarshalJSON()
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (r *Right) UnmarshalText(b []byte) (err error) {
	*r, err = ParseRight(string(b))
	return
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (r *Right) UnmarshalJSON(b []byte) error {
	return r.UnmarshalText(b[1 : len(b)-1])
}

// UnmarshalJSONPB implements jsonpb.JSONPBUnmarshaler interface.
func (r *Right) UnmarshalJSONPB(m *jsonpb.Unmarshaler, b []byte) error {
	return r.UnmarshalJSON(b)
}
