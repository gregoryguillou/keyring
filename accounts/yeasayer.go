package accounts

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/dontpanicdao/caigo/rpc"
	"github.com/dontpanicdao/caigo/rpc/types"
)

var _ rpc.AccountPlugin = &YeaSayerPlugin{}

type YeaSayerPlugin struct {
	classHash *big.Int
}

func WithYeaSayerPluginPlugin(pluginClassHash string) func() (rpc.AccountOption, error) {
	return func() (rpc.AccountOption, error) {
		plugin, ok := big.NewInt(0).SetString(pluginClassHash, 0)
		if !ok {
			return rpc.AccountOption{}, errors.New("could not convert plugin class hash")
		}
		return rpc.AccountOption{
			AccountPlugin: &YeaSayerPlugin{
				classHash: plugin,
			},
		}, nil
	}
}

func (plugin *YeaSayerPlugin) PluginCall(calls []types.FunctionCall) (types.FunctionCall, error) {
	data := []string{
		fmt.Sprintf("0x%s", plugin.classHash.Text(16)),
		"0x0", // empty (yeasayer), would have been: session_key
		"0x0", // empty (yeasayer), would have been: session_expires
		"0x0", // empty (yeasayer), would have been: merkle_root
		"0x1", // empty (yeasayer), would have been: proof_len
		"0x1", // empty (yeasayer), would have been: proofs with size = prool_len * call_array_len, i.e. 1
		"0x2", // empty (yeasayer), would have been: session_token_len
		"0x0", // empty (yeasayer), would have been: session_token[0]
		"0x0", // empty (yeasayer), would have been: session_token[1]
	}
	return types.FunctionCall{
		ContractAddress:    types.BigToHash(plugin.classHash),
		EntryPointSelector: "use_plugin",
		CallData:           data,
	}, nil
}
