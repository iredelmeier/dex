package tracedstorage

import (
	"time"

	"github.com/dexidp/dex/instrumentation"
	"github.com/dexidp/dex/storage"
	opentracing "github.com/opentracing/opentracing-go"
)

func TraceStorage(parent storage.Storage, c instrumentation.TracerConfig) storage.Storage {
	tracer := instrumentation.NewTracer(c, "storage")

	return &tracedStorage{
		parent: parent,
		tracer: tracer,
	}
}

type tracedStorage struct {
	parent storage.Storage
	tracer opentracing.Tracer
}

func (s *tracedStorage) Close() error {
	span := s.tracer.StartSpan("storage.Close")
	defer span.Finish()

	return s.parent.Close()
}

func (s *tracedStorage) CreateAuthRequest(a storage.AuthRequest) error {
	span := s.tracer.StartSpan("storage.CreateAuthRequest")
	defer span.Finish()

	return s.parent.CreateAuthRequest(a)
}

func (s *tracedStorage) CreateClient(c storage.Client) error {
	span := s.tracer.StartSpan("storage.CreateClient")
	defer span.Finish()

	return s.parent.CreateClient(c)
}

func (s *tracedStorage) CreateAuthCode(c storage.AuthCode) error {
	span := s.tracer.StartSpan("storage.CreateAuthCode")
	defer span.Finish()

	return s.parent.CreateAuthCode(c)
}

func (s *tracedStorage) CreateRefresh(r storage.RefreshToken) error {
	span := s.tracer.StartSpan("storage.CreateRefresh")
	defer span.Finish()

	return s.parent.CreateRefresh(r)
}

func (s *tracedStorage) CreatePassword(p storage.Password) error {
	span := s.tracer.StartSpan("storage.CreatePassword")
	defer span.Finish()

	return s.parent.CreatePassword(p)
}

func (s *tracedStorage) CreateOfflineSessions(os storage.OfflineSessions) error {
	span := s.tracer.StartSpan("storage.CreateOfflineSessions")
	defer span.Finish()

	return s.parent.CreateOfflineSessions(os)
}

func (s *tracedStorage) CreateConnector(c storage.Connector) error {
	span := s.tracer.StartSpan("storage.CreateConnector")
	defer span.Finish()

	return s.parent.CreateConnector(c)
}

func (s *tracedStorage) GetAuthRequest(id string) (storage.AuthRequest, error) {
	span := s.tracer.StartSpan("storage.GetAuthRequest")
	defer span.Finish()

	return s.parent.GetAuthRequest(id)
}

func (s *tracedStorage) GetAuthCode(id string) (storage.AuthCode, error) {
	span := s.tracer.StartSpan("storage.GetAuthCode")
	defer span.Finish()

	return s.parent.GetAuthCode(id)
}

func (s *tracedStorage) GetClient(id string) (storage.Client, error) {
	span := s.tracer.StartSpan("storage.GetClient")
	defer span.Finish()

	return s.parent.GetClient(id)
}

func (s *tracedStorage) GetKeys() (storage.Keys, error) {
	span := s.tracer.StartSpan("storage.GetKeys")
	defer span.Finish()

	return s.parent.GetKeys()
}

func (s *tracedStorage) GetRefresh(id string) (storage.RefreshToken, error) {
	span := s.tracer.StartSpan("storage.GetRefresh")
	defer span.Finish()

	return s.parent.GetRefresh(id)
}

func (s *tracedStorage) GetPassword(email string) (storage.Password, error) {
	span := s.tracer.StartSpan("storage.GetPassword")
	defer span.Finish()

	return s.parent.GetPassword(email)
}

func (s *tracedStorage) GetOfflineSessions(userID string, connID string) (storage.OfflineSessions, error) {
	span := s.tracer.StartSpan("storage.GetOfflineSessions")
	defer span.Finish()

	return s.parent.GetOfflineSessions(userID, connID)
}

func (s *tracedStorage) GetConnector(id string) (storage.Connector, error) {
	span := s.tracer.StartSpan("storage.GetConnector")
	defer span.Finish()

	return s.parent.GetConnector(id)
}

func (s *tracedStorage) ListClients() ([]storage.Client, error) {
	span := s.tracer.StartSpan("storage.ListClients")
	defer span.Finish()

	return s.parent.ListClients()
}

