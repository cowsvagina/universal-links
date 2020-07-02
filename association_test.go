package uni_links

import (
	"testing"
)

func TestAssociation_Marshal(t *testing.T) {
	cases := []struct {
		Association AssociationMarshaler
		Expect      string
	}{
		{
			Association: Association{
				AppLinks: AppLinks{
					Details: nil,
				},
				WebCredentials: nil,
			},
			Expect: `{"applinks":{"details":[]}}`,
		},
		{
			Association: Association{
				AppLinks: AppLinks{
					Details: nil,
				},
				WebCredentials: &WebCredentials{Apps: nil},
			},
			Expect: `{"applinks":{"details":[]},"webcredentials":{"apps":[]}}`,
		},
		{
			Association: Association{
				AppLinks: AppLinks{
					Details: []Detail{
						{
							AppIDs:     nil,
							Components: nil,
						},
					},
				},
			},
			Expect: `{"applinks":{"details":[{"appIDs":[],"components":[]}]}}`,
		},
		{
			Association: Association{
				AppLinks: AppLinks{
					Details: []Detail{
						{
							AppIDs: []string{"ABCDE12345.com.example.app"},
						},
					},
				},
			},
			Expect: `{"applinks":{"details":[{"appIDs":["ABCDE12345.com.example.app"],"components":[]}]}}`,
		},
		{
			Association: Association{
				AppLinks: AppLinks{
					Details: []Detail{
						{
							AppIDs: []string{"ABCDE12345.com.example.app", "ABCDE12345.com.example.app2"},
							Components: []Component{
								{
									Fragment: "no_universal_links",
									Exclude:  true,
									Comment:  "Matches any URL whose fragment equals no_universal_links and instructs the system not to open it as a universal link",
								},
								{
									Path:    "/buy/*",
									Comment: "Matches any URL whose path starts with /buy/",
								},
								{
									Path:    "/help/website/*",
									Exclude: true,
									Comment: "Matches any URL whose path starts with /help/website/ and instructs the system not to open it as a universal link",
								},
								{
									Path:    "/help/*",
									Query:   map[string]string{"articleNumber": "????"},
									Comment: "Matches any URL whose path starts with /help/ and which has a query item with name 'articleNumber' and a value of exactly 4 characters",
								},
							},
						},
					},
				},
			},
			Expect: `{"applinks":{"details":[{"appIDs":["ABCDE12345.com.example.app","ABCDE12345.com.example.app2"],"components":[{"#":"no_universal_links","exclude":true,"comment":"Matches any URL whose fragment equals no_universal_links and instructs the system not to open it as a universal link"},{"/":"/buy/*","comment":"Matches any URL whose path starts with /buy/"},{"/":"/help/website/*","exclude":true,"comment":"Matches any URL whose path starts with /help/website/ and instructs the system not to open it as a universal link"},{"/":"/help/*","?":{"articleNumber":"????"},"comment":"Matches any URL whose path starts with /help/ and which has a query item with name 'articleNumber' and a value of exactly 4 characters"}]}]}}`,
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
