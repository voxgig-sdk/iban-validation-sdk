# IbanValidation SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

IbanValidationUtility.registrar = ->(u) {
  u.clean = IbanValidationUtilities::Clean
  u.done = IbanValidationUtilities::Done
  u.make_error = IbanValidationUtilities::MakeError
  u.feature_add = IbanValidationUtilities::FeatureAdd
  u.feature_hook = IbanValidationUtilities::FeatureHook
  u.feature_init = IbanValidationUtilities::FeatureInit
  u.fetcher = IbanValidationUtilities::Fetcher
  u.make_fetch_def = IbanValidationUtilities::MakeFetchDef
  u.make_context = IbanValidationUtilities::MakeContext
  u.make_options = IbanValidationUtilities::MakeOptions
  u.make_request = IbanValidationUtilities::MakeRequest
  u.make_response = IbanValidationUtilities::MakeResponse
  u.make_result = IbanValidationUtilities::MakeResult
  u.make_point = IbanValidationUtilities::MakePoint
  u.make_spec = IbanValidationUtilities::MakeSpec
  u.make_url = IbanValidationUtilities::MakeUrl
  u.param = IbanValidationUtilities::Param
  u.prepare_auth = IbanValidationUtilities::PrepareAuth
  u.prepare_body = IbanValidationUtilities::PrepareBody
  u.prepare_headers = IbanValidationUtilities::PrepareHeaders
  u.prepare_method = IbanValidationUtilities::PrepareMethod
  u.prepare_params = IbanValidationUtilities::PrepareParams
  u.prepare_path = IbanValidationUtilities::PreparePath
  u.prepare_query = IbanValidationUtilities::PrepareQuery
  u.result_basic = IbanValidationUtilities::ResultBasic
  u.result_body = IbanValidationUtilities::ResultBody
  u.result_headers = IbanValidationUtilities::ResultHeaders
  u.transform_request = IbanValidationUtilities::TransformRequest
  u.transform_response = IbanValidationUtilities::TransformResponse
}
