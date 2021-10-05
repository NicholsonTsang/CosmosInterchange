// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendSellOrder } from "./types/ibcdex/tx";
import { MsgSendCreatePair } from "./types/ibcdex/tx";
const types = [
    ["/username.interchange.ibcdex.MsgSendSellOrder", MsgSendSellOrder],
    ["/username.interchange.ibcdex.MsgSendCreatePair", MsgSendCreatePair],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSendSellOrder: (data) => ({ typeUrl: "/username.interchange.ibcdex.MsgSendSellOrder", value: data }),
        msgSendCreatePair: (data) => ({ typeUrl: "/username.interchange.ibcdex.MsgSendCreatePair", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
