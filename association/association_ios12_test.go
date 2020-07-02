package association

import (
	"testing"
)

func TestAssociationIOS12_Marshal(t *testing.T) {
	cases := []struct {
		Association AssociationMarshaler
		Expect      string
	}{
		{
			Association: AssociationIOS12{
				AppLinks: appLinksIOS12{
					Apps:    nil,
					Details: nil,
				},
				WebCredentials: nil,
			},
			Expect: `{"applinks":{"apps":[],"details":[]}}`,
		},
		{
			Association: AssociationIOS12{
				AppLinks: appLinksIOS12{
					Apps:    nil,
					Details: nil,
				},
				WebCredentials: &webCredentials{Apps: nil},
			},
			Expect: `{"applinks":{"apps":[],"details":[]},"webcredentials":{"apps":[]}}`,
		},
		{
			Association: AssociationIOS12{
				AppLinks: appLinksIOS12{
					Apps: nil,
					Details: []detailIOS12{
						{
							AppID: "ABCDE12345.com.example.app",
						},
					},
				},
			},
			Expect: `{"applinks":{"apps":[],"details":[{"appID":"ABCDE12345.com.example.app","paths":[]}]}}`,
		},
		{
			Association: AssociationIOS12{
				AppLinks: appLinksIOS12{
					Apps: nil,
					Details: []detailIOS12{
						{
							AppID: "ABCDE12345.com.example.app",
							Paths: []string{"/buy/*", "NOT /help/website/*", "/help/*"},
						},
					},
				},
				WebCredentials: &webCredentials{
					Apps: []string{"ABCDE12345.com.example.app"},
				},
			},
			Expect: `{"applinks":{"apps":[],"details":[{"appID":"ABCDE12345.com.example.app","paths":["/buy/*","NOT /help/website/*","/help/*"]}]},"webcredentials":{"apps":["ABCDE12345.com.example.app"]}}`,
		},
	}

	for i, each := range cases {
		actual, err := each.Association.Marshal()
		if err != nil {
			t.Fatalf("%d - error: %s", i, err)
		}
		if actual != each.Expect {
			t.Fatalf("%d - expect: %s, got: %s", i, each.Expect, actual)
		}
	}
}
