# Event Logger Operator — Step-by-Step Setup

This guide walks through building, pushing, and deploying the Event Logger Operator  
so it can be run on a different machine without relying on `make deploy`.

---

## 1. Build and Push Operator Image

Run these on your **build machine**:

```bash
make docker-build IMG=eduardkh/event-logger-operator:v0.1.0
make docker-push IMG=eduardkh/event-logger-operator:v0.1.0
```

## 2. Generate CRDs

```bash
make manifests
```

## 3. Apply the CRDs (on the Target Machine)

```bash
kubectl apply -f config/crd/bases/
```

## 4. Deploy Operator Bundle (RBAC + Deployment)

```bash
kubectl apply -f operator-bundle.yaml
```

## 5. Restart Pod (Optional but Recommended)

```bash
kubectl delete pod -n event-logger-operator-system -l control-plane=controller-manager
```

## 6. Verify Deployment

```bash
kubectl get pods -n event-logger-operator-system
```

## Follow logs

```bash
kubectl logs -f deployment/event-logger-operator-controller-manager -n event-logger-operator-system
```
