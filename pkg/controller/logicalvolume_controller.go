package controller

import (
	"time"

	internalController "github.com/topolvm/topolvm/internal/controller"
	"github.com/topolvm/topolvm/pkg/lvmd/proto"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DefaultLVMOperationTimeout is how long each lvmd call made by the
// LogicalVolume reconciler may take unless WithLVMOperationTimeout is given.
const DefaultLVMOperationTimeout = internalController.DefaultLVMOperationTimeout

// LogicalVolumeReconcilerOption configures the LogicalVolume reconciler.
type LogicalVolumeReconcilerOption = internalController.LogicalVolumeReconcilerOption

// WithLVMOperationTimeout sets how long each lvmd call may take before it is
// cancelled. Zero or a negative value disables the timeout.
func WithLVMOperationTimeout(timeout time.Duration) LogicalVolumeReconcilerOption {
	return internalController.WithLVMOperationTimeout(timeout)
}

// SetupLogicalVolumeReconcilerWithServices creates LogicalVolumeReconciler and sets up with manager.
func SetupLogicalVolumeReconcilerWithServices(
	mgr ctrl.Manager,
	client client.Client,
	nodeName string,
	vgService proto.VGServiceClient,
	lvService proto.LVServiceClient,
	opts ...LogicalVolumeReconcilerOption,
) error {
	reconciler := internalController.NewLogicalVolumeReconcilerWithServices(
		client, nodeName, vgService, lvService, opts...)
	return reconciler.SetupWithManager(mgr)
}
