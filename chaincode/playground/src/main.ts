/// <reference path="../node_modules/fabric-shim/types/index.d.ts" />
import * as shim from 'fabric-shim';
import { MyChaincode } from './mycc';

console.info('=========== Starting prenda chaincode ===========');
shim.start(new MyChaincode());
console.info('========== WAITING FOR INVOKES ==========');
