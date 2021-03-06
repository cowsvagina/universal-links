package uni_links

// reference:
//		https://developer.apple.com/documentation/uikit/inter-process_communication/allowing_apps_and_websites_to_link_to_your_content/enabling_universal_links
//		https://developer.apple.com/documentation/uikit/inter-process_communication/allowing_apps_and_websites_to_link_to_your_content/handling_universal_links
//		https://developer.apple.com/documentation/safariservices/supporting_associated_domains_in_your_app

import (
	jsoniter "github.com/json-iterator/go"
)

type AssociationMarshaler interface {
	Marshal() (string, error)
}

// AssociationIOS12 iOS 12及以下版本的通用连接关联配置
// reference: https://developer.apple.com/documentation/uikit/inter-process_communication/allowing_apps_and_websites_to_link_to_your_content/enabling_universal_links#3381438
type Association struct {
	AppLinks       AppLinks        `json:"applinks"`
	WebCredentials *WebCredentials `json:"webcredentials,omitempty"`
}

func (a Association) Marshal() (string, error) {
	if a.AppLinks.Details == nil {
		a.AppLinks.Details = make([]Detail, 0)
	} else {
		for i, _ := range a.AppLinks.Details {
			if a.AppLinks.Details[i].AppIDs == nil {
				a.AppLinks.Details[i].AppIDs = make([]string, 0)
			}
			if a.AppLinks.Details[i].Components == nil {
				a.AppLinks.Details[i].Components = make([]Component, 0)
			}
		}
	}
	if a.WebCredentials != nil && a.WebCredentials.Apps == nil {
		a.WebCredentials.Apps = make([]string, 0)
	}

	v, err := jsoniter.Marshal(&a)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

// AppLinks applinks字段
// Apps应该始终为空数组
// see https://developer.apple.com/documentation/uikit/inter-process_communication/allowing_apps_and_websites_to_link_to_your_content/enabling_universal_links#3002229
// > apps:
// >     This key is not used for universal links, but it must be present and set to an empty array, as shown in Listing 1.
type AppLinks struct {
	Details []Detail `json:"details"`
}

// Detail detail信息
type Detail struct {
	AppIDs     []string    `json:"appIDs"`
	Components []Component `json:"components"`
}

type Component struct {
	Path     string            `json:"/,omitempty"`
	Query    map[string]string `json:"?,omitempty"`
	Fragment string            `json:"#,omitempty"`
	Exclude  bool              `json:"exclude,omitempty"`
	Comment  string            `json:"comment,omitempty"`
}

type WebCredentials struct {
	Apps []string `json:"apps"`
}
