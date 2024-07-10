/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "bze.tokenfactory.v1";

export interface DenomAdmin {
  /** Can be empty for no admin, or a valid osmosis address */
  admin: string;
}

const baseDenomAdmin: object = { admin: "" };

export const DenomAdmin = {
  encode(message: DenomAdmin, writer: Writer = Writer.create()): Writer {
    if (message.admin !== "") {
      writer.uint32(10).string(message.admin);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DenomAdmin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDenomAdmin } as DenomAdmin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.admin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DenomAdmin {
    const message = { ...baseDenomAdmin } as DenomAdmin;
    if (object.admin !== undefined && object.admin !== null) {
      message.admin = String(object.admin);
    } else {
      message.admin = "";
    }
    return message;
  },

  toJSON(message: DenomAdmin): unknown {
    const obj: any = {};
    message.admin !== undefined && (obj.admin = message.admin);
    return obj;
  },

  fromPartial(object: DeepPartial<DenomAdmin>): DenomAdmin {
    const message = { ...baseDenomAdmin } as DenomAdmin;
    if (object.admin !== undefined && object.admin !== null) {
      message.admin = object.admin;
    } else {
      message.admin = "";
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
