
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { IbanValidationSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await IbanValidationSDK.test()
    equal(null !== testsdk, true)
  })

})
