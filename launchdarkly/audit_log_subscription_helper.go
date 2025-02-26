package launchdarkly

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	strcase "github.com/stoewer/go-strcase"
)

var KEBAB_CASE_INTEGRATIONS = []string{"splunk"}

type IntegrationConfig map[string]FormVariable

type FormVariable struct {
	Type          string
	IsOptional    *bool
	AllowedValues *[]string
	DefaultValue  *interface{}
	IsSecret      *bool
}

func auditLogSubscriptionSchema(isDataSource bool) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		INTEGRATION_KEY: {
			// validated as part of the config validation
			Type:     schema.TypeString,
			Required: true,
			// we are omitting appdynamics for now because it requires oauth
			ValidateFunc: validation.StringNotInSlice([]string{"appdynamics"}, false),
			ForceNew:     true,
		},
		NAME: {
			Type:     schema.TypeString,
			Required: !isDataSource,
			Optional: isDataSource,
		},
		CONFIG: {
			Type:     schema.TypeMap,
			Required: !isDataSource,
			Optional: isDataSource,
		},
		STATEMENTS: policyStatementsSchema(policyStatementSchemaOptions{required: !isDataSource}),
		ON: {
			Type:     schema.TypeBool,
			Required: !isDataSource,
			Optional: isDataSource,
		},
		TAGS: tagsSchema(),
	}
}

func parseAuditLogSubscriptionConfigs() map[string]IntegrationConfig {
	// SUBSCRIPTION_CONFIGURATION_FIELDS can be found in audit_log_subscription_configs.go
	configs := make(map[string]IntegrationConfig, len(SUBSCRIPTION_CONFIGURATION_FIELDS))
	for integrationKey, rawVariables := range SUBSCRIPTION_CONFIGURATION_FIELDS {
		cfg := IntegrationConfig{}
		variables := rawVariables.(map[string]interface{})
		for k, v := range variables {
			variable := v.(map[string]interface{})
			formVariable := FormVariable{Type: variable["type"].(string)}
			if variable["isOptional"] != nil {
				isOptional := variable["isOptional"].(bool)
				formVariable.IsOptional = &isOptional
			}
			if variable["allowedValues"] != nil {
				rawValues := variable["allowedValues"].([]interface{})
				var allowedValues []string
				for _, value := range rawValues {
					allowedValues = append(allowedValues, value.(string))
				}
				formVariable.AllowedValues = &allowedValues
			}
			if variable["isSecret"] != nil {
				isSecret := variable["isSecret"].(bool)
				formVariable.IsSecret = &isSecret
			}
			if variable["defaultValue"] != nil {
				defaultValue := variable["defaultValue"]
				formVariable.DefaultValue = &defaultValue
			}
			cfg[k] = formVariable
		}
		configs[integrationKey] = cfg
	}
	return configs
}

func getConfigFieldKey(integrationKey, resourceKey string) string {
	// a select number of integrations take fields in kebab case, ex. "skip-ca-verification"
	// currently this only applies to splunk
	for _, integration := range KEBAB_CASE_INTEGRATIONS {
		if integrationKey == integration {
			return strcase.KebabCase(resourceKey)
		}
	}
	return strcase.LowerCamelCase(resourceKey)
}