func (s *tracedStorage) ListRefreshTokens() ([]storage.RefreshToken, error) {
	span := s.tracer.StartSpan("storage.ListRefreshTokens")
	defer span.Finish()

	return s.parent.ListRefreshTokens()
}

func (s *tracedStorage) ListPasswords() ([]storage.Password, error) {
	span := s.tracer.StartSpan("storage.ListPasswords")
	defer span.Finish()

	return s.parent.ListPasswords()
}

func (s *tracedStorage) ListConnectors() ([]storage.Connector, error) {
	span := s.tracer.StartSpan("storage.ListConnectors")
	defer span.Finish()

	return s.parent.ListConnectors()
}

func (s *tracedStorage) DeleteAuthRequest(id string) error {
	span := s.tracer.StartSpan("storage.DeleteAuthRequest")
	defer span.Finish()

	return s.parent.DeleteAuthRequest(id)
}

func (s *tracedStorage) DeleteAuthCode(code string) error {
	span := s.tracer.StartSpan("storage.DeleteAuthCode")
	defer span.Finish()

	return s.parent.DeleteAuthCode(code)
}

func (s *tracedStorage) DeleteClient(id string) error {
	span := s.tracer.StartSpan("storage.DeleteClient")
	defer span.Finish()

	return s.parent.DeleteClient(id)
}

func (s *tracedStorage) DeleteRefresh(id string) error {
	span := s.tracer.StartSpan("storage.DeleteRefresh")
	defer span.Finish()

	return s.parent.DeleteRefresh(id)
}

func (s *tracedStorage) DeletePassword(email string) error {
	span := s.tracer.StartSpan("storage.DeletePassword")
	defer span.Finish()

	return s.parent.DeletePassword(email)
}

func (s *tracedStorage) DeleteOfflineSessions(userID string, connID string) error {
	span := s.tracer.StartSpan("storage.DeleteOfflineSessions")
	defer span.Finish()

	return s.parent.DeleteOfflineSessions(userID, connID)
}

func (s *tracedStorage) DeleteConnector(id string) error {
	span := s.tracer.StartSpan("storage.DeleteConnector")
	defer span.Finish()

	return s.parent.DeleteConnector(id)
}

func (s *tracedStorage) UpdateClient(id string, updater func(storage.Client) (storage.Client, error)) error {
	span := s.tracer.StartSpan("storage.UpdateClient")
	defer span.Finish()

	return s.parent.UpdateClient(id, updater)
}

func (s *tracedStorage) UpdateKeys(updater func(storage.Keys) (storage.Keys, error)) error {
	span := s.tracer.StartSpan("storage.UpdateKeys")
	defer span.Finish()

	return s.parent.UpdateKeys(updater)
}

func (s *tracedStorage) UpdateAuthRequest(id string, updater func(storage.AuthRequest) (storage.AuthRequest, error)) error {
	span := s.tracer.StartSpan("storage.UpdateAuthRequest")
	defer span.Finish()

	return s.parent.UpdateAuthRequest(id, updater)
}

func (s *tracedStorage) UpdateRefreshToken(id string, updater func(storage.RefreshToken) (storage.RefreshToken, error)) error {
	span := s.tracer.StartSpan("storage.UpdateRefreshToken")
	defer span.Finish()

	return s.parent.UpdateRefreshToken(id, updater)
}

func (s *tracedStorage) UpdatePassword(email string, updater func(storage.Password) (storage.Password, error)) error {
	span := s.tracer.StartSpan("storage.UpdatePassword")
	defer span.Finish()

	return s.parent.UpdatePassword(email, updater)
}

func (s *tracedStorage) UpdateOfflineSessions(userID string, connID string, updater func(storage.OfflineSessions) (storage.OfflineSessions, error)) error {
	span := s.tracer.StartSpan("storage.UpdateOfflineSessions")
	defer span.Finish()

	return s.parent.UpdateOfflineSessions(userID, connID, updater)
}

func (s *tracedStorage) UpdateConnector(id string, updater func(storage.Connector) (storage.Connector, error)) error {
	span := s.tracer.StartSpan("storage.UpdateConnector")
	defer span.Finish()

	return s.parent.UpdateConnector(id, updater)
}

func (s *tracedStorage) GarbageCollect(now time.Time) (storage.GCResult, error) {
	span := s.tracer.StartSpan("storage.GarbageCollect")
	defer span.Finish()

	return s.parent.GarbageCollect(now)
}
