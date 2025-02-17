// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file zetachain/zetacore/observer/tx.proto (package zetachain.zetacore.observer, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { ObserverUpdateReason } from "./observer_pb.js";
import type { HeaderData } from "../pkg/proofs/proofs_pb.js";
import type { ChainParams } from "./params_pb.js";
import type { ConfirmationParams } from "./confirmation_params_pb.js";
import type { Blame } from "./blame_pb.js";
import type { ReceiveStatus } from "../pkg/chains/chains_pb.js";
import type { GasPriceIncreaseFlags } from "./crosschain_flags_pb.js";
import type { OperationalFlags } from "./operational_pb.js";

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateObserver
 */
export declare class MsgUpdateObserver extends Message<MsgUpdateObserver> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string old_observer_address = 2;
   */
  oldObserverAddress: string;

  /**
   * @generated from field: string new_observer_address = 3;
   */
  newObserverAddress: string;

  /**
   * @generated from field: zetachain.zetacore.observer.ObserverUpdateReason update_reason = 4;
   */
  updateReason: ObserverUpdateReason;

  constructor(data?: PartialMessage<MsgUpdateObserver>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateObserver";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateObserver;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateObserver;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateObserver;

  static equals(a: MsgUpdateObserver | PlainMessage<MsgUpdateObserver> | undefined, b: MsgUpdateObserver | PlainMessage<MsgUpdateObserver> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateObserverResponse
 */
export declare class MsgUpdateObserverResponse extends Message<MsgUpdateObserverResponse> {
  constructor(data?: PartialMessage<MsgUpdateObserverResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateObserverResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateObserverResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateObserverResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateObserverResponse;

  static equals(a: MsgUpdateObserverResponse | PlainMessage<MsgUpdateObserverResponse> | undefined, b: MsgUpdateObserverResponse | PlainMessage<MsgUpdateObserverResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgVoteBlockHeader
 */
export declare class MsgVoteBlockHeader extends Message<MsgVoteBlockHeader> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: bytes block_hash = 3;
   */
  blockHash: Uint8Array;

  /**
   * @generated from field: int64 height = 4;
   */
  height: bigint;

  /**
   * @generated from field: zetachain.zetacore.pkg.proofs.HeaderData header = 5;
   */
  header?: HeaderData;

  constructor(data?: PartialMessage<MsgVoteBlockHeader>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgVoteBlockHeader";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteBlockHeader;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteBlockHeader;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteBlockHeader;

  static equals(a: MsgVoteBlockHeader | PlainMessage<MsgVoteBlockHeader> | undefined, b: MsgVoteBlockHeader | PlainMessage<MsgVoteBlockHeader> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgVoteBlockHeaderResponse
 */
export declare class MsgVoteBlockHeaderResponse extends Message<MsgVoteBlockHeaderResponse> {
  /**
   * @generated from field: bool ballot_created = 1;
   */
  ballotCreated: boolean;

  /**
   * @generated from field: bool vote_finalized = 2;
   */
  voteFinalized: boolean;

  constructor(data?: PartialMessage<MsgVoteBlockHeaderResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgVoteBlockHeaderResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteBlockHeaderResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteBlockHeaderResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteBlockHeaderResponse;

  static equals(a: MsgVoteBlockHeaderResponse | PlainMessage<MsgVoteBlockHeaderResponse> | undefined, b: MsgVoteBlockHeaderResponse | PlainMessage<MsgVoteBlockHeaderResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateChainParams
 */
export declare class MsgUpdateChainParams extends Message<MsgUpdateChainParams> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: zetachain.zetacore.observer.ChainParams chainParams = 2;
   */
  chainParams?: ChainParams;

  constructor(data?: PartialMessage<MsgUpdateChainParams>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateChainParams";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateChainParams;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateChainParams;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateChainParams;

  static equals(a: MsgUpdateChainParams | PlainMessage<MsgUpdateChainParams> | undefined, b: MsgUpdateChainParams | PlainMessage<MsgUpdateChainParams> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateChainParamsResponse
 */
export declare class MsgUpdateChainParamsResponse extends Message<MsgUpdateChainParamsResponse> {
  constructor(data?: PartialMessage<MsgUpdateChainParamsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateChainParamsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateChainParamsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateChainParamsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateChainParamsResponse;

  static equals(a: MsgUpdateChainParamsResponse | PlainMessage<MsgUpdateChainParamsResponse> | undefined, b: MsgUpdateChainParamsResponse | PlainMessage<MsgUpdateChainParamsResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateOperationalChainParams
 */
export declare class MsgUpdateOperationalChainParams extends Message<MsgUpdateOperationalChainParams> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: uint64 gas_price_ticker = 3;
   */
  gasPriceTicker: bigint;

  /**
   * @generated from field: uint64 inbound_ticker = 4;
   */
  inboundTicker: bigint;

  /**
   * @generated from field: uint64 outbound_ticker = 5;
   */
  outboundTicker: bigint;

  /**
   * @generated from field: uint64 watch_utxo_ticker = 6;
   */
  watchUtxoTicker: bigint;

  /**
   * @generated from field: int64 outbound_schedule_interval = 7;
   */
  outboundScheduleInterval: bigint;

  /**
   * @generated from field: int64 outbound_schedule_lookahead = 8;
   */
  outboundScheduleLookahead: bigint;

  /**
   * @generated from field: zetachain.zetacore.observer.ConfirmationParams confirmation_params = 9;
   */
  confirmationParams?: ConfirmationParams;

  constructor(data?: PartialMessage<MsgUpdateOperationalChainParams>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateOperationalChainParams";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateOperationalChainParams;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateOperationalChainParams;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateOperationalChainParams;

  static equals(a: MsgUpdateOperationalChainParams | PlainMessage<MsgUpdateOperationalChainParams> | undefined, b: MsgUpdateOperationalChainParams | PlainMessage<MsgUpdateOperationalChainParams> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateOperationalChainParamsResponse
 */
export declare class MsgUpdateOperationalChainParamsResponse extends Message<MsgUpdateOperationalChainParamsResponse> {
  constructor(data?: PartialMessage<MsgUpdateOperationalChainParamsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateOperationalChainParamsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateOperationalChainParamsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateOperationalChainParamsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateOperationalChainParamsResponse;

  static equals(a: MsgUpdateOperationalChainParamsResponse | PlainMessage<MsgUpdateOperationalChainParamsResponse> | undefined, b: MsgUpdateOperationalChainParamsResponse | PlainMessage<MsgUpdateOperationalChainParamsResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgRemoveChainParams
 */
export declare class MsgRemoveChainParams extends Message<MsgRemoveChainParams> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  constructor(data?: PartialMessage<MsgRemoveChainParams>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgRemoveChainParams";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRemoveChainParams;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRemoveChainParams;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRemoveChainParams;

  static equals(a: MsgRemoveChainParams | PlainMessage<MsgRemoveChainParams> | undefined, b: MsgRemoveChainParams | PlainMessage<MsgRemoveChainParams> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgRemoveChainParamsResponse
 */
export declare class MsgRemoveChainParamsResponse extends Message<MsgRemoveChainParamsResponse> {
  constructor(data?: PartialMessage<MsgRemoveChainParamsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgRemoveChainParamsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRemoveChainParamsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRemoveChainParamsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRemoveChainParamsResponse;

  static equals(a: MsgRemoveChainParamsResponse | PlainMessage<MsgRemoveChainParamsResponse> | undefined, b: MsgRemoveChainParamsResponse | PlainMessage<MsgRemoveChainParamsResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgAddObserver
 */
export declare class MsgAddObserver extends Message<MsgAddObserver> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string observer_address = 2;
   */
  observerAddress: string;

  /**
   * @generated from field: string zetaclient_grantee_pubkey = 3;
   */
  zetaclientGranteePubkey: string;

  /**
   * @generated from field: bool add_node_account_only = 4;
   */
  addNodeAccountOnly: boolean;

  constructor(data?: PartialMessage<MsgAddObserver>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgAddObserver";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddObserver;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddObserver;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddObserver;

  static equals(a: MsgAddObserver | PlainMessage<MsgAddObserver> | undefined, b: MsgAddObserver | PlainMessage<MsgAddObserver> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgAddObserverResponse
 */
export declare class MsgAddObserverResponse extends Message<MsgAddObserverResponse> {
  constructor(data?: PartialMessage<MsgAddObserverResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgAddObserverResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddObserverResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddObserverResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddObserverResponse;

  static equals(a: MsgAddObserverResponse | PlainMessage<MsgAddObserverResponse> | undefined, b: MsgAddObserverResponse | PlainMessage<MsgAddObserverResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgVoteBlame
 */
export declare class MsgVoteBlame extends Message<MsgVoteBlame> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: zetachain.zetacore.observer.Blame blame_info = 3;
   */
  blameInfo?: Blame;

  constructor(data?: PartialMessage<MsgVoteBlame>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgVoteBlame";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteBlame;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteBlame;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteBlame;

  static equals(a: MsgVoteBlame | PlainMessage<MsgVoteBlame> | undefined, b: MsgVoteBlame | PlainMessage<MsgVoteBlame> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgVoteBlameResponse
 */
export declare class MsgVoteBlameResponse extends Message<MsgVoteBlameResponse> {
  constructor(data?: PartialMessage<MsgVoteBlameResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgVoteBlameResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteBlameResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteBlameResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteBlameResponse;

  static equals(a: MsgVoteBlameResponse | PlainMessage<MsgVoteBlameResponse> | undefined, b: MsgVoteBlameResponse | PlainMessage<MsgVoteBlameResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateKeygen
 */
export declare class MsgUpdateKeygen extends Message<MsgUpdateKeygen> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 block = 2;
   */
  block: bigint;

  constructor(data?: PartialMessage<MsgUpdateKeygen>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateKeygen";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateKeygen;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateKeygen;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateKeygen;

  static equals(a: MsgUpdateKeygen | PlainMessage<MsgUpdateKeygen> | undefined, b: MsgUpdateKeygen | PlainMessage<MsgUpdateKeygen> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateKeygenResponse
 */
export declare class MsgUpdateKeygenResponse extends Message<MsgUpdateKeygenResponse> {
  constructor(data?: PartialMessage<MsgUpdateKeygenResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateKeygenResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateKeygenResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateKeygenResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateKeygenResponse;

  static equals(a: MsgUpdateKeygenResponse | PlainMessage<MsgUpdateKeygenResponse> | undefined, b: MsgUpdateKeygenResponse | PlainMessage<MsgUpdateKeygenResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgResetChainNonces
 */
export declare class MsgResetChainNonces extends Message<MsgResetChainNonces> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: int64 chain_nonce_low = 3;
   */
  chainNonceLow: bigint;

  /**
   * @generated from field: int64 chain_nonce_high = 4;
   */
  chainNonceHigh: bigint;

  constructor(data?: PartialMessage<MsgResetChainNonces>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgResetChainNonces";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgResetChainNonces;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgResetChainNonces;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgResetChainNonces;

  static equals(a: MsgResetChainNonces | PlainMessage<MsgResetChainNonces> | undefined, b: MsgResetChainNonces | PlainMessage<MsgResetChainNonces> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgResetChainNoncesResponse
 */
export declare class MsgResetChainNoncesResponse extends Message<MsgResetChainNoncesResponse> {
  constructor(data?: PartialMessage<MsgResetChainNoncesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgResetChainNoncesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgResetChainNoncesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgResetChainNoncesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgResetChainNoncesResponse;

  static equals(a: MsgResetChainNoncesResponse | PlainMessage<MsgResetChainNoncesResponse> | undefined, b: MsgResetChainNoncesResponse | PlainMessage<MsgResetChainNoncesResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgVoteTSS
 */
export declare class MsgVoteTSS extends Message<MsgVoteTSS> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string tss_pubkey = 2;
   */
  tssPubkey: string;

  /**
   * @generated from field: int64 keygen_zeta_height = 3;
   */
  keygenZetaHeight: bigint;

  /**
   * @generated from field: zetachain.zetacore.pkg.chains.ReceiveStatus status = 4;
   */
  status: ReceiveStatus;

  constructor(data?: PartialMessage<MsgVoteTSS>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgVoteTSS";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteTSS;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteTSS;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteTSS;

  static equals(a: MsgVoteTSS | PlainMessage<MsgVoteTSS> | undefined, b: MsgVoteTSS | PlainMessage<MsgVoteTSS> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgVoteTSSResponse
 */
export declare class MsgVoteTSSResponse extends Message<MsgVoteTSSResponse> {
  /**
   * @generated from field: bool ballot_created = 1;
   */
  ballotCreated: boolean;

  /**
   * @generated from field: bool vote_finalized = 2;
   */
  voteFinalized: boolean;

  /**
   * @generated from field: bool keygen_success = 3;
   */
  keygenSuccess: boolean;

  constructor(data?: PartialMessage<MsgVoteTSSResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgVoteTSSResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteTSSResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteTSSResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteTSSResponse;

  static equals(a: MsgVoteTSSResponse | PlainMessage<MsgVoteTSSResponse> | undefined, b: MsgVoteTSSResponse | PlainMessage<MsgVoteTSSResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgEnableCCTX
 */
export declare class MsgEnableCCTX extends Message<MsgEnableCCTX> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: bool enableInbound = 2;
   */
  enableInbound: boolean;

  /**
   * @generated from field: bool enableOutbound = 3;
   */
  enableOutbound: boolean;

  constructor(data?: PartialMessage<MsgEnableCCTX>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgEnableCCTX";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgEnableCCTX;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgEnableCCTX;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgEnableCCTX;

  static equals(a: MsgEnableCCTX | PlainMessage<MsgEnableCCTX> | undefined, b: MsgEnableCCTX | PlainMessage<MsgEnableCCTX> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgEnableCCTXResponse
 */
export declare class MsgEnableCCTXResponse extends Message<MsgEnableCCTXResponse> {
  constructor(data?: PartialMessage<MsgEnableCCTXResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgEnableCCTXResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgEnableCCTXResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgEnableCCTXResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgEnableCCTXResponse;

  static equals(a: MsgEnableCCTXResponse | PlainMessage<MsgEnableCCTXResponse> | undefined, b: MsgEnableCCTXResponse | PlainMessage<MsgEnableCCTXResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgDisableCCTX
 */
export declare class MsgDisableCCTX extends Message<MsgDisableCCTX> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: bool disableInbound = 2;
   */
  disableInbound: boolean;

  /**
   * @generated from field: bool disableOutbound = 3;
   */
  disableOutbound: boolean;

  constructor(data?: PartialMessage<MsgDisableCCTX>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgDisableCCTX";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgDisableCCTX;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgDisableCCTX;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgDisableCCTX;

  static equals(a: MsgDisableCCTX | PlainMessage<MsgDisableCCTX> | undefined, b: MsgDisableCCTX | PlainMessage<MsgDisableCCTX> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgDisableCCTXResponse
 */
export declare class MsgDisableCCTXResponse extends Message<MsgDisableCCTXResponse> {
  constructor(data?: PartialMessage<MsgDisableCCTXResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgDisableCCTXResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgDisableCCTXResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgDisableCCTXResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgDisableCCTXResponse;

  static equals(a: MsgDisableCCTXResponse | PlainMessage<MsgDisableCCTXResponse> | undefined, b: MsgDisableCCTXResponse | PlainMessage<MsgDisableCCTXResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateGasPriceIncreaseFlags
 */
export declare class MsgUpdateGasPriceIncreaseFlags extends Message<MsgUpdateGasPriceIncreaseFlags> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: zetachain.zetacore.observer.GasPriceIncreaseFlags gasPriceIncreaseFlags = 2;
   */
  gasPriceIncreaseFlags?: GasPriceIncreaseFlags;

  constructor(data?: PartialMessage<MsgUpdateGasPriceIncreaseFlags>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateGasPriceIncreaseFlags";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateGasPriceIncreaseFlags;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateGasPriceIncreaseFlags;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateGasPriceIncreaseFlags;

  static equals(a: MsgUpdateGasPriceIncreaseFlags | PlainMessage<MsgUpdateGasPriceIncreaseFlags> | undefined, b: MsgUpdateGasPriceIncreaseFlags | PlainMessage<MsgUpdateGasPriceIncreaseFlags> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateGasPriceIncreaseFlagsResponse
 */
export declare class MsgUpdateGasPriceIncreaseFlagsResponse extends Message<MsgUpdateGasPriceIncreaseFlagsResponse> {
  constructor(data?: PartialMessage<MsgUpdateGasPriceIncreaseFlagsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateGasPriceIncreaseFlagsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateGasPriceIncreaseFlagsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateGasPriceIncreaseFlagsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateGasPriceIncreaseFlagsResponse;

  static equals(a: MsgUpdateGasPriceIncreaseFlagsResponse | PlainMessage<MsgUpdateGasPriceIncreaseFlagsResponse> | undefined, b: MsgUpdateGasPriceIncreaseFlagsResponse | PlainMessage<MsgUpdateGasPriceIncreaseFlagsResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateOperationalFlags
 */
export declare class MsgUpdateOperationalFlags extends Message<MsgUpdateOperationalFlags> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: zetachain.zetacore.observer.OperationalFlags operational_flags = 2;
   */
  operationalFlags?: OperationalFlags;

  constructor(data?: PartialMessage<MsgUpdateOperationalFlags>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateOperationalFlags";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateOperationalFlags;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateOperationalFlags;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateOperationalFlags;

  static equals(a: MsgUpdateOperationalFlags | PlainMessage<MsgUpdateOperationalFlags> | undefined, b: MsgUpdateOperationalFlags | PlainMessage<MsgUpdateOperationalFlags> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.MsgUpdateOperationalFlagsResponse
 */
export declare class MsgUpdateOperationalFlagsResponse extends Message<MsgUpdateOperationalFlagsResponse> {
  constructor(data?: PartialMessage<MsgUpdateOperationalFlagsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.MsgUpdateOperationalFlagsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateOperationalFlagsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateOperationalFlagsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateOperationalFlagsResponse;

  static equals(a: MsgUpdateOperationalFlagsResponse | PlainMessage<MsgUpdateOperationalFlagsResponse> | undefined, b: MsgUpdateOperationalFlagsResponse | PlainMessage<MsgUpdateOperationalFlagsResponse> | undefined): boolean;
}

