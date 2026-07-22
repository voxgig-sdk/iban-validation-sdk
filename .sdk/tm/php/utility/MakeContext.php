<?php
declare(strict_types=1);

// IbanValidation SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class IbanValidationMakeContext
{
    public static function call(array $ctxmap, ?IbanValidationContext $basectx): IbanValidationContext
    {
        return new IbanValidationContext($ctxmap, $basectx);
    }
}
