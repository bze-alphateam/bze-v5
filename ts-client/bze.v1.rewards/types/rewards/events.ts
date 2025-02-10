/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "bze.rewards.v1";

export interface StakingRewardCreateEvent {
  rewardId: string;
  prizeAmount: string;
  prizeDenom: string;
  stakingDenom: string;
  duration: number;
  minStake: number;
  lock: number;
}

export interface StakingRewardUpdateEvent {
  rewardId: string;
  duration: number;
}

export interface StakingRewardClaimEvent {
  rewardId: string;
  address: string;
  amount: string;
}

export interface StakingRewardJoinEvent {
  rewardId: string;
  address: string;
  amount: string;
}

export interface StakingRewardExitEvent {
  rewardId: string;
  address: string;
}

export interface StakingRewardFinishEvent {
  rewardId: string;
}

export interface StakingRewardDistributionEvent {
  rewardId: string;
  amount: string;
}

export interface TradingRewardCreateEvent {
  rewardId: string;
  /** the amount paid as prize for each slot */
  prizeAmount: string;
  /** the denom paid as prize */
  prizeDenom: string;
  duration: number;
  marketId: string;
  slots: number;
  creator: string;
}

export interface TradingRewardExpireEvent {
  rewardId: string;
}

export interface TradingRewardActivationEvent {
  rewardId: string;
}

export interface TradingRewardDistributionEvent {
  rewardId: string;
  prizeAmount: string;
  prizeDenom: string;
  winners: string[];
}

const baseStakingRewardCreateEvent: object = {
  rewardId: "",
  prizeAmount: "",
  prizeDenom: "",
  stakingDenom: "",
  duration: 0,
  minStake: 0,
  lock: 0,
};

export const StakingRewardCreateEvent = {
  encode(
    message: StakingRewardCreateEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    if (message.prizeAmount !== "") {
      writer.uint32(18).string(message.prizeAmount);
    }
    if (message.prizeDenom !== "") {
      writer.uint32(26).string(message.prizeDenom);
    }
    if (message.stakingDenom !== "") {
      writer.uint32(34).string(message.stakingDenom);
    }
    if (message.duration !== 0) {
      writer.uint32(40).uint32(message.duration);
    }
    if (message.minStake !== 0) {
      writer.uint32(48).uint64(message.minStake);
    }
    if (message.lock !== 0) {
      writer.uint32(56).uint32(message.lock);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): StakingRewardCreateEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseStakingRewardCreateEvent,
    } as StakingRewardCreateEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        case 2:
          message.prizeAmount = reader.string();
          break;
        case 3:
          message.prizeDenom = reader.string();
          break;
        case 4:
          message.stakingDenom = reader.string();
          break;
        case 5:
          message.duration = reader.uint32();
          break;
        case 6:
          message.minStake = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.lock = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StakingRewardCreateEvent {
    const message = {
      ...baseStakingRewardCreateEvent,
    } as StakingRewardCreateEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    if (object.prizeAmount !== undefined && object.prizeAmount !== null) {
      message.prizeAmount = String(object.prizeAmount);
    } else {
      message.prizeAmount = "";
    }
    if (object.prizeDenom !== undefined && object.prizeDenom !== null) {
      message.prizeDenom = String(object.prizeDenom);
    } else {
      message.prizeDenom = "";
    }
    if (object.stakingDenom !== undefined && object.stakingDenom !== null) {
      message.stakingDenom = String(object.stakingDenom);
    } else {
      message.stakingDenom = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = Number(object.duration);
    } else {
      message.duration = 0;
    }
    if (object.minStake !== undefined && object.minStake !== null) {
      message.minStake = Number(object.minStake);
    } else {
      message.minStake = 0;
    }
    if (object.lock !== undefined && object.lock !== null) {
      message.lock = Number(object.lock);
    } else {
      message.lock = 0;
    }
    return message;
  },

  toJSON(message: StakingRewardCreateEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    message.prizeAmount !== undefined &&
      (obj.prizeAmount = message.prizeAmount);
    message.prizeDenom !== undefined && (obj.prizeDenom = message.prizeDenom);
    message.stakingDenom !== undefined &&
      (obj.stakingDenom = message.stakingDenom);
    message.duration !== undefined && (obj.duration = message.duration);
    message.minStake !== undefined && (obj.minStake = message.minStake);
    message.lock !== undefined && (obj.lock = message.lock);
    return obj;
  },

  fromPartial(
    object: DeepPartial<StakingRewardCreateEvent>
  ): StakingRewardCreateEvent {
    const message = {
      ...baseStakingRewardCreateEvent,
    } as StakingRewardCreateEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    if (object.prizeAmount !== undefined && object.prizeAmount !== null) {
      message.prizeAmount = object.prizeAmount;
    } else {
      message.prizeAmount = "";
    }
    if (object.prizeDenom !== undefined && object.prizeDenom !== null) {
      message.prizeDenom = object.prizeDenom;
    } else {
      message.prizeDenom = "";
    }
    if (object.stakingDenom !== undefined && object.stakingDenom !== null) {
      message.stakingDenom = object.stakingDenom;
    } else {
      message.stakingDenom = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = object.duration;
    } else {
      message.duration = 0;
    }
    if (object.minStake !== undefined && object.minStake !== null) {
      message.minStake = object.minStake;
    } else {
      message.minStake = 0;
    }
    if (object.lock !== undefined && object.lock !== null) {
      message.lock = object.lock;
    } else {
      message.lock = 0;
    }
    return message;
  },
};

