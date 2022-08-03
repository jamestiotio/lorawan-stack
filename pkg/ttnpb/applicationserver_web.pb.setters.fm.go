// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *ApplicationWebhookIdentifiers) SetFields(src *ApplicationWebhookIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}
		case "webhook_id":
			if len(subs) > 0 {
				return fmt.Errorf("'webhook_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.WebhookId = src.WebhookId
			} else {
				var zero string
				dst.WebhookId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhookTemplateIdentifiers) SetFields(src *ApplicationWebhookTemplateIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "template_id":
			if len(subs) > 0 {
				return fmt.Errorf("'template_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TemplateId = src.TemplateId
			} else {
				var zero string
				dst.TemplateId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhookTemplateField) SetFields(src *ApplicationWebhookTemplateField, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "id":
			if len(subs) > 0 {
				return fmt.Errorf("'id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Id = src.Id
			} else {
				var zero string
				dst.Id = zero
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "description":
			if len(subs) > 0 {
				return fmt.Errorf("'description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Description = src.Description
			} else {
				var zero string
				dst.Description = zero
			}
		case "secret":
			if len(subs) > 0 {
				return fmt.Errorf("'secret' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Secret = src.Secret
			} else {
				var zero bool
				dst.Secret = zero
			}
		case "default_value":
			if len(subs) > 0 {
				return fmt.Errorf("'default_value' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DefaultValue = src.DefaultValue
			} else {
				var zero string
				dst.DefaultValue = zero
			}
		case "optional":
			if len(subs) > 0 {
				return fmt.Errorf("'optional' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Optional = src.Optional
			} else {
				var zero bool
				dst.Optional = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhookTemplate) SetFields(src *ApplicationWebhookTemplate, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplateIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &ApplicationWebhookTemplateIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "description":
			if len(subs) > 0 {
				return fmt.Errorf("'description' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Description = src.Description
			} else {
				var zero string
				dst.Description = zero
			}
		case "logo_url":
			if len(subs) > 0 {
				return fmt.Errorf("'logo_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LogoUrl = src.LogoUrl
			} else {
				var zero string
				dst.LogoUrl = zero
			}
		case "info_url":
			if len(subs) > 0 {
				return fmt.Errorf("'info_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.InfoUrl = src.InfoUrl
			} else {
				var zero string
				dst.InfoUrl = zero
			}
		case "documentation_url":
			if len(subs) > 0 {
				return fmt.Errorf("'documentation_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DocumentationUrl = src.DocumentationUrl
			} else {
				var zero string
				dst.DocumentationUrl = zero
			}
		case "base_url":
			if len(subs) > 0 {
				return fmt.Errorf("'base_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BaseUrl = src.BaseUrl
			} else {
				var zero string
				dst.BaseUrl = zero
			}
		case "headers":
			if len(subs) > 0 {
				return fmt.Errorf("'headers' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Headers = src.Headers
			} else {
				dst.Headers = nil
			}
		case "format":
			if len(subs) > 0 {
				return fmt.Errorf("'format' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Format = src.Format
			} else {
				var zero string
				dst.Format = zero
			}
		case "fields":
			if len(subs) > 0 {
				return fmt.Errorf("'fields' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Fields = src.Fields
			} else {
				dst.Fields = nil
			}
		case "create_downlink_api_key":
			if len(subs) > 0 {
				return fmt.Errorf("'create_downlink_api_key' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreateDownlinkApiKey = src.CreateDownlinkApiKey
			} else {
				var zero bool
				dst.CreateDownlinkApiKey = zero
			}
		case "uplink_message":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.UplinkMessage == nil) && dst.UplinkMessage == nil {
					continue
				}
				if src != nil {
					newSrc = src.UplinkMessage
				}
				if dst.UplinkMessage != nil {
					newDst = dst.UplinkMessage
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.UplinkMessage = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UplinkMessage = src.UplinkMessage
				} else {
					dst.UplinkMessage = nil
				}
			}
		case "join_accept":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.JoinAccept == nil) && dst.JoinAccept == nil {
					continue
				}
				if src != nil {
					newSrc = src.JoinAccept
				}
				if dst.JoinAccept != nil {
					newDst = dst.JoinAccept
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.JoinAccept = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.JoinAccept = src.JoinAccept
				} else {
					dst.JoinAccept = nil
				}
			}
		case "downlink_ack":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.DownlinkAck == nil) && dst.DownlinkAck == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkAck
				}
				if dst.DownlinkAck != nil {
					newDst = dst.DownlinkAck
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.DownlinkAck = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkAck = src.DownlinkAck
				} else {
					dst.DownlinkAck = nil
				}
			}
		case "downlink_nack":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.DownlinkNack == nil) && dst.DownlinkNack == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkNack
				}
				if dst.DownlinkNack != nil {
					newDst = dst.DownlinkNack
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.DownlinkNack = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkNack = src.DownlinkNack
				} else {
					dst.DownlinkNack = nil
				}
			}
		case "downlink_sent":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.DownlinkSent == nil) && dst.DownlinkSent == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkSent
				}
				if dst.DownlinkSent != nil {
					newDst = dst.DownlinkSent
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.DownlinkSent = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkSent = src.DownlinkSent
				} else {
					dst.DownlinkSent = nil
				}
			}
		case "downlink_failed":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.DownlinkFailed == nil) && dst.DownlinkFailed == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkFailed
				}
				if dst.DownlinkFailed != nil {
					newDst = dst.DownlinkFailed
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.DownlinkFailed = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkFailed = src.DownlinkFailed
				} else {
					dst.DownlinkFailed = nil
				}
			}
		case "downlink_queued":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.DownlinkQueued == nil) && dst.DownlinkQueued == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkQueued
				}
				if dst.DownlinkQueued != nil {
					newDst = dst.DownlinkQueued
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.DownlinkQueued = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkQueued = src.DownlinkQueued
				} else {
					dst.DownlinkQueued = nil
				}
			}
		case "downlink_queue_invalidated":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.DownlinkQueueInvalidated == nil) && dst.DownlinkQueueInvalidated == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkQueueInvalidated
				}
				if dst.DownlinkQueueInvalidated != nil {
					newDst = dst.DownlinkQueueInvalidated
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.DownlinkQueueInvalidated = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkQueueInvalidated = src.DownlinkQueueInvalidated
				} else {
					dst.DownlinkQueueInvalidated = nil
				}
			}
		case "location_solved":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.LocationSolved == nil) && dst.LocationSolved == nil {
					continue
				}
				if src != nil {
					newSrc = src.LocationSolved
				}
				if dst.LocationSolved != nil {
					newDst = dst.LocationSolved
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.LocationSolved = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.LocationSolved = src.LocationSolved
				} else {
					dst.LocationSolved = nil
				}
			}
		case "service_data":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplate_Message
				if (src == nil || src.ServiceData == nil) && dst.ServiceData == nil {
					continue
				}
				if src != nil {
					newSrc = src.ServiceData
				}
				if dst.ServiceData != nil {
					newDst = dst.ServiceData
				} else {
					newDst = &ApplicationWebhookTemplate_Message{}
					dst.ServiceData = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ServiceData = src.ServiceData
				} else {
					dst.ServiceData = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhookTemplates) SetFields(src *ApplicationWebhookTemplates, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "templates":
			if len(subs) > 0 {
				return fmt.Errorf("'templates' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Templates = src.Templates
			} else {
				dst.Templates = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhookHealth) SetFields(src *ApplicationWebhookHealth, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {

		case "status":
			if len(subs) == 0 && src == nil {
				dst.Status = nil
				continue
			} else if len(subs) == 0 {
				dst.Status = src.Status
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "healthy":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.Status.(*ApplicationWebhookHealth_Healthy)
					}
					if srcValid := srcTypeOk || src == nil || src.Status == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'healthy', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Status.(*ApplicationWebhookHealth_Healthy)
					if dstValid := dstTypeOk || dst.Status == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'healthy', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ApplicationWebhookHealth_WebhookHealthStatusHealthy
						if srcTypeOk {
							newSrc = src.Status.(*ApplicationWebhookHealth_Healthy).Healthy
						}
						if dstTypeOk {
							newDst = dst.Status.(*ApplicationWebhookHealth_Healthy).Healthy
						} else if srcTypeOk {
							newDst = &ApplicationWebhookHealth_WebhookHealthStatusHealthy{}
							dst.Status = &ApplicationWebhookHealth_Healthy{Healthy: newDst}
						} else {
							dst.Status = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Status = src.Status
						} else {
							dst.Status = nil
						}
					}
				case "unhealthy":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.Status.(*ApplicationWebhookHealth_Unhealthy)
					}
					if srcValid := srcTypeOk || src == nil || src.Status == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'unhealthy', while different oneof is set in source")
					}
					_, dstTypeOk := dst.Status.(*ApplicationWebhookHealth_Unhealthy)
					if dstValid := dstTypeOk || dst.Status == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'unhealthy', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ApplicationWebhookHealth_WebhookHealthStatusUnhealthy
						if srcTypeOk {
							newSrc = src.Status.(*ApplicationWebhookHealth_Unhealthy).Unhealthy
						}
						if dstTypeOk {
							newDst = dst.Status.(*ApplicationWebhookHealth_Unhealthy).Unhealthy
						} else if srcTypeOk {
							newDst = &ApplicationWebhookHealth_WebhookHealthStatusUnhealthy{}
							dst.Status = &ApplicationWebhookHealth_Unhealthy{Unhealthy: newDst}
						} else {
							dst.Status = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.Status = src.Status
						} else {
							dst.Status = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhook) SetFields(src *ApplicationWebhook, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &ApplicationWebhookIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				dst.CreatedAt = nil
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				dst.UpdatedAt = nil
			}
		case "base_url":
			if len(subs) > 0 {
				return fmt.Errorf("'base_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BaseUrl = src.BaseUrl
			} else {
				var zero string
				dst.BaseUrl = zero
			}
		case "headers":
			if len(subs) > 0 {
				return fmt.Errorf("'headers' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Headers = src.Headers
			} else {
				dst.Headers = nil
			}
		case "format":
			if len(subs) > 0 {
				return fmt.Errorf("'format' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Format = src.Format
			} else {
				var zero string
				dst.Format = zero
			}
		case "template_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplateIdentifiers
				if (src == nil || src.TemplateIds == nil) && dst.TemplateIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.TemplateIds
				}
				if dst.TemplateIds != nil {
					newDst = dst.TemplateIds
				} else {
					newDst = &ApplicationWebhookTemplateIdentifiers{}
					dst.TemplateIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.TemplateIds = src.TemplateIds
				} else {
					dst.TemplateIds = nil
				}
			}
		case "template_fields":
			if len(subs) > 0 {
				return fmt.Errorf("'template_fields' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TemplateFields = src.TemplateFields
			} else {
				dst.TemplateFields = nil
			}
		case "downlink_api_key":
			if len(subs) > 0 {
				return fmt.Errorf("'downlink_api_key' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DownlinkApiKey = src.DownlinkApiKey
			} else {
				var zero string
				dst.DownlinkApiKey = zero
			}
		case "uplink_message":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.UplinkMessage == nil) && dst.UplinkMessage == nil {
					continue
				}
				if src != nil {
					newSrc = src.UplinkMessage
				}
				if dst.UplinkMessage != nil {
					newDst = dst.UplinkMessage
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.UplinkMessage = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UplinkMessage = src.UplinkMessage
				} else {
					dst.UplinkMessage = nil
				}
			}
		case "join_accept":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.JoinAccept == nil) && dst.JoinAccept == nil {
					continue
				}
				if src != nil {
					newSrc = src.JoinAccept
				}
				if dst.JoinAccept != nil {
					newDst = dst.JoinAccept
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.JoinAccept = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.JoinAccept = src.JoinAccept
				} else {
					dst.JoinAccept = nil
				}
			}
		case "downlink_ack":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.DownlinkAck == nil) && dst.DownlinkAck == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkAck
				}
				if dst.DownlinkAck != nil {
					newDst = dst.DownlinkAck
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.DownlinkAck = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkAck = src.DownlinkAck
				} else {
					dst.DownlinkAck = nil
				}
			}
		case "downlink_nack":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.DownlinkNack == nil) && dst.DownlinkNack == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkNack
				}
				if dst.DownlinkNack != nil {
					newDst = dst.DownlinkNack
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.DownlinkNack = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkNack = src.DownlinkNack
				} else {
					dst.DownlinkNack = nil
				}
			}
		case "downlink_sent":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.DownlinkSent == nil) && dst.DownlinkSent == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkSent
				}
				if dst.DownlinkSent != nil {
					newDst = dst.DownlinkSent
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.DownlinkSent = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkSent = src.DownlinkSent
				} else {
					dst.DownlinkSent = nil
				}
			}
		case "downlink_failed":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.DownlinkFailed == nil) && dst.DownlinkFailed == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkFailed
				}
				if dst.DownlinkFailed != nil {
					newDst = dst.DownlinkFailed
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.DownlinkFailed = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkFailed = src.DownlinkFailed
				} else {
					dst.DownlinkFailed = nil
				}
			}
		case "downlink_queued":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.DownlinkQueued == nil) && dst.DownlinkQueued == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkQueued
				}
				if dst.DownlinkQueued != nil {
					newDst = dst.DownlinkQueued
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.DownlinkQueued = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkQueued = src.DownlinkQueued
				} else {
					dst.DownlinkQueued = nil
				}
			}
		case "downlink_queue_invalidated":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.DownlinkQueueInvalidated == nil) && dst.DownlinkQueueInvalidated == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkQueueInvalidated
				}
				if dst.DownlinkQueueInvalidated != nil {
					newDst = dst.DownlinkQueueInvalidated
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.DownlinkQueueInvalidated = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkQueueInvalidated = src.DownlinkQueueInvalidated
				} else {
					dst.DownlinkQueueInvalidated = nil
				}
			}
		case "location_solved":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.LocationSolved == nil) && dst.LocationSolved == nil {
					continue
				}
				if src != nil {
					newSrc = src.LocationSolved
				}
				if dst.LocationSolved != nil {
					newDst = dst.LocationSolved
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.LocationSolved = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.LocationSolved = src.LocationSolved
				} else {
					dst.LocationSolved = nil
				}
			}
		case "service_data":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook_Message
				if (src == nil || src.ServiceData == nil) && dst.ServiceData == nil {
					continue
				}
				if src != nil {
					newSrc = src.ServiceData
				}
				if dst.ServiceData != nil {
					newDst = dst.ServiceData
				} else {
					newDst = &ApplicationWebhook_Message{}
					dst.ServiceData = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ServiceData = src.ServiceData
				} else {
					dst.ServiceData = nil
				}
			}
		case "health_status":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookHealth
				if (src == nil || src.HealthStatus == nil) && dst.HealthStatus == nil {
					continue
				}
				if src != nil {
					newSrc = src.HealthStatus
				}
				if dst.HealthStatus != nil {
					newDst = dst.HealthStatus
				} else {
					newDst = &ApplicationWebhookHealth{}
					dst.HealthStatus = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.HealthStatus = src.HealthStatus
				} else {
					dst.HealthStatus = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhooks) SetFields(src *ApplicationWebhooks, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "webhooks":
			if len(subs) > 0 {
				return fmt.Errorf("'webhooks' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Webhooks = src.Webhooks
			} else {
				dst.Webhooks = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhookFormats) SetFields(src *ApplicationWebhookFormats, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "formats":
			if len(subs) > 0 {
				return fmt.Errorf("'formats' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Formats = src.Formats
			} else {
				dst.Formats = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetApplicationWebhookRequest) SetFields(src *GetApplicationWebhookRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &ApplicationWebhookIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListApplicationWebhooksRequest) SetFields(src *ListApplicationWebhooksRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SetApplicationWebhookRequest) SetFields(src *SetApplicationWebhookRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "webhook":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhook
				if (src == nil || src.Webhook == nil) && dst.Webhook == nil {
					continue
				}
				if src != nil {
					newSrc = src.Webhook
				}
				if dst.Webhook != nil {
					newDst = dst.Webhook
				} else {
					newDst = &ApplicationWebhook{}
					dst.Webhook = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Webhook = src.Webhook
				} else {
					dst.Webhook = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetApplicationWebhookTemplateRequest) SetFields(src *GetApplicationWebhookTemplateRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationWebhookTemplateIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &ApplicationWebhookTemplateIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListApplicationWebhookTemplatesRequest) SetFields(src *ListApplicationWebhookTemplatesRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhookTemplate_Message) SetFields(src *ApplicationWebhookTemplate_Message, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "path":
			if len(subs) > 0 {
				return fmt.Errorf("'path' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Path = src.Path
			} else {
				var zero string
				dst.Path = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhookHealth_WebhookHealthStatusHealthy) SetFields(src *ApplicationWebhookHealth_WebhookHealthStatusHealthy, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message ApplicationWebhookHealth_WebhookHealthStatusHealthy has no fields, but paths %s were specified", paths)
	}
	return nil
}

func (dst *ApplicationWebhookHealth_WebhookHealthStatusUnhealthy) SetFields(src *ApplicationWebhookHealth_WebhookHealthStatusUnhealthy, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "failed_attempts":
			if len(subs) > 0 {
				return fmt.Errorf("'failed_attempts' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FailedAttempts = src.FailedAttempts
			} else {
				var zero uint64
				dst.FailedAttempts = zero
			}
		case "last_failed_attempt_at":
			if len(subs) > 0 {
				return fmt.Errorf("'last_failed_attempt_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LastFailedAttemptAt = src.LastFailedAttemptAt
			} else {
				dst.LastFailedAttemptAt = nil
			}
		case "last_failed_attempt_details":
			if len(subs) > 0 {
				var newDst, newSrc *ErrorDetails
				if (src == nil || src.LastFailedAttemptDetails == nil) && dst.LastFailedAttemptDetails == nil {
					continue
				}
				if src != nil {
					newSrc = src.LastFailedAttemptDetails
				}
				if dst.LastFailedAttemptDetails != nil {
					newDst = dst.LastFailedAttemptDetails
				} else {
					newDst = &ErrorDetails{}
					dst.LastFailedAttemptDetails = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.LastFailedAttemptDetails = src.LastFailedAttemptDetails
				} else {
					dst.LastFailedAttemptDetails = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationWebhook_Message) SetFields(src *ApplicationWebhook_Message, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "path":
			if len(subs) > 0 {
				return fmt.Errorf("'path' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Path = src.Path
			} else {
				var zero string
				dst.Path = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
