/*
Copyright 2021 The Karmada Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	componentbaseconfig "k8s.io/component-base/config"

	"github.com/karmada-io/karmada/pkg/features"
	"github.com/karmada-io/karmada/pkg/sharedcli/profileflag"
	"github.com/karmada-io/karmada/pkg/sharedcli/ratelimiterflag"
	"github.com/karmada-io/karmada/pkg/util"
	"github.com/karmada-io/karmada/pkg/util/names"
)

const (
	// DefaultKarmadaClusterNamespace defines the default namespace where the member cluster secrets are stored.
	DefaultKarmadaClusterNamespace = "karmada-cluster"
)

var (
	defaultElectionLeaseDuration = metav1.Duration{Duration: 15 * time.Second}
	defaultElectionRenewDeadline = metav1.Duration{Duration: 10 * time.Second}
	defaultElectionRetryPeriod   = metav1.Duration{Duration: 2 * time.Second}
)

// Options contains everything necessary to create and run controller-manager.
type Options struct {
	// Controllers contains all controller names.
	Controllers       []string
	LeaderElection    componentbaseconfig.LeaderElectionConfiguration
	KarmadaKubeConfig string
	// ClusterContext is the name of the cluster context in control plane KUBECONFIG file.
	// Default value is the current-context.
	KarmadaContext string
	ClusterName    string
	// ClusterNamespace holds the namespace name where the member cluster secrets are stored.
	ClusterNamespace string
	// ClusterStatusUpdateFrequency is the frequency that controller computes and report cluster status.
	// It must work with ClusterMonitorGracePeriod(--cluster-monitor-grace-period) in karmada-controller-manager.
	ClusterStatusUpdateFrequency metav1.Duration
	// ClusterLeaseDuration is a duration that candidates for a lease need to wait to force acquire it.
	// This is measure against time of last observed RenewTime.
	ClusterLeaseDuration metav1.Duration
	// ClusterLeaseRenewIntervalFraction is a fraction coordinated with ClusterLeaseDuration that
	// how long the current holder of a lease has last updated the lease.
	ClusterLeaseRenewIntervalFraction float64
	// ClusterSuccessThreshold is the duration of successes for the cluster to be considered healthy after recovery.
	ClusterSuccessThreshold metav1.Duration
	// ClusterFailureThreshold is the duration of failure for the cluster to be considered unhealthy.
	ClusterFailureThreshold metav1.Duration
	// ClusterAPIQPS is the QPS to use while talking with cluster kube-apiserver.
	ClusterAPIQPS float32
	// ClusterAPIBurst is the burst to allow while talking with cluster kube-apiserver.
	ClusterAPIBurst int
	// KubeAPIQPS is the QPS to use while talking with karmada-apiserver.
	KubeAPIQPS float32
	// KubeAPIBurst is the burst to allow while talking with karmada-apiserver.
	KubeAPIBurst int

	ClusterCacheSyncTimeout metav1.Duration
	// ResyncPeriod is the base frequency the informers are resynced.
	// Defaults to 0, which means the created informer will never do resyncs.
	ResyncPeriod metav1.Duration

	// ClusterAPIEndpoint holds the apiEndpoint of the cluster.
	ClusterAPIEndpoint string
	// ProxyServerAddress holds the proxy server address that is used to proxy to the cluster.
	ProxyServerAddress string
	// ConcurrentClusterSyncs is the number of cluster objects that are
	// allowed to sync concurrently.
	ConcurrentClusterSyncs int
	// ConcurrentWorkSyncs is the number of work objects that are
	// allowed to sync concurrently.
	ConcurrentWorkSyncs int
	// MetricsBindAddress is the TCP address that the controller should bind to
	// for serving prometheus metrics.
	// It can be set to "0" to disable the metrics serving.
	// Defaults to ":8080".
	MetricsBindAddress string
	// HealthProbeBindAddress is the TCP address that the controller should bind to
	// for serving health probes
	// It can be set to "0" to disable serving the health probe.
	// Defaults to ":10357".
	HealthProbeBindAddress string

	RateLimiterOpts ratelimiterflag.Options

	ProfileOpts profileflag.Options

	// ReportSecrets specifies the secrets that are allowed to be reported to the Karmada control plane
	// during registering.
	// Valid values are:
	// - "None": Don't report any secrets.
	// - "KubeCredentials": Report the secret that contains mandatory credentials to access the member cluster.
	// - "KubeImpersonator": Report the secret that contains the token of impersonator.
	// - "KubeCredentials,KubeImpersonator": Report both KubeCredentials and KubeImpersonator.
	// Defaults to "KubeCredentials,KubeImpersonator".
	ReportSecrets []string

	// ClusterProvider is the cluster's provider.
	ClusterProvider string

	// ClusterRegion represents the region of the cluster locate in.
	ClusterRegion string

	// ClusterZones represents the zones of the cluster locate in.
	ClusterZones []string

	// EnableClusterResourceModeling indicates if enable cluster resource modeling.
	// The resource modeling might be used by the scheduler to make scheduling decisions
	// in scenario of dynamic replica assignment based on cluster free resources.
	// Disable if it does not fit your cases for better performance.
	EnableClusterResourceModeling bool

	// CertRotationCheckingInterval defines the interval of checking if the certificate need to be rotated.
	CertRotationCheckingInterval time.Duration
	// CertRotationRemainingTimeThreshold defines the threshold of remaining time of the valid certificate.
	// If the ratio of remaining time to total time is less than or equal to this threshold, the certificate rotation starts.
	CertRotationRemainingTimeThreshold float64
	// KarmadaKubeconfigNamespace is the namespace of the secret containing karmada-agent certificate.
	KarmadaKubeconfigNamespace string
}

// NewOptions builds an default scheduler options.
func NewOptions() *Options {
	return &Options{
		LeaderElection: componentbaseconfig.LeaderElectionConfiguration{
			LeaderElect:       true,
			ResourceLock:      resourcelock.LeasesResourceLock,
			ResourceNamespace: names.NamespaceKarmadaSystem,
		},
	}
}

// AddFlags adds flags of scheduler to the specified FlagSet
func (o *Options) AddFlags(fs *pflag.FlagSet, allControllers []string) {
	if o == nil {
		return
	}

	fs.StringSliceVar(&o.Controllers, "controllers", []string{"*"}, fmt.Sprintf(
		"A list of controllers to enable. '*' enables all on-by-default controllers, 'foo' enables the controller named 'foo', '-foo' disables the controller named 'foo'. All controllers: %s.",
		strings.Join(allControllers, ", "),
	))
	fs.BoolVar(&o.LeaderElection.LeaderElect, "leader-elect", true, "Start a leader election client and gain leadership before executing the main loop. Enable this when running replicated components for high availability.")
	fs.StringVar(&o.LeaderElection.ResourceNamespace, "leader-elect-resource-namespace", names.NamespaceKarmadaSystem, "The namespace of resource object that is used for locking during leader election.")
	fs.DurationVar(&o.LeaderElection.LeaseDuration.Duration, "leader-elect-lease-duration", defaultElectionLeaseDuration.Duration, ""+
		"The duration that non-leader candidates will wait after observing a leadership "+
		"renewal until attempting to acquire leadership of a led but unrenewed leader "+
		"slot. This is effectively the maximum duration that a leader can be stopped "+
		"before it is replaced by another candidate. This is only applicable if leader "+
		"election is enabled.")
	fs.DurationVar(&o.LeaderElection.RenewDeadline.Duration, "leader-elect-renew-deadline", defaultElectionRenewDeadline.Duration, ""+
		"The interval between attempts by the acting master to renew a leadership slot "+
		"before it stops leading. This must be less than or equal to the lease duration. "+
		"This is only applicable if leader election is enabled.")
	fs.DurationVar(&o.LeaderElection.RetryPeriod.Duration, "leader-elect-retry-period", defaultElectionRetryPeriod.Duration, ""+
		"The duration the clients should wait between attempting acquisition and renewal "+
		"of a leadership. This is only applicable if leader election is enabled.")
	fs.StringVar(&o.KarmadaKubeConfig, "karmada-kubeconfig", o.KarmadaKubeConfig, "Path to karmada control plane kubeconfig file.")
	fs.StringVar(&o.KarmadaContext, "karmada-context", "", "Name of the cluster context in karmada control plane kubeconfig file.")
	fs.StringVar(&o.ClusterName, "cluster-name", o.ClusterName, "Name of member cluster that the agent serves for.")
	fs.StringVar(&o.ClusterNamespace, "cluster-namespace", DefaultKarmadaClusterNamespace, "Namespace in the control plane where member cluster secrets are stored.")
	fs.DurationVar(&o.ClusterStatusUpdateFrequency.Duration, "cluster-status-update-frequency", 10*time.Second, "Specifies how often karmada-agent posts cluster status to karmada-apiserver. Note: be cautious when changing the constant, it must work with ClusterMonitorGracePeriod in karmada-controller-manager.")
	fs.DurationVar(&o.ClusterLeaseDuration.Duration, "cluster-lease-duration", 40*time.Second,
		"Specifies the expiration period of a cluster lease.")
	fs.Float64Var(&o.ClusterLeaseRenewIntervalFraction, "cluster-lease-renew-interval-fraction", 0.25,
		"Specifies the cluster lease renew interval fraction.")
	fs.DurationVar(&o.ClusterSuccessThreshold.Duration, "cluster-success-threshold", 30*time.Second, "The duration of successes for the cluster to be considered healthy after recovery.")
	fs.DurationVar(&o.ClusterFailureThreshold.Duration, "cluster-failure-threshold", 30*time.Second, "The duration of failure for the cluster to be considered unhealthy.")
	fs.Float32Var(&o.ClusterAPIQPS, "cluster-api-qps", 40.0, "QPS to use while talking with cluster kube-apiserver.")
	fs.IntVar(&o.ClusterAPIBurst, "cluster-api-burst", 60, "Burst to use while talking with cluster kube-apiserver.")
	fs.Float32Var(&o.KubeAPIQPS, "kube-api-qps", 40.0, "QPS to use while talking with karmada-apiserver.")
	fs.IntVar(&o.KubeAPIBurst, "kube-api-burst", 60, "Burst to use while talking with karmada-apiserver.")
	fs.DurationVar(&o.ClusterCacheSyncTimeout.Duration, "cluster-cache-sync-timeout", util.CacheSyncTimeout, "Timeout period waiting for cluster cache to sync.")
	fs.StringVar(&o.ClusterAPIEndpoint, "cluster-api-endpoint", o.ClusterAPIEndpoint, "APIEndpoint of the cluster.")
	fs.StringVar(&o.ProxyServerAddress, "proxy-server-address", o.ProxyServerAddress, "Address of the proxy server that is used to proxy to the cluster.")
	fs.DurationVar(&o.ResyncPeriod.Duration, "resync-period", 0, "Base frequency the informers are resynced.")
	fs.IntVar(&o.ConcurrentClusterSyncs, "concurrent-cluster-syncs", 5, "The number of Clusters that are allowed to sync concurrently.")
	fs.IntVar(&o.ConcurrentWorkSyncs, "concurrent-work-syncs", 5, "The number of Works that are allowed to sync concurrently.")
	fs.StringSliceVar(&o.ReportSecrets, "report-secrets", []string{"KubeCredentials", "KubeImpersonator"}, "The secrets that are allowed to be reported to the Karmada control plane during registering. Valid values are 'KubeCredentials', 'KubeImpersonator' and 'None'. e.g 'KubeCredentials,KubeImpersonator' or 'None'.")
	fs.StringVar(&o.MetricsBindAddress, "metrics-bind-address", ":8080", "The TCP address that the controller should bind to for serving prometheus metrics(e.g. 127.0.0.1:8080, :8080). It can be set to \"0\" to disable the metrics serving.")
	fs.StringVar(&o.HealthProbeBindAddress, "health-probe-bind-address", ":10357", "The TCP address that the controller should bind to for serving health probes(e.g. 127.0.0.1:10357, :10357). It can be set to \"0\" to disable serving the health probe. Defaults to 0.0.0.0:10357.")
	fs.StringVar(&o.ClusterProvider, "cluster-provider", "", "Provider of the joining cluster. The Karmada scheduler can use this information to spread workloads across providers for higher availability.")
	fs.StringVar(&o.ClusterRegion, "cluster-region", "", "The region of the joining cluster. The Karmada scheduler can use this information to spread workloads across regions for higher availability.")
	fs.StringSliceVar(&o.ClusterZones, "cluster-zones", []string{}, "The zones of the joining cluster. The Karmada scheduler can use this information to spread workloads across zones for higher availability.")
	fs.BoolVar(&o.EnableClusterResourceModeling, "enable-cluster-resource-modeling", true, "Enable means controller would build resource modeling for each cluster by syncing Nodes and Pods resources.\n"+
		"The resource modeling might be used by the scheduler to make scheduling decisions in scenario of dynamic replica assignment based on cluster free resources.\n"+
		"Disable if it does not fit your cases for better performance.")
	fs.DurationVar(&o.CertRotationCheckingInterval, "cert-rotation-checking-interval", 5*time.Minute, "The interval of checking if the certificate need to be rotated. This is only applicable if cert rotation is enabled")
	fs.Float64Var(&o.CertRotationRemainingTimeThreshold, "cert-rotation-remaining-time-threshold", 0.2, "The threshold of remaining time of the valid certificate. This is only applicable if cert rotation is enabled.")
	fs.StringVar(&o.KarmadaKubeconfigNamespace, "karmada-kubeconfig-namespace", "karmada-system", "Namespace of the secret containing karmada-agent certificate. This is only applicable if cert rotation is enabled.")
	o.RateLimiterOpts.AddFlags(fs)
	features.FeatureGate.AddFlag(fs)
	o.ProfileOpts.AddFlags(fs)
}
