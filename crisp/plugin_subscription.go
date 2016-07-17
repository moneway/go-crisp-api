// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// PluginSubscriptionListData mapping
type PluginSubscriptionListData struct {
  Data  *[]PluginSubscription  `json:"data,omitempty"`
}

// PluginSubscriptionData mapping
type PluginSubscriptionData struct {
  Data  *PluginSubscription  `json:"data,omitempty"`
}

// PluginSubscription mapping
type PluginSubscription struct {
  ID           *string    `json:"id,omitempty"`
  URN          *string    `json:"urn,omitempty"`
  Type         *string    `json:"type,omitempty"`
  Name         *string    `json:"name,omitempty"`
  Description  *string    `json:"description,omitempty"`
  Features     *[]string  `json:"features,omitempty"`
  Showcase     *[]string  `json:"showcase,omitempty"`
  Price        *uint      `json:"price,omitempty"`
  Color        *string    `json:"color,omitempty"`
  Icon         *string    `json:"icon,omitempty"`
  Banner       *string    `json:"banner,omitempty"`
  Since        *string    `json:"since,omitempty"`
  Active       *bool      `json:"active,omitempty"`
  WebsiteID    *string    `json:"website_id,omitempty"`
  CardID       *string    `json:"card_id,omitempty"`
}

// PluginSubscriptionCreate mapping
type PluginSubscriptionCreate struct {
  PluginID  *string  `json:"plugin_id,omitempty"`
}

// PluginSubscriptionSettingsData mapping
type PluginSubscriptionSettingsData struct {
  Data  *PluginSubscriptionSettings  `json:"data,omitempty"`
}

// PluginSubscriptionSettings mapping
type PluginSubscriptionSettings struct {
  PluginID   *string       `json:"plugin_id,omitempty"`
  WebsiteID  *string       `json:"website_id,omitempty"`
  Schema     *interface{}  `json:"schema,omitempty"`
  Settings   *interface{}  `json:"settings,omitempty"`
}


// ListAllActiveSubscriptions lists all active plugin subscriptions on all websites, linked to payment methods owned by the user.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-subscription-get
func (service *PluginService) ListAllActiveSubscriptions() (*[]PluginSubscription, *Response, error) {
  url := "plugins/subscription"
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginSubscriptionListData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// ListSubscriptionsForWebsite lists plugin subscriptions for given website.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-subscription-get-1
func (service *PluginService) ListSubscriptionsForWebsite(websiteID string) (*PluginSubscription, *Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginSubscriptionData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// GetSubscriptionDetails resolves details on a given subscription.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-subscription-get-2
func (service *PluginService) GetSubscriptionDetails(websiteID string, pluginID string) (*PluginSubscription, *Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s", websiteID, pluginID)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginSubscriptionData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// SubscribeWebsiteToPlugin subscribes a given website to a given plugin.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-subscription-post
func (service *PluginService) SubscribeWebsiteToPlugin(websiteID string, pluginID string) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, PluginSubscriptionCreate{PluginID: &pluginID})

  return service.client.Do(req, nil)
}


// UnsubscribePluginFromWebsite unsubscribes a given plugin from a given website.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-subscription-delete
func (service *PluginService) UnsubscribePluginFromWebsite(websiteID string, pluginID string) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s", websiteID, pluginID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// GetSubscriptionSettings resolves plugin subscription settings. Used to read plugin configuration on a given website.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-subscription-get-3
func (service *PluginService) GetSubscriptionSettings(websiteID string, pluginID string) (*PluginSubscriptionSettings, *Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/settings", websiteID, pluginID)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginSubscriptionSettingsData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// SaveSubscriptionSettings saves plugin subscription settings. Used to configure a given plugin on a given website.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-subscription-patch
func (service *PluginService) SaveSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/settings", websiteID, pluginID)
  req, _ := service.client.NewRequest("PATCH", url, settings)

  return service.client.Do(req, nil)
}
