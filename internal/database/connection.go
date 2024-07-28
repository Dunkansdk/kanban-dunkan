package database

import "database/sql"

type IConnection interface {
	GetConnection() *sql.DB
}

type ConnectionHandler struct {
	instance IConnection
}

func CreateConnection(connection IConnection) *ConnectionHandler {
	return &ConnectionHandler{instance: connection}
}

func (handler *ConnectionHandler) Connection() *sql.DB {
	return handler.instance.GetConnection()
}
