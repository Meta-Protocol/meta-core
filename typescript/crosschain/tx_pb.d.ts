// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file crosschain/tx.proto (package zetachain.zetacore.crosschain, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { CoinType, Proof, ReceiveStatus } from "../common/common_pb.js";

/**
 * @generated from message zetachain.zetacore.crosschain.MsgCreateTSSVoter
 */
export declare class MsgCreateTSSVoter extends Message<MsgCreateTSSVoter> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string tss_pubkey = 2;
   */
  tssPubkey: string;

  /**
   * @generated from field: int64 keyGenZetaHeight = 3;
   */
  keyGenZetaHeight: bigint;

  /**
   * @generated from field: common.ReceiveStatus status = 4;
   */
  status: ReceiveStatus;

  constructor(data?: PartialMessage<MsgCreateTSSVoter>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgCreateTSSVoter";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgCreateTSSVoter;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgCreateTSSVoter;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgCreateTSSVoter;

  static equals(a: MsgCreateTSSVoter | PlainMessage<MsgCreateTSSVoter> | undefined, b: MsgCreateTSSVoter | PlainMessage<MsgCreateTSSVoter> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgCreateTSSVoterResponse
 */
export declare class MsgCreateTSSVoterResponse extends Message<MsgCreateTSSVoterResponse> {
  constructor(data?: PartialMessage<MsgCreateTSSVoterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgCreateTSSVoterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgCreateTSSVoterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgCreateTSSVoterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgCreateTSSVoterResponse;

  static equals(a: MsgCreateTSSVoterResponse | PlainMessage<MsgCreateTSSVoterResponse> | undefined, b: MsgCreateTSSVoterResponse | PlainMessage<MsgCreateTSSVoterResponse> | undefined): boolean;
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
 * @generated from message zetachain.zetacore.crosschain.MsgAddToInTxTracker
 */
export declare class MsgAddToInTxTracker extends Message<MsgAddToInTxTracker> {
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
   * @generated from field: common.CoinType coin_type = 4;
   */
  coinType: CoinType;

  /**
   * @generated from field: common.Proof proof = 5;
   */
  proof?: Proof;

  /**
   * @generated from field: string block_hash = 6;
   */
  blockHash: string;

  /**
   * @generated from field: int64 tx_index = 7;
   */
  txIndex: bigint;

  constructor(data?: PartialMessage<MsgAddToInTxTracker>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAddToInTxTracker";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddToInTxTracker;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddToInTxTracker;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddToInTxTracker;

  static equals(a: MsgAddToInTxTracker | PlainMessage<MsgAddToInTxTracker> | undefined, b: MsgAddToInTxTracker | PlainMessage<MsgAddToInTxTracker> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgAddToInTxTrackerResponse
 */
export declare class MsgAddToInTxTrackerResponse extends Message<MsgAddToInTxTrackerResponse> {
  constructor(data?: PartialMessage<MsgAddToInTxTrackerResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAddToInTxTrackerResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddToInTxTrackerResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddToInTxTrackerResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddToInTxTrackerResponse;

  static equals(a: MsgAddToInTxTrackerResponse | PlainMessage<MsgAddToInTxTrackerResponse> | undefined, b: MsgAddToInTxTrackerResponse | PlainMessage<MsgAddToInTxTrackerResponse> | undefined): boolean;
}

/**
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
 * @generated from message zetachain.zetacore.crosschain.MsgAddToOutTxTracker
 */
export declare class MsgAddToOutTxTracker extends Message<MsgAddToOutTxTracker> {
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
   * @generated from field: common.Proof proof = 5;
   */
  proof?: Proof;

  /**
   * @generated from field: string block_hash = 6;
   */
  blockHash: string;

  /**
   * @generated from field: int64 tx_index = 7;
   */
  txIndex: bigint;

  constructor(data?: PartialMessage<MsgAddToOutTxTracker>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAddToOutTxTracker";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddToOutTxTracker;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddToOutTxTracker;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddToOutTxTracker;

  static equals(a: MsgAddToOutTxTracker | PlainMessage<MsgAddToOutTxTracker> | undefined, b: MsgAddToOutTxTracker | PlainMessage<MsgAddToOutTxTracker> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgAddToOutTxTrackerResponse
 */
export declare class MsgAddToOutTxTrackerResponse extends Message<MsgAddToOutTxTrackerResponse> {
  /**
   * if the tx was removed from the tracker due to no pending cctx
   *
   * @generated from field: bool is_removed = 1;
   */
  isRemoved: boolean;

  constructor(data?: PartialMessage<MsgAddToOutTxTrackerResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgAddToOutTxTrackerResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgAddToOutTxTrackerResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgAddToOutTxTrackerResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgAddToOutTxTrackerResponse;

  static equals(a: MsgAddToOutTxTrackerResponse | PlainMessage<MsgAddToOutTxTrackerResponse> | undefined, b: MsgAddToOutTxTrackerResponse | PlainMessage<MsgAddToOutTxTrackerResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgRemoveFromOutTxTracker
 */
export declare class MsgRemoveFromOutTxTracker extends Message<MsgRemoveFromOutTxTracker> {
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

  constructor(data?: PartialMessage<MsgRemoveFromOutTxTracker>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgRemoveFromOutTxTracker";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRemoveFromOutTxTracker;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRemoveFromOutTxTracker;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRemoveFromOutTxTracker;

  static equals(a: MsgRemoveFromOutTxTracker | PlainMessage<MsgRemoveFromOutTxTracker> | undefined, b: MsgRemoveFromOutTxTracker | PlainMessage<MsgRemoveFromOutTxTracker> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgRemoveFromOutTxTrackerResponse
 */
export declare class MsgRemoveFromOutTxTrackerResponse extends Message<MsgRemoveFromOutTxTrackerResponse> {
  constructor(data?: PartialMessage<MsgRemoveFromOutTxTrackerResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgRemoveFromOutTxTrackerResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgRemoveFromOutTxTrackerResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgRemoveFromOutTxTrackerResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgRemoveFromOutTxTrackerResponse;

  static equals(a: MsgRemoveFromOutTxTrackerResponse | PlainMessage<MsgRemoveFromOutTxTrackerResponse> | undefined, b: MsgRemoveFromOutTxTrackerResponse | PlainMessage<MsgRemoveFromOutTxTrackerResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgGasPriceVoter
 */
export declare class MsgGasPriceVoter extends Message<MsgGasPriceVoter> {
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
   * @generated from field: uint64 block_number = 4;
   */
  blockNumber: bigint;

  /**
   * @generated from field: string supply = 5;
   */
  supply: string;

  constructor(data?: PartialMessage<MsgGasPriceVoter>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgGasPriceVoter";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgGasPriceVoter;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgGasPriceVoter;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgGasPriceVoter;

  static equals(a: MsgGasPriceVoter | PlainMessage<MsgGasPriceVoter> | undefined, b: MsgGasPriceVoter | PlainMessage<MsgGasPriceVoter> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgGasPriceVoterResponse
 */
export declare class MsgGasPriceVoterResponse extends Message<MsgGasPriceVoterResponse> {
  constructor(data?: PartialMessage<MsgGasPriceVoterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgGasPriceVoterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgGasPriceVoterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgGasPriceVoterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgGasPriceVoterResponse;

  static equals(a: MsgGasPriceVoterResponse | PlainMessage<MsgGasPriceVoterResponse> | undefined, b: MsgGasPriceVoterResponse | PlainMessage<MsgGasPriceVoterResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteOnObservedOutboundTx
 */
export declare class MsgVoteOnObservedOutboundTx extends Message<MsgVoteOnObservedOutboundTx> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string cctx_hash = 2;
   */
  cctxHash: string;

  /**
   * @generated from field: string observed_outTx_hash = 3;
   */
  observedOutTxHash: string;

  /**
   * @generated from field: uint64 observed_outTx_blockHeight = 4;
   */
  observedOutTxBlockHeight: bigint;

  /**
   * @generated from field: uint64 observed_outTx_gas_used = 10;
   */
  observedOutTxGasUsed: bigint;

  /**
   * @generated from field: string observed_outTx_effective_gas_price = 11;
   */
  observedOutTxEffectiveGasPrice: string;

  /**
   * @generated from field: uint64 observed_outTx_effective_gas_limit = 12;
   */
  observedOutTxEffectiveGasLimit: bigint;

  /**
   * @generated from field: string value_received = 5;
   */
  valueReceived: string;

  /**
   * @generated from field: common.ReceiveStatus status = 6;
   */
  status: ReceiveStatus;

  /**
   * @generated from field: int64 outTx_chain = 7;
   */
  outTxChain: bigint;

  /**
   * @generated from field: uint64 outTx_tss_nonce = 8;
   */
  outTxTssNonce: bigint;

  /**
   * @generated from field: common.CoinType coin_type = 9;
   */
  coinType: CoinType;

  constructor(data?: PartialMessage<MsgVoteOnObservedOutboundTx>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteOnObservedOutboundTx";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteOnObservedOutboundTx;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteOnObservedOutboundTx;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteOnObservedOutboundTx;

  static equals(a: MsgVoteOnObservedOutboundTx | PlainMessage<MsgVoteOnObservedOutboundTx> | undefined, b: MsgVoteOnObservedOutboundTx | PlainMessage<MsgVoteOnObservedOutboundTx> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteOnObservedOutboundTxResponse
 */
export declare class MsgVoteOnObservedOutboundTxResponse extends Message<MsgVoteOnObservedOutboundTxResponse> {
  constructor(data?: PartialMessage<MsgVoteOnObservedOutboundTxResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteOnObservedOutboundTxResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteOnObservedOutboundTxResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteOnObservedOutboundTxResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteOnObservedOutboundTxResponse;

  static equals(a: MsgVoteOnObservedOutboundTxResponse | PlainMessage<MsgVoteOnObservedOutboundTxResponse> | undefined, b: MsgVoteOnObservedOutboundTxResponse | PlainMessage<MsgVoteOnObservedOutboundTxResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteOnObservedInboundTx
 */
export declare class MsgVoteOnObservedInboundTx extends Message<MsgVoteOnObservedInboundTx> {
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
   * @generated from field: string in_tx_hash = 9;
   */
  inTxHash: string;

  /**
   * @generated from field: uint64 in_block_height = 10;
   */
  inBlockHeight: bigint;

  /**
   * @generated from field: uint64 gas_limit = 11;
   */
  gasLimit: bigint;

  /**
   * @generated from field: common.CoinType coin_type = 12;
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

  constructor(data?: PartialMessage<MsgVoteOnObservedInboundTx>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteOnObservedInboundTx";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteOnObservedInboundTx;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteOnObservedInboundTx;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteOnObservedInboundTx;

  static equals(a: MsgVoteOnObservedInboundTx | PlainMessage<MsgVoteOnObservedInboundTx> | undefined, b: MsgVoteOnObservedInboundTx | PlainMessage<MsgVoteOnObservedInboundTx> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.MsgVoteOnObservedInboundTxResponse
 */
export declare class MsgVoteOnObservedInboundTxResponse extends Message<MsgVoteOnObservedInboundTxResponse> {
  constructor(data?: PartialMessage<MsgVoteOnObservedInboundTxResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.MsgVoteOnObservedInboundTxResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MsgVoteOnObservedInboundTxResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MsgVoteOnObservedInboundTxResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MsgVoteOnObservedInboundTxResponse;

  static equals(a: MsgVoteOnObservedInboundTxResponse | PlainMessage<MsgVoteOnObservedInboundTxResponse> | undefined, b: MsgVoteOnObservedInboundTxResponse | PlainMessage<MsgVoteOnObservedInboundTxResponse> | undefined): boolean;
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

