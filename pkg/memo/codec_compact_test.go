package memo_test

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/node/pkg/memo"
	"github.com/zeta-chain/node/testutil/sample"
)

func Test_NewCodecCompact(t *testing.T) {
	tests := []struct {
		name        string
		encodingFmt uint8
		fail        bool
	}{
		{
			name:        "create codec compact successfully",
			encodingFmt: memo.EncodingFmtCompactShort,
		},
		{
			name:        "create codec compact failed on invalid encoding format",
			encodingFmt: 0b11,
			fail:        true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			codec, err := memo.NewCodecCompact(tc.encodingFmt)
			if tc.fail {
				require.Error(t, err)
				require.Nil(t, codec)
			} else {
				require.NoError(t, err)
				require.NotNil(t, codec)
			}
		})
	}
}

func Test_CodecCompact_AddArguments(t *testing.T) {
	codec, err := memo.NewCodecCompact(memo.EncodingFmtCompactLong)
	require.NoError(t, err)
	require.NotNil(t, codec)

	address := sample.EthAddress()
	codec.AddArguments(memo.ArgReceiver(address))

	// attempt to pack the arguments, result should not be nil
	packedData, err := codec.PackArguments()
	require.NoError(t, err)
	require.True(t, len(packedData) == common.AddressLength)
}

func Test_CodecCompact_PackArguments(t *testing.T) {
	// create sample arguments
	argAddress := sample.EthAddress()
	argBytes := []byte("here is a bytes argument")
	argString := "some other string argument"

	// test cases
	tests := []struct {
		name        string
		encodingFmt uint8
		args        []memo.CodecArg
		expectedLen int
		errMsg      string
	}{
		{
			name:        "pack arguments of [address, bytes, string] in compact-short format",
			encodingFmt: memo.EncodingFmtCompactShort,
			args: []memo.CodecArg{
				memo.ArgReceiver(argAddress),
				memo.ArgPayload(argBytes),
				memo.ArgRevertAddress(argString),
			},
			expectedLen: 20 + 1 + len(argBytes) + 1 + len([]byte(argString)),
		},
		{
			name:        "pack arguments of [string, address, bytes] in compact-long format",
			encodingFmt: memo.EncodingFmtCompactLong,
			args: []memo.CodecArg{
				memo.ArgRevertAddress(argString),
				memo.ArgReceiver(argAddress),
				memo.ArgPayload(argBytes),
			},
			expectedLen: 2 + len([]byte(argString)) + 20 + 2 + len(argBytes),
		},
		{
			name:        "pack long string (> 255 bytes) with compact-long format",
			encodingFmt: memo.EncodingFmtCompactLong,
			args: []memo.CodecArg{
				memo.ArgPayload([]byte(sample.StringRandom(sample.Rand(), 256))),
			},
			expectedLen: 2 + 256,
		},
		{
			name:        "pack long string (> 255 bytes) with compact-short format should fail",
			encodingFmt: memo.EncodingFmtCompactShort,
			args: []memo.CodecArg{
				memo.ArgPayload([]byte(sample.StringRandom(sample.Rand(), 256))),
			},
			errMsg: "exceeds 255 bytes",
		},
		{
			name:        "pack long string (> 65535 bytes) with compact-long format should fail",
			encodingFmt: memo.EncodingFmtCompactLong,
			args: []memo.CodecArg{
				memo.ArgPayload([]byte(sample.StringRandom(sample.Rand(), 65536))),
			},
			errMsg: "exceeds 65535 bytes",
		},
		{
			name:        "pack empty byte array and string arguments",
			encodingFmt: memo.EncodingFmtCompactShort,
			args: []memo.CodecArg{
				memo.ArgPayload([]byte{}),
				memo.ArgRevertAddress(""),
			},
			expectedLen: 2,
		},
		{
			name:        "failed to pack bytes argument if string is passed",
			encodingFmt: memo.EncodingFmtCompactShort,
			args: []memo.CodecArg{
				memo.ArgPayload(argString), // expect bytes type, but passed string
			},
			errMsg: "argument is not of type []byte",
		},
		{
			name:        "failed to pack address argument if bytes is passed",
			encodingFmt: memo.EncodingFmtCompactShort,
			args: []memo.CodecArg{
				memo.ArgReceiver(argBytes), // expect address type, but passed bytes
			},
			errMsg: "argument is not of type common.Address",
		},
		{
			name:        "failed to pack string argument if bytes is passed",
			encodingFmt: memo.EncodingFmtCompactShort,
			args: []memo.CodecArg{
				memo.ArgRevertAddress(argBytes), // expect string type, but passed bytes
			},
			errMsg: "argument is not of type string",
		},
		{
			name:        "failed to pack unsupported argument type",
			encodingFmt: memo.EncodingFmtCompactShort,
			args: []memo.CodecArg{
				memo.NewArg("receiver", memo.ArgType("unknown"), nil),
			},
			errMsg: "unsupported argument (receiver) type",
		},
	}

	// loop through each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// create a new compact codec and add arguments
			codec, err := memo.NewCodecCompact(tc.encodingFmt)
			require.NoError(t, err)
			codec.AddArguments(tc.args...)

			// pack arguments
			packedData, err := codec.PackArguments()

			if tc.errMsg != "" {
				require.ErrorContains(t, err, tc.errMsg)
				require.Nil(t, packedData)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectedLen, len(packedData))

			// calc expected data for comparison
			expectedData := sample.CompactPack(tc.encodingFmt, tc.args...)

			// validate the packed data
			require.True(t, bytes.Equal(expectedData, packedData), "compact encoded data mismatch")
		})
	}
}

