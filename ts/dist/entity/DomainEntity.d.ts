import { IbanValidationEntityBase } from '../IbanValidationEntityBase';
import type { IbanValidationSDK } from '../IbanValidationSDK';
import type { Control } from '../types';
import type { Domain, DomainListMatch } from '../IbanValidationTypes';
declare class DomainEntity extends IbanValidationEntityBase<Domain> {
    constructor(client: IbanValidationSDK, entopts: any);
    make(this: DomainEntity): DomainEntity;
    list(this: any, reqmatch?: DomainListMatch, ctrl?: Control): Promise<DomainEntity[]>;
}
export { DomainEntity };
