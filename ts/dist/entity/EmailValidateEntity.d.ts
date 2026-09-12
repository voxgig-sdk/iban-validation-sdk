import { IbanValidationEntityBase } from '../IbanValidationEntityBase';
import type { IbanValidationSDK } from '../IbanValidationSDK';
import type { Control } from '../types';
import type { EmailValidate, EmailValidateLoadMatch } from '../IbanValidationTypes';
declare class EmailValidateEntity extends IbanValidationEntityBase<EmailValidate> {
    constructor(client: IbanValidationSDK, entopts: any);
    make(this: EmailValidateEntity): EmailValidateEntity;
    load(this: any, reqmatch?: EmailValidateLoadMatch, ctrl?: Control): Promise<EmailValidateEntity>;
}
export { EmailValidateEntity };
