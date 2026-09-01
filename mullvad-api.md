# API Route Map

This repository contains two API surfaces:

1. The outbound HTTPS REST API used by the daemon to communicate with Mullvad services.
2. The local gRPC management API used by the app frontend to communicate with the daemon.

## Outbound HTTPS REST API

The base URL is constructed as `https://{api-host}/{path}` in [`mullvad-api/src/rest.rs`](mullvad-api/src/rest.rs).

| Method         | Route                                  | Purpose                           | Implementation                                                   |
| -------------- | -------------------------------------- | --------------------------------- | ---------------------------------------------------------------- |
| `POST`         | `/auth/v1/token`                       | Obtain an account access token    | [`mullvad-api/src/access.rs`](mullvad-api/src/access.rs)         |
| `GET`          | `/accounts/v1/accounts/me`             | Get account data                  | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `POST`         | `/accounts/v1/accounts`                | Create an account                 | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `DELETE`       | `/accounts/v1/accounts/me`             | Delete an account                 | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `GET`          | `/accounts/v1/devices`                 | List devices                      | [`mullvad-api/src/device.rs`](mullvad-api/src/device.rs)         |
| `POST`         | `/accounts/v1/devices`                 | Create a device                   | [`mullvad-api/src/device.rs`](mullvad-api/src/device.rs)         |
| `GET`          | `/accounts/v1/devices/{id}`            | Get a device                      | [`mullvad-api/src/device.rs`](mullvad-api/src/device.rs)         |
| `DELETE`       | `/accounts/v1/devices/{id}`            | Remove a device                   | [`mullvad-api/src/device.rs`](mullvad-api/src/device.rs)         |
| `PUT`          | `/accounts/v1/devices/{id}/pubkey`     | Rotate a WireGuard key            | [`mullvad-api/src/device.rs`](mullvad-api/src/device.rs)         |
| `GET`          | `/app/v1/relays`                       | Fetch the relay list              | [`mullvad-api/src/relay_list.rs`](mullvad-api/src/relay_list.rs) |
| `GET`          | `/app/releases/{platform}.json`        | Check app versions                | [`mullvad-api/src/version.rs`](mullvad-api/src/version.rs)       |
| `GET`          | `/app/releases/android.json`           | Check Android versions            | [`mullvad-api/src/version.rs`](mullvad-api/src/version.rs)       |
| `POST`         | `/app/v1/submit-voucher`               | Submit a voucher                  | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `POST`         | `/app/v1/www-auth-token`               | Get a web authentication token    | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `POST`         | `/app/v1/problem-report`               | Submit a problem report           | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `GET` / `HEAD` | `/app/v1/api-addrs`                    | Retrieve or check API addresses   | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `POST`         | `/payments/apple/v2/init`              | Initialize an Apple payment       | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `POST`         | `/payments/apple/v2/check`             | Verify an Apple payment           | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `POST`         | `/payments/google-play/v1/init`        | Initialize a Google Play purchase | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |
| `POST`         | `/payments/google-play/v1/acknowledge` | Verify a Google Play purchase     | [`mullvad-api/src/lib.rs`](mullvad-api/src/lib.rs)               |

The Apple payment routes are compiled for iOS, and the Google Play routes are compiled for Android.

### Route not present in this codebase

`/www/relays/all/` is not referenced by the current source tree and is not an active route in this client implementation. The current relay-list request is:

```text
GET /app/v1/relays
```

If `/www/relays/all/` is a legacy or web-facing endpoint, it may belong to a different Mullvad service or an older API version rather than this app’s API client.

### Connection-check endpoints

The daemon also queries these endpoints to determine the current public IP and Mullvad location:

```text
GET https://ipv4.am.i.mullvad.net/json
GET https://ipv6.am.i.mullvad.net/json
```

The IPv6 request is made only when IPv6 checking is enabled. See [`mullvad-daemon/src/geoip.rs`](mullvad-daemon/src/geoip.rs).

## Local daemon management API

The app frontend communicates with the daemon through gRPC over the local RPC socket. The gRPC services are registered in [`mullvad-management-interface/src/lib.rs`](mullvad-management-interface/src/lib.rs).

### Services

- `ManagementService`
- `RelaySelectorService`

The complete method definitions are in [`mullvad-management-interface/proto/management_interface.proto`](mullvad-management-interface/proto/management_interface.proto) and [`mullvad-management-interface/proto/relay_selector.proto`](mullvad-management-interface/proto/relay_selector.proto).

Representative gRPC routes include:

```text
/ManagementService/ConnectTunnel
/ManagementService/DisconnectTunnel
/ManagementService/ReconnectTunnel
/ManagementService/GetTunnelState
/ManagementService/EventsListen
/ManagementService/GetSettings
/ManagementService/SetRelaySettings
/ManagementService/GetRelayLocations
/ManagementService/GetBridges
/ManagementService/CreateNewAccount
/ManagementService/LoginAccount
/ManagementService/LogoutAccount
/ManagementService/GetAccountData
/ManagementService/ListDevices
/ManagementService/RemoveDevice
/ManagementService/RotateWireguardKey
/ManagementService/AddApiAccessMethod
/ManagementService/SetApiAccessMethod
/ManagementService/GetCurrentApiAccessMethod
/ManagementService/GetSplitTunnelProcesses
/ManagementService/AddSplitTunnelProcess
/ManagementService/RemoveSplitTunnelProcess
/ManagementService/LogListen
/ManagementService/AppUpgrade
/ManagementService/AppUpgradeEventsListen
/RelaySelectorService/PartitionRelays
```

## Architectural distinction

| API surface      | Direction                   | Transport                  | Definition                           |
| ---------------- | --------------------------- | -------------------------- | ------------------------------------ |
| Mullvad REST API | Daemon → Mullvad services   | HTTPS/JSON                 | `mullvad-api/src`                    |
| Management API   | App frontend → local daemon | gRPC over local RPC socket | `mullvad-management-interface/proto` |
