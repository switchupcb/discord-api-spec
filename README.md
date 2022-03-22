# Discord API Spec
Discord does **NOT** provide an API Specification. This repository aims to provide a human and machine readable [Discord API](https://discord.com/developers/docs/reference) Specification. A [generator](CONTRIBUTING.md) is used to generate the specification (which is split into folders).

| File          | Description                                                                                      |
| :------------ | :----------------------------------------------------------------------------------------------- |
| `json`        | Contains valid JSON.                                                                             |
| `go`          | Contains converted Go structs _(inline/outline)_ from validated JSON.                            |
| `unconverted` | Unconverted TypeScript from [discord-api-types](https://github.com/discordjs/discord-api-types). |

## How is it Created?

Read [Contributing](CONTRIBUTING.md) for more information.

## What can I do with this?

A machine readable specification will allow you to generate language types _(i.e Go Structs)_ and other related code programmatically. This makes creating and updating a feature-complete [API Wrapper Library](https://discord.com/developers/docs/topics/community-resources#libraries-discord-libraries) low-effort and simple.

## What is the alternative?

The current alternative Discord provides is use an outdated contributor-led documentation. Discord API documentation does not provide JSON responses, and there is no guarantee that it's `Object` tables are up-to-date.

## License

The [Mozilla Public License 2.0](LICENSE) requires preservation of copyright and license notices (where the source is distributed). Modifications to this project must remain open sourced, but larger works that contain it may remain proprietary.