# IbanValidation SDK utility: make_context

from ibanvalidation_sdk.core.context import IbanValidationContext


def make_context_util(ctxmap, basectx):
    return IbanValidationContext(ctxmap, basectx)
