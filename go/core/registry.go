package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewDnsResultEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

var NewDomainEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

var NewEmailValidateEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

var NewGenerateEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

var NewGrammarEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

var NewIpnEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

var NewRedactEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

var NewSslEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

var NewUtilityEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

var NewWhoiEntityFunc func(client *IbanValidationSDK, entopts map[string]any) IbanValidationEntity

