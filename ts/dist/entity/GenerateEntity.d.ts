import { IbanValidationEntityBase } from '../IbanValidationEntityBase';
import type { IbanValidationSDK } from '../IbanValidationSDK';
import type { Control } from '../types';
import type { Generate, GenerateLoadMatch } from '../IbanValidationTypes';
declare class GenerateEntity extends IbanValidationEntityBase<Generate> {
    constructor(client: IbanValidationSDK, entopts: any);
    make(this: GenerateEntity): GenerateEntity;
    load(this: any, reqmatch?: GenerateLoadMatch, ctrl?: Control): Promise<GenerateEntity>;
}
export { GenerateEntity };
