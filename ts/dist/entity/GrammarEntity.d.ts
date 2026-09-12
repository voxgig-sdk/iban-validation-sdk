import { IbanValidationEntityBase } from '../IbanValidationEntityBase';
import type { IbanValidationSDK } from '../IbanValidationSDK';
import type { Control } from '../types';
import type { Grammar, GrammarCreateData } from '../IbanValidationTypes';
declare class GrammarEntity extends IbanValidationEntityBase<Grammar> {
    constructor(client: IbanValidationSDK, entopts: any);
    make(this: GrammarEntity): GrammarEntity;
    create(this: any, reqdata?: GrammarCreateData, ctrl?: Control): Promise<GrammarEntity>;
}
export { GrammarEntity };