// configFromResourceData uses the configuration generated into audit_log_subscription_config.json
// to validate and generate the config the API expects
func configFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	// TODO: refactor to return list of diags warnings with all formatting errors
	integrationKey := d.Get(INTEGRATION_KEY).(string)
	config := d.Get(CONFIG).(map[string]interface{})
	configMap := parseAuditLogSubscriptionConfigs()
	configFormat, ok := configMap[integrationKey]
	if !ok {
		return config, fmt.Errorf("%s is not a valid integration_key for audit log subscriptions", integrationKey)
	}
	for k := range config {
		// error if an incorrect config variable has been set
		key := getConfigFieldKey(integrationKey, k) // convert casing to compare to required config format
		if integrationKey == "datadog" && key == "hostUrl" {
			// this is a one-off for now
			key = "hostURL"
		}
		if _, ok := configFormat[key]; !ok {
			return config, fmt.Errorf("config variable %s not valid for integration type %s", k, integrationKey)
		}
	}
	convertedConfig := make(map[string]interface{}, len(config))
	for k, v := range configFormat {
		key := strcase.SnakeCase(k) // convert to snake case to validate user config
		rawValue, ok := config[key]
		if !ok {
			if !*v.IsOptional {
				return config, fmt.Errorf("config variable %s must be set", key)
			}
			// we will let the API handle default configs for now since it otherwise messes
			// up the plan if we set an attribute a user has not set on a non-computed attribute
			continue
		}
		// type will be one of ["string", "boolean", "uri", "enum", "oauth", "dynamicEnum"]
		// for now we do not need to handle oauth or dynamicEnum
		switch v.Type {
		case "string", "uri":
			// we'll let the API handle the URI validation for now
			value := rawValue.(string)
			convertedConfig[k] = value
		case "boolean":
			value, err := strconv.ParseBool(rawValue.(string)) // map values may only be one type, so all non-string types have to be converted
			if err != nil {
				return config, fmt.Errorf("config value %s for %v must be of type bool", rawValue, k)
			}
			convertedConfig[k] = value
		case "enum":
			value := rawValue.(string)
			if !stringInSlice(value, *v.AllowedValues) {
				return config, fmt.Errorf("config value %s for %v must be one of the following approved string values: %v", rawValue, k, *v.AllowedValues)
			}
			convertedConfig[k] = value
		default:
			// just set to the existing value
			convertedConfig[k] = rawValue
		}
	}
	return convertedConfig, nil
}

func configToResourceData(d *schema.ResourceData, config map[string]interface{}) (map[string]interface{}, error) {
	integrationKey := d.Get(INTEGRATION_KEY).(string)
	configMap := parseAuditLogSubscriptionConfigs()
	configFormat, ok := configMap[integrationKey]
	if !ok {
		return config, fmt.Errorf("%s is not a currently supported integration_key for audit log subscriptions", integrationKey)
	}
	originalConfig := d.Get(CONFIG).(map[string]interface{})
	convertedConfig := make(map[string]interface{}, len(config))
	for k, v := range config {
		key := strcase.SnakeCase(k)
		// some attributes have defaults that the API will return and terraform will complain since config
		// is not a computed attribute (cannot be both required & computed)
		// TODO: handle this in a SuppressDiff function
		if _, setByUser := originalConfig[key]; !setByUser {
			continue
		}
		convertedConfig[key] = v
		if value, isBool := v.(bool); isBool {
			convertedConfig[key] = strconv.FormatBool(value)
		}
		if *configFormat[k].IsSecret {
			// if the user didn't put it in as obfuscated, we don't want to set it as obfuscated
			convertedConfig[key] = originalConfig[key]
		}
	}
	return convertedConfig, nil
}

func auditLogSubscriptionRead(ctx context.Context, d *schema.ResourceData, metaRaw interface{}, isDataSource bool) diag.Diagnostics {
	var diags diag.Diagnostics
	client := metaRaw.(*Client)
	var id string
	if isDataSource {
		id = d.Get(ID).(string)
	} else {
		id = d.Id()
	}
	integrationKey := d.Get(INTEGRATION_KEY).(string)

	sub, res, err := client.ld.IntegrationAuditLogSubscriptionsApi.GetSubscriptionByID(client.ctx, integrationKey, id).Execute()

	if isStatusNotFound(res) && !isDataSource {
		log.Printf("[WARN] failed to find integration with ID %q, removing from state if present", id)
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("[WARN] failed to find integration with ID %q, removing from state if present", id),
		})
		d.SetId("")
		return diags
	}
	if err != nil {
		return diag.Errorf("failed to get integration with ID %q: %v", id, err)
	}

	if isDataSource {
		d.SetId(*sub.Id)
	}

	_ = d.Set(INTEGRATION_KEY, sub.Kind)
	_ = d.Set(NAME, sub.Name)
	_ = d.Set(ON, sub.On)
	cfg, err := configToResourceData(d, *sub.Config)
	if err != nil {
		return diag.Errorf("failed to set config on integration with id %q: %v", *sub.Id, err)
	}
	err = d.Set(CONFIG, cfg)
	if err != nil {
		return diag.Errorf("failed to set config on integration with id %q: %v", *sub.Id, err)
	}
	err = d.Set(STATEMENTS, policyStatementsToResourceData(*sub.Statements))
	if err != nil {
		return diag.Errorf("failed to set statements on integration with id %q: %v", *sub.Id, err)
	}
	err = d.Set(TAGS, sub.Tags)
	if err != nil {
		return diag.Errorf("failed to set tags on integration with id %q: %v", *sub.Id, err)
	}
	return diags
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
