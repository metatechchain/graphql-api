/*
Package rpc implements bridge to Meta full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an MetaTech/Meta node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Meta RPC interface for remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Meta RPC interface with connection limited to specified endpoints.

We strongly discourage opening Meta RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum"
)

// mtcHeadsObserverSubscribeTick represents the time between subscription attempts.
const mtcHeadsObserverSubscribeTick = 30 * time.Second

// observeBlocks collects new blocks from the blockchain network
// and posts them into the proxy channel for processing.
func (mtc *MtcBridge) observeBlocks() {
	var sub ethereum.Subscription
	defer func() {
		if sub != nil {
			sub.Unsubscribe()
		}
		mtc.log.Noticef("block observer done")
		mtc.wg.Done()
	}()

	sub = mtc.blockSubscription()
	for {
		// re-subscribe if the subscription ref is not valid
		if sub == nil {
			tm := time.NewTimer(mtcHeadsObserverSubscribeTick)
			select {
			case <-mtc.sigClose:
				return
			case <-tm.C:
				sub = mtc.blockSubscription()
				continue
			}
		}

		// use the subscriptions
		select {
		case <-mtc.sigClose:
			return
		case err := <-sub.Err():
			mtc.log.Errorf("block subscription failed; %s", err.Error())
			sub = nil
		}
	}
}

// blockSubscription provides a subscription for new blocks received
// by the connected blockchain node.
func (mtc *MtcBridge) blockSubscription() ethereum.Subscription {
	sub, err := mtc.rpc.EthSubscribe(context.Background(), mtc.headers, "newHeads")
	if err != nil {
		mtc.log.Criticalf("can not observe new blocks; %s", err.Error())
		return nil
	}
	return sub
}
