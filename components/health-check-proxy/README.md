# Health check proxy

## Overview

The Health check proxy is designed to check liveness or readiness endpoints in the applications with mTLS enabled which are built from the scratch image.

The Kubernetes API server cannot reach Pods with mTLS enabled via httpGet, because of that it cannot make a direct call to the liveness or readiness probe for status.

In this case, the best way to check the status of our application is to provide a light-weight binary to our application image, which will proxy our application status to the Kubernetes API server.

## Usage

To build image locally, use command:

```bash
make build-local
```

To attach the Health check proxy binary into your scratch image, just build from it in your dockerfile:

```dockerfile
FROM eu.gcr.io/kyma-project/health-check-proxy:TAG
```

The health-check binary is located under `/health-check` path.

## Parameters

You can configure the Health check proxy binary with the following flags:

| Parameter | Description | Default value |
|-----------|-------------|---------------|
| **path** | Defines an URL path to endpoint which exposes a status. |  |
| **host** | Defines the Host address. | `localhost` |
| **statusPort** | Specifies port of the status endpoint. |  |
| **retry** | Specifies a number of retries for calling the given endpoint. | `1` |

## Example

See the example on how to use the Health check proxy to check the status of your application:

```yaml
readinessProbe:
  exec:
    command:
      - "/healthcheck"
      - "-path=ready"
      - "-statusPort={{ .Values.broker.statusPort }}"
livenessProbe:
  exec:
    command:
      - "/healthcheck"
      - "-path=live"
      - "-statusPort={{ .Values.broker.statusPort }}"
```