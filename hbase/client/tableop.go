package client

import "github.com/liyuan35023/utils/hbase/model"

// TableOp is interface for query, insert, remove, scan rows
type TableOp interface {
	Query(get *model.Get) (*model.HBaseRow, error)
	Insert(put *model.Put) error
	Remove(del *model.Delete) error
	GetScanner(scan *model.Scan) (*model.Scanner, error)
	GetName() string
	Close()
}
