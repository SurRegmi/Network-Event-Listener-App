// Copyright (c) 2017 company Inc. All rights reserved.

// Package with all the constants required by the listener library.

package lib

const (
  // Webhook URLs
  CreateWebhook = "/api/company/v3/webhooks"
  ListWebhooks = "/api/company/v3/webhooks/list"
  GetWebhook = "/api/company/v3/webhooks/{uuid}"
  UpdateWebhook = "/api/company/v3/webhooks/"
  DeleteWebhook = "/api/company/v3/webhooks/{uuid}"

  // Auth Check URL
  GetCurrentUser = "/api/company/v3/users/me"

  // Listener Defaults
  DefaultListenerPort = "8080"
  ListenerCallbackURL = "/listener/callback"
  EventConsumerCallbackMethod = "OnEvent"
  WebhookNamePrefix = "company_Listener_Webhook_"
  WebhookKind = "webhook"

  // Events (Can be taken from the YAML config later)
  VM_CREATE = "VM.CREATE"
  VM_DELETE = "VM.DELETE"
  VM_ON = "VM.ON"
  VM_OFF = "VM.OFF"
  VM_UPDATE = "VM.UPDATE"
  VM_MIGRATE = "VM.MIGRATE"
  VM_NIC_PLUG = "VM.NIC_PLUG"
  VM_NIC_UNPLUG = "VM.NIC_UNPLUG"
)
