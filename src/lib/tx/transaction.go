package tx

// TransactionStatus
// To keep the consistence among series of middleware operations.
type TransactionStatus struct {
}

type TransactionManager struct {
}

func (tx *TransactionStatus) commit(status TransactionStatus) error {
	return nil
}

func (tx *TransactionStatus) rollback(status TransactionStatus) error {
	return nil
}

func (tx *TransactionManager) commit(status TransactionStatus) error {
	return nil
}

func (tx *TransactionManager) rollback(status TransactionStatus) error {
	return nil
}
