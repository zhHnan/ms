package errs

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hnz.com/ms_serve/ms-common"
)

func GrpcError(err *BError) error {
	return status.Error(codes.Code(err.Code), err.Msg)
}
func ParseGrpcError(err error) (common.BusinessCode, string) {
	fromError, _ := status.FromError(err)
	return common.BusinessCode(fromError.Code()), fromError.Message()
}
