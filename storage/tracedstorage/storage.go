package tracedstorage

import (
	"context"
	"time"

	"github.com/dexidp/dex/storage"
	opentracing "github.com/opentracing/opentracing-go"
)

func TraceStorage(parent storage.Storage) storage.Storage {
	return &tracedStorage{
		parent: parent,
	}
}

type tracedStorage struct {
	parent storage.Storage
}

func (s *tracedStorage) Close(ctx context.Context) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.Close")
	defer span.Finish()

	return s.parent.Close(ctx)
}

func (s *tracedStorage) CreateAuthRequest(ctx context.Context, a storage.AuthRequest) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.CreateAuthRequest")
	defer span.Finish()

	return s.parent.CreateAuthRequest(ctx, a)
}

func (s *tracedStorage) CreateClient(ctx context.Context, c storage.Client) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.CreateClient")
	defer span.Finish()

	return s.parent.CreateClient(ctx, c)
}

func (s *tracedStorage) CreateAuthCode(ctx context.Context, c storage.AuthCode) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.CreateAuthCode")
	defer span.Finish()

	return s.parent.CreateAuthCode(ctx, c)
}

func (s *tracedStorage) CreateRefresh(ctx context.Context, r storage.RefreshToken) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.CreateRefresh")
	defer span.Finish()

	return s.parent.CreateRefresh(ctx, r)
}

func (s *tracedStorage) CreatePassword(ctx context.Context, p storage.Password) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.CreatePassword")
	defer span.Finish()

	return s.parent.CreatePassword(ctx, p)
}

func (s *tracedStorage) CreateOfflineSessions(ctx context.Context, os storage.OfflineSessions) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.CreateOfflineSessions")
	defer span.Finish()

	return s.parent.CreateOfflineSessions(ctx, os)
}

func (s *tracedStorage) CreateConnector(ctx context.Context, c storage.Connector) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.CreateConnector")
	defer span.Finish()

	return s.parent.CreateConnector(ctx, c)
}

func (s *tracedStorage) GetAuthRequest(ctx context.Context, id string) (storage.AuthRequest, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.GetAuthRequest")
	defer span.Finish()

	return s.parent.GetAuthRequest(ctx, id)
}

func (s *tracedStorage) GetAuthCode(ctx context.Context, id string) (storage.AuthCode, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.GetAuthCode")
	defer span.Finish()

	return s.parent.GetAuthCode(ctx, id)
}

func (s *tracedStorage) GetClient(ctx context.Context, id string) (storage.Client, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.GetClient")
	defer span.Finish()

	return s.parent.GetClient(ctx, id)
}

func (s *tracedStorage) GetKeys(ctx context.Context) (storage.Keys, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.GetKeys")
	defer span.Finish()

	return s.parent.GetKeys(ctx)
}

func (s *tracedStorage) GetRefresh(ctx context.Context, id string) (storage.RefreshToken, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.GetRefresh")
	defer span.Finish()

	return s.parent.GetRefresh(ctx, id)
}

func (s *tracedStorage) GetPassword(ctx context.Context, email string) (storage.Password, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.GetPassword")
	defer span.Finish()

	return s.parent.GetPassword(ctx, email)
}

func (s *tracedStorage) GetOfflineSessions(ctx context.Context, userID string, connID string) (storage.OfflineSessions, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.GetOfflineSessions")
	defer span.Finish()

	return s.parent.GetOfflineSessions(ctx, userID, connID)
}

func (s *tracedStorage) GetConnector(ctx context.Context, id string) (storage.Connector, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.GetConnector")
	defer span.Finish()

	return s.parent.GetConnector(ctx, id)
}

func (s *tracedStorage) ListClients(ctx context.Context) ([]storage.Client, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.ListClients")
	defer span.Finish()

	return s.parent.ListClients(ctx)
}

func (s *tracedStorage) ListRefreshTokens(ctx context.Context) ([]storage.RefreshToken, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.ListRefreshTokens")
	defer span.Finish()

	return s.parent.ListRefreshTokens(ctx)
}

