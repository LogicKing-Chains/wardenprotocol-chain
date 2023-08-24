package main

import (
	"context"
	"crypto/elliptic"
	"log"
	"log/slog"

	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.qredo.com/qrdochain/fusionchain/go-client"
	"gitlab.qredo.com/qrdochain/fusionchain/x/treasury/types"
)

// MockKeyRequestsHandler implements KeyRequestsHandler.
// It uses an in-memory database to store the generated keys.
type MockKeyRequestsHandler struct {
	KeyDB       *InMemoryKeyDB
	QueryClient *client.QueryClient
	TxClient    *client.TxClient
	Logger      *slog.Logger
}

func (h *MockKeyRequestsHandler) HandleKeyRequests(ctx context.Context, pendingRequests []*types.KeyRequest) error {
	// In this mock implementation we process each request one by one.
	// Each of them will generate a separate transaction to Fusion Chain as result.
	// A production implementation might want to bundle all the ApproveKeyRequest messages into a single transaction.
	for _, request := range pendingRequests {
		h.processReq(ctx, request)
	}
	return nil
}

func (h *MockKeyRequestsHandler) processReq(ctx context.Context, request *types.KeyRequest) {
	l := h.Logger.With("request_id", request.Id)
	l.InfoContext(ctx, "received")

	// generate new key
	sk, err := crypto.GenerateKey()
	if err != nil {
		l.ErrorContext(ctx, "error", err)
		return
	}

	pk := elliptic.Marshal(sk.PublicKey, sk.PublicKey.X, sk.PublicKey.Y)

	// approve the user request, provide the generated public key
	err = h.TxClient.FulfilKeyRequest(ctx, request.Id, pk)
	if err != nil {
		l.ErrorContext(ctx, "fulfilling request", err)
		return
	}

	// fetch again the request to get the newly created key id
	updatedRequest, err := h.QueryClient.GetKeyRequest(ctx, request.Id)
	if err != nil {
		log.Printf("KeyRequest[%d] error: %s\n", request.Id, err)
		l.ErrorContext(ctx, "getting updated request", err)
		return
	}
	if updatedRequest.Status != types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED {
		l.ErrorContext(ctx, "request is not an approved request")
		return
	}
	keyID := updatedRequest.GetSuccessKeyId()

	// store the generated secret key in our database, will be used when user requests signatures
	err = h.KeyDB.Set(keyID, sk)
	if err != nil {
		l.ErrorContext(ctx, "storing key", err)
		return
	}

	l.InfoContext(ctx, "fulfilled", "key_id", keyID)
}
