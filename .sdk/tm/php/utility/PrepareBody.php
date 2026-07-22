<?php
declare(strict_types=1);

// IbanValidation SDK utility: prepare_body

class IbanValidationPrepareBody
{
    public static function call(IbanValidationContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
