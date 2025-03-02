package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/pkg/errors"
)

// Client Sui client.
type Client struct {
	sui.ISuiAPI
}

const DefaultEventsLimit = 50

const filterMoveEventModule = "MoveEventModule"

// NewFromEndpoint Client constructor based on endpoint string.
func NewFromEndpoint(endpoint string) *Client {
	return New(sui.NewSuiClient(endpoint))
}

// New Client constructor.
func New(client sui.ISuiAPI) *Client {
	return &Client{ISuiAPI: client}
}

// HealthCheck queries latest checkpoint and returns its timestamp.
func (c *Client) HealthCheck(ctx context.Context) (time.Time, error) {
	checkpoint, err := c.GetLatestCheckpoint(ctx)
	if err != nil {
		return time.Time{}, errors.Wrapf(err, "unable to get latest checkpoint")
	}

	ts, err := strconv.ParseInt(checkpoint.TimestampMs, 10, 64)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "unable to parse checkpoint timestamp")
	}

	return time.UnixMilli(ts).UTC(), nil
}

// GetLatestCheckpoint returns the latest checkpoint.
// See https://docs.sui.io/concepts/cryptography/system/checkpoint-verification
func (c *Client) GetLatestCheckpoint(ctx context.Context) (models.CheckpointResponse, error) {
	seqNum, err := c.SuiGetLatestCheckpointSequenceNumber(ctx)
	if err != nil {
		return models.CheckpointResponse{}, errors.Wrap(err, "unable to get latest seq num")
	}

	return c.SuiGetCheckpoint(ctx, models.SuiGetCheckpointRequest{
		CheckpointID: fmt.Sprintf("%d", seqNum),
	})
}

// EventQuery represents pagination options
type EventQuery struct {
	PackageID string
	Module    string
	Cursor    string
	Limit     uint64
}

// QueryModuleEvents queries module events. Return events and the next pagination cursor.
// If cursor is empty, then the end of scroll reached.
func (c *Client) QueryModuleEvents(ctx context.Context, q EventQuery) ([]models.SuiEventResponse, string, error) {
	if q.Limit == 0 {
		q.Limit = DefaultEventsLimit
	}

	if err := q.validate(); err != nil {
		return nil, "", errors.Wrap(err, "invalid request")
	}

	req, err := q.asRequest()
	if err != nil {
		return nil, "", errors.Wrap(err, "unable to create request")
	}

	res, err := c.SuiXQueryEvents(ctx, req)
	switch {
	case err != nil:
		return nil, "", errors.Wrap(err, "unable to query events")
	case !res.HasNextPage:
		return res.Data, "", nil
	default:
		return res.Data, EncodeCursor(res.NextCursor), nil
	}
}

func (p *EventQuery) validate() error {
	switch {
	case p.PackageID == "":
		return errors.New("package id is empty")
	case p.Module == "":
		return errors.New("module is empty")
	case p.Limit == 0:
		return errors.New("limit is empty")
	case p.Limit > 1000:
		return errors.New("limit exceeded")
	default:
		return nil
	}
}

func (p *EventQuery) asRequest() (models.SuiXQueryEventsRequest, error) {
	filter := map[string]any{
		filterMoveEventModule: map[string]any{
			"package": p.PackageID,
			"module":  p.Module,
		},
	}

	cursor, err := DecodeCursor(p.Cursor)
	if err != nil {
		return models.SuiXQueryEventsRequest{}, err
	}

	return models.SuiXQueryEventsRequest{
		SuiEventFilter:  filter,
		Cursor:          cursor,
		Limit:           p.Limit,
		DescendingOrder: false,
	}, nil
}

// GetOwnedObjectID returns the first owned object ID by owner address and struct type.
// If no objects found or multiple objects found, returns error.
func (c *Client) GetOwnedObjectID(ctx context.Context, ownerAddress, structType string) (string, error) {
	res, err := c.SuiXGetOwnedObjects(ctx, models.SuiXGetOwnedObjectsRequest{
		Address: ownerAddress,
		Query: models.SuiObjectResponseQuery{
			Filter: map[string]any{
				"StructType": structType,
			},
		},
		Limit: 1,
	})

	switch {
	case err != nil:
		return "", errors.Wrap(err, "unable to get owned objects")
	case len(res.Data) == 0:
		return "", errors.New("no objects found")
	case len(res.Data) > 1:
		return "", errors.New("multiple objects found")
	}

	return res.Data[0].Data.ObjectId, nil
}

// EncodeCursor encodes event ID into cursor.
func EncodeCursor(id models.EventId) string {
	return fmt.Sprintf("%s,%s", id.TxDigest, id.EventSeq)
}

// DecodeCursor decodes cursor into event ID.
func DecodeCursor(cursor string) (*models.EventId, error) {
	if cursor == "" {
		return nil, nil
	}

	parts := strings.Split(cursor, ",")
	if len(parts) != 2 {
		return nil, errors.New("invalid cursor format")
	}

	return &models.EventId{
		TxDigest: parts[0],
		EventSeq: parts[1],
	}, nil
}
