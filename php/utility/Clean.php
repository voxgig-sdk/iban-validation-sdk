<?php
declare(strict_types=1);

// IbanValidation SDK utility: clean

class IbanValidationClean
{
    public static function call(IbanValidationContext $ctx, mixed $val): mixed
    {
        return $val;
    }
}
