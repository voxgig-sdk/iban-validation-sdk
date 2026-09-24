
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { IbanValidationSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = IbanValidationSDK.test()
    equal(testsdk instanceof IbanValidationSDK, true,
      'IbanValidationSDK.test() must return a client synchronously')
  })

})
