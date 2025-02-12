// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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

import React from 'react'
import { Container, Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'

import PageTitle from '@ttn-lw/components/page-title'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import NotFoundRoute from '@ttn-lw/lib/components/not-found-route'

import WebhookAdd from '@console/containers/webhook-add'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

const m = defineMessages({
  addCustomWebhook: 'Add custom webhook',
  addWebhookViaTemplate: 'Add webhook for {templateName}',
  customWebhook: 'Custom webhook',
})

const ApplicationWebhookAddForm = props => {
  const { templateId, isCustom, webhookTemplate, appId } = props

  let breadcrumbContent = m.customWebhook
  if (!templateId) {
    breadcrumbContent = sharedMessages.add
  } else if (!isCustom && webhookTemplate.name) {
    breadcrumbContent = webhookTemplate.name
  }

  useBreadcrumbs(
    'apps.single.integrations.webhooks.various.add',
    <Breadcrumb
      path={`/applications/${appId}/integrations/webhooks/add/template/${templateId}`}
      content={breadcrumbContent}
    />,
  )

  let pageTitle = m.addCustomWebhook
  if (!webhookTemplate) {
    pageTitle = sharedMessages.addWebhook
  } else if (isCustom) {
    pageTitle = {
      ...m.addWebhookViaTemplate,
      values: {
        templateName: webhookTemplate.name,
      },
    }
  }

  // Render Not Found error when the template was not found.
  if (!isCustom && templateId && !webhookTemplate) {
    return <NotFoundRoute />
  }

  return (
    <Container>
      <PageTitle title={pageTitle} className="mb-0" hideHeading={Boolean(webhookTemplate)} />
      <Row>
        <Col lg={8} md={12}>
          <WebhookAdd appId={appId} templateId={templateId} webhookTemplate={webhookTemplate} />
        </Col>
      </Row>
    </Container>
  )
}

ApplicationWebhookAddForm.propTypes = {
  appId: PropTypes.string.isRequired,
  isCustom: PropTypes.bool.isRequired,
  templateId: PropTypes.string,
  webhookTemplate: PropTypes.webhookTemplate,
}

ApplicationWebhookAddForm.defaultProps = {
  templateId: undefined,
  webhookTemplate: undefined,
}

export default ApplicationWebhookAddForm
