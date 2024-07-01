// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file zetachain/zetacore/pkg/chains/chains.proto (package zetachain.zetacore.pkg.chains, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * ReceiveStatus represents the status of an outbound
 * TODO: Rename and move
 * https://github.com/zeta-chain/node/issues/2257
 *
 * @generated from enum zetachain.zetacore.pkg.chains.ReceiveStatus
 */
export declare enum ReceiveStatus {
  /**
   * Created is used for inbounds
   *
   * @generated from enum value: created = 0;
   */
  created = 0,

  /**
   * @generated from enum value: success = 1;
   */
  success = 1,

  /**
   * @generated from enum value: failed = 2;
   */
  failed = 2,
}

/**
 * ChainName represents the name of the chain
 *
 * @generated from enum zetachain.zetacore.pkg.chains.ChainName
 */
export declare enum ChainName {
  /**
   * @generated from enum value: empty = 0;
   */
  empty = 0,

  /**
   * @generated from enum value: eth_mainnet = 1;
   */
  eth_mainnet = 1,

  /**
   * @generated from enum value: zeta_mainnet = 2;
   */
  zeta_mainnet = 2,

  /**
   * @generated from enum value: btc_mainnet = 3;
   */
  btc_mainnet = 3,

  /**
   * @generated from enum value: polygon_mainnet = 4;
   */
  polygon_mainnet = 4,

  /**
   * @generated from enum value: bsc_mainnet = 5;
   */
  bsc_mainnet = 5,

  /**
   * @generated from enum value: goerli_testnet = 6;
   */
  goerli_testnet = 6,

  /**
   * @generated from enum value: mumbai_testnet = 7;
   */
  mumbai_testnet = 7,

  /**
   * @generated from enum value: bsc_testnet = 10;
   */
  bsc_testnet = 10,

  /**
   * @generated from enum value: zeta_testnet = 11;
   */
  zeta_testnet = 11,

  /**
   * @generated from enum value: btc_testnet = 12;
   */
  btc_testnet = 12,

  /**
   * @generated from enum value: sepolia_testnet = 13;
   */
  sepolia_testnet = 13,

  /**
   * @generated from enum value: goerli_localnet = 14;
   */
  goerli_localnet = 14,

  /**
   * @generated from enum value: btc_regtest = 15;
   */
  btc_regtest = 15,

  /**
   * @generated from enum value: amoy_testnet = 16;
   */
  amoy_testnet = 16,

  /**
   * @generated from enum value: optimism_mainnet = 17;
   */
  optimism_mainnet = 17,

  /**
   * @generated from enum value: optimism_sepolia = 18;
   */
  optimism_sepolia = 18,

  /**
   * @generated from enum value: base_mainnet = 19;
   */
  base_mainnet = 19,

  /**
   * @generated from enum value: base_sepolia = 20;
   */
  base_sepolia = 20,

  /**
   * @generated from enum value: solana_mainnet = 21;
   */
  solana_mainnet = 21,

  /**
   * @generated from enum value: solana_testnet = 22;
   */
  solana_testnet = 22,

  /**
   * @generated from enum value: solana_localnet = 23;
   */
  solana_localnet = 23,
}

/**
 * Network represents the network of the chain
 * there is a single instance of the network on mainnet
 * then the network can have eventual testnets or devnets
 *
 * @generated from enum zetachain.zetacore.pkg.chains.Network
 */
export declare enum Network {
  /**
   * @generated from enum value: eth = 0;
   */
  eth = 0,

  /**
   * @generated from enum value: zeta = 1;
   */
  zeta = 1,

  /**
   * @generated from enum value: btc = 2;
   */
  btc = 2,

  /**
   * @generated from enum value: polygon = 3;
   */
  polygon = 3,

  /**
   * @generated from enum value: bsc = 4;
   */
  bsc = 4,

  /**
   * @generated from enum value: optimism = 5;
   */
  optimism = 5,

  /**
   * @generated from enum value: base = 6;
   */
  base = 6,

  /**
   * @generated from enum value: solana = 7;
   */
  solana = 7,
}

/**
 * NetworkType represents the network type of the chain
 * Mainnet, Testnet, Privnet, Devnet
 *
 * @generated from enum zetachain.zetacore.pkg.chains.NetworkType
 */
