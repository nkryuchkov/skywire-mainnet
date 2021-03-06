package node

import (
	"net"
	"net/rpc"
	"time"

	"github.com/SkycoinProject/skycoin/src/util/logging"
	"github.com/spf13/cobra"

	"github.com/SkycoinProject/skywire-mainnet/pkg/visor"
)

var log = logging.MustGetLogger("skywire-cli")

var rpcAddr string

func init() {
	RootCmd.PersistentFlags().StringVarP(&rpcAddr, "rpc", "", "localhost:3435", "RPC server address")
}

// RootCmd contains commands that interact with the skywire-visor
var RootCmd = &cobra.Command{
	Use:   "node",
	Short: "Contains sub-commands that interact with the local Skywire Visor",
}

func rpcClient() visor.RPCClient {
	conn, err := net.DialTimeout("tcp", rpcAddr, rpcDialTimeout)
	if err != nil {
		log.Fatal("RPC connection failed:", err)
	}
	if err := conn.SetDeadline(time.Now().Add(rpcConnDuration)); err != nil {
		log.Fatal("RPC connection failed:", err)
	}
	return visor.NewRPCClient(rpc.NewClient(conn), visor.RPCPrefix)
}

const (
	rpcDialTimeout  = time.Second * 5
	rpcConnDuration = time.Second * 60
)
