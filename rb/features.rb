# IbanValidation SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module IbanValidationFeatures
  def self.make_feature(name)
    case name
    when "base"
      IbanValidationBaseFeature.new
    when "ratelimit"
      IbanValidationRatelimitFeature.new
    when "retry"
      IbanValidationRetryFeature.new
    when "test"
      IbanValidationTestFeature.new
    when "timeout"
      IbanValidationTimeoutFeature.new
    else
      IbanValidationBaseFeature.new
    end
  end
end
