# Observability

## Status

Accepted.

## Context

To observe the system properly by collecting telemetry to show how workloads
behave under real conditions. It should provide visibility into below
cross-cutting concerns.

* Reliability
* Performance
* Security
* Cost

A well-architected monitoring stack enables below, which forms the foundation
for proactive management and continuous improvement.

* Early issue detection
* Effective incident response
* Informed operational decisions

## Decision

### Frontend

Sentry. Refer to below setup instructions and guides.

* [Kotlin Multiplatform](https://docs.sentry.io/platforms/kotlin/guides/kotlin-multiplatform/)
* [Guides](https://docs.sentry.io/get-started/guides/)

### Backend

OpenTelemetry to collect telemetry data, then utilized by a Grafana dashboard.

* Traces: Each request should have a trace with distributed trace ID to allow
  easy tracing. Use spans to provide further insights to specific area.
* Metrics: Numerical values for tracking.
* Logs: To help detect and investigate anomalies.

Each microservice should utilize OTel SDK to send telemetry data to telemetry
store. Then, Grafana will implement the dashboard to allow users view and
filter the data, further aggregate to provide more insights.

On AWS, utilize EKS + ADOT to easily collect telemetry per pod without much
configuration, and is more resource efficient.

## Consequences

* Setup overhead.
* Incur cost to host telemetry data.
* Introduce maintenance effort on observability system.

## Compliance

None needed.

## Notes

Author: bingbeann
Approved: bingbeann, 2026-09-21
Last Updated: 2026-09-21

## References

* [Architecture strategies for designing a monitoring system](https://learn.microsoft.com/en-us/azure/well-architected/operational-excellence/observability)