export declare enum NetworkType {
  /**
   * @generated from enum value: mainnet = 0;
   */
  mainnet = 0,

  /**
   * @generated from enum value: testnet = 1;
   */
  testnet = 1,

  /**
   * @generated from enum value: privnet = 2;
   */
  privnet = 2,

  /**
   * @generated from enum value: devnet = 3;
   */
  devnet = 3,
}

/**
 * Vm represents the virtual machine type of the chain to support smart
 * contracts
 *
 * @generated from enum zetachain.zetacore.pkg.chains.Vm
 */
export declare enum Vm {
  /**
   * @generated from enum value: no_vm = 0;
   */
  no_vm = 0,

  /**
   * @generated from enum value: evm = 1;
   */
  evm = 1,

  /**
   * @generated from enum value: svm = 2;
   */
  svm = 2,
}

/**
 * Consensus represents the consensus algorithm used by the chain
 * this can represent the consensus of a L1
 * this can also represent the solution of a L2
 *
 * @generated from enum zetachain.zetacore.pkg.chains.Consensus
 */
export declare enum Consensus {
  /**
   * @generated from enum value: ethereum = 0;
   */
  ethereum = 0,

  /**
   * @generated from enum value: tendermint = 1;
   */
  tendermint = 1,

  /**
   * @generated from enum value: bitcoin = 2;
   */
  bitcoin = 2,

  /**
   * @generated from enum value: op_stack = 3;
   */
  op_stack = 3,

  /**
   * @generated from enum value: solana_consensus = 4;
   */
  solana_consensus = 4,
}

/**
 * CCTXGateway describes for the chain the gateway used to handle CCTX outbounds
 *
 * @generated from enum zetachain.zetacore.pkg.chains.CCTXGateway
 */
export declare enum CCTXGateway {
  /**
   * zevm is the internal CCTX gateway to process outbound on the ZEVM and read
   * inbound events from the ZEVM only used for ZetaChain chains
   *
   * @generated from enum value: zevm = 0;
   */
  zevm = 0,

  /**
   * observers is the CCTX gateway for chains relying on the observer set to
   * observe inbounds and TSS for outbounds
   *
   * @generated from enum value: observers = 1;
   */
  observers = 1,
}

/**
 * Chain represents static data about a blockchain network
 * it is identified by a unique chain ID
 *
 * @generated from message zetachain.zetacore.pkg.chains.Chain
 */
export declare class Chain extends Message<Chain> {
  /**
   * ChainId is the unique identifier of the chain
   *
   * @generated from field: int64 chain_id = 2;
   */
  chainId: bigint;

  /**
   * ChainName is the name of the chain
   *
   * @generated from field: zetachain.zetacore.pkg.chains.ChainName chain_name = 1;
   */
  chainName: ChainName;

  /**
   * Network is the network of the chain
   *
   * @generated from field: zetachain.zetacore.pkg.chains.Network network = 3;
   */
  network: Network;

  /**
   * NetworkType is the network type of the chain: mainnet, testnet, etc..
   *
   * @generated from field: zetachain.zetacore.pkg.chains.NetworkType network_type = 4;
   */
  networkType: NetworkType;

  /**
   * Vm is the virtual machine used in the chain
   *
   * @generated from field: zetachain.zetacore.pkg.chains.Vm vm = 5;
   */
  vm: Vm;

  /**
   * Consensus is the underlying consensus algorithm used by the chain
   *
   * @generated from field: zetachain.zetacore.pkg.chains.Consensus consensus = 6;
   */
  consensus: Consensus;

  /**
   * IsExternal describe if the chain is ZetaChain or external
   *
   * @generated from field: bool is_external = 7;
   */
  isExternal: boolean;

  /**
   * CCTXGateway is the gateway used to handle CCTX outbounds
   *
   * @generated from field: zetachain.zetacore.pkg.chains.CCTXGateway cctx_gateway = 8;
   */
  cctxGateway: CCTXGateway;

  constructor(data?: PartialMessage<Chain>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.pkg.chains.Chain";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Chain;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Chain;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Chain;

  static equals(a: Chain | PlainMessage<Chain> | undefined, b: Chain | PlainMessage<Chain> | undefined): boolean;
}

