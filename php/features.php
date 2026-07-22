<?php
declare(strict_types=1);

// IbanValidation SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class IbanValidationFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new IbanValidationBaseFeature();
            case "test":
                return new IbanValidationTestFeature();
            default:
                return new IbanValidationBaseFeature();
        }
    }
}
