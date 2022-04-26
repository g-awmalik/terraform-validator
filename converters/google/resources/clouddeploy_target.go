// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

const CloudDeployTargetAssetType string = "clouddeploy.googleapis.com/Target"

func resourceConverterCloudDeployTarget() ResourceConverter {
	return ResourceConverter{
		AssetType: CloudDeployTargetAssetType,
		Convert:   GetCloudDeployTargetCaiObject,
	}
}

func GetCloudDeployTargetCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//clouddeploy.googleapis.com/projects/{{project}}/locations/{{region}}/targets/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetCloudDeployTargetApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: CloudDeployTargetAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/clouddeploy/v1/rest",
				DiscoveryName:        "Target",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetCloudDeployTargetApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandCloudDeployTargetDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCloudDeployTargetLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandCloudDeployTargetAnnotations(d.Get("annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("annotations"); !isEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}
	requireApprovalProp, err := expandCloudDeployTargetRequireApproval(d.Get("require_approval"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("require_approval"); !isEmptyValue(reflect.ValueOf(requireApprovalProp)) && (ok || !reflect.DeepEqual(v, requireApprovalProp)) {
		obj["requireApproval"] = requireApprovalProp
	}
	gkeProp, err := expandCloudDeployTargetGke(d.Get("gke"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gke"); !isEmptyValue(reflect.ValueOf(gkeProp)) && (ok || !reflect.DeepEqual(v, gkeProp)) {
		obj["gke"] = gkeProp
	}
	anthosClusterProp, err := expandCloudDeployTargetAnthosCluster(d.Get("anthos_cluster"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("anthos_cluster"); !isEmptyValue(reflect.ValueOf(anthosClusterProp)) && (ok || !reflect.DeepEqual(v, anthosClusterProp)) {
		obj["anthosCluster"] = anthosClusterProp
	}
	executionConfigsProp, err := expandCloudDeployTargetExecutionConfigs(d.Get("execution_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("execution_configs"); !isEmptyValue(reflect.ValueOf(executionConfigsProp)) && (ok || !reflect.DeepEqual(v, executionConfigsProp)) {
		obj["executionConfigs"] = executionConfigsProp
	}

	return obj, nil
}

func expandCloudDeployTargetDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployTargetLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudDeployTargetAnnotations(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudDeployTargetRequireApproval(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployTargetGke(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCluster, err := expandCloudDeployTargetGkeCluster(original["cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCluster); val.IsValid() && !isEmptyValue(val) {
		transformed["cluster"] = transformedCluster
	}

	transformedInternalIp, err := expandCloudDeployTargetGkeInternalIp(original["internal_ip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInternalIp); val.IsValid() && !isEmptyValue(val) {
		transformed["internalIp"] = transformedInternalIp
	}

	return transformed, nil
}

func expandCloudDeployTargetGkeCluster(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployTargetGkeInternalIp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployTargetAnthosCluster(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMembership, err := expandCloudDeployTargetAnthosClusterMembership(original["membership"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMembership); val.IsValid() && !isEmptyValue(val) {
		transformed["membership"] = transformedMembership
	}

	return transformed, nil
}

func expandCloudDeployTargetAnthosClusterMembership(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployTargetExecutionConfigs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedUsages, err := expandCloudDeployTargetExecutionConfigsUsages(original["usages"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUsages); val.IsValid() && !isEmptyValue(val) {
			transformed["usages"] = transformedUsages
		}

		transformedServiceAccount, err := expandCloudDeployTargetExecutionConfigsServiceAccount(original["service_account"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedServiceAccount); val.IsValid() && !isEmptyValue(val) {
			transformed["serviceAccount"] = transformedServiceAccount
		}

		transformedArtifactStorage, err := expandCloudDeployTargetExecutionConfigsArtifactStorage(original["artifact_storage"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArtifactStorage); val.IsValid() && !isEmptyValue(val) {
			transformed["artifactStorage"] = transformedArtifactStorage
		}

		transformedWorkerPool, err := expandCloudDeployTargetExecutionConfigsWorkerPool(original["worker_pool"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedWorkerPool); val.IsValid() && !isEmptyValue(val) {
			transformed["workerPool"] = transformedWorkerPool
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudDeployTargetExecutionConfigsUsages(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployTargetExecutionConfigsServiceAccount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployTargetExecutionConfigsArtifactStorage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudDeployTargetExecutionConfigsWorkerPool(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}