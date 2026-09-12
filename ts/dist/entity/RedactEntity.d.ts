import { IbanValidationEntityBase } from '../IbanValidationEntityBase';
import type { IbanValidationSDK } from '../IbanValidationSDK';
import type { Control } from '../types';
import type { Redact, RedactCreateData } from '../IbanValidationTypes';
declare class RedactEntity extends IbanValidationEntityBase<Redact> {
    constructor(client: IbanValidationSDK, entopts: any);
    make(this: RedactEntity): RedactEntity;
    create(this: any, reqdata?: RedactCreateData, ctrl?: Control): Promise<RedactEntity>;
}
export { RedactEntity };
