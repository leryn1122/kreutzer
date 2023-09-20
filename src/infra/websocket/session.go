package ws

import "sync"

type Session = string

type SessionManager struct {
	mutex    sync.Mutex
	Sessions map[Session]*Connection
}

func NewSessionManager() SessionManager {
	return SessionManager{
		mutex:    sync.Mutex{},
		Sessions: make(map[Session]*Connection),
	}
}

func (m *SessionManager) AddSession(conn *Connection) {
	m.mutex.TryLock()
	defer m.mutex.Unlock()
	m.Sessions[conn.Session] = conn
}

func (m *SessionManager) RemoveSession(conn *Connection) {
	m.mutex.TryLock()
	defer m.mutex.Unlock()
	delete(m.Sessions, conn.Session)
}
