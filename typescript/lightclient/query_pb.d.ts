// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file lightclient/query.proto (package zetachain.zetacore.lightclient, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination_pb.js";
import type { BlockHeader, Proof } from "../pkg/proofs/proofs_pb.js";
import type { ChainState } from "./chain_state_pb.js";

/**
 * @generated from message zetachain.zetacore.lightclient.QueryAllBlockHeaderRequest
 */
export declare class QueryAllBlockHeaderRequest extends Message<QueryAllBlockHeaderRequest> {
  /**
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 1;
   */
  pagination?: PageRequest;

  constructor(data?: PartialMessage<QueryAllBlockHeaderRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryAllBlockHeaderRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAllBlockHeaderRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAllBlockHeaderRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAllBlockHeaderRequest;

  static equals(a: QueryAllBlockHeaderRequest | PlainMessage<QueryAllBlockHeaderRequest> | undefined, b: QueryAllBlockHeaderRequest | PlainMessage<QueryAllBlockHeaderRequest> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.lightclient.QueryAllBlockHeaderResponse
 */
export declare class QueryAllBlockHeaderResponse extends Message<QueryAllBlockHeaderResponse> {
  /**
   * @generated from field: repeated proofs.BlockHeader block_headers = 1;
   */
  blockHeaders: BlockHeader[];

  /**
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;

  constructor(data?: PartialMessage<QueryAllBlockHeaderResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryAllBlockHeaderResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAllBlockHeaderResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAllBlockHeaderResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAllBlockHeaderResponse;

  static equals(a: QueryAllBlockHeaderResponse | PlainMessage<QueryAllBlockHeaderResponse> | undefined, b: QueryAllBlockHeaderResponse | PlainMessage<QueryAllBlockHeaderResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.lightclient.QueryGetBlockHeaderRequest
 */
export declare class QueryGetBlockHeaderRequest extends Message<QueryGetBlockHeaderRequest> {
  /**
   * @generated from field: bytes block_hash = 1;
   */
  blockHash: Uint8Array;

  constructor(data?: PartialMessage<QueryGetBlockHeaderRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryGetBlockHeaderRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetBlockHeaderRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetBlockHeaderRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetBlockHeaderRequest;

  static equals(a: QueryGetBlockHeaderRequest | PlainMessage<QueryGetBlockHeaderRequest> | undefined, b: QueryGetBlockHeaderRequest | PlainMessage<QueryGetBlockHeaderRequest> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.lightclient.QueryGetBlockHeaderResponse
 */
export declare class QueryGetBlockHeaderResponse extends Message<QueryGetBlockHeaderResponse> {
  /**
   * @generated from field: proofs.BlockHeader block_header = 1;
   */
  blockHeader?: BlockHeader;

  constructor(data?: PartialMessage<QueryGetBlockHeaderResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryGetBlockHeaderResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetBlockHeaderResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetBlockHeaderResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetBlockHeaderResponse;

  static equals(a: QueryGetBlockHeaderResponse | PlainMessage<QueryGetBlockHeaderResponse> | undefined, b: QueryGetBlockHeaderResponse | PlainMessage<QueryGetBlockHeaderResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.lightclient.QueryAllChainStateRequest
 */
export declare class QueryAllChainStateRequest extends Message<QueryAllChainStateRequest> {
  /**
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 1;
   */
  pagination?: PageRequest;

  constructor(data?: PartialMessage<QueryAllChainStateRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryAllChainStateRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAllChainStateRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAllChainStateRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAllChainStateRequest;

  static equals(a: QueryAllChainStateRequest | PlainMessage<QueryAllChainStateRequest> | undefined, b: QueryAllChainStateRequest | PlainMessage<QueryAllChainStateRequest> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.lightclient.QueryAllChainStateResponse
 */
export declare class QueryAllChainStateResponse extends Message<QueryAllChainStateResponse> {
  /**
   * @generated from field: repeated zetachain.zetacore.lightclient.ChainState chain_state = 1;
   */
  chainState: ChainState[];

  /**
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;

  constructor(data?: PartialMessage<QueryAllChainStateResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryAllChainStateResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryAllChainStateResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryAllChainStateResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryAllChainStateResponse;

  static equals(a: QueryAllChainStateResponse | PlainMessage<QueryAllChainStateResponse> | undefined, b: QueryAllChainStateResponse | PlainMessage<QueryAllChainStateResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.lightclient.QueryGetChainStateRequest
 */
export declare class QueryGetChainStateRequest extends Message<QueryGetChainStateRequest> {
  /**
   * @generated from field: int64 chain_id = 1;
   */
  chainId: bigint;

  constructor(data?: PartialMessage<QueryGetChainStateRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryGetChainStateRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetChainStateRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetChainStateRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetChainStateRequest;

  static equals(a: QueryGetChainStateRequest | PlainMessage<QueryGetChainStateRequest> | undefined, b: QueryGetChainStateRequest | PlainMessage<QueryGetChainStateRequest> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.lightclient.QueryGetChainStateResponse
 */
export declare class QueryGetChainStateResponse extends Message<QueryGetChainStateResponse> {
  /**
   * @generated from field: zetachain.zetacore.lightclient.ChainState chain_state = 1;
   */
  chainState?: ChainState;

  constructor(data?: PartialMessage<QueryGetChainStateResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryGetChainStateResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryGetChainStateResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryGetChainStateResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryGetChainStateResponse;

  static equals(a: QueryGetChainStateResponse | PlainMessage<QueryGetChainStateResponse> | undefined, b: QueryGetChainStateResponse | PlainMessage<QueryGetChainStateResponse> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.lightclient.QueryProveRequest
 */
export declare class QueryProveRequest extends Message<QueryProveRequest> {
  /**
   * @generated from field: int64 chain_id = 1;
   */
  chainId: bigint;

  /**
   * @generated from field: string tx_hash = 2;
   */
  txHash: string;

  /**
   * @generated from field: proofs.Proof proof = 3;
   */
  proof?: Proof;

  /**
   * @generated from field: string block_hash = 4;
   */
  blockHash: string;

  /**
   * @generated from field: int64 tx_index = 5;
   */
  txIndex: bigint;

  constructor(data?: PartialMessage<QueryProveRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryProveRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryProveRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryProveRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryProveRequest;

  static equals(a: QueryProveRequest | PlainMessage<QueryProveRequest> | undefined, b: QueryProveRequest | PlainMessage<QueryProveRequest> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.lightclient.QueryProveResponse
 */
export declare class QueryProveResponse extends Message<QueryProveResponse> {
  /**
   * @generated from field: bool valid = 1;
   */
  valid: boolean;

  constructor(data?: PartialMessage<QueryProveResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.lightclient.QueryProveResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryProveResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryProveResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryProveResponse;

  static equals(a: QueryProveResponse | PlainMessage<QueryProveResponse> | undefined, b: QueryProveResponse | PlainMessage<QueryProveResponse> | undefined): boolean;
}

