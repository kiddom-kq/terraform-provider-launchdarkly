package launchdarkly

// to get the updated SUBSCRIPTION_CONFIGURATION_FIELDS value, paste the generated json in
// audit_log_subscription_configs.json into https://rodrigo-brito.github.io/json-to-go-map/

// TODO: generate this automatically
// func parseAuditLogSubscriptionConfigsFromJson() (map[string]IntegrationConfig, error) {
// 	var configs map[string]IntegrationConfig
// 	file, err := ioutil.ReadFile(CONFIG_FILE)
// 	if err != nil {
// 		return configs, err
// 	}

// 	err = json.Unmarshal([]byte(file), &configs)
// 	if err != nil {
// 		return configs, err
// 	}
// 	return configs, nil
// }

var SUBSCRIPTION_CONFIGURATION_FIELDS = map[string]interface{}{
	"appdynamics": map[string]interface{}{
		"account": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
		"applicationID": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
	},
	"datadog": map[string]interface{}{
		"apiKey": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      true,
		},
		"hostURL": map[string]interface{}{
			"type":       "enum",
			"isOptional": true,
			"allowedValues": []interface{}{
				"https://api.datadoghq.com",
				"https://api.datadoghq.eu",
			},
			"defaultValue": "https://api.datadoghq.com",
			"isSecret":     false,
		},
	},
	"dynatrace": map[string]interface{}{
		"apiToken": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      true,
		},
		"url": map[string]interface{}{
			"type":          "uri",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
		"entity": map[string]interface{}{
			"type":       "enum",
			"isOptional": true,
			"allowedValues": []interface{}{
				"APPLICATION",
				"APPLICATION_METHOD",
				"APPLICATION_METHOD_GROUP",
				"AUTO_SCALING_GROUP",
				"AUXILIARY_SYNTHETIC_TEST",
				"AWS_APPLICATION_LOAD_BALANCER",
				"AWS_AVAILABILITY_ZONE",
				"AWS_CREDENTIALS",
				"AWS_LAMBDA_FUNCTION",
				"AWS_NETWORK_LOAD_BALANCER",
				"AZURE_API_MANAGEMENT_SERVICE",
				"AZURE_APPLICATION_GATEWAY",
				"AZURE_COSMOS_DB",
				"AZURE_CREDENTIALS",
				"AZURE_EVENT_HUB",
				"AZURE_EVENT_HUB_NAMESPACE",
				"AZURE_FUNCTION_APP",
				"AZURE_IOT_HUB",
				"AZURE_LOAD_BALANCER",
				"AZURE_MGMT_GROUP",
				"AZURE_REDIS_CACHE",
				"AZURE_REGION",
				"AZURE_SERVICE_BUS_NAMESPACE",
				"AZURE_SERVICE_BUS_QUEUE",
				"AZURE_SERVICE_BUS_TOPIC",
				"AZURE_SQL_DATABASE",
				"AZURE_SQL_ELASTIC_POOL",
				"AZURE_SQL_SERVER",
				"AZURE_STORAGE_ACCOUNT",
				"AZURE_SUBSCRIPTION",
				"AZURE_TENANT",
				"AZURE_VM",
				"AZURE_VM_SCALE_SET",
				"AZURE_WEB_APP",
				"CF_APPLICATION",
				"CF_FOUNDATION",
				"CINDER_VOLUME",
				"CLOUD_APPLICATION",
				"CLOUD_APPLICATION_INSTANCE",
				"CLOUD_APPLICATION_NAMESPACE",
				"CONTAINER_GROUP",
				"CONTAINER_GROUP_INSTANCE",
				"CUSTOM_APPLICATION",
				"CUSTOM_DEVICE",
				"CUSTOM_DEVICE_GROUP",
				"DCRUM_APPLICATION",
				"DCRUM_SERVICE",
				"DCRUM_SERVICE_INSTANCE",
				"DEVICE_APPLICATION_METHOD",
				"DISK",
				"DOCKER_CONTAINER_GROUP_INSTANCE",
				"DYNAMO_DB_TABLE",
				"EBS_VOLUME",
				"EC2_INSTANCE",
				"ELASTIC_LOAD_BALANCER",
				"ENVIRONMENT",
				"EXTERNAL_SYNTHETIC_TEST_STEP",
				"GCP_ZONE",
				"GEOLOCATION",
				"GEOLOC_SITE",
				"GOOGLE_COMPUTE_ENGINE",
				"HOST",
				"HOST_GROUP",
				"HTTP_CHECK",
				"HTTP_CHECK_STEP",
				"HYPERVISOR",
				"KUBERNETES_CLUSTER",
				"KUBERNETES_NODE",
				"MOBILE_APPLICATION",
				"NETWORK_INTERFACE",
				"NEUTRON_SUBNET",
				"OPENSTACK_PROJECT",
				"OPENSTACK_REGION",
				"OPENSTACK_VM",
				"OS",
				"PROCESS_GROUP",
				"PROCESS_GROUP_INSTANCE",
				"RELATIONAL_DATABASE_SERVICE",
				"SERVICE",
				"SERVICE_INSTANCE",
				"SERVICE_METHOD",
				"SERVICE_METHOD_GROUP",
				"SWIFT_CONTAINER",
				"SYNTHETIC_LOCATION",
				"SYNTHETIC_TEST",
				"SYNTHETIC_TEST_STEP",
				"VIRTUALMACHINE",
				"VMWARE_DATACENTER",
			},
			"defaultValue": "APPLICATION",
			"isSecret":     false,
		},
	},
	"elastic": map[string]interface{}{
		"url": map[string]interface{}{
			"type":          "uri",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
		"token": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      true,
		},
		"index": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
	},
	"honeycomb": map[string]interface{}{
		"datasetName": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
		"apiKey": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      true,
		},
	},
	"logdna": map[string]interface{}{
		"ingestionKey": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      true,
		},
		"level": map[string]interface{}{
			"type":          "string",
			"isOptional":    true,
			"allowedValues": nil,
			"defaultValue":  "INFO",
			"isSecret":      false,
		},
	},
	"msteams": map[string]interface{}{
		"url": map[string]interface{}{
			"type":          "uri",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
	},
	"new-relic-apm": map[string]interface{}{
		"apiKey": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      true,
		},
		"applicationId": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
		"domain": map[string]interface{}{
			"type":       "enum",
			"isOptional": true,
			"allowedValues": []interface{}{
				"api.newrelic.com",
				"api.eu.newrelic.com",
			},
			"defaultValue": "api.newrelic.com",
			"isSecret":     false,
		},
	},
	"signalfx": map[string]interface{}{
		"accessToken": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      true,
		},
		"realm": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
	},
	"splunk": map[string]interface{}{
		"base-url": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
		"token": map[string]interface{}{
			"type":          "string",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      true,
		},
		"skip-ca-verification": map[string]interface{}{
			"type":          "boolean",
			"isOptional":    false,
			"allowedValues": nil,
			"defaultValue":  nil,
			"isSecret":      false,
		},
	},
}
