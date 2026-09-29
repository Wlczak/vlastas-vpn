# Vlasta'S VPN

The backend service for the Vlasta's VPN project or VPN

## Docs

- [API](https://wlczak.github.io/vlastas-vpn/)

## The road to the MVP

The current goal is to get the simplest version of this project as is realistically practical. The MVP aims for a release which supports live fetching and parsing of the Mullvad server list as well as creating a tunnel between connected devices and the Mullvad servers. The connection will be configurable from a simple HTML control panel or the public APi and will be accompanied by a release of the android app. The next stage of the project will likely aim at broadening platform support and user authentication for any practical use.

### MVP Update

The first goal was to release the VPN backend along with native clients, but since the whole project is built around Wireguard, which has famously a lot of native clients. I will be focusing on this backend until it is ready for release.

### 1st stage

- [ ] Working tunneling from devices to target VPN server (code implemented but not in use)
- [ ] Basic unstyled server dashboard for switching target server/device IP location (rework with some js)
- [ ] Fetching of Mullvad servers (make into timed cache instead of fetch calls)
- [ ] Public user API for changing the target Mullvad server (needs to be pretty solid)

### 2nd stage

- [ ] user accounts + auth (no user roles)

### 3rd stage

- [ ] Admin accounts
- [ ] user limits
- [ ] switching quota/points system
- [ ] support for multiple mullvad connections

## Client projects

### Android

- Vlasta's VPN for Android: https://github.com/Wlczak/vlastas-vpn-android (on hold until server is up to functional standard)
