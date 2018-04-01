# http-mirror

This is a **very** simple http proxy. Useful to mirror pages on your own domain.

More features may be added in the future, but none are planned.

## Usage

    http-mirror URL

- You have to specify an URL to connect to, otherwise it will use
  `https://example.com`.
- The URL has to include the protocol (`http://`, `https://`...), otherwise it
  won't work.

Example:

    http-mirror https://github.com
