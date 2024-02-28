/* eslint-disable */
import { Params } from "../tradebin/params";
import { Market } from "../tradebin/market";
import { QueueMessage } from "../tradebin/queue_message";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "bze.tradebin.v1";

/** GenesisState defines the tradebin module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  marketList: Market[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  queueMessageList: QueueMessage[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.marketList) {
      Market.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.queueMessageList) {
      QueueMessage.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.marketList = [];
    message.queueMessageList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.marketList.push(Market.decode(reader, reader.uint32()));
          break;
        case 3:
          message.queueMessageList.push(
            QueueMessage.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.marketList = [];
    message.queueMessageList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.marketList !== undefined && object.marketList !== null) {
      for (const e of object.marketList) {
        message.marketList.push(Market.fromJSON(e));
      }
    }
    if (
      object.queueMessageList !== undefined &&
      object.queueMessageList !== null
    ) {
      for (const e of object.queueMessageList) {
        message.queueMessageList.push(QueueMessage.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.marketList) {
      obj.marketList = message.marketList.map((e) =>
        e ? Market.toJSON(e) : undefined
      );
    } else {
      obj.marketList = [];
    }
    if (message.queueMessageList) {
      obj.queueMessageList = message.queueMessageList.map((e) =>
        e ? QueueMessage.toJSON(e) : undefined
      );
    } else {
      obj.queueMessageList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.marketList = [];
    message.queueMessageList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.marketList !== undefined && object.marketList !== null) {
      for (const e of object.marketList) {
        message.marketList.push(Market.fromPartial(e));
      }
    }
    if (
      object.queueMessageList !== undefined &&
      object.queueMessageList !== null
    ) {
      for (const e of object.queueMessageList) {
        message.queueMessageList.push(QueueMessage.fromPartial(e));
      }
    }
    return message;
  },
};

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