func (s *tracedStorage) ListPasswords(ctx context.Context) ([]storage.Password, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.ListPasswords")
	defer span.Finish()

	return s.parent.ListPasswords(ctx)
}

func (s *tracedStorage) ListConnectors(ctx context.Context) ([]storage.Connector, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.ListConnectors")
	defer span.Finish()

	return s.parent.ListConnectors(ctx)
}

func (s *tracedStorage) DeleteAuthRequest(ctx context.Context, id string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.DeleteAuthRequest")
	defer span.Finish()

	return s.parent.DeleteAuthRequest(ctx, id)
}

func (s *tracedStorage) DeleteAuthCode(ctx context.Context, code string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.DeleteAuthCode")
	defer span.Finish()

	return s.parent.DeleteAuthCode(ctx, code)
}

func (s *tracedStorage) DeleteClient(ctx context.Context, id string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.DeleteClient")
	defer span.Finish()

	return s.parent.DeleteClient(ctx, id)
}

func (s *tracedStorage) DeleteRefresh(ctx context.Context, id string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.DeleteRefresh")
	defer span.Finish()

	return s.parent.DeleteRefresh(ctx, id)
}

func (s *tracedStorage) DeletePassword(ctx context.Context, email string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.DeletePassword")
	defer span.Finish()

	return s.parent.DeletePassword(ctx, email)
}

func (s *tracedStorage) DeleteOfflineSessions(ctx context.Context, userID string, connID string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.DeleteOfflineSessions")
	defer span.Finish()

	return s.parent.DeleteOfflineSessions(ctx, userID, connID)
}

func (s *tracedStorage) DeleteConnector(ctx context.Context, id string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.DeleteConnector")
	defer span.Finish()

	return s.parent.DeleteConnector(ctx, id)
}

func (s *tracedStorage) UpdateClient(ctx context.Context, id string, updater func(storage.Client) (storage.Client, error)) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.UpdateClient")
	defer span.Finish()

	return s.parent.UpdateClient(ctx, id, updater)
}

func (s *tracedStorage) UpdateKeys(ctx context.Context, updater func(storage.Keys) (storage.Keys, error)) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.UpdateKeys")
	defer span.Finish()

	return s.parent.UpdateKeys(ctx, updater)
}

func (s *tracedStorage) UpdateAuthRequest(ctx context.Context, id string, updater func(storage.AuthRequest) (storage.AuthRequest, error)) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.UpdateAuthRequest")
	defer span.Finish()

	return s.parent.UpdateAuthRequest(ctx, id, updater)
}

func (s *tracedStorage) UpdateRefreshToken(ctx context.Context, id string, updater func(storage.RefreshToken) (storage.RefreshToken, error)) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.UpdateRefreshToken")
	defer span.Finish()

	return s.parent.UpdateRefreshToken(ctx, id, updater)
}

func (s *tracedStorage) UpdatePassword(ctx context.Context, email string, updater func(storage.Password) (storage.Password, error)) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.UpdatePassword")
	defer span.Finish()

	return s.parent.UpdatePassword(ctx, email, updater)
}

func (s *tracedStorage) UpdateOfflineSessions(ctx context.Context, userID string, connID string, updater func(storage.OfflineSessions) (storage.OfflineSessions, error)) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.UpdateOfflineSessions")
	defer span.Finish()

	return s.parent.UpdateOfflineSessions(ctx, userID, connID, updater)
}

func (s *tracedStorage) UpdateConnector(ctx context.Context, id string, updater func(storage.Connector) (storage.Connector, error)) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.UpdateConnector")
	defer span.Finish()

	return s.parent.UpdateConnector(ctx, id, updater)
}

func (s *tracedStorage) GarbageCollect(ctx context.Context, now time.Time) (storage.GCResult, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "storage.GarbageCollect")
	defer span.Finish()

	return s.parent.GarbageCollect(ctx, now)
}
