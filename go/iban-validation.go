package voxgigibanvalidationsdk

import (
	"github.com/voxgig-sdk/iban-validation-sdk/go/core"
	"github.com/voxgig-sdk/iban-validation-sdk/go/entity"
	"github.com/voxgig-sdk/iban-validation-sdk/go/feature"
	_ "github.com/voxgig-sdk/iban-validation-sdk/go/utility"
)

// Type aliases preserve external API.
type IbanValidationSDK = core.IbanValidationSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type IbanValidationEntity = core.IbanValidationEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type IbanValidationError = core.IbanValidationError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewDnsResultEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewDnsResultEntity(client, entopts)
	}
	core.NewDomainEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewDomainEntity(client, entopts)
	}
	core.NewEmailValidateEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewEmailValidateEntity(client, entopts)
	}
	core.NewGenerateEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewGenerateEntity(client, entopts)
	}
	core.NewGrammarEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewGrammarEntity(client, entopts)
	}
	core.NewIpnEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewIpnEntity(client, entopts)
	}
	core.NewRedactEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewRedactEntity(client, entopts)
	}
	core.NewSslEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewSslEntity(client, entopts)
	}
	core.NewUtilityEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewUtilityEntity(client, entopts)
	}
	core.NewWhoiEntityFunc = func(client *core.IbanValidationSDK, entopts map[string]any) core.IbanValidationEntity {
		return entity.NewWhoiEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewIbanValidationSDK = core.NewIbanValidationSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewIbanValidationSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *IbanValidationSDK  { return NewIbanValidationSDK(nil) }
func Test() *IbanValidationSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
