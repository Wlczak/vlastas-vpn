# vlastas-vpn

The backend for the Vlasta's VPN service.

## The road to the MVP

The current goal is to get the simples version of this project as is realistically practical. The MVP aims for a release which supports live fetching and parsing of the Mullvad server list as well as creating a tunnel between connected devices and the Mullvad servers. The connection will be configurable from a simple HTML control panel or the public APi and will be accompanied by a release of the android app. The next stage of the project will likely aim at broadening platform support and user authentication for any practical use.

- [x] Working tunneling from devices to target VPN server
- [ ] Basic unstyled server dashboard for switching target server/device IP location
- [ ] Fetching of Mullvad servers
- [ ] Public user API for changing the target Mullvad server