func Test_CodecCompact_UnpackArguments(t *testing.T) {
	// create sample arguments
	argAddress := sample.EthAddress()
	argBytes := []byte("some test bytes argument")
	argString := "some other string argument"

	// test cases
	tests := []struct {
		name        string
		encodingFmt uint8
		data        []byte
		expected    []memo.CodecArg
		errMsg      string
	}{
		{
			name:        "unpack arguments of [address, bytes, string] in compact-short format",
			encodingFmt: memo.EncodingFmtCompactShort,
			data: sample.CompactPack(
				memo.EncodingFmtCompactShort,
				memo.ArgReceiver(argAddress),
				memo.ArgPayload(argBytes),
				memo.ArgRevertAddress(argString),
			),
			expected: []memo.CodecArg{
				memo.ArgReceiver(argAddress),
				memo.ArgPayload(argBytes),
				memo.ArgRevertAddress(argString),
			},
		},
		{
			name:        "unpack arguments of [string, address, bytes] in compact-long format",
			encodingFmt: memo.EncodingFmtCompactLong,
			data: sample.CompactPack(
				memo.EncodingFmtCompactLong,
				memo.ArgRevertAddress(argString),
				memo.ArgReceiver(argAddress),
				memo.ArgPayload(argBytes),
			),
			expected: []memo.CodecArg{
				memo.ArgRevertAddress(argString),
				memo.ArgReceiver(argAddress),
				memo.ArgPayload(argBytes),
			},
		},
		{
			name:        "unpack empty byte array and string argument",
			encodingFmt: memo.EncodingFmtCompactShort,
			data: sample.CompactPack(
				memo.EncodingFmtCompactShort,
				memo.ArgPayload([]byte{}),
				memo.ArgRevertAddress(""),
			),
			expected: []memo.CodecArg{
				memo.ArgPayload([]byte{}),
				memo.ArgRevertAddress(""),
			},
		},
		{
			name:        "failed to unpack address if data length < 20 bytes",
			encodingFmt: memo.EncodingFmtCompactShort,
			data:        []byte{0x01, 0x02, 0x03, 0x04, 0x05},
			expected: []memo.CodecArg{
				memo.ArgReceiver(argAddress),
			},
			errMsg: "expected address, got 5 bytes",
		},
		{
			name:        "failed to unpack string if data length < 1 byte",
			encodingFmt: memo.EncodingFmtCompactShort,
			data:        []byte{},
			expected: []memo.CodecArg{
				memo.ArgRevertAddress(argString),
			},
			errMsg: "expected 1 bytes to decode length",
		},
		{
			name:        "failed to unpack string if actual data is less than decoded length",
			encodingFmt: memo.EncodingFmtCompactShort,
			data:        []byte{0x05, 0x0a, 0x0b, 0x0c, 0x0d}, // length = 5, but only 4 bytes provided
			expected: []memo.CodecArg{
				memo.ArgPayload(argBytes),
			},
			errMsg: "expected 5 bytes, got 4",
		},
		{
			name:        "failed to unpack bytes argument if string is passed",
			encodingFmt: memo.EncodingFmtCompactShort,
			data: sample.CompactPack(
				memo.EncodingFmtCompactShort,
				memo.ArgPayload(argBytes),
			),
			expected: []memo.CodecArg{
				memo.ArgPayload(argString), // expect bytes type, but passed string
			},
			errMsg: "argument is not of type *[]byte",
		},
		{
			name:        "failed to unpack address argument if bytes is passed",
			encodingFmt: memo.EncodingFmtCompactShort,
			data: sample.CompactPack(
				memo.EncodingFmtCompactShort,
				memo.ArgReceiver(argAddress),
			),
			expected: []memo.CodecArg{
				memo.ArgReceiver(argBytes), // expect address type, but passed bytes
			},
			errMsg: "argument is not of type *common.Address",
		},
		{
			name:        "failed to unpack string argument if address is passed",
			encodingFmt: memo.EncodingFmtCompactShort,
			data: sample.CompactPack(
				memo.EncodingFmtCompactShort,
				memo.ArgRevertAddress(argString),
			),
			expected: []memo.CodecArg{
				memo.ArgRevertAddress(argAddress), // expect string type, but passed address
			},
			errMsg: "argument is not of type *string",
		},
		{
			name:        "failed to unpack unsupported argument type",
			encodingFmt: memo.EncodingFmtCompactShort,
			data:        []byte{},
			expected: []memo.CodecArg{
				memo.NewArg("payload", memo.ArgType("unknown"), nil),
			},
			errMsg: "unsupported argument (payload) type",
		},
		{
			name:        "unpacking should fail if not all data is consumed",
			encodingFmt: memo.EncodingFmtCompactShort,
			data: func() []byte {
				data := sample.CompactPack(
					memo.EncodingFmtCompactShort,
					memo.ArgReceiver(argAddress),
					memo.ArgPayload(argBytes),
				)
				// append 1 extra byte
				return append(data, 0x00)
			}(),
			expected: []memo.CodecArg{
				memo.ArgReceiver(argAddress),
				memo.ArgPayload(argBytes),
			},
			errMsg: "consumed bytes (45) != total bytes (46)",
		},
	}

	// loop through each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// create a new compact codec and add arguments
			codec, err := memo.NewCodecCompact(tc.encodingFmt)
			require.NoError(t, err)

			// add output arguments
			output := make([]memo.CodecArg, len(tc.expected))
			for i, arg := range tc.expected {
				output[i] = memo.NewArg(arg.Name, arg.Type, newArgInstance(arg.Arg))
			}
			codec.AddArguments(output...)

			// unpack arguments from compact-encoded data
			err = codec.UnpackArguments(tc.data)

			// validate error message
			if tc.errMsg != "" {
				require.ErrorContains(t, err, tc.errMsg)
				return
			}

			// validate the unpacked arguments values
			require.NoError(t, err)
			for i, arg := range tc.expected {
				ensureArgEquality(t, arg.Arg, output[i].Arg)
			}
		})
	}
}
