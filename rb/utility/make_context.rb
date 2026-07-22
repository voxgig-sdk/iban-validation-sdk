# IbanValidation SDK utility: make_context
require_relative '../core/context'
module IbanValidationUtilities
  MakeContext = ->(ctxmap, basectx) {
    IbanValidationContext.new(ctxmap, basectx)
  }
end
