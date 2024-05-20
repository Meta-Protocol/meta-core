// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file zetachain/zetacore/observer/crosschain_flags.proto (package zetachain.zetacore.observer, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message zetachain.zetacore.observer.GasPriceIncreaseFlags
 */
export declare class GasPriceIncreaseFlags extends Message<GasPriceIncreaseFlags> {
  /**
   * @generated from field: int64 epochLength = 1;
   */
  epochLength: bigint;

  /**
   * @generated from field: google.protobuf.Duration retryInterval = 2;
   */
  retryInterval?: Duration;

  /**
   * @generated from field: uint32 gasPriceIncreasePercent = 3;
   */
  gasPriceIncreasePercent: number;

  /**
   * Maximum gas price increase in percent of the median gas price
   * Default is used if 0
   *
   * @generated from field: uint32 gasPriceIncreaseMax = 4;
   */
  gasPriceIncreaseMax: number;

  /**
   * Maximum number of pending crosschain transactions to check for gas price
   * increase
   *
   * @generated from field: uint32 maxPendingCctxs = 5;
   */
  maxPendingCctxs: number;

  constructor(data?: PartialMessage<GasPriceIncreaseFlags>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.GasPriceIncreaseFlags";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GasPriceIncreaseFlags;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GasPriceIncreaseFlags;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GasPriceIncreaseFlags;

  static equals(a: GasPriceIncreaseFlags | PlainMessage<GasPriceIncreaseFlags> | undefined, b: GasPriceIncreaseFlags | PlainMessage<GasPriceIncreaseFlags> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.CrosschainFlags
 */
export declare class CrosschainFlags extends Message<CrosschainFlags> {
  /**
   * @generated from field: bool isInboundEnabled = 1;
   */
  isInboundEnabled: boolean;

  /**
   * @generated from field: bool isOutboundEnabled = 2;
   */
  isOutboundEnabled: boolean;

  /**
   * @generated from field: zetachain.zetacore.observer.GasPriceIncreaseFlags gasPriceIncreaseFlags = 3;
   */
  gasPriceIncreaseFlags?: GasPriceIncreaseFlags;

  constructor(data?: PartialMessage<CrosschainFlags>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.CrosschainFlags";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CrosschainFlags;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CrosschainFlags;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CrosschainFlags;

  static equals(a: CrosschainFlags | PlainMessage<CrosschainFlags> | undefined, b: CrosschainFlags | PlainMessage<CrosschainFlags> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.observer.LegacyCrosschainFlags
 */
export declare class LegacyCrosschainFlags extends Message<LegacyCrosschainFlags> {
  /**
   * @generated from field: bool isInboundEnabled = 1;
   */
  isInboundEnabled: boolean;

  /**
   * @generated from field: bool isOutboundEnabled = 2;
   */
  isOutboundEnabled: boolean;

  /**
   * @generated from field: zetachain.zetacore.observer.GasPriceIncreaseFlags gasPriceIncreaseFlags = 3;
   */
  gasPriceIncreaseFlags?: GasPriceIncreaseFlags;

  constructor(data?: PartialMessage<LegacyCrosschainFlags>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.LegacyCrosschainFlags";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LegacyCrosschainFlags;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LegacyCrosschainFlags;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LegacyCrosschainFlags;

  static equals(a: LegacyCrosschainFlags | PlainMessage<LegacyCrosschainFlags> | undefined, b: LegacyCrosschainFlags | PlainMessage<LegacyCrosschainFlags> | undefined): boolean;
}

