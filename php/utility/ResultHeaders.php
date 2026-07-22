<?php
declare(strict_types=1);

// IbanValidation SDK utility: result_headers

class IbanValidationResultHeaders
{
    public static function call(IbanValidationContext $ctx): ?IbanValidationResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
