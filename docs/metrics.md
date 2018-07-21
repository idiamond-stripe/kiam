# Metrics

Kiam exports both Prometheus and StatsD metrics to determine the health of the
system, check the timing of each RPC call, and monitor the size of the
credentials cache. By default, Prometheus metrics are exported on
`localhost:9620` and StatsD metrics are sent to `127.0.0.1:8125`. StatsD
metrics are not aggregated and flushed every 100ms.

## Metrics configuration

- The `statsd` flag controls the address to which to send StatsD metrics. This
  is by default `127.0.0..1:8125`. If this is blank, StatsD metrics will be
  silenced.
- The `statsd-prefix` flag controls the initial prefix that will be appended to
  Kiam's StatsD metrics. This is by default `kiam`.
- The `statsd-interval` flag controls how frequently the in-memory metrics
  buffer will be flushed to the specified StatsD endpoint. Metrics are
  not aggregated in this buffer and the raw counts will be flushed to the
  underlying StatsD sink. This is by default `100ms`.
- The `prometheus-listen-addr` controls which address Kiam should create a
  Prometheus endpoint on. This is by default `localhost:9620`. The metrics
  themselves can be accessed at `<prometheus-listen-addr>/metrics`.
- The `prometheus-sync-interval` flag controls how frequently Prometheus
  metrics should be updated. This is by default `5s`.

## Emitted Metrics

### Prometheus
#### Metadata Subsystem
- `handler_latency_milliseconds` - Bucketed histogram of handler timings. Tagged by handler
- `credential_fetch_error` - Number of errors fetching the credentials for a pod
- `credential_encode_error` - Number of errors encoding credentials for a pod
- `find_role_error_total` - Number of errors finding the role for a pod
- `empty_role_total` - Number of empty roles returned
- `success_total` - Number of successful responses from a handler
- `responses_total` - Responses from mocked out metadata handlers

#### STS Subsystem
- `cache_hit_total` - Number of cache hits to the metadata cache
- `cache_miss_total` - Number of cache misses to the metadata cache
- `error_issuing_count` - Number of errors issuing credentials
- `assumerole_timing_milliseconds` - Bucketed histogram of assumeRole timings
- `assume_role_executing_total` - Number of assume role calls currently executing

#### K8s Subsystem
- `dropped_pods_total` - Number of dropped pods because of full cache

### StatsD Timing metrics
- `gateway.rpc.GetRole` - Observed client side latency of GetRole RPC
- `gateway.rpc.GetCredentials` - Observed client side latency of GetCredentials RPC
- `server.rpc.GetRoleCredentials` - Observed server side latency of GetRoleCredentials RPC
- `server.rpc.IsAllowedAssumeRole` - Observed server side latency of IsAllowedAssumeRole RPC
- `server.rpc.GetHealth` - Observed server side latency of GetHealth RPC
- `server.rpc.GetPodRole` - Observed server side latency of GetPodRole RPC
- `server.rpc.GetRoleCredentials` - Observed server side latency of GetRoleCredentials RPC
- `handler.role_name` - Observed latency of role_name handler
- `handler.health` - Observed latency of health handler
- `handler.credentials` - Observed latency of credentials handler
- `aws.assume_role` - Observed latency of aws assume role request