// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file zetachain/zetacore/crosschain/tx.proto (package zetachain.zetacore.crosschain, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { CoinType } from "../pkg/coin/coin_pb.js";
import type { Proof } from "../pkg/proofs/proofs_pb.js";
import type { ReceiveStatus } from "../pkg/chains/chains_pb.js";
import type { CallOptions, ProtocolContractVersion, RevertOptions } from "./cross_chain_tx_pb.js";
import type { RateLimiterFlags } from "./rate_limiter_flags_pb.js";

/**
 * InboundStatus represents the status of an observed inbound
 *
 * @generated from enum zetachain.zetacore.crosschain.InboundStatus
 */
export declare enum InboundStatus {
  /**
   * @generated from enum value: success = 0;
   */
  success = 0,

  /**
   * @generated from enum value: insufficient_depositor_fee = 1;
   */
  insufficient_depositor_fee = 1,
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgMigrateTssFunds
 */
export declare class MsgMigrateTssFunds extends Message<MsgMigrateTssFunds> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: string amount = 3;
   */
  amount: string;

  constructor(data?: PartialMessage<MsgMigrateTssFunds>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgMigrateTssFunds";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgMigrateTssFunds;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgMigrateTssFunds;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgMigrateTssFunds;

  static equals(a: MsgMigrateTssFunds | PlainMessage<MsgMigrateTssFunds> | undefined, b: MsgMigrateTssFunds | PlainMessage<MsgMigrateTssFunds> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgMigrateTssFundsResponse
 */
export declare class MsgMigrateTssFundsResponse extends Message<MsgMigrateTssFundsResponse> {
  constructor(data?: PartialMessage<MsgMigrateTssFundsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgMigrateTssFundsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgMigrateTssFundsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgMigrateTssFundsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgMigrateTssFundsResponse;

  static equals(a: MsgMigrateTssFundsResponse | PlainMessage<MsgMigrateTssFundsResponse> | undefined, b: MsgMigrateTssFundsResponse | PlainMessage<MsgMigrateTssFundsResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgUpdateTssAddress
 */
export declare class MsgUpdateTssAddress extends Message<MsgUpdateTssAddress> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string tss_pubkey = 2;
   */
  tssPubkey: string;

  constructor(data?: PartialMessage<MsgUpdateTssAddress>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgUpdateTssAddress";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateTssAddress;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateTssAddress;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateTssAddress;

  static equals(a: MsgUpdateTssAddress | PlainMessage<MsgUpdateTssAddress> | undefined, b: MsgUpdateTssAddress | PlainMessage<MsgUpdateTssAddress> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgUpdateTssAddressResponse
 */
export declare class MsgUpdateTssAddressResponse extends Message<MsgUpdateTssAddressResponse> {
  constructor(data?: PartialMessage<MsgUpdateTssAddressResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgUpdateTssAddressResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateTssAddressResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateTssAddressResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateTssAddressResponse;

  static equals(a: MsgUpdateTssAddressResponse | PlainMessage<MsgUpdateTssAddressResponse> | undefined, b: MsgUpdateTssAddressResponse | PlainMessage<MsgUpdateTssAddressResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgAddInboundTracker
 */
export declare class MsgAddInboundTracker extends Message<MsgAddInboundTracker> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: string tx_hash = 3;
   */
  txHash: string;

  /**
   * @generated from field: zetachain.zetacore.pkg.coin.CoinType coin_type = 4;
   */
  coinType: CoinType;

  /**
   * @generated from field: zetachain.zetacore.pkg.proofs.Proof proof = 5 [deprecated = true];
   * @deprecated
   */
  proof?: Proof;

  /**
   * @generated from field: string block_hash = 6 [deprecated = true];
   * @deprecated
   */
  blockHash: string;

  /**
   * @generated from field: int64 tx_index = 7 [deprecated = true];
   * @deprecated
   */
  txIndex: bigint;

  constructor(data?: PartialMessage<MsgAddInboundTracker>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAddInboundTracker";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddInboundTracker;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddInboundTracker;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddInboundTracker;

  static equals(a: MsgAddInboundTracker | PlainMessage<MsgAddInboundTracker> | undefined, b: MsgAddInboundTracker | PlainMessage<MsgAddInboundTracker> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgAddInboundTrackerResponse
 */
export declare class MsgAddInboundTrackerResponse extends Message<MsgAddInboundTrackerResponse> {
  constructor(data?: PartialMessage<MsgAddInboundTrackerResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAddInboundTrackerResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddInboundTrackerResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddInboundTrackerResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddInboundTrackerResponse;

  static equals(a: MsgAddInboundTrackerResponse | PlainMessage<MsgAddInboundTrackerResponse> | undefined, b: MsgAddInboundTrackerResponse | PlainMessage<MsgAddInboundTrackerResponse> | undefined): boolean;
}

/**
 * TODO: https://github.com/zeta-chain/node/issues/3083
 *
 * @generated from message zetachain.zetacore.crosschain.MsgWhitelistERC20
 */
export declare class MsgWhitelistERC20 extends Message<MsgWhitelistERC20> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string erc20_address = 2;
   */
  erc20Address: string;

  /**
   * @generated from field: int64 chain_id = 3;
   */
  chainId: bigint;

  /**
   * @generated from field: string name = 4;
   */
  name: string;

  /**
   * @generated from field: string symbol = 5;
   */
  symbol: string;

  /**
   * @generated from field: uint32 decimals = 6;
   */
  decimals: number;

  /**
   * @generated from field: int64 gas_limit = 7;
   */
  gasLimit: bigint;

  constructor(data?: PartialMessage<MsgWhitelistERC20>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgWhitelistERC20";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgWhitelistERC20;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgWhitelistERC20;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgWhitelistERC20;

  static equals(a: MsgWhitelistERC20 | PlainMessage<MsgWhitelistERC20> | undefined, b: MsgWhitelistERC20 | PlainMessage<MsgWhitelistERC20> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgWhitelistERC20Response
 */
export declare class MsgWhitelistERC20Response extends Message<MsgWhitelistERC20Response> {
  /**
   * @generated from field: string zrc20_address = 1;
   */
  zrc20Address: string;

  /**
   * @generated from field: string cctx_index = 2;
   */
  cctxIndex: string;

  constructor(data?: PartialMessage<MsgWhitelistERC20Response>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgWhitelistERC20Response";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgWhitelistERC20Response;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgWhitelistERC20Response;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgWhitelistERC20Response;

  static equals(a: MsgWhitelistERC20Response | PlainMessage<MsgWhitelistERC20Response> | undefined, b: MsgWhitelistERC20Response | PlainMessage<MsgWhitelistERC20Response> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgAddOutboundTracker
 */
export declare class MsgAddOutboundTracker extends Message<MsgAddOutboundTracker> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: uint64 nonce = 3;
   */
  nonce: bigint;

  /**
   * @generated from field: string tx_hash = 4;
   */
  txHash: string;

  /**
   * @generated from field: zetachain.zetacore.pkg.proofs.Proof proof = 5 [deprecated = true];
   * @deprecated
   */
  proof?: Proof;

  /**
   * @generated from field: string block_hash = 6 [deprecated = true];
   * @deprecated
   */
  blockHash: string;

  /**
   * @generated from field: int64 tx_index = 7 [deprecated = true];
   * @deprecated
   */
  txIndex: bigint;

  constructor(data?: PartialMessage<MsgAddOutboundTracker>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAddOutboundTracker";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddOutboundTracker;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddOutboundTracker;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddOutboundTracker;

  static equals(a: MsgAddOutboundTracker | PlainMessage<MsgAddOutboundTracker> | undefined, b: MsgAddOutboundTracker | PlainMessage<MsgAddOutboundTracker> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgAddOutboundTrackerResponse
 */
export declare class MsgAddOutboundTrackerResponse extends Message<MsgAddOutboundTrackerResponse> {
  /**
   * if the tx was removed from the tracker due to no pending cctx
   *
   * @generated from field: bool is_removed = 1;
   */
  isRemoved: boolean;

  constructor(data?: PartialMessage<MsgAddOutboundTrackerResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAddOutboundTrackerResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddOutboundTrackerResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddOutboundTrackerResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddOutboundTrackerResponse;

  static equals(a: MsgAddOutboundTrackerResponse | PlainMessage<MsgAddOutboundTrackerResponse> | undefined, b: MsgAddOutboundTrackerResponse | PlainMessage<MsgAddOutboundTrackerResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgRemoveOutboundTracker
 */
export declare class MsgRemoveOutboundTracker extends Message<MsgRemoveOutboundTracker> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: uint64 nonce = 3;
   */
  nonce: bigint;

  constructor(data?: PartialMessage<MsgRemoveOutboundTracker>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgRemoveOutboundTracker";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRemoveOutboundTracker;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRemoveOutboundTracker;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRemoveOutboundTracker;

  static equals(a: MsgRemoveOutboundTracker | PlainMessage<MsgRemoveOutboundTracker> | undefined, b: MsgRemoveOutboundTracker | PlainMessage<MsgRemoveOutboundTracker> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgRemoveOutboundTrackerResponse
 */
export declare class MsgRemoveOutboundTrackerResponse extends Message<MsgRemoveOutboundTrackerResponse> {
  constructor(data?: PartialMessage<MsgRemoveOutboundTrackerResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgRemoveOutboundTrackerResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRemoveOutboundTrackerResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRemoveOutboundTrackerResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRemoveOutboundTrackerResponse;

  static equals(a: MsgRemoveOutboundTrackerResponse | PlainMessage<MsgRemoveOutboundTrackerResponse> | undefined, b: MsgRemoveOutboundTrackerResponse | PlainMessage<MsgRemoveOutboundTrackerResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteGasPrice
 */
export declare class MsgVoteGasPrice extends Message<MsgVoteGasPrice> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: uint64 price = 3;
   */
  price: bigint;

  /**
   * @generated from field: uint64 priority_fee = 6;
   */
  priorityFee: bigint;

  /**
   * @generated from field: uint64 block_number = 4;
   */
  blockNumber: bigint;

  /**
   * @generated from field: string supply = 5 [deprecated = true];
   * @deprecated
   */
  supply: string;

  constructor(data?: PartialMessage<MsgVoteGasPrice>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteGasPrice";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteGasPrice;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteGasPrice;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteGasPrice;

  static equals(a: MsgVoteGasPrice | PlainMessage<MsgVoteGasPrice> | undefined, b: MsgVoteGasPrice | PlainMessage<MsgVoteGasPrice> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteGasPriceResponse
 */
export declare class MsgVoteGasPriceResponse extends Message<MsgVoteGasPriceResponse> {
  constructor(data?: PartialMessage<MsgVoteGasPriceResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteGasPriceResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteGasPriceResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteGasPriceResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteGasPriceResponse;

  static equals(a: MsgVoteGasPriceResponse | PlainMessage<MsgVoteGasPriceResponse> | undefined, b: MsgVoteGasPriceResponse | PlainMessage<MsgVoteGasPriceResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteOutbound
 */
export declare class MsgVoteOutbound extends Message<MsgVoteOutbound> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string cctx_hash = 2;
   */
  cctxHash: string;

  /**
   * @generated from field: string observed_outbound_hash = 3;
   */
  observedOutboundHash: string;

  /**
   * @generated from field: uint64 observed_outbound_block_height = 4;
   */
  observedOutboundBlockHeight: bigint;

  /**
   * @generated from field: uint64 observed_outbound_gas_used = 10;
   */
  observedOutboundGasUsed: bigint;

  /**
   * @generated from field: string observed_outbound_effective_gas_price = 11;
   */
  observedOutboundEffectiveGasPrice: string;

  /**
   * @generated from field: uint64 observed_outbound_effective_gas_limit = 12;
   */
  observedOutboundEffectiveGasLimit: bigint;

  /**
   * @generated from field: string value_received = 5;
   */
  valueReceived: string;

  /**
   * @generated from field: zetachain.zetacore.pkg.chains.ReceiveStatus status = 6;
   */
  status: ReceiveStatus;

  /**
   * @generated from field: int64 outbound_chain = 7;
   */
  outboundChain: bigint;

  /**
   * @generated from field: uint64 outbound_tss_nonce = 8;
   */
  outboundTssNonce: bigint;

  /**
   * @generated from field: zetachain.zetacore.pkg.coin.CoinType coin_type = 9;
   */
  coinType: CoinType;

  constructor(data?: PartialMessage<MsgVoteOutbound>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteOutbound";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteOutbound;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteOutbound;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteOutbound;

  static equals(a: MsgVoteOutbound | PlainMessage<MsgVoteOutbound> | undefined, b: MsgVoteOutbound | PlainMessage<MsgVoteOutbound> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteOutboundResponse
 */
export declare class MsgVoteOutboundResponse extends Message<MsgVoteOutboundResponse> {
  constructor(data?: PartialMessage<MsgVoteOutboundResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteOutboundResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteOutboundResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteOutboundResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteOutboundResponse;

  static equals(a: MsgVoteOutboundResponse | PlainMessage<MsgVoteOutboundResponse> | undefined, b: MsgVoteOutboundResponse | PlainMessage<MsgVoteOutboundResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteInbound
 */
export declare class MsgVoteInbound extends Message<MsgVoteInbound> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string sender = 2;
   */
  sender: string;

  /**
   * @generated from field: int64 sender_chain_id = 3;
   */
  senderChainId: bigint;

  /**
   * @generated from field: string receiver = 4;
   */
  receiver: string;

  /**
   * @generated from field: int64 receiver_chain = 5;
   */
  receiverChain: bigint;

  /**
   *  string zeta_burnt = 6;
   *
   * @generated from field: string amount = 6;
   */
  amount: string;

  /**
   *  string mMint = 7;
   *
   * @generated from field: string message = 8;
   */
  message: string;

  /**
   * @generated from field: string inbound_hash = 9;
   */
  inboundHash: string;

  /**
   * @generated from field: uint64 inbound_block_height = 10;
   */
  inboundBlockHeight: bigint;

  /**
   * Deprecated (v21), use CallOptions
   *
   * @generated from field: uint64 gas_limit = 11;
   */
  gasLimit: bigint;

  /**
   * @generated from field: zetachain.zetacore.pkg.coin.CoinType coin_type = 12;
   */
  coinType: CoinType;

  /**
   * @generated from field: string tx_origin = 13;
   */
  txOrigin: string;

  /**
   * @generated from field: string asset = 14;
   */
  asset: string;

  /**
   * event index of the sent asset in the observed tx
   *
   * @generated from field: uint64 event_index = 15;
   */
  eventIndex: bigint;

  /**
   * protocol contract version to use for the cctx workflow
   *
   * @generated from field: zetachain.zetacore.crosschain.ProtocolContractVersion protocol_contract_version = 16;
   */
  protocolContractVersion: ProtocolContractVersion;

  /**
   * revert options provided by the sender
   *
   * @generated from field: zetachain.zetacore.crosschain.RevertOptions revert_options = 17;
   */
  revertOptions?: RevertOptions;

  /**
   * @generated from field: zetachain.zetacore.crosschain.CallOptions call_options = 18;
   */
  callOptions?: CallOptions;

  /**
   * define if a smart contract call should be made with asset
   *
   * @generated from field: bool is_cross_chain_call = 19;
   */
  isCrossChainCall: boolean;

  /**
   * success or failure
   *
   * @generated from field: zetachain.zetacore.crosschain.InboundStatus status = 20;
   */
  status: InboundStatus;

  constructor(data?: PartialMessage<MsgVoteInbound>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteInbound";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteInbound;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteInbound;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteInbound;

  static equals(a: MsgVoteInbound | PlainMessage<MsgVoteInbound> | undefined, b: MsgVoteInbound | PlainMessage<MsgVoteInbound> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteInboundResponse
 */
export declare class MsgVoteInboundResponse extends Message<MsgVoteInboundResponse> {
  constructor(data?: PartialMessage<MsgVoteInboundResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteInboundResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteInboundResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteInboundResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteInboundResponse;

  static equals(a: MsgVoteInboundResponse | PlainMessage<MsgVoteInboundResponse> | undefined, b: MsgVoteInboundResponse | PlainMessage<MsgVoteInboundResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgAbortStuckCCTX
 */
export declare class MsgAbortStuckCCTX extends Message<MsgAbortStuckCCTX> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string cctx_index = 2;
   */
  cctxIndex: string;

  constructor(data?: PartialMessage<MsgAbortStuckCCTX>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAbortStuckCCTX";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAbortStuckCCTX;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAbortStuckCCTX;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAbortStuckCCTX;

  static equals(a: MsgAbortStuckCCTX | PlainMessage<MsgAbortStuckCCTX> | undefined, b: MsgAbortStuckCCTX | PlainMessage<MsgAbortStuckCCTX> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgAbortStuckCCTXResponse
 */
export declare class MsgAbortStuckCCTXResponse extends Message<MsgAbortStuckCCTXResponse> {
  constructor(data?: PartialMessage<MsgAbortStuckCCTXResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAbortStuckCCTXResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAbortStuckCCTXResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAbortStuckCCTXResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAbortStuckCCTXResponse;

  static equals(a: MsgAbortStuckCCTXResponse | PlainMessage<MsgAbortStuckCCTXResponse> | undefined, b: MsgAbortStuckCCTXResponse | PlainMessage<MsgAbortStuckCCTXResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgRefundAbortedCCTX
 */
export declare class MsgRefundAbortedCCTX extends Message<MsgRefundAbortedCCTX> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string cctx_index = 2;
   */
  cctxIndex: string;

  /**
   * if not provided, the refund will be sent to the sender/txOrgin
   *
   * @generated from field: string refund_address = 3;
   */
  refundAddress: string;

  constructor(data?: PartialMessage<MsgRefundAbortedCCTX>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgRefundAbortedCCTX";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRefundAbortedCCTX;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRefundAbortedCCTX;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRefundAbortedCCTX;

  static equals(a: MsgRefundAbortedCCTX | PlainMessage<MsgRefundAbortedCCTX> | undefined, b: MsgRefundAbortedCCTX | PlainMessage<MsgRefundAbortedCCTX> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgRefundAbortedCCTXResponse
 */
export declare class MsgRefundAbortedCCTXResponse extends Message<MsgRefundAbortedCCTXResponse> {
  constructor(data?: PartialMessage<MsgRefundAbortedCCTXResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgRefundAbortedCCTXResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRefundAbortedCCTXResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRefundAbortedCCTXResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRefundAbortedCCTXResponse;

  static equals(a: MsgRefundAbortedCCTXResponse | PlainMessage<MsgRefundAbortedCCTXResponse> | undefined, b: MsgRefundAbortedCCTXResponse | PlainMessage<MsgRefundAbortedCCTXResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgUpdateRateLimiterFlags
 */
export declare class MsgUpdateRateLimiterFlags extends Message<MsgUpdateRateLimiterFlags> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: zetachain.zetacore.crosschain.RateLimiterFlags rate_limiter_flags = 2;
   */
  rateLimiterFlags?: RateLimiterFlags;

  constructor(data?: PartialMessage<MsgUpdateRateLimiterFlags>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgUpdateRateLimiterFlags";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateRateLimiterFlags;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateRateLimiterFlags;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateRateLimiterFlags;

  static equals(a: MsgUpdateRateLimiterFlags | PlainMessage<MsgUpdateRateLimiterFlags> | undefined, b: MsgUpdateRateLimiterFlags | PlainMessage<MsgUpdateRateLimiterFlags> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgUpdateRateLimiterFlagsResponse
 */
export declare class MsgUpdateRateLimiterFlagsResponse extends Message<MsgUpdateRateLimiterFlagsResponse> {
  constructor(data?: PartialMessage<MsgUpdateRateLimiterFlagsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgUpdateRateLimiterFlagsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateRateLimiterFlagsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateRateLimiterFlagsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateRateLimiterFlagsResponse;

  static equals(a: MsgUpdateRateLimiterFlagsResponse | PlainMessage<MsgUpdateRateLimiterFlagsResponse> | undefined, b: MsgUpdateRateLimiterFlagsResponse | PlainMessage<MsgUpdateRateLimiterFlagsResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgMigrateERC20CustodyFunds
 */
export declare class MsgMigrateERC20CustodyFunds extends Message<MsgMigrateERC20CustodyFunds> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * @generated from field: string new_custody_address = 3;
   */
  newCustodyAddress: string;

  /**
   * @generated from field: string erc20_address = 4;
   */
  erc20Address: string;

  /**
   * @generated from field: string amount = 5;
   */
  amount: string;

  constructor(data?: PartialMessage<MsgMigrateERC20CustodyFunds>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgMigrateERC20CustodyFunds";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgMigrateERC20CustodyFunds;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgMigrateERC20CustodyFunds;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgMigrateERC20CustodyFunds;

  static equals(a: MsgMigrateERC20CustodyFunds | PlainMessage<MsgMigrateERC20CustodyFunds> | undefined, b: MsgMigrateERC20CustodyFunds | PlainMessage<MsgMigrateERC20CustodyFunds> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgMigrateERC20CustodyFundsResponse
 */
export declare class MsgMigrateERC20CustodyFundsResponse extends Message<MsgMigrateERC20CustodyFundsResponse> {
  /**
   * @generated from field: string cctx_index = 1;
   */
  cctxIndex: string;

  constructor(data?: PartialMessage<MsgMigrateERC20CustodyFundsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgMigrateERC20CustodyFundsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgMigrateERC20CustodyFundsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgMigrateERC20CustodyFundsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgMigrateERC20CustodyFundsResponse;

  static equals(a: MsgMigrateERC20CustodyFundsResponse | PlainMessage<MsgMigrateERC20CustodyFundsResponse> | undefined, b: MsgMigrateERC20CustodyFundsResponse | PlainMessage<MsgMigrateERC20CustodyFundsResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgUpdateERC20CustodyPauseStatus
 */
export declare class MsgUpdateERC20CustodyPauseStatus extends Message<MsgUpdateERC20CustodyPauseStatus> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * pause or unpause
   * true = pause, false = unpause
   *
   * @generated from field: bool pause = 3;
   */
  pause: boolean;

  constructor(data?: PartialMessage<MsgUpdateERC20CustodyPauseStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgUpdateERC20CustodyPauseStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateERC20CustodyPauseStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateERC20CustodyPauseStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateERC20CustodyPauseStatus;

  static equals(a: MsgUpdateERC20CustodyPauseStatus | PlainMessage<MsgUpdateERC20CustodyPauseStatus> | undefined, b: MsgUpdateERC20CustodyPauseStatus | PlainMessage<MsgUpdateERC20CustodyPauseStatus> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgUpdateERC20CustodyPauseStatusResponse
 */
export declare class MsgUpdateERC20CustodyPauseStatusResponse extends Message<MsgUpdateERC20CustodyPauseStatusResponse> {
  /**
   * @generated from field: string cctx_index = 1;
   */
  cctxIndex: string;

  constructor(data?: PartialMessage<MsgUpdateERC20CustodyPauseStatusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgUpdateERC20CustodyPauseStatusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgUpdateERC20CustodyPauseStatusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgUpdateERC20CustodyPauseStatusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgUpdateERC20CustodyPauseStatusResponse;

  static equals(a: MsgUpdateERC20CustodyPauseStatusResponse | PlainMessage<MsgUpdateERC20CustodyPauseStatusResponse> | undefined, b: MsgUpdateERC20CustodyPauseStatusResponse | PlainMessage<MsgUpdateERC20CustodyPauseStatusResponse> | undefined): boolean;
}

