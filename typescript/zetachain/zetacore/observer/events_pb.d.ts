// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file zetachain/zetacore/observer/events.proto (package zetachain.zetacore.observer, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Voter } from "./ballot_pb.js";
import type { GasPriceIncreaseFlags } from "./crosschain_flags_pb.js";

/**
 * @generated from message zetachain.zetacore.observer.EventBallotCreated
 */
export declare class EventBallotCreated extends Message<EventBallotCreated> {
  /**
   * @generated from field: string msg_type_url = 1;
   */
  msgTypeUrl: string;

  /**
   * @generated from field: string ballot_identifier = 2;
   */
  ballotIdentifier: string;

  /**
   * @generated from field: string observation_hash = 3;
   */
  observationHash: string;

  /**
   * @generated from field: string observation_chain = 4;
   */
  observationChain: string;

  /**
   * @generated from field: string ballot_type = 5;
   */
  ballotType: string;

  constructor(data?: PartialMessage<EventBallotCreated>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.EventBallotCreated";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventBallotCreated;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventBallotCreated;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventBallotCreated;

  static equals(a: EventBallotCreated | PlainMessage<EventBallotCreated> | undefined, b: EventBallotCreated | PlainMessage<EventBallotCreated> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.EventBallotDeleted
 */
export declare class EventBallotDeleted extends Message<EventBallotDeleted> {
  /**
   * @generated from field: string ballot_identifier = 1;
   */
  ballotIdentifier: string;

  /**
   * @generated from field: string ballot_type = 2;
   */
  ballotType: string;

  /**
   * @generated from field: repeated zetachain.zetacore.observer.Voter voters = 3;
   */
  voters: Voter[];

  constructor(data?: PartialMessage<EventBallotDeleted>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.EventBallotDeleted";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventBallotDeleted;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventBallotDeleted;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventBallotDeleted;

  static equals(a: EventBallotDeleted | PlainMessage<EventBallotDeleted> | undefined, b: EventBallotDeleted | PlainMessage<EventBallotDeleted> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.EventKeygenBlockUpdated
 */
export declare class EventKeygenBlockUpdated extends Message<EventKeygenBlockUpdated> {
  /**
   * @generated from field: string msg_type_url = 1;
   */
  msgTypeUrl: string;

  /**
   * @generated from field: string keygen_block = 2;
   */
  keygenBlock: string;

  /**
   * @generated from field: string keygen_pubkeys = 3;
   */
  keygenPubkeys: string;

  constructor(data?: PartialMessage<EventKeygenBlockUpdated>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.EventKeygenBlockUpdated";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventKeygenBlockUpdated;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventKeygenBlockUpdated;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventKeygenBlockUpdated;

  static equals(a: EventKeygenBlockUpdated | PlainMessage<EventKeygenBlockUpdated> | undefined, b: EventKeygenBlockUpdated | PlainMessage<EventKeygenBlockUpdated> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.EventNewObserverAdded
 */
export declare class EventNewObserverAdded extends Message<EventNewObserverAdded> {
  /**
   * @generated from field: string msg_type_url = 1;
   */
  msgTypeUrl: string;

  /**
   * @generated from field: string observer_address = 2;
   */
  observerAddress: string;

  /**
   * @generated from field: string zetaclient_grantee_address = 3;
   */
  zetaclientGranteeAddress: string;

  /**
   * @generated from field: string zetaclient_grantee_pubkey = 4;
   */
  zetaclientGranteePubkey: string;

  /**
   * @generated from field: uint64 observer_last_block_count = 5;
   */
  observerLastBlockCount: bigint;

  constructor(data?: PartialMessage<EventNewObserverAdded>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.EventNewObserverAdded";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventNewObserverAdded;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventNewObserverAdded;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventNewObserverAdded;

  static equals(a: EventNewObserverAdded | PlainMessage<EventNewObserverAdded> | undefined, b: EventNewObserverAdded | PlainMessage<EventNewObserverAdded> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.EventCCTXDisabled
 */
export declare class EventCCTXDisabled extends Message<EventCCTXDisabled> {
  /**
   * @generated from field: string msg_type_url = 1;
   */
  msgTypeUrl: string;

  /**
   * @generated from field: bool isInboundEnabled = 2;
   */
  isInboundEnabled: boolean;

  /**
   * @generated from field: bool isOutboundEnabled = 3;
   */
  isOutboundEnabled: boolean;

  constructor(data?: PartialMessage<EventCCTXDisabled>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.EventCCTXDisabled";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventCCTXDisabled;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventCCTXDisabled;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventCCTXDisabled;

  static equals(a: EventCCTXDisabled | PlainMessage<EventCCTXDisabled> | undefined, b: EventCCTXDisabled | PlainMessage<EventCCTXDisabled> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.EventCCTXEnabled
 */
export declare class EventCCTXEnabled extends Message<EventCCTXEnabled> {
  /**
   * @generated from field: string msg_type_url = 1;
   */
  msgTypeUrl: string;

  /**
   * @generated from field: bool isInboundEnabled = 2;
   */
  isInboundEnabled: boolean;

  /**
   * @generated from field: bool isOutboundEnabled = 3;
   */
  isOutboundEnabled: boolean;

  constructor(data?: PartialMessage<EventCCTXEnabled>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.EventCCTXEnabled";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventCCTXEnabled;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventCCTXEnabled;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventCCTXEnabled;

  static equals(a: EventCCTXEnabled | PlainMessage<EventCCTXEnabled> | undefined, b: EventCCTXEnabled | PlainMessage<EventCCTXEnabled> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.EventGasPriceIncreaseFlagsUpdated
 */
export declare class EventGasPriceIncreaseFlagsUpdated extends Message<EventGasPriceIncreaseFlagsUpdated> {
  /**
   * @generated from field: string msg_type_url = 1;
   */
  msgTypeUrl: string;

  /**
   * @generated from field: zetachain.zetacore.observer.GasPriceIncreaseFlags gasPriceIncreaseFlags = 2;
   */
  gasPriceIncreaseFlags?: GasPriceIncreaseFlags;

  constructor(data?: PartialMessage<EventGasPriceIncreaseFlagsUpdated>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.EventGasPriceIncreaseFlagsUpdated";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventGasPriceIncreaseFlagsUpdated;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventGasPriceIncreaseFlagsUpdated;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventGasPriceIncreaseFlagsUpdated;

  static equals(a: EventGasPriceIncreaseFlagsUpdated | PlainMessage<EventGasPriceIncreaseFlagsUpdated> | undefined, b: EventGasPriceIncreaseFlagsUpdated | PlainMessage<EventGasPriceIncreaseFlagsUpdated> | undefined): boolean;
}