const baseStakingRewardUpdateEvent: object = { rewardId: "", duration: 0 };

export const StakingRewardUpdateEvent = {
  encode(
    message: StakingRewardUpdateEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    if (message.duration !== 0) {
      writer.uint32(16).uint32(message.duration);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): StakingRewardUpdateEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseStakingRewardUpdateEvent,
    } as StakingRewardUpdateEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        case 2:
          message.duration = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StakingRewardUpdateEvent {
    const message = {
      ...baseStakingRewardUpdateEvent,
    } as StakingRewardUpdateEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = Number(object.duration);
    } else {
      message.duration = 0;
    }
    return message;
  },

  toJSON(message: StakingRewardUpdateEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    message.duration !== undefined && (obj.duration = message.duration);
    return obj;
  },

  fromPartial(
    object: DeepPartial<StakingRewardUpdateEvent>
  ): StakingRewardUpdateEvent {
    const message = {
      ...baseStakingRewardUpdateEvent,
    } as StakingRewardUpdateEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = object.duration;
    } else {
      message.duration = 0;
    }
    return message;
  },
};

const baseStakingRewardClaimEvent: object = {
  rewardId: "",
  address: "",
  amount: "",
};

export const StakingRewardClaimEvent = {
  encode(
    message: StakingRewardClaimEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StakingRewardClaimEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseStakingRewardClaimEvent,
    } as StakingRewardClaimEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StakingRewardClaimEvent {
    const message = {
      ...baseStakingRewardClaimEvent,
    } as StakingRewardClaimEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: StakingRewardClaimEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    message.address !== undefined && (obj.address = message.address);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<StakingRewardClaimEvent>
  ): StakingRewardClaimEvent {
    const message = {
      ...baseStakingRewardClaimEvent,
    } as StakingRewardClaimEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseStakingRewardJoinEvent: object = {
  rewardId: "",
  address: "",
  amount: "",
};

export const StakingRewardJoinEvent = {
  encode(
    message: StakingRewardJoinEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StakingRewardJoinEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStakingRewardJoinEvent } as StakingRewardJoinEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StakingRewardJoinEvent {
    const message = { ...baseStakingRewardJoinEvent } as StakingRewardJoinEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: StakingRewardJoinEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    message.address !== undefined && (obj.address = message.address);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<StakingRewardJoinEvent>
  ): StakingRewardJoinEvent {
    const message = { ...baseStakingRewardJoinEvent } as StakingRewardJoinEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseStakingRewardExitEvent: object = { rewardId: "", address: "" };

export const StakingRewardExitEvent = {
  encode(
    message: StakingRewardExitEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StakingRewardExitEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStakingRewardExitEvent } as StakingRewardExitEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StakingRewardExitEvent {
    const message = { ...baseStakingRewardExitEvent } as StakingRewardExitEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: StakingRewardExitEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<StakingRewardExitEvent>
  ): StakingRewardExitEvent {
    const message = { ...baseStakingRewardExitEvent } as StakingRewardExitEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseStakingRewardFinishEvent: object = { rewardId: "" };

export const StakingRewardFinishEvent = {
  encode(
    message: StakingRewardFinishEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): StakingRewardFinishEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseStakingRewardFinishEvent,
    } as StakingRewardFinishEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StakingRewardFinishEvent {
    const message = {
      ...baseStakingRewardFinishEvent,
    } as StakingRewardFinishEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    return message;
  },

  toJSON(message: StakingRewardFinishEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<StakingRewardFinishEvent>
  ): StakingRewardFinishEvent {
    const message = {
      ...baseStakingRewardFinishEvent,
    } as StakingRewardFinishEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    return message;
  },
};

const baseStakingRewardDistributionEvent: object = { rewardId: "", amount: "" };

export const StakingRewardDistributionEvent = {
  encode(
    message: StakingRewardDistributionEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): StakingRewardDistributionEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseStakingRewardDistributionEvent,
    } as StakingRewardDistributionEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
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

  fromJSON(object: any): StakingRewardDistributionEvent {
    const message = {
      ...baseStakingRewardDistributionEvent,
    } as StakingRewardDistributionEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: StakingRewardDistributionEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<StakingRewardDistributionEvent>
  ): StakingRewardDistributionEvent {
    const message = {
      ...baseStakingRewardDistributionEvent,
    } as StakingRewardDistributionEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseTradingRewardCreateEvent: object = {
  rewardId: "",
  prizeAmount: "",
  prizeDenom: "",
  duration: 0,
  marketId: "",
  slots: 0,
  creator: "",
};

export const TradingRewardCreateEvent = {
  encode(
    message: TradingRewardCreateEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    if (message.prizeAmount !== "") {
      writer.uint32(18).string(message.prizeAmount);
    }
    if (message.prizeDenom !== "") {
      writer.uint32(26).string(message.prizeDenom);
    }
    if (message.duration !== 0) {
      writer.uint32(32).uint32(message.duration);
    }
    if (message.marketId !== "") {
      writer.uint32(42).string(message.marketId);
    }
    if (message.slots !== 0) {
      writer.uint32(48).uint32(message.slots);
    }
    if (message.creator !== "") {
      writer.uint32(58).string(message.creator);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): TradingRewardCreateEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseTradingRewardCreateEvent,
    } as TradingRewardCreateEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        case 2:
          message.prizeAmount = reader.string();
          break;
        case 3:
          message.prizeDenom = reader.string();
          break;
        case 4:
          message.duration = reader.uint32();
          break;
        case 5:
          message.marketId = reader.string();
          break;
        case 6:
          message.slots = reader.uint32();
          break;
        case 7:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TradingRewardCreateEvent {
    const message = {
      ...baseTradingRewardCreateEvent,
    } as TradingRewardCreateEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    if (object.prizeAmount !== undefined && object.prizeAmount !== null) {
      message.prizeAmount = String(object.prizeAmount);
    } else {
      message.prizeAmount = "";
    }
    if (object.prizeDenom !== undefined && object.prizeDenom !== null) {
      message.prizeDenom = String(object.prizeDenom);
    } else {
      message.prizeDenom = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = Number(object.duration);
    } else {
      message.duration = 0;
    }
    if (object.marketId !== undefined && object.marketId !== null) {
      message.marketId = String(object.marketId);
    } else {
      message.marketId = "";
    }
    if (object.slots !== undefined && object.slots !== null) {
      message.slots = Number(object.slots);
    } else {
      message.slots = 0;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: TradingRewardCreateEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    message.prizeAmount !== undefined &&
      (obj.prizeAmount = message.prizeAmount);
    message.prizeDenom !== undefined && (obj.prizeDenom = message.prizeDenom);
    message.duration !== undefined && (obj.duration = message.duration);
    message.marketId !== undefined && (obj.marketId = message.marketId);
    message.slots !== undefined && (obj.slots = message.slots);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<TradingRewardCreateEvent>
  ): TradingRewardCreateEvent {
    const message = {
      ...baseTradingRewardCreateEvent,
    } as TradingRewardCreateEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    if (object.prizeAmount !== undefined && object.prizeAmount !== null) {
      message.prizeAmount = object.prizeAmount;
    } else {
      message.prizeAmount = "";
    }
    if (object.prizeDenom !== undefined && object.prizeDenom !== null) {
      message.prizeDenom = object.prizeDenom;
    } else {
      message.prizeDenom = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = object.duration;
    } else {
      message.duration = 0;
    }
    if (object.marketId !== undefined && object.marketId !== null) {
      message.marketId = object.marketId;
    } else {
      message.marketId = "";
    }
    if (object.slots !== undefined && object.slots !== null) {
      message.slots = object.slots;
    } else {
      message.slots = 0;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseTradingRewardExpireEvent: object = { rewardId: "" };

export const TradingRewardExpireEvent = {
  encode(
    message: TradingRewardExpireEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): TradingRewardExpireEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseTradingRewardExpireEvent,
    } as TradingRewardExpireEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TradingRewardExpireEvent {
    const message = {
      ...baseTradingRewardExpireEvent,
    } as TradingRewardExpireEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    return message;
  },

  toJSON(message: TradingRewardExpireEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<TradingRewardExpireEvent>
  ): TradingRewardExpireEvent {
    const message = {
      ...baseTradingRewardExpireEvent,
    } as TradingRewardExpireEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    return message;
  },
};

const baseTradingRewardActivationEvent: object = { rewardId: "" };

export const TradingRewardActivationEvent = {
  encode(
    message: TradingRewardActivationEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): TradingRewardActivationEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseTradingRewardActivationEvent,
    } as TradingRewardActivationEvent;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TradingRewardActivationEvent {
    const message = {
      ...baseTradingRewardActivationEvent,
    } as TradingRewardActivationEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    return message;
  },

  toJSON(message: TradingRewardActivationEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<TradingRewardActivationEvent>
  ): TradingRewardActivationEvent {
    const message = {
      ...baseTradingRewardActivationEvent,
    } as TradingRewardActivationEvent;
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    return message;
  },
};

const baseTradingRewardDistributionEvent: object = {
  rewardId: "",
  prizeAmount: "",
  prizeDenom: "",
  winners: "",
};

export const TradingRewardDistributionEvent = {
  encode(
    message: TradingRewardDistributionEvent,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.rewardId !== "") {
      writer.uint32(10).string(message.rewardId);
    }
    if (message.prizeAmount !== "") {
      writer.uint32(18).string(message.prizeAmount);
    }
    if (message.prizeDenom !== "") {
      writer.uint32(26).string(message.prizeDenom);
    }
    for (const v of message.winners) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): TradingRewardDistributionEvent {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseTradingRewardDistributionEvent,
    } as TradingRewardDistributionEvent;
    message.winners = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rewardId = reader.string();
          break;
        case 2:
          message.prizeAmount = reader.string();
          break;
        case 3:
          message.prizeDenom = reader.string();
          break;
        case 4:
          message.winners.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TradingRewardDistributionEvent {
    const message = {
      ...baseTradingRewardDistributionEvent,
    } as TradingRewardDistributionEvent;
    message.winners = [];
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = String(object.rewardId);
    } else {
      message.rewardId = "";
    }
    if (object.prizeAmount !== undefined && object.prizeAmount !== null) {
      message.prizeAmount = String(object.prizeAmount);
    } else {
      message.prizeAmount = "";
    }
    if (object.prizeDenom !== undefined && object.prizeDenom !== null) {
      message.prizeDenom = String(object.prizeDenom);
    } else {
      message.prizeDenom = "";
    }
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: TradingRewardDistributionEvent): unknown {
    const obj: any = {};
    message.rewardId !== undefined && (obj.rewardId = message.rewardId);
    message.prizeAmount !== undefined &&
      (obj.prizeAmount = message.prizeAmount);
    message.prizeDenom !== undefined && (obj.prizeDenom = message.prizeDenom);
    if (message.winners) {
      obj.winners = message.winners.map((e) => e);
    } else {
      obj.winners = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<TradingRewardDistributionEvent>
  ): TradingRewardDistributionEvent {
    const message = {
      ...baseTradingRewardDistributionEvent,
    } as TradingRewardDistributionEvent;
    message.winners = [];
    if (object.rewardId !== undefined && object.rewardId !== null) {
      message.rewardId = object.rewardId;
    } else {
      message.rewardId = "";
    }
    if (object.prizeAmount !== undefined && object.prizeAmount !== null) {
      message.prizeAmount = object.prizeAmount;
    } else {
      message.prizeAmount = "";
    }
    if (object.prizeDenom !== undefined && object.prizeDenom !== null) {
      message.prizeDenom = object.prizeDenom;
    } else {
      message.prizeDenom = "";
    }
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(e);
      }
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
