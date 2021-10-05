import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "username.interchange.ibcdex";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCancelSellOrder {
    creator: string;
    port: string;
    channel: string;
    amountDenom: string;
    priceDenom: string;
    orderID: number;
}
export interface MsgCancelSellOrderResponse {
}
export interface MsgSendBuyOrder {
    sender: string;
    port: string;
    channelID: string;
    timeoutTimestamp: number;
    amountDenom: string;
    amount: number;
    priceDenom: string;
    price: number;
}
export interface MsgSendBuyOrderResponse {
}
export interface MsgSendSellOrder {
    sender: string;
    port: string;
    channelID: string;
    timeoutTimestamp: number;
    amountDenom: string;
    amount: number;
    priceDenom: string;
    price: number;
}
export interface MsgSendSellOrderResponse {
}
export interface MsgSendCreatePair {
    sender: string;
    port: string;
    channelID: string;
    timeoutTimestamp: number;
    sourceDenom: string;
    targetDenom: string;
}
export interface MsgSendCreatePairResponse {
}
export declare const MsgCancelSellOrder: {
    encode(message: MsgCancelSellOrder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCancelSellOrder;
    fromJSON(object: any): MsgCancelSellOrder;
    toJSON(message: MsgCancelSellOrder): unknown;
    fromPartial(object: DeepPartial<MsgCancelSellOrder>): MsgCancelSellOrder;
};
export declare const MsgCancelSellOrderResponse: {
    encode(_: MsgCancelSellOrderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCancelSellOrderResponse;
    fromJSON(_: any): MsgCancelSellOrderResponse;
    toJSON(_: MsgCancelSellOrderResponse): unknown;
    fromPartial(_: DeepPartial<MsgCancelSellOrderResponse>): MsgCancelSellOrderResponse;
};
export declare const MsgSendBuyOrder: {
    encode(message: MsgSendBuyOrder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendBuyOrder;
    fromJSON(object: any): MsgSendBuyOrder;
    toJSON(message: MsgSendBuyOrder): unknown;
    fromPartial(object: DeepPartial<MsgSendBuyOrder>): MsgSendBuyOrder;
};
export declare const MsgSendBuyOrderResponse: {
    encode(_: MsgSendBuyOrderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendBuyOrderResponse;
    fromJSON(_: any): MsgSendBuyOrderResponse;
    toJSON(_: MsgSendBuyOrderResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendBuyOrderResponse>): MsgSendBuyOrderResponse;
};
export declare const MsgSendSellOrder: {
    encode(message: MsgSendSellOrder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendSellOrder;
    fromJSON(object: any): MsgSendSellOrder;
    toJSON(message: MsgSendSellOrder): unknown;
    fromPartial(object: DeepPartial<MsgSendSellOrder>): MsgSendSellOrder;
};
export declare const MsgSendSellOrderResponse: {
    encode(_: MsgSendSellOrderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendSellOrderResponse;
    fromJSON(_: any): MsgSendSellOrderResponse;
    toJSON(_: MsgSendSellOrderResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendSellOrderResponse>): MsgSendSellOrderResponse;
};
export declare const MsgSendCreatePair: {
    encode(message: MsgSendCreatePair, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendCreatePair;
    fromJSON(object: any): MsgSendCreatePair;
    toJSON(message: MsgSendCreatePair): unknown;
    fromPartial(object: DeepPartial<MsgSendCreatePair>): MsgSendCreatePair;
};
export declare const MsgSendCreatePairResponse: {
    encode(_: MsgSendCreatePairResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendCreatePairResponse;
    fromJSON(_: any): MsgSendCreatePairResponse;
    toJSON(_: MsgSendCreatePairResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendCreatePairResponse>): MsgSendCreatePairResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CancelSellOrder(request: MsgCancelSellOrder): Promise<MsgCancelSellOrderResponse>;
    SendBuyOrder(request: MsgSendBuyOrder): Promise<MsgSendBuyOrderResponse>;
    SendSellOrder(request: MsgSendSellOrder): Promise<MsgSendSellOrderResponse>;
    SendCreatePair(request: MsgSendCreatePair): Promise<MsgSendCreatePairResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CancelSellOrder(request: MsgCancelSellOrder): Promise<MsgCancelSellOrderResponse>;
    SendBuyOrder(request: MsgSendBuyOrder): Promise<MsgSendBuyOrderResponse>;
    SendSellOrder(request: MsgSendSellOrder): Promise<MsgSendSellOrderResponse>;
    SendCreatePair(request: MsgSendCreatePair): Promise<MsgSendCreatePairResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
