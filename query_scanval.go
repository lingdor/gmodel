package gmodel

type ScanParam interface {
	GetPointer() any
	GetVal() any
}

type DefaultScanParam[T any] struct {
	v *T
}

func (d *DefaultScanParam[T]) GetPointer() any {
	return d.v
}

func (d *DefaultScanParam[T]) GetVal() any {
	if d.v == nil {
		return nil
	}
	return *d.v
}
