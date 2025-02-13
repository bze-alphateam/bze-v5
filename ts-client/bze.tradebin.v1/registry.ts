import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgFillOrders } from "./types/tradebin/tx";
import { MsgCreateOrder } from "./types/tradebin/tx";
import { MsgCreateMarket } from "./types/tradebin/tx";
import { MsgCancelOrder } from "./types/tradebin/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/bze.tradebin.v1.MsgFillOrders", MsgFillOrders],
    ["/bze.tradebin.v1.MsgCreateOrder", MsgCreateOrder],
    ["/bze.tradebin.v1.MsgCreateMarket", MsgCreateMarket],
    ["/bze.tradebin.v1.MsgCancelOrder", MsgCancelOrder],
    
];

export { msgTypes }