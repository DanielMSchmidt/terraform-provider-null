package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func RequiresReplaceIfValuesNotNull() tfsdk.AttributePlanModifier {
	return requiresReplaceIfValuesNotNullModifier{}
}

type requiresReplaceIfValuesNotNullModifier struct{}

func (r requiresReplaceIfValuesNotNullModifier) Modify(ctx context.Context, req tfsdk.ModifyAttributePlanRequest, resp *tfsdk.ModifyAttributePlanResponse) {
	if req.AttributeConfig == nil || req.AttributePlan == nil || req.AttributeState == nil {
		// shouldn't happen, but let's not panic if it does
		return
	}

	if req.State.Raw.IsNull() {
		// if we're creating the resource, no need to delete and
		// recreate it
		return
	}

	if req.Plan.Raw.IsNull() {
		// if we're deleting the resource, no need to delete and
		// recreate it
		return
	}

	// If there are no differences, do not mark the resource for replacement
	// and ensure the plan matches the configuration.
	if req.AttributeConfig.Equal(req.AttributeState) {
		return
	}

	if req.AttributeState.IsNull() {
		// terraform-plugin-sdk would store maps as null if all keys had null
		// values. To prevent unintentional replacement plans when migrating
		// to terraform-plugin-framework, only trigger replacement when the
		// prior state (map) is null and when there are not null map values.
		allNullValues := true

		configMap, ok := req.AttributeConfig.(types.Map)

		if !ok {
			return
		}

		for _, configValue := range configMap.Elements() {
			if !configValue.IsNull() {
				allNullValues = false
			}
		}

		if allNullValues {
			return
		}
	} else {
		// terraform-plugin-sdk would completely omit storing map keys with
		// null values, so this also must prevent unintentional replacement
		// in that case as well.
		allNewNullValues := true

		configMap, ok := req.AttributeConfig.(types.Map)

		if !ok {
			return
		}

		stateMap, ok := req.AttributeState.(types.Map)

		if !ok {
			return
		}

		for configKey, configValue := range configMap.Elements() {
			stateValue, ok := stateMap.Elements()[configKey]

			// If the key doesn't exist in state and the config value is
			// null, do not trigger replacement.
			if !ok && configValue.IsNull() {
				continue
			}

			// If the state value exists and it is equal to the config value,
			// do not trigger replacement.
			if configValue.Equal(stateValue) {
				continue
			}

			allNewNullValues = false
			break
		}

		for stateKey := range stateMap.Elements() {
			_, ok := configMap.Elements()[stateKey]

			// If the key doesn't exist in the config, but there is a state
			// value, trigger replacement.
			if !ok {
				allNewNullValues = false
				break
			}
		}

		if allNewNullValues {
			return
		}
	}

	resp.RequiresReplace = true
}

// Description returns a human-readable description of the plan modifier.
func (r requiresReplaceIfValuesNotNullModifier) Description(ctx context.Context) string {
	return "If the value of this attribute changes, Terraform will destroy and recreate the resource."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (r requiresReplaceIfValuesNotNullModifier) MarkdownDescription(ctx context.Context) string {
	return "If the value of this attribute changes, Terraform will destroy and recreate the resource."
}
