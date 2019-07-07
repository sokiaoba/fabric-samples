/// <reference path="../node_modules/fabric-shim/types/index.d.ts" />
import * as shim from 'fabric-shim';
import { Shim, ChaincodeStub, ChaincodeInterface, ChaincodeResponse } from 'fabric-shim';

export class MyChaincode implements ChaincodeInterface {
    async Init(stub: ChaincodeStub): Promise<ChaincodeResponse> {
        console.info('=========== Instantiated my chaincode ===========');
        return shim.success();
    }

    async Invoke(stub: ChaincodeStub): Promise<ChaincodeResponse> {
        console.info('=========== Invoked my chaincode ===========');
        return shim.success();
    }     
}
