// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file zetachain/zetacore/crosschain/cross_chain_tx.proto (package zetachain.zetacore.crosschain, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { CoinType } from "../pkg/coin/coin_pb.js";

/**
 * @generated from enum zetachain.zetacore.crosschain.CctxStatus
 */
export declare enum CctxStatus {
  /**
   * some observer sees inbound tx
   *
   * @generated from enum value: PendingInbound = 0;
   */
  PendingInbound = 0,

  /**
   * super majority observer see inbound tx
   *
   * @generated from enum value: PendingOutbound = 1;
   */
  PendingOutbound = 1,

  /**
   * the corresponding outbound tx is mined
   *
   * @generated from enum value: OutboundMined = 3;
   */
  OutboundMined = 3,

  /**
   * outbound cannot succeed; should revert inbound
   *
   * @generated from enum value: PendingRevert = 4;
   */
  PendingRevert = 4,

  /**
   * inbound reverted.
   *
   * @generated from enum value: Reverted = 5;
   */
  Reverted = 5,

  /**
   * inbound tx error or invalid paramters and cannot revert; just abort.
   *
   * @generated from enum value: Aborted = 6;
   */
  Aborted = 6,
}

/**
 * @generated from enum zetachain.zetacore.crosschain.TxFinalizationStatus
 */
export declare enum TxFinalizationStatus {
  /**
   * the corresponding tx is not finalized
   *
   * @generated from enum value: NotFinalized = 0;
   */
  NotFinalized = 0,

  /**
   * the corresponding tx is finalized but not executed yet
   *
   * @generated from enum value: Finalized = 1;
   */
  Finalized = 1,

  /**
   * the corresponding tx is executed
   *
   * @generated from enum value: Executed = 2;
   */
  Executed = 2,
}

/**
 * @generated from enum zetachain.zetacore.crosschain.ConfirmationMode
 */
export declare enum ConfirmationMode {
  /**
   * an inbound/outbound is confirmed using safe confirmation count
   *
   * @generated from enum value: SAFE = 0;
   */
  SAFE = 0,

  /**
   * an inbound/outbound is confirmed using fast confirmation count
   *
   * @generated from enum value: FAST = 1;
   */
  FAST = 1,
}

/**
 * InboundStatus represents the status of an observed inbound
 *
 * @generated from enum zetachain.zetacore.crosschain.InboundStatus
 */
export declare enum InboundStatus {
  /**
   * @generated from enum value: SUCCESS = 0;
   */
  SUCCESS = 0,

  /**
   * this field is specifically for Bitcoin when the deposit amount is less than
   * depositor fee
   *
   * @generated from enum value: INSUFFICIENT_DEPOSITOR_FEE = 1;
   */
  INSUFFICIENT_DEPOSITOR_FEE = 1,

  /**
   * the receiver address parsed from the inbound is invalid
   *
   * @generated from enum value: INVALID_RECEIVER_ADDRESS = 2;
   */
  INVALID_RECEIVER_ADDRESS = 2,

  /**
   * parse memo is invalid
   *
   * @generated from enum value: INVALID_MEMO = 3;
   */
  INVALID_MEMO = 3,
}

/**
 * ProtocolContractVersion represents the version of the protocol contract used
 * for cctx workflow
 *
 * @generated from enum zetachain.zetacore.crosschain.ProtocolContractVersion
 */
export declare enum ProtocolContractVersion {
  /**
   * @generated from enum value: V1 = 0;
   */
  V1 = 0,

  /**
   * @generated from enum value: V2 = 1;
   */
  V2 = 1,
}

/**
 * @generated from message zetachain.zetacore.crosschain.InboundParams
 */
export declare class InboundParams extends Message<InboundParams> {
  /**
   * this address is the immediate contract/EOA that calls
   *
   * @generated from field: string sender = 1;
   */
  sender: string;

  /**
   * the Connector.send()
   *
   * @generated from field: int64 sender_chain_id = 2;
   */
  senderChainId: bigint;

  /**
   * this address is the EOA that signs the inbound tx
   *
   * @generated from field: string tx_origin = 3;
   */
  txOrigin: string;

  /**
   * @generated from field: zetachain.zetacore.pkg.coin.CoinType coin_type = 4;
   */
  coinType: CoinType;

  /**
   * for ERC20 coin type, the asset is an address of the ERC20 contract
   *
   * @generated from field: string asset = 5;
   */
  asset: string;

  /**
   * @generated from field: string amount = 6;
   */
  amount: string;

  /**
   * @generated from field: string observed_hash = 7;
   */
  observedHash: string;

