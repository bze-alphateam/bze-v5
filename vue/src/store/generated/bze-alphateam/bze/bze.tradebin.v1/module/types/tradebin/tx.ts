/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "bze.tradebin.v1";

export interface MsgCreateMarket {
  creator: string;
  base: string;
  quote: string;
}

export interface MsgCreateMarketResponse {}

export interface MsgCreateOrder {
  creator: string;
  order_type: string;
  amount: string;
  price: string;
  marketId: string;
}

export interface MsgCreateOrderResponse {}

export interface MsgCancelOrder {
  creator: string;
  marketId: string;
  orderId: string;
  order_type: string;
}

export interface MsgCancelOrderResponse {}

export interface FillOrderItem {
  price: string;
  amount: string;
}

export interface MsgFillOrders {
  creator: string;
  marketId: string;
  order_type: string;
  orders: FillOrderItem[];
}

export interface MsgFillOrdersResponse {}

const baseMsgCreateMarket: object = { creator: "", base: "", quote: "" };

export const MsgCreateMarket = {
  encode(message: MsgCreateMarket, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.base !== "") {
      writer.uint32(18).string(message.base);
    }
    if (message.quote !== "") {
      writer.uint32(26).string(message.quote);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateMarket {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateMarket } as MsgCreateMarket;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.base = reader.string();
          break;
        case 3:
          message.quote = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateMarket {
    const message = { ...baseMsgCreateMarket } as MsgCreateMarket;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.base !== undefined && object.base !== null) {
      message.base = String(object.base);
    } else {
      message.base = "";
    }
    if (object.quote !== undefined && object.quote !== null) {
      message.quote = String(object.quote);
    } else {
      message.quote = "";
    }
    return message;
  },

  toJSON(message: MsgCreateMarket): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.base !== undefined && (obj.base = message.base);
    message.quote !== undefined && (obj.quote = message.quote);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateMarket>): MsgCreateMarket {
    const message = { ...baseMsgCreateMarket } as MsgCreateMarket;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.base !== undefined && object.base !== null) {
      message.base = object.base;
    } else {
      message.base = "";
    }
    if (object.quote !== undefined && object.quote !== null) {
      message.quote = object.quote;
    } else {
      message.quote = "";
    }
    return message;
  },
};

const baseMsgCreateMarketResponse: object = {};

export const MsgCreateMarketResponse = {
  encode(_: MsgCreateMarketResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateMarketResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateMarketResponse,
    } as MsgCreateMarketResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateMarketResponse {
    const message = {
      ...baseMsgCreateMarketResponse,
    } as MsgCreateMarketResponse;
    return message;
  },

  toJSON(_: MsgCreateMarketResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateMarketResponse>
  ): MsgCreateMarketResponse {
    const message = {
      ...baseMsgCreateMarketResponse,
    } as MsgCreateMarketResponse;
    return message;
  },
};

const baseMsgCreateOrder: object = {
  creator: "",
  order_type: "",
  amount: "",
  price: "",
  marketId: "",
};

