# IbanValidation SDK feature factory

from ibanvalidation_sdk.feature.base_feature import IbanValidationBaseFeature
from ibanvalidation_sdk.feature.test_feature import IbanValidationTestFeature


def _make_feature(name):
    features = {
        "base": lambda: IbanValidationBaseFeature(),
        "test": lambda: IbanValidationTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
