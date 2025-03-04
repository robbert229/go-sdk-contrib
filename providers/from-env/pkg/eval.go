package from_env

import (
	"github.com/open-feature/go-sdk/pkg/openfeature"
)

type StoredFlag struct {
	DefaultVariant string    `json:"defaultVariant"`
	Variants       []Variant `json:"variant"`
}

type Variant struct {
	Criteria     []Criteria  `json:"criteria"`
	TargetingKey string      `json:"targetingKey"`
	Value        interface{} `json:"value"`
	Name         string      `json:"name"`
}

type Criteria struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (f *StoredFlag) evaluate(evalCtx map[string]interface{}) (string, openfeature.Reason, interface{}, error) {
	var defaultVariant *Variant
	for _, variant := range f.Variants {
		if variant.Name == f.DefaultVariant {
			defaultVariant = &variant
		}
		if variant.TargetingKey != "" && variant.TargetingKey != evalCtx["targetingKey"] {
			continue
		}
		match := true
		for _, criteria := range variant.Criteria {
			val, ok := evalCtx[criteria.Key]
			if !ok || val != criteria.Value {
				match = false
				break
			}
		}
		if match {
			return variant.Name, openfeature.TargetingMatchReason, variant.Value, nil
		}
	}
	if defaultVariant == nil {
		return "", openfeature.ErrorReason, nil, openfeature.NewParseErrorResolutionError("")
	}
	return defaultVariant.Name, openfeature.DefaultReason, defaultVariant.Value, nil
}
