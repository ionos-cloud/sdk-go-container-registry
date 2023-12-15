# CHANGELOG

## 1.1.0
### Changes
- Rename `name` ro `repositoryName` in RegistriesRepositoriesDelete
### Features:
Added support for **Container Registry Vulnerability Scanning**:
- New APIs added: `api_artifacts.go` and `api_vulnerabilities.go
- Added new methods to repositories_api: `RegistriesRepositoriesFindByName` and `RegistriesRepositoriesGet`

## 1.0.1
- `Bytes` field updated from `int32` to `int64`

## 1.0.0
This is the first release of the SDK Python Container Registry. Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls.
