# cityu-cp37-keep-online

Login [CityU HK CP37 captive portal] to keep your host online

[CityU HK CP37 captive portal]: https://cp37.cs.cityu.edu.hk/cp

## Get started

First of all, you need to use the webpage of the [CityU HK CP37 captive portal] in GUI to login once, because it requires **2FA** to login.
While logging in, check the option to no longer requiring 2FA for this device, so that for this device you can just use the username and password to login later.
The "no 2FA later" mechanism is remembered by the remote SAML server and works until you login your account on another device **for the CP37 captive portal**, or say the unlocking is per-account-plus-domain.

Then install this package.
For Debian/Ubuntu, you can use the DEB package included in releases.
For other distributions, you can install this package as a [standard Golang project].

[standard Golang project]: https://github.com/golang-standards/project-layout

And install a Chromium/Chrome browser as a dependency.
You can either use the system package manager, e.g., for Debian/Ubuntu:

```bash
sudo apt-get install chromium-browser
```

Or check [the official Chromium download page] for the documentation.

[the official Chromium download page]: https://www.chromium.org/getting-involved/download-chromium/

One more comfirmation to make, which is usually nothing to do, is:
You need to make sure the installed systemd service `cp37-login.service` can find the Chromium/Chrome executable.
Check the `PATH` variable in the systemd service file `cp37-login.service`.

Finally reload the systemd services and start the timer, which will login every 5 minutes:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now cp37-login.timer
```

To check the status/logs of the login:

```bash
sudo systemctl status cp37-login.service
# Or more detailed logs:
sudo journalctl -u cp37-login.service
```

## Options

All options are set by environment variables.
Dotenv is also loaded from the `.env` file in the working directory.
Since the package is installed with systemd daemonization, you can override these environment variables in `/etc/default/cp37-login` or directly in the systemd service file `cp37-login.service`.

| Env name              | Description                                                                  | Default value or required                               |
| --------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------------- |
| `CP37_ROD_CTL_URL`    | [rod] control URL to connect to a Chromium/Chrome                            | Defaults to start a new Chromium/Chrome session by rod  |
| `CP37_ROD_BIN_PATH`   | Path to the Chromium/Chrome executable used to get a control URL.            | Defaults to lookup PATH. Used after `CP37_ROD_CTL_URL`. |
| `CP37_CITYU_USERNAME` | CityU CP37 captive portal username, which is the same as the [AIMS] username | Required                                                |
| `CP37_CITYU_PASSWORD` | CityU CP37 captive portal password, which is the same as the AIMS password   | Required                                                |
| `CP37_LOGIN_TIMEOUT`  | Timeout in the Golang duration string format for login                       | Defaults to `30s`, which is 30 seconds                  |

[rod]: https://github.com/go-rod/rod
[AIMS]: https://banweb.cityu.edu.hk

<!-- ## Development -->
<!-- TODO -->

## License

SPDX-License-Identifier: Apache-2.0 OR MIT

Copyright (C) myl7
