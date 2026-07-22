-- IbanValidation SDK error

local IbanValidationError = {}
IbanValidationError.__index = IbanValidationError


function IbanValidationError.new(code, msg, ctx)
  local self = setmetatable({}, IbanValidationError)
  self.is_sdk_error = true
  self.sdk = "IbanValidation"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function IbanValidationError:error()
  return self.msg
end


function IbanValidationError:__tostring()
  return self.msg
end


return IbanValidationError
