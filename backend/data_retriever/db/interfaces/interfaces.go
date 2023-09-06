package interfaces

type CrudDbClient[T any, Key any] interface {
    CreateRecord(recordStruct T) error
    RetrieveRecordById(id Key) (T, error)
}
