/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "bze.v1.rewards";

export interface ActivateTradingRewardProposal {
  title: string;
  description: string;
  rewardId: string;
}

const baseActivateTradingRewardProposal: object = {
  title: "",
  description: "",
  rewardId: "",
};

export const ActivateTradingRewardProposal = {
  encode(
    message: ActivateTradingRewardProposal,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.rewardId !== "") {
      writer.uint32(26).string(message.rewardId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): ActivateTradingRewardProposal {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseActivateTradingRewardProposal,
    } as ActivateTradingRewardProposal;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.rewardId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ActivateTradingRewardProposal {
    const message = {
      ...baseActivateTradingRewardProposal,
    } as ActivateTradingRewardProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    return message;
  },

  toJSON(message: ActivateTradingRewardProposal): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ActivateTradingRewardProposal>
  ): ActivateTradingRewardProposal {
    const message = {
      ...baseActivateTradingRewardProposal,
    } as ActivateTradingRewardProposal;
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
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
