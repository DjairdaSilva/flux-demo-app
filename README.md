# Flux CD Demo App

Demo app for Latitude K8s assessment. Deployed via Flux CD GitOps.

## How it works

1. Flux monitors this repo (every 30s)
2. Changes to `k8s/` are automatically applied to the cluster
3. To update the app: edit `k8s/overlays/production/kustomization.yaml` and push

## Quick changes to test CD

Change the app version:
```bash
# Edit k8s/overlays/production/kustomization.yaml
# Change value: "v1.0.0" to "v2.0.0"
# Change color: "#16a34a" (green) to "#dc2626" (red)
git add . && git commit -m "deploy v2.0.0" && git push
```

Flux will detect the change and deploy automatically within 30 seconds.

## URLs

- App: https://demo.64.34.93.17.sslip.io
- Headlamp: https://headlamp.64.34.93.17.sslip.io
- SigNoz: https://signoz.64.34.93.17.sslip.io
