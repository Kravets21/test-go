# Go API client for AirEml

airEml - API

## Overview
- API version: 1.0.0
- Package version: 1.0.0

## Installation

Install the following dependencies:

```shell
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import AirEml "github.com/pdffiller/eml-go-api"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `AirEml.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), AirEml.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `AirEml.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), AirEml.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `AirEml.ContextOperationServerIndices` and `AirEml.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), AirEml.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), AirEml.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation For Authorization

Authentication schemes defined for the API:


- **Type**: HTTP Bearer token authentication

AirEml api client will automatically refresh tokens and perform user authentication by your client_id and client_secret
You need only to set your credentials

Example:
```golang
cfg := airEml.NewConfiguration()
cfg.ClientID = "9ac87c6c-****-****-****-1389e2975c8c"
cfg.ClientSecret = "qF********************zCeZRmDcC8"

apiClient := airEml.NewAPIClient(cfg)
ctx := context.Background()

resp, r, err := apiClient.SendingAPI.SendSmsNotification(ctx).SmsNotificationSendRequest(airEml.SmsNotificationSendRequest{}).Execute()
```


## Documentation for API Endpoints

All URIs are relative to *https://aireml.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AudienceAPI* | [**AttachAudienceRecipientContact**](docs/AudienceAPI.md#attachaudiencerecipientcontact) | **Post** /api/v1/audience/recipients/{audience_recipient_id}/contacts/{audience_contact_type} | Attach a Contact to the Audience Recipient
*AudienceAPI* | [**AttachAudienceRecipientTag**](docs/AudienceAPI.md#attachaudiencerecipienttag) | **Post** /api/v1/audience/recipients/{audience_recipient_id}/tags/{audience_tag_name} | Attach a Tag to the Audience Recipient
*AudienceAPI* | [**CreateAudienceRecipient**](docs/AudienceAPI.md#createaudiencerecipient) | **Post** /api/v1/audience/recipients | Create a New Recipient
*AudienceAPI* | [**CreateAudienceSegment**](docs/AudienceAPI.md#createaudiencesegment) | **Post** /api/v1/audience/segments | Create a New Segment
*AudienceAPI* | [**CreateAudienceSegmentFilter**](docs/AudienceAPI.md#createaudiencesegmentfilter) | **Post** /api/v1/audience/segments/{audience_segment_id}/filters | Create a New Segment Filter
*AudienceAPI* | [**CreateAudienceTag**](docs/AudienceAPI.md#createaudiencetag) | **Post** /api/v1/audience/tags | Create a New Tag
*AudienceAPI* | [**DeleteAudienceRecipient**](docs/AudienceAPI.md#deleteaudiencerecipient) | **Delete** /api/v1/audience/recipients/{audience_recipient_id} | Delete a Recipient
*AudienceAPI* | [**DeleteAudienceRecipientContact**](docs/AudienceAPI.md#deleteaudiencerecipientcontact) | **Delete** /api/v1/audience/recipients/{audience_recipient_id}/contacts/{audience_contact_type} | Delete a Contact from the Audience Recipient
*AudienceAPI* | [**DeleteAudienceRecipientTag**](docs/AudienceAPI.md#deleteaudiencerecipienttag) | **Delete** /api/v1/audience/recipients/{audience_recipient_id}/tags/{audience_tag_name} | Delete a Tag from the Audience Recipient
*AudienceAPI* | [**DeleteAudienceSegment**](docs/AudienceAPI.md#deleteaudiencesegment) | **Delete** /api/v1/audience/segments/{audience_segment_id} | Delete a Segment
*AudienceAPI* | [**DeleteAudienceSegmentFilter**](docs/AudienceAPI.md#deleteaudiencesegmentfilter) | **Delete** /api/v1/audience/segments/{audience_segment_id}/filters/{audience_filter_id} | Delete a Filter from the Audience Segment
*AudienceAPI* | [**DeleteAudienceTag**](docs/AudienceAPI.md#deleteaudiencetag) | **Delete** /api/v1/audience/tags/{audience_tag_id} | Delete a Tag
*AudienceAPI* | [**GetAudienceRecipient**](docs/AudienceAPI.md#getaudiencerecipient) | **Get** /api/v1/audience/recipients/{audience_recipient_id} | Retrieve an Existing Recipient
*AudienceAPI* | [**GetAudienceSegment**](docs/AudienceAPI.md#getaudiencesegment) | **Get** /api/v1/audience/segments/{audience_segment_id} | Retrieve an Existing Segment
*AudienceAPI* | [**GetAudienceSegmentFilter**](docs/AudienceAPI.md#getaudiencesegmentfilter) | **Get** /api/v1/audience/segments/{audience_segment_id}/filters/{audience_filter_id} | Retrieve an Existing Filter from the Audience Segment
*AudienceAPI* | [**GetAudienceTag**](docs/AudienceAPI.md#getaudiencetag) | **Get** /api/v1/audience/tags/{audience_tag_id} | Retrieve an Existing Tag
*AudienceAPI* | [**ListAudienceRecipientSegments**](docs/AudienceAPI.md#listaudiencerecipientsegments) | **Get** /api/v1/audience/recipients/{audience_recipient_id}/segments | List All Segments from the Recipient
*AudienceAPI* | [**ListAudienceRecipients**](docs/AudienceAPI.md#listaudiencerecipients) | **Get** /api/v1/audience/recipients | List All Recipients
*AudienceAPI* | [**ListAudienceSegmentFilters**](docs/AudienceAPI.md#listaudiencesegmentfilters) | **Get** /api/v1/audience/segments/{audience_segment_id}/filters | List All Filters from the Segment
*AudienceAPI* | [**ListAudienceSegments**](docs/AudienceAPI.md#listaudiencesegments) | **Get** /api/v1/audience/segments | List All Segments
*AudienceAPI* | [**ListAudienceTags**](docs/AudienceAPI.md#listaudiencetags) | **Get** /api/v1/audience/tags | List All Tags
*AudienceAPI* | [**UpdateAssignActiveStatusToAudienceSegment**](docs/AudienceAPI.md#updateassignactivestatustoaudiencesegment) | **Put** /api/v1/audience/segments/{audience_segment_id}/activate | Activate a Segment
*AudienceAPI* | [**UpdateAssignInactiveStatusToAudienceSegment**](docs/AudienceAPI.md#updateassigninactivestatustoaudiencesegment) | **Put** /api/v1/audience/segments/{audience_segment_id}/deactivate | Deactivate a Segment
*AudienceAPI* | [**UpdateAudienceRecipient**](docs/AudienceAPI.md#updateaudiencerecipient) | **Put** /api/v1/audience/recipients/{audience_recipient_id} | Update a Recipient
*AudienceAPI* | [**UpdateAudienceSegment**](docs/AudienceAPI.md#updateaudiencesegment) | **Put** /api/v1/audience/segments/{audience_segment_id} | Update a Segment
*AudienceAPI* | [**UpdateAudienceSegmentFilter**](docs/AudienceAPI.md#updateaudiencesegmentfilter) | **Put** /api/v1/audience/segments/{audience_segment_id}/filters/{audience_filter_id} | Update a Filter in the Audience Segment
*AudienceAPI* | [**UpdateAudienceTag**](docs/AudienceAPI.md#updateaudiencetag) | **Put** /api/v1/audience/tags/{audience_tag_id} | Update a Tag
*AuthorizationAPI* | [**CreateAuthTokenForServiceAccount**](docs/AuthorizationAPI.md#createauthtokenforserviceaccount) | **Post** /api/v1/auth/tokens/service-account | Create Access Token for Service Account
*EmailSettingsAPI* | [**CreateEmailChannel**](docs/EmailSettingsAPI.md#createemailchannel) | **Post** /api/v1/email/channels | Create a New Email Notification Channel
*EmailSettingsAPI* | [**CreateEmailChannelRoutingRule**](docs/EmailSettingsAPI.md#createemailchannelroutingrule) | **Post** /api/v1/email/channels/{channel_id}/routing-rules | Create a Routing-rule in the Email Notification Channel
*EmailSettingsAPI* | [**CreateEmailSender**](docs/EmailSettingsAPI.md#createemailsender) | **Post** /api/v1/email/senders | Create a New Email Notification Sender
*EmailSettingsAPI* | [**CreateEmailTransport**](docs/EmailSettingsAPI.md#createemailtransport) | **Post** /api/v1/email/transports | Create a New Email Notification Transport
*EmailSettingsAPI* | [**DeleteEmailChannel**](docs/EmailSettingsAPI.md#deleteemailchannel) | **Delete** /api/v1/email/channels/{channel_id} | Delete a Email Notification Channel
*EmailSettingsAPI* | [**DeleteEmailChannelRoutingRule**](docs/EmailSettingsAPI.md#deleteemailchannelroutingrule) | **Delete** /api/v1/email/channels/{channel_id}/routing-rules/{routing_rule_id} | Delete a Routing-rule from the Email Notification Channel
*EmailSettingsAPI* | [**DeleteEmailSender**](docs/EmailSettingsAPI.md#deleteemailsender) | **Delete** /api/v1/email/senders/{sender_id} | Delete a Email Notification Sender
*EmailSettingsAPI* | [**DeleteEmailTransport**](docs/EmailSettingsAPI.md#deleteemailtransport) | **Delete** /api/v1/email/transports/{transport_id} | Delete a Email Notification Transport
*EmailSettingsAPI* | [**GetEmailChannel**](docs/EmailSettingsAPI.md#getemailchannel) | **Get** /api/v1/email/channels/{channel_id} | Retrieve an Existing Email Notification Channel
*EmailSettingsAPI* | [**GetEmailChannelRoutingRule**](docs/EmailSettingsAPI.md#getemailchannelroutingrule) | **Get** /api/v1/email/channels/{channel_id}/routing-rules/{routing_rule_id} | Retrieve an Existing Routing-rule from the Email Notification Channel
*EmailSettingsAPI* | [**GetEmailSender**](docs/EmailSettingsAPI.md#getemailsender) | **Get** /api/v1/email/senders/{sender_id} | Retrieve an Existing Email Notification Sender
*EmailSettingsAPI* | [**GetEmailTransport**](docs/EmailSettingsAPI.md#getemailtransport) | **Get** /api/v1/email/transports/{transport_id} | Retrieve an Existing Email Notification Transport
*EmailSettingsAPI* | [**ListEmailChannelRoutingRules**](docs/EmailSettingsAPI.md#listemailchannelroutingrules) | **Get** /api/v1/email/channels/{channel_id}/routing-rules | List All Routing-rules from the Email Notification Channel
*EmailSettingsAPI* | [**ListEmailChannels**](docs/EmailSettingsAPI.md#listemailchannels) | **Get** /api/v1/email/channels | List All Email Notification Channels
*EmailSettingsAPI* | [**ListEmailSenders**](docs/EmailSettingsAPI.md#listemailsenders) | **Get** /api/v1/email/senders | List All Email Notification Senders
*EmailSettingsAPI* | [**ListEmailTransports**](docs/EmailSettingsAPI.md#listemailtransports) | **Get** /api/v1/email/transports | List All Email Notification Transports
*EmailSettingsAPI* | [**UpdateEmailChannel**](docs/EmailSettingsAPI.md#updateemailchannel) | **Put** /api/v1/email/channels/{channel_id} | Update a Email Notification Channel
*EmailSettingsAPI* | [**UpdateEmailChannelRoutingRule**](docs/EmailSettingsAPI.md#updateemailchannelroutingrule) | **Put** /api/v1/email/channels/{channel_id}/routing-rules/{routing_rule_id} | Update a Routing-rule in the Email Notification Channel
*EmailSettingsAPI* | [**UpdateEmailSender**](docs/EmailSettingsAPI.md#updateemailsender) | **Put** /api/v1/email/senders/{sender_id} | Update a Email Notification Sender
*EmailSettingsAPI* | [**UpdateEmailTransport**](docs/EmailSettingsAPI.md#updateemailtransport) | **Put** /api/v1/email/transports/{transport_id} | Update a Email Notification Transport
*EmailTemplatesAPI* | [**CreateEmailTemplate**](docs/EmailTemplatesAPI.md#createemailtemplate) | **Post** /api/v1/email/templates | Create a New Email Notification Template
*EmailTemplatesAPI* | [**CreateEmailTemplateContent**](docs/EmailTemplatesAPI.md#createemailtemplatecontent) | **Post** /api/v1/email/templates/{template_id}/contents/{template_content_locale} | Create a Content in the Email Notification Template
*EmailTemplatesAPI* | [**DeleteEmailTemplate**](docs/EmailTemplatesAPI.md#deleteemailtemplate) | **Delete** /api/v1/email/templates/{template_id} | Delete a Email Notification Template
*EmailTemplatesAPI* | [**DeleteEmailTemplateContent**](docs/EmailTemplatesAPI.md#deleteemailtemplatecontent) | **Delete** /api/v1/email/templates/{template_id}/contents/{template_content_locale} | Delete a Content from the Email Notification Template
*EmailTemplatesAPI* | [**GetEmailTemplate**](docs/EmailTemplatesAPI.md#getemailtemplate) | **Get** /api/v1/email/templates/{template_id} | Retrieve an Existing Email Notification Template
*EmailTemplatesAPI* | [**GetEmailTemplateContent**](docs/EmailTemplatesAPI.md#getemailtemplatecontent) | **Get** /api/v1/email/templates/{template_id}/contents/{template_content_locale} | Retrieve an Existing Content from the Email Notification Template
*EmailTemplatesAPI* | [**ListEmailTemplateContents**](docs/EmailTemplatesAPI.md#listemailtemplatecontents) | **Get** /api/v1/email/templates/{template_id}/contents | List All Contents from the Email Notification Template
*EmailTemplatesAPI* | [**ListEmailTemplates**](docs/EmailTemplatesAPI.md#listemailtemplates) | **Get** /api/v1/email/templates | List All Email Notification Templates
*EmailTemplatesAPI* | [**UpdateEmailTemplate**](docs/EmailTemplatesAPI.md#updateemailtemplate) | **Put** /api/v1/email/templates/{template_id} | Update a Email Notification Template
*EmailTemplatesAPI* | [**UpdateEmailTemplateContent**](docs/EmailTemplatesAPI.md#updateemailtemplatecontent) | **Put** /api/v1/email/templates/{template_id}/contents/{template_content_locale} | Update a Content in the Email Notification Template
*NotificationsAPI* | [**GetEmailNotification**](docs/NotificationsAPI.md#getemailnotification) | **Get** /api/v1/email/notifications/{notification_id} | Retrieve an Existing Email Notification
*NotificationsAPI* | [**GetEmailNotificationStates**](docs/NotificationsAPI.md#getemailnotificationstates) | **Get** /api/v1/email/notifications/{notification_id}/states | Retrieve an Existing Email Notification States
*NotificationsAPI* | [**GetPushNotification**](docs/NotificationsAPI.md#getpushnotification) | **Get** /api/v1/push/notifications/{notification_id} | Retrieve an Existing Push Notification
*NotificationsAPI* | [**GetPushNotificationStates**](docs/NotificationsAPI.md#getpushnotificationstates) | **Get** /api/v1/push/notifications/{notification_id}/states | Retrieve an Existing Push Notification States
*NotificationsAPI* | [**GetSmsNotification**](docs/NotificationsAPI.md#getsmsnotification) | **Get** /api/v1/sms/notifications/{notification_id} | Retrieve an Existing SMS Notification
*NotificationsAPI* | [**GetSmsNotificationStates**](docs/NotificationsAPI.md#getsmsnotificationstates) | **Get** /api/v1/sms/notifications/{notification_id}/states | Retrieve an Existing Sms Notification States
*NotificationsAPI* | [**ListEmailNotifications**](docs/NotificationsAPI.md#listemailnotifications) | **Get** /api/v1/email/notifications | List All Email Notifications
*NotificationsAPI* | [**ListPushNotifications**](docs/NotificationsAPI.md#listpushnotifications) | **Get** /api/v1/push/notifications | List All Push Notifications
*NotificationsAPI* | [**ListSmsNotifications**](docs/NotificationsAPI.md#listsmsnotifications) | **Get** /api/v1/sms/notifications | List All SMS Notifications
*PushSettingsAPI* | [**CreatePushPlatform**](docs/PushSettingsAPI.md#createpushplatform) | **Post** /api/v1/push/platforms | Create a New Push Notification Platform
*PushSettingsAPI* | [**ListPushPlatforms**](docs/PushSettingsAPI.md#listpushplatforms) | **Get** /api/v1/push/platforms | List All Push Notification Platforms
*PushTemplatesAPI* | [**CreatePushTemplate**](docs/PushTemplatesAPI.md#createpushtemplate) | **Post** /api/v1/push/templates | Create a New Push Notification Template
*PushTemplatesAPI* | [**CreatePushTemplateContent**](docs/PushTemplatesAPI.md#createpushtemplatecontent) | **Post** /api/v1/push/templates/{template_id}/contents/{template_content_locale} | Create a Content in the Push Notification Template
*PushTemplatesAPI* | [**DeletePushTemplate**](docs/PushTemplatesAPI.md#deletepushtemplate) | **Delete** /api/v1/push/templates/{template_id} | Delete a Push Notification Template
*PushTemplatesAPI* | [**DeletePushTemplateContent**](docs/PushTemplatesAPI.md#deletepushtemplatecontent) | **Delete** /api/v1/push/templates/{template_id}/contents/{template_content_locale} | Delete a Content from the Push Notification Template
*PushTemplatesAPI* | [**GetPushTemplate**](docs/PushTemplatesAPI.md#getpushtemplate) | **Get** /api/v1/push/templates/{template_id} | Retrieve an Existing Push Notification Template
*PushTemplatesAPI* | [**GetPushTemplateContent**](docs/PushTemplatesAPI.md#getpushtemplatecontent) | **Get** /api/v1/push/templates/{template_id}/contents/{template_content_locale} | Retrieve an Existing Content from the Push Notification Template
*PushTemplatesAPI* | [**ListPushTemplateContents**](docs/PushTemplatesAPI.md#listpushtemplatecontents) | **Get** /api/v1/push/templates/{template_id}/contents | List All Contents from the Push Notification Template
*PushTemplatesAPI* | [**ListPushTemplates**](docs/PushTemplatesAPI.md#listpushtemplates) | **Get** /api/v1/push/templates | List All Push Notification Templates
*PushTemplatesAPI* | [**UpdatePushTemplate**](docs/PushTemplatesAPI.md#updatepushtemplate) | **Put** /api/v1/push/templates/{template_id} | Update a Push Notification Template
*PushTemplatesAPI* | [**UpdatePushTemplateContent**](docs/PushTemplatesAPI.md#updatepushtemplatecontent) | **Put** /api/v1/push/templates/{template_id}/contents/{template_content_locale} | Update a Content in the Push Notification Template
*SMSSettingsAPI* | [**CreateSmsChannel**](docs/SMSSettingsAPI.md#createsmschannel) | **Post** /api/v1/sms/channels | Create a New SMS Channel
*SMSSettingsAPI* | [**CreateSmsTransport**](docs/SMSSettingsAPI.md#createsmstransport) | **Post** /api/v1/sms/transports | Create a New SMS Transport
*SMSSettingsAPI* | [**DeleteSmsChannel**](docs/SMSSettingsAPI.md#deletesmschannel) | **Delete** /api/v1/sms/channels/{channel_id} | Delete a SMS Channel
*SMSSettingsAPI* | [**DeleteSmsRoutingRule**](docs/SMSSettingsAPI.md#deletesmsroutingrule) | **Delete** /api/v1/sms/channels/{channel_id}/rules/{routing_rule_id} | Delete a SMS Channel Rule
*SMSSettingsAPI* | [**DeleteSmsTransport**](docs/SMSSettingsAPI.md#deletesmstransport) | **Delete** /api/v1/sms/transports/{transport_id} | Delete a SMS Transport
*SMSSettingsAPI* | [**GetSmsChannel**](docs/SMSSettingsAPI.md#getsmschannel) | **Get** /api/v1/sms/channels/{channel_id} | Retrieve an Existing SMS Channel
*SMSSettingsAPI* | [**GetSmsRoutingRule**](docs/SMSSettingsAPI.md#getsmsroutingrule) | **Get** /api/v1/sms/channels/{channel_id}/rules/{routing_rule_id} | Retrieve an Existing SMS Channel Rule
*SMSSettingsAPI* | [**GetSmsSender**](docs/SMSSettingsAPI.md#getsmssender) | **Get** /api/v1/sms/transports/{transport_id}/senders/{sender_id} | Retrieve an Existing Number from the SMS Transport
*SMSSettingsAPI* | [**GetSmsTransport**](docs/SMSSettingsAPI.md#getsmstransport) | **Get** /api/v1/sms/transports/{transport_id} | Retrieve an Existing SMS Transport
*SMSSettingsAPI* | [**GetWebhookByTransport**](docs/SMSSettingsAPI.md#getwebhookbytransport) | **Get** /api/v1/sms/transports/{transport_id}/generate-webhook | Webhook url generated by transport
*SMSSettingsAPI* | [**ListAllSmsSenders**](docs/SMSSettingsAPI.md#listallsmssenders) | **Get** /api/v1/sms/transports/all/senders | List All SMS Transport Numbers
*SMSSettingsAPI* | [**ListSmsChannels**](docs/SMSSettingsAPI.md#listsmschannels) | **Get** /api/v1/sms/channels | List All SMS Channels
*SMSSettingsAPI* | [**ListSmsChannelsRules**](docs/SMSSettingsAPI.md#listsmschannelsrules) | **Get** /api/v1/sms/channels/{channel_id}/rules | List of all Rules from the SMS Channel
*SMSSettingsAPI* | [**ListSmsSenders**](docs/SMSSettingsAPI.md#listsmssenders) | **Get** /api/v1/sms/transports/{transport_id}/senders | List All Numbers from the SMS Transport
*SMSSettingsAPI* | [**ListSmsTransports**](docs/SMSSettingsAPI.md#listsmstransports) | **Get** /api/v1/sms/transports | List All SMS Transports
*SMSSettingsAPI* | [**SaveSmsRoutingRule**](docs/SMSSettingsAPI.md#savesmsroutingrule) | **Post** /api/v1/sms/channels/{channel_id}/rules | Save a parameters SMS Channel Rule
*SMSSettingsAPI* | [**UpdateSmsChannel**](docs/SMSSettingsAPI.md#updatesmschannel) | **Put** /api/v1/sms/channels/{channel_id} | Update a SMS Channel
*SMSSettingsAPI* | [**UpdateSmsRoutingRule**](docs/SMSSettingsAPI.md#updatesmsroutingrule) | **Put** /api/v1/sms/channels/{channel_id}/rules/{routing_rule_id} | Update a SMS Channel Rule
*SMSSettingsAPI* | [**UpdateSmsTransport**](docs/SMSSettingsAPI.md#updatesmstransport) | **Put** /api/v1/sms/transports/{transport_id} | Update a SMS Transport
*SMSTemplatesAPI* | [**CreateSmsTemplate**](docs/SMSTemplatesAPI.md#createsmstemplate) | **Post** /api/v1/sms/templates | Create a New SMS Template
*SMSTemplatesAPI* | [**DeleteSmsTemplate**](docs/SMSTemplatesAPI.md#deletesmstemplate) | **Delete** /api/v1/sms/templates/{template_id} | Delete a SMS Template
*SMSTemplatesAPI* | [**DeleteSmsTemplateContent**](docs/SMSTemplatesAPI.md#deletesmstemplatecontent) | **Delete** /api/v1/sms/templates/{template_id}/contents/{template_content_locale} | Delete a Content from the SMS Template
*SMSTemplatesAPI* | [**GetSmsTemplate**](docs/SMSTemplatesAPI.md#getsmstemplate) | **Get** /api/v1/sms/templates/{template_id} | Retrieve an Existing SMS Template
*SMSTemplatesAPI* | [**GetSmsTemplateContent**](docs/SMSTemplatesAPI.md#getsmstemplatecontent) | **Get** /api/v1/sms/templates/{template_id}/contents/{template_content_locale} | Retrieve an Existing Content from the SMS Template
*SMSTemplatesAPI* | [**ListSmsTemplateContents**](docs/SMSTemplatesAPI.md#listsmstemplatecontents) | **Get** /api/v1/sms/templates/{template_id}/contents | List All Contents from the SMS Template
*SMSTemplatesAPI* | [**ListSmsTemplates**](docs/SMSTemplatesAPI.md#listsmstemplates) | **Get** /api/v1/sms/templates | List All SMS Templates
*SMSTemplatesAPI* | [**SaveSmsTemplateContent**](docs/SMSTemplatesAPI.md#savesmstemplatecontent) | **Post** /api/v1/sms/templates/{template_id}/contents/{template_content_locale} | Save a Content in the SMS Template
*SMSTemplatesAPI* | [**UpdateSmsTemplate**](docs/SMSTemplatesAPI.md#updatesmstemplate) | **Put** /api/v1/sms/templates/{template_id} | Update a SMS Template
*SMSTemplatesAPI* | [**UpdateSmsTemplateContent**](docs/SMSTemplatesAPI.md#updatesmstemplatecontent) | **Put** /api/v1/sms/templates/{template_id}/contents/{template_content_locale} | Update a Content in the SMS Notification Template
*SendingAPI* | [**SendEmail**](docs/SendingAPI.md#sendemail) | **Post** /api/v1/email/send | Send Email
*SendingAPI* | [**SendEmailBatch**](docs/SendingAPI.md#sendemailbatch) | **Post** /api/v1/email/send/batch | Send Email Batch
*SendingAPI* | [**SendPush**](docs/SendingAPI.md#sendpush) | **Post** /api/v1/push/send | Send Push Notification
*SendingAPI* | [**SendSmsBatchNotification**](docs/SendingAPI.md#sendsmsbatchnotification) | **Post** /api/v1/sms/send/batch | Send SMS Batch Notification
*SendingAPI* | [**SendSmsNotification**](docs/SendingAPI.md#sendsmsnotification) | **Post** /api/v1/sms/send | Send SMS Notification
*WebhooksAPI* | [**CreateWebhookSubscription**](docs/WebhooksAPI.md#createwebhooksubscription) | **Post** /api/v1/webhook/subscriptions | Create a New Subscription
*WebhooksAPI* | [**DeleteWebhookSubscription**](docs/WebhooksAPI.md#deletewebhooksubscription) | **Delete** /api/v1/webhook/subscriptions/{webhook_subscription_id} | Delete a webhook subscription
*WebhooksAPI* | [**GetWebhookSubscription**](docs/WebhooksAPI.md#getwebhooksubscription) | **Get** /api/v1/webhook/subscriptions/{webhook_subscription_id} | Retrieve an Existing webhook subscription
*WebhooksAPI* | [**ListWebhookEvents**](docs/WebhooksAPI.md#listwebhookevents) | **Get** /api/v1/webhook/events | List All Webhook events
*WebhooksAPI* | [**ListWebhookSubscriptions**](docs/WebhooksAPI.md#listwebhooksubscriptions) | **Get** /api/v1/webhook/subscriptions | List All webhook Subscriptions
*WebhooksAPI* | [**UpdateWebhookSubscription**](docs/WebhooksAPI.md#updatewebhooksubscription) | **Put** /api/v1/webhook/subscriptions/{webhook_subscription_id} | Update a webhook subscription


## Documentation For Models

 - [AudienceFilter](docs/AudienceFilter.md)
 - [AudienceFilterCollectionResponse](docs/AudienceFilterCollectionResponse.md)
 - [AudienceFilterRequest](docs/AudienceFilterRequest.md)
 - [AudienceFilterResponse](docs/AudienceFilterResponse.md)
 - [AudienceRecipient](docs/AudienceRecipient.md)
 - [AudienceRecipientCollectionResponse](docs/AudienceRecipientCollectionResponse.md)
 - [AudienceRecipientContact](docs/AudienceRecipientContact.md)
 - [AudienceRecipientContactRequest](docs/AudienceRecipientContactRequest.md)
 - [AudienceRecipientRequest](docs/AudienceRecipientRequest.md)
 - [AudienceRecipientResponse](docs/AudienceRecipientResponse.md)
 - [AudienceRecipientSegment](docs/AudienceRecipientSegment.md)
 - [AudienceRecipientSegmentCollectionResponse](docs/AudienceRecipientSegmentCollectionResponse.md)
 - [AudienceRecipientTag](docs/AudienceRecipientTag.md)
 - [AudienceRule](docs/AudienceRule.md)
 - [AudienceRuleProperty](docs/AudienceRuleProperty.md)
 - [AudienceSegment](docs/AudienceSegment.md)
 - [AudienceSegmentCollectionResponse](docs/AudienceSegmentCollectionResponse.md)
 - [AudienceSegmentRequest](docs/AudienceSegmentRequest.md)
 - [AudienceSegmentResponse](docs/AudienceSegmentResponse.md)
 - [AudienceTag](docs/AudienceTag.md)
 - [AudienceTagCollectionResponse](docs/AudienceTagCollectionResponse.md)
 - [AudienceTagRequest](docs/AudienceTagRequest.md)
 - [AudienceTagResponse](docs/AudienceTagResponse.md)
 - [AuthToken](docs/AuthToken.md)
 - [AuthTokenForServiceAccountRequest](docs/AuthTokenForServiceAccountRequest.md)
 - [AuthTokenResponse](docs/AuthTokenResponse.md)
 - [Collection](docs/Collection.md)
 - [CollectionLinks](docs/CollectionLinks.md)
 - [CollectionMeta](docs/CollectionMeta.md)
 - [EmailChannel](docs/EmailChannel.md)
 - [EmailChannelCollectionResponse](docs/EmailChannelCollectionResponse.md)
 - [EmailChannelRequest](docs/EmailChannelRequest.md)
 - [EmailChannelResponse](docs/EmailChannelResponse.md)
 - [EmailChannelRoutingRuleResponse](docs/EmailChannelRoutingRuleResponse.md)
 - [EmailChannelRoutingRules](docs/EmailChannelRoutingRules.md)
 - [EmailChannelRoutingRulesCollectionResponse](docs/EmailChannelRoutingRulesCollectionResponse.md)
 - [EmailNotification](docs/EmailNotification.md)
 - [EmailNotificationClick](docs/EmailNotificationClick.md)
 - [EmailNotificationCollectionResponse](docs/EmailNotificationCollectionResponse.md)
 - [EmailNotificationDetails](docs/EmailNotificationDetails.md)
 - [EmailNotificationOpen](docs/EmailNotificationOpen.md)
 - [EmailNotificationResponse](docs/EmailNotificationResponse.md)
 - [EmailNotificationStateInner](docs/EmailNotificationStateInner.md)
 - [EmailNotificationStateResponse](docs/EmailNotificationStateResponse.md)
 - [EmailRoutingRuleCreateRequest](docs/EmailRoutingRuleCreateRequest.md)
 - [EmailRoutingRuleUpdateRequest](docs/EmailRoutingRuleUpdateRequest.md)
 - [EmailSendBatchRequest](docs/EmailSendBatchRequest.md)
 - [EmailSendBatchResponse](docs/EmailSendBatchResponse.md)
 - [EmailSendBatchResponseDataInner](docs/EmailSendBatchResponseDataInner.md)
 - [EmailSendRequest](docs/EmailSendRequest.md)
 - [EmailSendRequestAttachmentsInner](docs/EmailSendRequestAttachmentsInner.md)
 - [EmailSendRequestFrom](docs/EmailSendRequestFrom.md)
 - [EmailSendRequestTemplate](docs/EmailSendRequestTemplate.md)
 - [EmailSendRequestTo](docs/EmailSendRequestTo.md)
 - [EmailSendResponse](docs/EmailSendResponse.md)
 - [EmailSendResponseData](docs/EmailSendResponseData.md)
 - [EmailSender](docs/EmailSender.md)
 - [EmailSenderCollectionResponse](docs/EmailSenderCollectionResponse.md)
 - [EmailSenderCreateRequest](docs/EmailSenderCreateRequest.md)
 - [EmailSenderResponse](docs/EmailSenderResponse.md)
 - [EmailSenderUpdateRequest](docs/EmailSenderUpdateRequest.md)
 - [EmailTemplate](docs/EmailTemplate.md)
 - [EmailTemplateCollectionResponse](docs/EmailTemplateCollectionResponse.md)
 - [EmailTemplateContent](docs/EmailTemplateContent.md)
 - [EmailTemplateContentCollectionResponse](docs/EmailTemplateContentCollectionResponse.md)
 - [EmailTemplateContentCreateRequest](docs/EmailTemplateContentCreateRequest.md)
 - [EmailTemplateContentResponse](docs/EmailTemplateContentResponse.md)
 - [EmailTemplateContentUpdateRequest](docs/EmailTemplateContentUpdateRequest.md)
 - [EmailTemplateContentVariablesSettingsInner](docs/EmailTemplateContentVariablesSettingsInner.md)
 - [EmailTemplateContentVariablesSettingsInnerRecipientProperty](docs/EmailTemplateContentVariablesSettingsInnerRecipientProperty.md)
 - [EmailTemplateCreateRequest](docs/EmailTemplateCreateRequest.md)
 - [EmailTemplateResponse](docs/EmailTemplateResponse.md)
 - [EmailTemplateUpdateRequest](docs/EmailTemplateUpdateRequest.md)
 - [EmailTransport](docs/EmailTransport.md)
 - [EmailTransportCollectionResponse](docs/EmailTransportCollectionResponse.md)
 - [EmailTransportCreateRequest](docs/EmailTransportCreateRequest.md)
 - [EmailTransportResponse](docs/EmailTransportResponse.md)
 - [EmailTransportUpdateRequest](docs/EmailTransportUpdateRequest.md)
 - [Error](docs/Error.md)
 - [ErrorMultiple](docs/ErrorMultiple.md)
 - [Event](docs/Event.md)
 - [EventResponse](docs/EventResponse.md)
 - [LocaleCode](docs/LocaleCode.md)
 - [PushNotification](docs/PushNotification.md)
 - [PushNotificationCollectionResponse](docs/PushNotificationCollectionResponse.md)
 - [PushNotificationResponse](docs/PushNotificationResponse.md)
 - [PushNotificationSendRequest](docs/PushNotificationSendRequest.md)
 - [PushNotificationSendRequestTemplate](docs/PushNotificationSendRequestTemplate.md)
 - [PushNotificationSendRequestTo](docs/PushNotificationSendRequestTo.md)
 - [PushNotificationSentResponse](docs/PushNotificationSentResponse.md)
 - [PushNotificationStateResponse](docs/PushNotificationStateResponse.md)
 - [PushPlatform](docs/PushPlatform.md)
 - [PushPlatformCollectionResponse](docs/PushPlatformCollectionResponse.md)
 - [PushPlatformRequest](docs/PushPlatformRequest.md)
 - [PushPlatformResponse](docs/PushPlatformResponse.md)
 - [PushTemplate](docs/PushTemplate.md)
 - [PushTemplateCollectionResponse](docs/PushTemplateCollectionResponse.md)
 - [PushTemplateContent](docs/PushTemplateContent.md)
 - [PushTemplateContentCollectionResponse](docs/PushTemplateContentCollectionResponse.md)
 - [PushTemplateContentRequest](docs/PushTemplateContentRequest.md)
 - [PushTemplateContentResponse](docs/PushTemplateContentResponse.md)
 - [PushTemplateRequest](docs/PushTemplateRequest.md)
 - [PushTemplateResponse](docs/PushTemplateResponse.md)
 - [PushTemplateUpdateRequest](docs/PushTemplateUpdateRequest.md)
 - [SmsChannel](docs/SmsChannel.md)
 - [SmsChannelCollectionResponse](docs/SmsChannelCollectionResponse.md)
 - [SmsChannelRequest](docs/SmsChannelRequest.md)
 - [SmsChannelResponse](docs/SmsChannelResponse.md)
 - [SmsMultipleNotificationSentResponse](docs/SmsMultipleNotificationSentResponse.md)
 - [SmsNotification](docs/SmsNotification.md)
 - [SmsNotificationCollectionResponse](docs/SmsNotificationCollectionResponse.md)
 - [SmsNotificationResponse](docs/SmsNotificationResponse.md)
 - [SmsNotificationSendMultipleRequest](docs/SmsNotificationSendMultipleRequest.md)
 - [SmsNotificationSendRequest](docs/SmsNotificationSendRequest.md)
 - [SmsNotificationSendRequestTemplate](docs/SmsNotificationSendRequestTemplate.md)
 - [SmsNotificationSendRequestTo](docs/SmsNotificationSendRequestTo.md)
 - [SmsNotificationSentResponse](docs/SmsNotificationSentResponse.md)
 - [SmsNotificationStateResponse](docs/SmsNotificationStateResponse.md)
 - [SmsRoutingRule](docs/SmsRoutingRule.md)
 - [SmsRoutingRuleCollectionResponse](docs/SmsRoutingRuleCollectionResponse.md)
 - [SmsRoutingRuleRequest](docs/SmsRoutingRuleRequest.md)
 - [SmsRoutingRuleResponse](docs/SmsRoutingRuleResponse.md)
 - [SmsSender](docs/SmsSender.md)
 - [SmsSenderCollectionResponse](docs/SmsSenderCollectionResponse.md)
 - [SmsSenderResponse](docs/SmsSenderResponse.md)
 - [SmsTemplate](docs/SmsTemplate.md)
 - [SmsTemplateCollectionResponse](docs/SmsTemplateCollectionResponse.md)
 - [SmsTemplateContent](docs/SmsTemplateContent.md)
 - [SmsTemplateContentCollectionResponse](docs/SmsTemplateContentCollectionResponse.md)
 - [SmsTemplateContentRequest](docs/SmsTemplateContentRequest.md)
 - [SmsTemplateContentResponse](docs/SmsTemplateContentResponse.md)
 - [SmsTemplateRequest](docs/SmsTemplateRequest.md)
 - [SmsTemplateResponse](docs/SmsTemplateResponse.md)
 - [SmsTransport](docs/SmsTransport.md)
 - [SmsTransportCollectionResponse](docs/SmsTransportCollectionResponse.md)
 - [SmsTransportCreateRequest](docs/SmsTransportCreateRequest.md)
 - [SmsTransportResponse](docs/SmsTransportResponse.md)
 - [SmsTransportUpdateRequest](docs/SmsTransportUpdateRequest.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionCollectionResponse](docs/SubscriptionCollectionResponse.md)
 - [SubscriptionCreateRequest](docs/SubscriptionCreateRequest.md)
 - [SubscriptionResponse](docs/SubscriptionResponse.md)
 - [SubscriptionUpdateRequest](docs/SubscriptionUpdateRequest.md)
 - [TransportWebhookGenerateResponse](docs/TransportWebhookGenerateResponse.md)
 - [TransportWebhookGenerateResponseData](docs/TransportWebhookGenerateResponseData.md)


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