export const MsgCreateOrder = {
  encode(message: MsgCreateOrder, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.order_type !== "") {
      writer.uint32(18).string(message.order_type);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    if (message.price !== "") {
      writer.uint32(34).string(message.price);
    }
    if (message.marketId !== "") {
      writer.uint32(42).string(message.marketId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateOrder } as MsgCreateOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.order_type = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        case 4:
          message.price = reader.string();
          break;
        case 5:
          message.marketId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateOrder {
    const message = { ...baseMsgCreateOrder } as MsgCreateOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.order_type !== undefined && object.order_type !== null) {
      message.order_type = String(object.order_type);
    } else {
      message.order_type = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.marketId !== undefined && object.marketId !== null) {
      message.marketId = String(object.marketId);
    } else {
      message.marketId = "";
    }
    return message;
  },

  toJSON(message: MsgCreateOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.order_type !== undefined && (obj.order_type = message.order_type);
    message.amount !== undefined && (obj.amount = message.amount);
    message.price !== undefined && (obj.price = message.price);
    message.marketId !== undefined && (obj.marketId = message.marketId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateOrder>): MsgCreateOrder {
    const message = { ...baseMsgCreateOrder } as MsgCreateOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.order_type !== undefined && object.order_type !== null) {
      message.order_type = object.order_type;
    } else {
      message.order_type = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.marketId !== undefined && object.marketId !== null) {
      message.marketId = object.marketId;
    } else {
      message.marketId = "";
    }
    return message;
  },
};

const baseMsgCreateOrderResponse: object = {};

export const MsgCreateOrderResponse = {
  encode(_: MsgCreateOrderResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateOrderResponse } as MsgCreateOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateOrderResponse {
    const message = { ...baseMsgCreateOrderResponse } as MsgCreateOrderResponse;
    return message;
  },

  toJSON(_: MsgCreateOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateOrderResponse>): MsgCreateOrderResponse {
    const message = { ...baseMsgCreateOrderResponse } as MsgCreateOrderResponse;
    return message;
  },
};

const baseMsgCancelOrder: object = {
  creator: "",
  marketId: "",
  orderId: "",
  order_type: "",
};

export const MsgCancelOrder = {
  encode(message: MsgCancelOrder, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.marketId !== "") {
      writer.uint32(18).string(message.marketId);
    }
    if (message.orderId !== "") {
      writer.uint32(26).string(message.orderId);
    }
    if (message.order_type !== "") {
      writer.uint32(34).string(message.order_type);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelOrder } as MsgCancelOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.marketId = reader.string();
          break;
        case 3:
          message.orderId = reader.string();
          break;
        case 4:
          message.order_type = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelOrder {
    const message = { ...baseMsgCancelOrder } as MsgCancelOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.marketId !== undefined && object.marketId !== null) {
      message.marketId = String(object.marketId);
    } else {
      message.marketId = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = String(object.orderId);
    } else {
      message.orderId = "";
    }
    if (object.order_type !== undefined && object.order_type !== null) {
      message.order_type = String(object.order_type);
    } else {
      message.order_type = "";
    }
    return message;
  },

  toJSON(message: MsgCancelOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.marketId !== undefined && (obj.marketId = message.marketId);
    message.orderId !== undefined && (obj.orderId = message.orderId);
    message.order_type !== undefined && (obj.order_type = message.order_type);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCancelOrder>): MsgCancelOrder {
    const message = { ...baseMsgCancelOrder } as MsgCancelOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.marketId !== undefined && object.marketId !== null) {
      message.marketId = object.marketId;
    } else {
      message.marketId = "";
    }
    if (object.orderId !== undefined && object.orderId !== null) {
      message.orderId = object.orderId;
    } else {
      message.orderId = "";
    }
    if (object.order_type !== undefined && object.order_type !== null) {
      message.order_type = object.order_type;
    } else {
      message.order_type = "";
    }
    return message;
  },
};

const baseMsgCancelOrderResponse: object = {};

export const MsgCancelOrderResponse = {
  encode(_: MsgCancelOrderResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelOrderResponse } as MsgCancelOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCancelOrderResponse {
    const message = { ...baseMsgCancelOrderResponse } as MsgCancelOrderResponse;
    return message;
  },

  toJSON(_: MsgCancelOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCancelOrderResponse>): MsgCancelOrderResponse {
    const message = { ...baseMsgCancelOrderResponse } as MsgCancelOrderResponse;
    return message;
  },
};

const baseFillOrderItem: object = { price: "", amount: "" };

