import { IbanValidationEntityBase } from '../IbanValidationEntityBase';
import type { IbanValidationSDK } from '../IbanValidationSDK';
import type { Control } from '../types';
import type { Ssl, SslListMatch } from '../IbanValidationTypes';
declare class SslEntity extends IbanValidationEntityBase<Ssl> {
    constructor(client: IbanValidationSDK, entopts: any);
    make(this: SslEntity): SslEntity;
    list(this: any, reqmatch?: SslListMatch, ctrl?: Control): Promise<SslEntity[]>;
}
export { SslEntity };
