/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../burner/params";
import { BurnedCoins } from "../burner/burned_coins";
import { Raffle, RaffleWinner, RaffleParticipant } from "../burner/raffle";

export const protobufPackage = "bze.burner.v1";

/** GenesisState defines the burner module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  burnedCoinsList: BurnedCoins[];
  raffleList: Raffle[];
  raffleWinnersList: RaffleWinner[];
  raffleParticipantsList: RaffleParticipant[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  raffleParticipantCounter: number;
}

const baseGenesisState: object = { raffleParticipantCounter: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.burnedCoinsList) {
      BurnedCoins.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.raffleList) {
      Raffle.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.raffleWinnersList) {
      RaffleWinner.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.raffleParticipantsList) {
      RaffleParticipant.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.raffleParticipantCounter !== 0) {
      writer.uint32(48).uint64(message.raffleParticipantCounter);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.burnedCoinsList = [];
    message.raffleList = [];
    message.raffleWinnersList = [];
    message.raffleParticipantsList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.burnedCoinsList.push(
            BurnedCoins.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.raffleList.push(Raffle.decode(reader, reader.uint32()));
          break;
        case 4:
          message.raffleWinnersList.push(
            RaffleWinner.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.raffleParticipantsList.push(
            RaffleParticipant.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.raffleParticipantCounter = longToNumber(
            reader.uint64() as Long
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
    message.burnedCoinsList = [];
    message.raffleList = [];
    message.raffleWinnersList = [];
    message.raffleParticipantsList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.burnedCoinsList !== undefined &&
      object.burnedCoinsList !== null
    ) {
      for (const e of object.burnedCoinsList) {
        message.burnedCoinsList.push(BurnedCoins.fromJSON(e));
      }
    }
    if (object.raffleList !== undefined && object.raffleList !== null) {
      for (const e of object.raffleList) {
        message.raffleList.push(Raffle.fromJSON(e));
      }
    }
    if (
      object.raffleWinnersList !== undefined &&
      object.raffleWinnersList !== null
    ) {
      for (const e of object.raffleWinnersList) {
        message.raffleWinnersList.push(RaffleWinner.fromJSON(e));
      }
    }
    if (
      object.raffleParticipantsList !== undefined &&
      object.raffleParticipantsList !== null
    ) {
      for (const e of object.raffleParticipantsList) {
        message.raffleParticipantsList.push(RaffleParticipant.fromJSON(e));
      }
    }
    if (
      object.raffleParticipantCounter !== undefined &&
      object.raffleParticipantCounter !== null
    ) {
      message.raffleParticipantCounter = Number(
        object.raffleParticipantCounter
      );
    } else {
      message.raffleParticipantCounter = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.burnedCoinsList) {
      obj.burnedCoinsList = message.burnedCoinsList.map((e) =>
        e ? BurnedCoins.toJSON(e) : undefined
      );
    } else {
      obj.burnedCoinsList = [];
    }
    if (message.raffleList) {
      obj.raffleList = message.raffleList.map((e) =>
        e ? Raffle.toJSON(e) : undefined
      );
    } else {
      obj.raffleList = [];
    }
    if (message.raffleWinnersList) {
      obj.raffleWinnersList = message.raffleWinnersList.map((e) =>
        e ? RaffleWinner.toJSON(e) : undefined
      );
    } else {
      obj.raffleWinnersList = [];
    }
    if (message.raffleParticipantsList) {
      obj.raffleParticipantsList = message.raffleParticipantsList.map((e) =>
        e ? RaffleParticipant.toJSON(e) : undefined
      );
    } else {
      obj.raffleParticipantsList = [];
    }
    message.raffleParticipantCounter !== undefined &&
      (obj.raffleParticipantCounter = message.raffleParticipantCounter);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.burnedCoinsList = [];
    message.raffleList = [];
    message.raffleWinnersList = [];
    message.raffleParticipantsList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.burnedCoinsList !== undefined &&
      object.burnedCoinsList !== null
    ) {
      for (const e of object.burnedCoinsList) {
        message.burnedCoinsList.push(BurnedCoins.fromPartial(e));
      }
    }
    if (object.raffleList !== undefined && object.raffleList !== null) {
      for (const e of object.raffleList) {
        message.raffleList.push(Raffle.fromPartial(e));
      }
    }
    if (
      object.raffleWinnersList !== undefined &&
      object.raffleWinnersList !== null
    ) {
      for (const e of object.raffleWinnersList) {
        message.raffleWinnersList.push(RaffleWinner.fromPartial(e));
      }
    }
    if (
      object.raffleParticipantsList !== undefined &&
      object.raffleParticipantsList !== null
    ) {
      for (const e of object.raffleParticipantsList) {
        message.raffleParticipantsList.push(RaffleParticipant.fromPartial(e));
      }
    }
    if (
      object.raffleParticipantCounter !== undefined &&
      object.raffleParticipantCounter !== null
    ) {
      message.raffleParticipantCounter = object.raffleParticipantCounter;
    } else {
      message.raffleParticipantCounter = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