export const FillOrderItem = {
  encode(message: FillOrderItem, writer: Writer = Writer.create()): Writer {
    if (message.price !== "") {
      writer.uint32(10).string(message.price);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): FillOrderItem {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFillOrderItem } as FillOrderItem;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.price = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FillOrderItem {
    const message = { ...baseFillOrderItem } as FillOrderItem;
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: FillOrderItem): unknown {
    const obj: any = {};
    message.price !== undefined && (obj.price = message.price);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<FillOrderItem>): FillOrderItem {
    const message = { ...baseFillOrderItem } as FillOrderItem;
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseMsgFillOrders: object = { creator: "", marketId: "", order_type: "" };

export const MsgFillOrders = {
  encode(message: MsgFillOrders, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.marketId !== "") {
      writer.uint32(18).string(message.marketId);
    }
    if (message.order_type !== "") {
      writer.uint32(26).string(message.order_type);
    }
    for (const v of message.orders) {
      FillOrderItem.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFillOrders {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFillOrders } as MsgFillOrders;
    message.orders = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.marketId = reader.string();
          break;
        case 3:
          message.order_type = reader.string();
          break;
        case 4:
          message.orders.push(FillOrderItem.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFillOrders {
    const message = { ...baseMsgFillOrders } as MsgFillOrders;
    message.orders = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.marketId !== undefined && object.marketId !== null) {
      message.marketId = String(object.marketId);
    } else {
      message.marketId = "";
    }
    if (object.order_type !== undefined && object.order_type !== null) {
      message.order_type = String(object.order_type);
    } else {
      message.order_type = "";
    }
    if (object.orders !== undefined && object.orders !== null) {
      for (const e of object.orders) {
        message.orders.push(FillOrderItem.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgFillOrders): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.marketId !== undefined && (obj.marketId = message.marketId);
    message.order_type !== undefined && (obj.order_type = message.order_type);
    if (message.orders) {
      obj.orders = message.orders.map((e) =>
        e ? FillOrderItem.toJSON(e) : undefined
      );
    } else {
      obj.orders = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgFillOrders>): MsgFillOrders {
    const message = { ...baseMsgFillOrders } as MsgFillOrders;
    message.orders = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.marketId !== undefined && object.marketId !== null) {
      message.marketId = object.marketId;
    } else {
      message.marketId = "";
    }
    if (object.order_type !== undefined && object.order_type !== null) {
      message.order_type = object.order_type;
    } else {
      message.order_type = "";
    }
    if (object.orders !== undefined && object.orders !== null) {
      for (const e of object.orders) {
        message.orders.push(FillOrderItem.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgFillOrdersResponse: object = {};

export const MsgFillOrdersResponse = {
  encode(_: MsgFillOrdersResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFillOrdersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFillOrdersResponse } as MsgFillOrdersResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgFillOrdersResponse {
    const message = { ...baseMsgFillOrdersResponse } as MsgFillOrdersResponse;
    return message;
  },

  toJSON(_: MsgFillOrdersResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgFillOrdersResponse>): MsgFillOrdersResponse {
    const message = { ...baseMsgFillOrdersResponse } as MsgFillOrdersResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateMarket(request: MsgCreateMarket): Promise<MsgCreateMarketResponse>;
  CreateOrder(request: MsgCreateOrder): Promise<MsgCreateOrderResponse>;
  CancelOrder(request: MsgCancelOrder): Promise<MsgCancelOrderResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  FillOrders(request: MsgFillOrders): Promise<MsgFillOrdersResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateMarket(request: MsgCreateMarket): Promise<MsgCreateMarketResponse> {
    const data = MsgCreateMarket.encode(request).finish();
    const promise = this.rpc.request(
      "bze.tradebin.v1.Msg",
      "CreateMarket",
      data
    );
    return promise.then((data) =>
      MsgCreateMarketResponse.decode(new Reader(data))
    );
  }

  CreateOrder(request: MsgCreateOrder): Promise<MsgCreateOrderResponse> {
    const data = MsgCreateOrder.encode(request).finish();
    const promise = this.rpc.request(
      "bze.tradebin.v1.Msg",
      "CreateOrder",
      data
    );
    return promise.then((data) =>
      MsgCreateOrderResponse.decode(new Reader(data))
    );
  }

  CancelOrder(request: MsgCancelOrder): Promise<MsgCancelOrderResponse> {
    const data = MsgCancelOrder.encode(request).finish();
    const promise = this.rpc.request(
      "bze.tradebin.v1.Msg",
      "CancelOrder",
      data
    );
    return promise.then((data) =>
      MsgCancelOrderResponse.decode(new Reader(data))
    );
  }

  FillOrders(request: MsgFillOrders): Promise<MsgFillOrdersResponse> {
    const data = MsgFillOrders.encode(request).finish();
    const promise = this.rpc.request("bze.tradebin.v1.Msg", "FillOrders", data);
    return promise.then((data) =>
      MsgFillOrdersResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
