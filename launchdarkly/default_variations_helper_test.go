package launchdarkly

import (
	"errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ldapi "github.com/launchdarkly/api-client-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultVariationsFromResourceData(t *testing.T) {
	testCases := []struct {
		name        string
		vars        map[string]interface{}
		expected    *ldapi.Defaults
		expectedErr error
	}{
		{
			name: "no defaults on boolean",
			vars: map[string]interface{}{
				VARIATIONS: []interface{}{
					map[string]interface{}{
						VALUE: "true",
					},
					map[string]interface{}{
						VALUE: "false",
					},
				},
			},
			expected: &ldapi.Defaults{
				OnVariation:  0,
				OffVariation: 1,
			},
		},
		{
			name: "basic defaults",
			vars: map[string]interface{}{
				VARIATIONS: []interface{}{
					map[string]interface{}{
						VALUE: "true",
					},
					map[string]interface{}{
						VALUE: "false",
					},
				},
				DEFAULTS: []interface{}{
					map[string]interface{}{
						ON_VARIATION:  0,
						OFF_VARIATION: 1,
					}},
			},
			expected: &ldapi.Defaults{
				OnVariation:  0,
				OffVariation: 1,
			},
		},
		{
			name: "invalid default on value",
			vars: map[string]interface{}{
				VARIATION_TYPE: "boolean",
				VARIATIONS: []interface{}{
					map[string]interface{}{
						VALUE: "true",
					},
					map[string]interface{}{
						VALUE: "false",
					},
				},
				DEFAULTS: []interface{}{
					map[string]interface{}{
						ON_VARIATION:  2,
						OFF_VARIATION: 1,
					}},
			},
			expectedErr: errors.New(`default on_variation 2 is out of range, must be between 0 and 1 inclusive`),
		},
		{
			name: "invalid default off value",
			vars: map[string]interface{}{
				VARIATION_TYPE: "boolean",
				VARIATIONS: []interface{}{
					map[string]interface{}{
						VALUE: "true",
					},
					map[string]interface{}{
						VALUE: "false",
					},
				},
				DEFAULTS: []interface{}{
					map[string]interface{}{
						ON_VARIATION:  0,
						OFF_VARIATION: 5,
					}},
			},
			expectedErr: errors.New(`default off_variation 5 is out of range, must be between 0 and 1 inclusive`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resourceData := schema.TestResourceDataRaw(t,
				map[string]*schema.Schema{VARIATION_TYPE: variationTypeSchema(), VARIATIONS: variationsSchema(),
					DEFAULTS: {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								ON_VARIATION: {
									Type:     schema.TypeInt,
									Required: true,
								},
								OFF_VARIATION: {
									Type:     schema.TypeInt,
									Required: true,
								},
							},
						},
					}},
				tc.vars,
			)

			actual, err := defaultVariationsFromResourceData(resourceData)
			require.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
