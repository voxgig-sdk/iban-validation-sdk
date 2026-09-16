# IbanValidation SDK feature factory

from ibanvalidation_sdk.feature.base_feature import IbanValidationBaseFeature
from ibanvalidation_sdk.feature.ratelimit_feature import IbanValidationRatelimitFeature
from ibanvalidation_sdk.feature.retry_feature import IbanValidationRetryFeature
from ibanvalidation_sdk.feature.test_feature import IbanValidationTestFeature
from ibanvalidation_sdk.feature.timeout_feature import IbanValidationTimeoutFeature


_FEATURES = {
    "base": lambda: IbanValidationBaseFeature(),
    "ratelimit": lambda: IbanValidationRatelimitFeature(),
    "retry": lambda: IbanValidationRetryFeature(),
    "test": lambda: IbanValidationTestFeature(),
    "timeout": lambda: IbanValidationTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
