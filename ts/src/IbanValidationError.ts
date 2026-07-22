
import { Context } from './Context'


class IbanValidationError extends Error {

  isIbanValidationError = true

  sdk = 'IbanValidation'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  IbanValidationError
}

