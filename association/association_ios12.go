package association

import (
	jsoniter "github.com/json-iterator/go"
)

// AssociationIOS12 iOS 12及以下版本的通用连接关联配置
// reference: https://developer.apple.com/documentation/uikit/inter-process_communication/allowing_apps_and_websites_to_link_to_your_content/enabling_universal_links#3381438
type AssociationIOS12 struct {
	AppLinks       appLinksIOS12   `json:"applinks"`
	WebCredentials *webCredentials `json:"webcredentials,omitempty"`
}

func (a AssociationIOS12) Marshal() (string, error) {
	if a.AppLinks.Apps == nil {
		a.AppLinks.Apps = make([]string, 0)
	}
	if a.AppLinks.Details == nil {
		a.AppLinks.Details = make([]detailIOS12, 0)
	} else {
		for i, _ := range a.AppLinks.Details {
			if a.AppLinks.Details[i].Paths == nil {
				a.AppLinks.Details[i].Paths = make([]string, 0)
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

// appLinksIOS12 applinks字段
// Apps应该始终为空数组
// see https://developer.apple.com/documentation/uikit/inter-process_communication/allowing_apps_and_websites_to_link_to_your_content/enabling_universal_links#3002229
// > apps:
// >     This key is not used for universal links, but it must be present and set to an empty array, as shown in Listing 1.
type appLinksIOS12 struct {
	Apps    []string      `json:"apps"`
	Details []detailIOS12 `json:"details"`
}

// detailIOS12 detail信息
type detailIOS12 struct {
	AppID string   `json:"appID"`
	Paths []string `json:"paths"`
}
