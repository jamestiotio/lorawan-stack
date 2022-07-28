// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.4.0
// - protoc             v3.9.1
// source: lorawan-stack/api/keys.proto

package ttnpb

import (
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// MarshalProtoJSON marshals the KeyEnvelope message to JSON.
func (x *KeyEnvelope) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Key) > 0 || s.HasField("key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("key")
		types.MarshalHEXBytes(s.WithField("key"), x.Key)
	}
	if x.KekLabel != "" || s.HasField("kek_label") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("kek_label")
		s.WriteString(x.KekLabel)
	}
	if len(x.EncryptedKey) > 0 || s.HasField("encrypted_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("encrypted_key")
		s.WriteBytes(x.EncryptedKey)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the KeyEnvelope to JSON.
func (x *KeyEnvelope) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the KeyEnvelope message from JSON.
func (x *KeyEnvelope) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "key":
			s.AddField("key")
			x.Key = types.Unmarshal16Bytes(s.WithField("key", false))
		case "kek_label", "kekLabel":
			s.AddField("kek_label")
			x.KekLabel = s.ReadString()
		case "encrypted_key", "encryptedKey":
			s.AddField("encrypted_key")
			x.EncryptedKey = s.ReadBytes()
		}
	})
}

// UnmarshalJSON unmarshals the KeyEnvelope from JSON.
func (x *KeyEnvelope) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the RootKeys message to JSON.
func (x *RootKeys) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.RootKeyId != "" || s.HasField("root_key_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("root_key_id")
		s.WriteString(x.RootKeyId)
	}
	if x.AppKey != nil || s.HasField("app_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("app_key")
		x.AppKey.MarshalProtoJSON(s.WithField("app_key"))
	}
	if x.NwkKey != nil || s.HasField("nwk_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("nwk_key")
		x.NwkKey.MarshalProtoJSON(s.WithField("nwk_key"))
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the RootKeys to JSON.
func (x *RootKeys) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the RootKeys message from JSON.
func (x *RootKeys) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "root_key_id", "rootKeyId":
			s.AddField("root_key_id")
			x.RootKeyId = s.ReadString()
		case "app_key", "appKey":
			if s.ReadNil() {
				x.AppKey = nil
				return
			}
			x.AppKey = &KeyEnvelope{}
			x.AppKey.UnmarshalProtoJSON(s.WithField("app_key", true))
		case "nwk_key", "nwkKey":
			if s.ReadNil() {
				x.NwkKey = nil
				return
			}
			x.NwkKey = &KeyEnvelope{}
			x.NwkKey.UnmarshalProtoJSON(s.WithField("nwk_key", true))
		}
	})
}

// UnmarshalJSON unmarshals the RootKeys from JSON.
func (x *RootKeys) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the SessionKeys message to JSON.
func (x *SessionKeys) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.SessionKeyId) > 0 || s.HasField("session_key_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("session_key_id")
		s.WriteBytes(x.SessionKeyId)
	}
	if x.FNwkSIntKey != nil || s.HasField("f_nwk_s_int_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("f_nwk_s_int_key")
		x.FNwkSIntKey.MarshalProtoJSON(s.WithField("f_nwk_s_int_key"))
	}
	if x.SNwkSIntKey != nil || s.HasField("s_nwk_s_int_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("s_nwk_s_int_key")
		x.SNwkSIntKey.MarshalProtoJSON(s.WithField("s_nwk_s_int_key"))
	}
	if x.NwkSEncKey != nil || s.HasField("nwk_s_enc_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("nwk_s_enc_key")
		x.NwkSEncKey.MarshalProtoJSON(s.WithField("nwk_s_enc_key"))
	}
	if x.AppSKey != nil || s.HasField("app_s_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("app_s_key")
		x.AppSKey.MarshalProtoJSON(s.WithField("app_s_key"))
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the SessionKeys to JSON.
func (x *SessionKeys) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the SessionKeys message from JSON.
func (x *SessionKeys) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "session_key_id", "sessionKeyId":
			s.AddField("session_key_id")
			x.SessionKeyId = s.ReadBytes()
		case "f_nwk_s_int_key", "fNwkSIntKey":
			if s.ReadNil() {
				x.FNwkSIntKey = nil
				return
			}
			x.FNwkSIntKey = &KeyEnvelope{}
			x.FNwkSIntKey.UnmarshalProtoJSON(s.WithField("f_nwk_s_int_key", true))
		case "s_nwk_s_int_key", "sNwkSIntKey":
			if s.ReadNil() {
				x.SNwkSIntKey = nil
				return
			}
			x.SNwkSIntKey = &KeyEnvelope{}
			x.SNwkSIntKey.UnmarshalProtoJSON(s.WithField("s_nwk_s_int_key", true))
		case "nwk_s_enc_key", "nwkSEncKey":
			if s.ReadNil() {
				x.NwkSEncKey = nil
				return
			}
			x.NwkSEncKey = &KeyEnvelope{}
			x.NwkSEncKey.UnmarshalProtoJSON(s.WithField("nwk_s_enc_key", true))
		case "app_s_key", "appSKey":
			if s.ReadNil() {
				x.AppSKey = nil
				return
			}
			x.AppSKey = &KeyEnvelope{}
			x.AppSKey.UnmarshalProtoJSON(s.WithField("app_s_key", true))
		}
	})
}

// UnmarshalJSON unmarshals the SessionKeys from JSON.
func (x *SessionKeys) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
