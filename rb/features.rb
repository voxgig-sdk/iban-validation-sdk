# IbanValidation SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module IbanValidationFeatures
  def self.make_feature(name)
    case name
    when "base"
      IbanValidationBaseFeature.new
    when "test"
      IbanValidationTestFeature.new
    else
      IbanValidationBaseFeature.new
    end
  end
end
