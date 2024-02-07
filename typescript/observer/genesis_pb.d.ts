// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file observer/genesis.proto (package zetachain.zetacore.observer, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Ballot } from "./ballot_pb.js";
import type { LastObserverCount, ObserverSet } from "./observer_pb.js";
import type { NodeAccount } from "./node_account_pb.js";
import type { CrosschainFlags } from "./crosschain_flags_pb.js";
import type { ChainParamsList, Params } from "./params_pb.js";
import type { Keygen } from "./keygen_pb.js";
import type { TSS } from "./tss_pb.js";
import type { TssFundMigratorInfo } from "./tss_funds_migrator_pb.js";
import type { Blame } from "./blame_pb.js";
import type { PendingNonces } from "./pending_nonces_pb.js";
import type { ChainNonces } from "./chain_nonces_pb.js";
import type { NonceToCctx } from "./nonce_to_cctx_pb.js";

/**
 * @generated from message zetachain.zetacore.observer.GenesisState
 */
export declare class GenesisState extends Message<GenesisState> {
  /**
   * @generated from field: repeated zetachain.zetacore.observer.Ballot ballots = 1;
   */
  ballots: Ballot[];

  /**
   * @generated from field: zetachain.zetacore.observer.ObserverSet observers = 2;
   */
  observers?: ObserverSet;

  /**
   * @generated from field: repeated zetachain.zetacore.observer.NodeAccount nodeAccountList = 3;
   */
  nodeAccountList: NodeAccount[];

  /**
   * @generated from field: zetachain.zetacore.observer.CrosschainFlags crosschain_flags = 4;
   */
  crosschainFlags?: CrosschainFlags;

  /**
   * @generated from field: zetachain.zetacore.observer.Params params = 5;
   */
  params?: Params;

  /**
   * @generated from field: zetachain.zetacore.observer.Keygen keygen = 6;
   */
  keygen?: Keygen;

  /**
   * @generated from field: zetachain.zetacore.observer.LastObserverCount last_observer_count = 7;
   */
  lastObserverCount?: LastObserverCount;

  /**
   * @generated from field: zetachain.zetacore.observer.ChainParamsList chain_params_list = 8;
   */
  chainParamsList?: ChainParamsList;

  /**
   * @generated from field: zetachain.zetacore.observer.TSS tss = 9;
   */
  tss?: TSS;

  /**
   * @generated from field: repeated zetachain.zetacore.observer.TSS tss_history = 10;
   */
  tssHistory: TSS[];

  /**
   * @generated from field: repeated zetachain.zetacore.observer.TssFundMigratorInfo tss_fund_migrators = 11;
   */
  tssFundMigrators: TssFundMigratorInfo[];

  /**
   * @generated from field: repeated zetachain.zetacore.observer.Blame blame_list = 12;
   */
  blameList: Blame[];

  /**
   * @generated from field: repeated zetachain.zetacore.observer.PendingNonces pending_nonces = 13;
   */
  pendingNonces: PendingNonces[];

  /**
   * @generated from field: repeated zetachain.zetacore.observer.ChainNonces chain_nonces = 14;
   */
  chainNonces: ChainNonces[];

  /**
   * @generated from field: repeated zetachain.zetacore.observer.NonceToCctx nonce_to_cctx = 15;
   */
  nonceToCctx: NonceToCctx[];

  constructor(data?: PartialMessage<GenesisState>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.GenesisState";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GenesisState;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GenesisState;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GenesisState;

  static equals(a: GenesisState | PlainMessage<GenesisState> | undefined, b: GenesisState | PlainMessage<GenesisState> | undefined): boolean;
}

