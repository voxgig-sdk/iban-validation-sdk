# IbanValidation SDK exists test

import pytest
from ibanvalidation_sdk import IbanValidationSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = IbanValidationSDK.test(None, None)
        assert testsdk is not None