  /**
   * @generated from field: uint64 observed_external_height = 8;
   */
  observedExternalHeight: bigint;

  /**
   * @generated from field: string ballot_index = 9;
   */
  ballotIndex: string;

  /**
   * @generated from field: uint64 finalized_zeta_height = 10;
   */
  finalizedZetaHeight: bigint;

  /**
   * @generated from field: zetachain.zetacore.crosschain.TxFinalizationStatus tx_finalization_status = 11;
   */
  txFinalizationStatus: TxFinalizationStatus;

  /**
   * this field describes if a smart contract call should be made for a inbound
   * with assets only used for protocol contract version 2
   *
   * @generated from field: bool is_cross_chain_call = 12;
   */
  isCrossChainCall: boolean;

  /**
   * status of the inbound observation
   *
   * @generated from field: zetachain.zetacore.crosschain.InboundStatus status = 20;
   */
  status: InboundStatus;

  /**
   * confirmation mode used for the inbound
   *
   * @generated from field: zetachain.zetacore.crosschain.ConfirmationMode confirmation_mode = 21;
   */
  confirmationMode: ConfirmationMode;

  constructor(data?: PartialMessage<InboundParams>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.InboundParams";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InboundParams;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InboundParams;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InboundParams;

  static equals(a: InboundParams | PlainMessage<InboundParams> | undefined, b: InboundParams | PlainMessage<InboundParams> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.ZetaAccounting
 */
export declare class ZetaAccounting extends Message<ZetaAccounting> {
  /**
   * aborted_zeta_amount stores the total aborted amount for cctx of coin-type
   * ZETA
   *
   * @generated from field: string aborted_zeta_amount = 1;
   */
  abortedZetaAmount: string;

  constructor(data?: PartialMessage<ZetaAccounting>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.ZetaAccounting";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ZetaAccounting;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ZetaAccounting;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ZetaAccounting;

  static equals(a: ZetaAccounting | PlainMessage<ZetaAccounting> | undefined, b: ZetaAccounting | PlainMessage<ZetaAccounting> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.CallOptions
 */
export declare class CallOptions extends Message<CallOptions> {
  /**
   * @generated from field: uint64 gas_limit = 1;
   */
  gasLimit: bigint;

  /**
   * @generated from field: bool is_arbitrary_call = 2;
   */
  isArbitraryCall: boolean;

  constructor(data?: PartialMessage<CallOptions>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.CallOptions";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CallOptions;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CallOptions;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CallOptions;

  static equals(a: CallOptions | PlainMessage<CallOptions> | undefined, b: CallOptions | PlainMessage<CallOptions> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.OutboundParams
 */
export declare class OutboundParams extends Message<OutboundParams> {
  /**
   * @generated from field: string receiver = 1;
   */
  receiver: string;

  /**
   * @generated from field: int64 receiver_chainId = 2;
   */
  receiverChainId: bigint;

  /**
   * @generated from field: zetachain.zetacore.pkg.coin.CoinType coin_type = 3;
   */
  coinType: CoinType;

  /**
   * @generated from field: string amount = 4;
   */
  amount: string;

  /**
   * @generated from field: uint64 tss_nonce = 5;
   */
  tssNonce: bigint;

  /**
   * Deprecated (v21), use CallOptions
   *
   * @generated from field: uint64 gas_limit = 6;
   */
  gasLimit: bigint;

  /**
   * @generated from field: string gas_price = 7;
   */
  gasPrice: string;

  /**
   * @generated from field: string gas_priority_fee = 23;
   */
  gasPriorityFee: string;

  /**
   * the above are commands for zetaclients
   * the following fields are used when the outbound tx is mined
   *
   * @generated from field: string hash = 8;
   */
  hash: string;

  /**
   * @generated from field: string ballot_index = 9;
   */
  ballotIndex: string;

  /**
   * @generated from field: uint64 observed_external_height = 10;
   */
  observedExternalHeight: bigint;

  /**
   * @generated from field: uint64 gas_used = 20;
   */
  gasUsed: bigint;

  /**
   * @generated from field: string effective_gas_price = 21;
   */
  effectiveGasPrice: string;

  /**
   * @generated from field: uint64 effective_gas_limit = 22;
   */
  effectiveGasLimit: bigint;

  /**
   * @generated from field: string tss_pubkey = 11;
   */
  tssPubkey: string;

  /**
   * @generated from field: zetachain.zetacore.crosschain.TxFinalizationStatus tx_finalization_status = 12;
   */
  txFinalizationStatus: TxFinalizationStatus;

  /**
   * @generated from field: zetachain.zetacore.crosschain.CallOptions call_options = 24;
   */
  callOptions?: CallOptions;

  /**
   * confirmation mode used for the outbound
   *
   * @generated from field: zetachain.zetacore.crosschain.ConfirmationMode confirmation_mode = 25;
   */
  confirmationMode: ConfirmationMode;

  constructor(data?: PartialMessage<OutboundParams>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.OutboundParams";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OutboundParams;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OutboundParams;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OutboundParams;

  static equals(a: OutboundParams | PlainMessage<OutboundParams> | undefined, b: OutboundParams | PlainMessage<OutboundParams> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.Status
 */
export declare class Status extends Message<Status> {
  /**
   * @generated from field: zetachain.zetacore.crosschain.CctxStatus status = 1;
   */
  status: CctxStatus;

  /**
   * status_message carries information about the status transitions:
   * why they were triggered, old and new status.
   *
   * @generated from field: string status_message = 2;
   */
  statusMessage: string;

  /**
   * error_message carries information about the error that caused the tx
   * to be PendingRevert, Reverted or Aborted.
   *
   * @generated from field: string error_message = 6;
   */
  errorMessage: string;

  /**
   * @generated from field: int64 lastUpdate_timestamp = 3;
   */
  lastUpdateTimestamp: bigint;

  /**
   * @generated from field: bool isAbortRefunded = 4;
   */
  isAbortRefunded: boolean;

  /**
   * when the CCTX was created. only populated on new transactions.
   *
   * @generated from field: int64 created_timestamp = 5;
   */
  createdTimestamp: bigint;

  /**
   * error_message_revert carries information about the revert outbound tx ,
   * which is created if the first outbound tx fails
   *
   * @generated from field: string error_message_revert = 7;
   */
  errorMessageRevert: string;

  /**
   * error_message_abort carries information when aborting the CCTX fails
   *
   * @generated from field: string error_message_abort = 8;
   */
  errorMessageAbort: string;

  constructor(data?: PartialMessage<Status>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.Status";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Status;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Status;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Status;

  static equals(a: Status | PlainMessage<Status> | undefined, b: Status | PlainMessage<Status> | undefined): boolean;
}

/**
 * RevertOptions represents the options for reverting a cctx
 *
 * @generated from message zetachain.zetacore.crosschain.RevertOptions
 */
export declare class RevertOptions extends Message<RevertOptions> {
  /**
   * @generated from field: string revert_address = 1;
   */
  revertAddress: string;

  /**
   * @generated from field: bool call_on_revert = 2;
   */
  callOnRevert: boolean;

  /**
   * @generated from field: string abort_address = 3;
   */
  abortAddress: string;

  /**
   * @generated from field: bytes revert_message = 4;
   */
  revertMessage: Uint8Array;

  /**
   * @generated from field: string revert_gas_limit = 5;
   */
  revertGasLimit: string;

  constructor(data?: PartialMessage<RevertOptions>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.RevertOptions";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RevertOptions;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RevertOptions;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RevertOptions;

  static equals(a: RevertOptions | PlainMessage<RevertOptions> | undefined, b: RevertOptions | PlainMessage<RevertOptions> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.CrossChainTx
 */
export declare class CrossChainTx extends Message<CrossChainTx> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string index = 2;
   */
  index: string;

  /**
   * @generated from field: string zeta_fees = 5;
   */
  zetaFees: string;

  /**
   * Not used by protocol , just relayed across
   *
   * @generated from field: string relayed_message = 6;
   */
  relayedMessage: string;

  /**
   * @generated from field: zetachain.zetacore.crosschain.Status cctx_status = 8;
   */
  cctxStatus?: Status;

  /**
   * @generated from field: zetachain.zetacore.crosschain.InboundParams inbound_params = 9;
   */
  inboundParams?: InboundParams;

  /**
   * @generated from field: repeated zetachain.zetacore.crosschain.OutboundParams outbound_params = 10;
   */
  outboundParams: OutboundParams[];

  /**
   * @generated from field: zetachain.zetacore.crosschain.ProtocolContractVersion protocol_contract_version = 11;
   */
  protocolContractVersion: ProtocolContractVersion;

  /**
   * @generated from field: zetachain.zetacore.crosschain.RevertOptions revert_options = 12;
   */
  revertOptions?: RevertOptions;

  constructor(data?: PartialMessage<CrossChainTx>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.CrossChainTx";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CrossChainTx;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CrossChainTx;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CrossChainTx;

  static equals(a: CrossChainTx | PlainMessage<CrossChainTx> | undefined, b: CrossChainTx | PlainMessage<CrossChainTx> | undefined): boolean;
}

