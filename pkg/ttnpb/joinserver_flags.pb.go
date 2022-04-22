// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.0.0
// - protoc              v3.9.1
// source: lorawan-stack/api/joinserver.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	gogo "github.com/TheThingsIndustries/protoc-gen-go-flags/gogo"
	pflag "github.com/spf13/pflag"
	customflags "go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/customflags"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// AddSelectFlagsForApplicationActivationSettings adds flags to select fields in ApplicationActivationSettings.
func AddSelectFlagsForApplicationActivationSettings(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("kek-label", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("kek-label", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("kek", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("kek", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("kek", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("home-net-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("home-net-id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("application-server-id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("application-server-id", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forApplicationActivationSettings message from select flags.
func PathsFromSelectFlagsForApplicationActivationSettings(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("kek_label", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("kek_label", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("kek", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("kek", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForKeyEnvelope(flags, flagsplugin.Prefix("kek", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("home_net_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("home_net_id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("application_server_id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("application_server_id", prefix))
	}
	return paths, nil
}

// AddSetFlagsForApplicationActivationSettings adds flags to select fields in ApplicationActivationSettings.
func AddSetFlagsForApplicationActivationSettings(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("kek-label", prefix), "", flagsplugin.WithHidden(hidden)))
	AddSetFlagsForKeyEnvelope(flags, flagsplugin.Prefix("kek", prefix), hidden)
	flags.AddFlag(customflags.New3BytesFlag(flagsplugin.Prefix("home-net-id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("application-server-id", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the ApplicationActivationSettings message from flags.
func (m *ApplicationActivationSettings) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("kek_label", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.KekLabel = val
		paths = append(paths, flagsplugin.Prefix("kek_label", prefix))
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("kek", prefix)); changed {
		m.Kek = &KeyEnvelope{}
		if setPaths, err := m.Kek.SetFromFlags(flags, flagsplugin.Prefix("kek", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := types.GetNetID(flags, flagsplugin.Prefix("home_net_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.HomeNetId = &val
		paths = append(paths, flagsplugin.Prefix("home_net_id", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("application_server_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.ApplicationServerId = val
		paths = append(paths, flagsplugin.Prefix("application_server_id", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForSetApplicationActivationSettingsRequest adds flags to select fields in SetApplicationActivationSettingsRequest.
func AddSelectFlagsForSetApplicationActivationSettingsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("application-ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("application-ids", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("settings", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("settings", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForApplicationActivationSettings(flags, flagsplugin.Prefix("settings", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("field-mask", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("field-mask", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forSetApplicationActivationSettingsRequest message from select flags.
func PathsFromSelectFlagsForSetApplicationActivationSettingsRequest(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("application_ids", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("settings", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("settings", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForApplicationActivationSettings(flags, flagsplugin.Prefix("settings", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSetApplicationActivationSettingsRequest adds flags to select fields in SetApplicationActivationSettingsRequest.
func AddSetFlagsForSetApplicationActivationSettingsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("application-ids", prefix), hidden)
	AddSetFlagsForApplicationActivationSettings(flags, flagsplugin.Prefix("settings", prefix), hidden)
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the SetApplicationActivationSettingsRequest message from flags.
func (m *SetApplicationActivationSettingsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("application_ids", prefix)); changed {
		m.ApplicationIds = &ApplicationIdentifiers{}
		if setPaths, err := m.ApplicationIds.SetFromFlags(flags, flagsplugin.Prefix("application_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("settings", prefix)); changed {
		m.Settings = &ApplicationActivationSettings{}
		if setPaths, err := m.Settings.SetFromFlags(flags, flagsplugin.Prefix("settings", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FieldMask = gogo.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	return paths, nil
}
