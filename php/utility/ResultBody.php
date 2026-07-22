<?php
declare(strict_types=1);

// IbanValidation SDK utility: result_body

class IbanValidationResultBody
{
    public static function call(IbanValidationContext $ctx): ?IbanValidationResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
