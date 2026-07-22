<?php
declare(strict_types=1);

// IbanValidation SDK base feature

class IbanValidationBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(IbanValidationContext $ctx, array $options): void {}
    public function PostConstruct(IbanValidationContext $ctx): void {}
    public function PostConstructEntity(IbanValidationContext $ctx): void {}
    public function SetData(IbanValidationContext $ctx): void {}
    public function GetData(IbanValidationContext $ctx): void {}
    public function GetMatch(IbanValidationContext $ctx): void {}
    public function SetMatch(IbanValidationContext $ctx): void {}
    public function PrePoint(IbanValidationContext $ctx): void {}
    public function PreSpec(IbanValidationContext $ctx): void {}
    public function PreRequest(IbanValidationContext $ctx): void {}
    public function PreResponse(IbanValidationContext $ctx): void {}
    public function PreResult(IbanValidationContext $ctx): void {}
    public function PreDone(IbanValidationContext $ctx): void {}
    public function PreUnexpected(IbanValidationContext $ctx): void {}
}
